package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	db, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Fatal(err)
	}

	nameList := []string{"ゆき", "みーくん", "くるみ", "りーさん", "めぐねえ", "太郎丸"}
	for i, v := range nameList {
		db.Do("SET", "character"+strconv.Itoa(i), v)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		randIndex := rand.Intn(len(nameList))
		name, _ := redis.String(db.Do("GET", "character"+strconv.Itoa(randIndex)))
		fmt.Fprintf(w, name)
	})
	http.ListenAndServe(":9000", nil)
}
