package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("hello world")
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye")
	})

	http.ListenAndServe(":9090", nil)

}
