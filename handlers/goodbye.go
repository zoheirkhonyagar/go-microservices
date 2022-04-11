package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	logger *log.Logger
}

func NewGoodbye(logger *log.Logger) *Goodbye {
	return &Goodbye{logger}
}

func (g *Goodbye) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	g.logger.Println("goodbye")
}
