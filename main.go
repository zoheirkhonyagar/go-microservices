package main

import (
	"log"
	"net/http"
)

func main() {

	hh := handlers
	// http.HandleFunc("/", handlers.Hello{})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye")
	})

	http.ListenAndServe(":9090", nil)

}
