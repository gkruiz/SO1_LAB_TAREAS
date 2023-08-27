package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/get_memory_data", get_data)
	log.Fatal(http.ListenAndServe(":10000", nil))

}

func get_data(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	ram := repe()

	fmt.Fprintf(w, ram)

}

func repe() string {

	ram, err := ioutil.ReadFile("/modulos/ram_201603009")

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(ram))

	return string(ram)
}
