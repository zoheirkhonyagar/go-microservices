package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		log.Println("hello world")

		data, error := ioutil.ReadAll(request.Body)

		if error != nil {
			http.Error(responseWriter, "Oops", http.StatusBadRequest)
			return
		}

		log.Printf("Data %s", data)

		fmt.Fprintf(responseWriter, "Hello %s", data)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye")
	})

	http.ListenAndServe(":9090", nil)

}
