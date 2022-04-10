package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{logger: l}
}

func (h *Hello) ServeHTTP(responseWriter http.ResponseWriter, request http.Request) {
	h.logger.Println("hello world")

	data, error := ioutil.ReadAll(request.Body)

	if error != nil {
		http.Error(responseWriter, "Oops", http.StatusBadRequest)
		return
	}

	log.Printf("Data %s", data)

	fmt.Fprintf(responseWriter, "Hello %s", data)
}
