package main

import (
	"log"
	"net/http"
	"os"

	"github.com/zoheirkhonyagar/go-microservices/handlers"
)

func main() {

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(logger)

	gh := handlers.NewGoodbye(logger)

	sm := http.NewServeMux()

	sm.Handle("/", hh)

	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9090", sm)

}
