package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
	endpoints()
	log.Fatal(http.ListenAndServe(":10000", nil))

}

func endpoints() {
	http.HandleFunc("/insertJSON", get_data)
}

func get_data(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		datos := string(body)

		suscribeRedis(datos)

		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

func suscribeRedis(datos string) {

	if err := redisClient.Publish(ctx, "pruebas", datos).Err(); err != nil {
		panic(err)
	}

}
