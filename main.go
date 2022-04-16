package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/zoheirkhonyagar/go-microservices/handlers"
)

func main() {

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	productsHandler := handlers.NewProducts(logger)

	gh := handlers.NewGoodbye(logger)

	sm := http.NewServeMux()

	sm.Handle("/", productsHandler)

	sm.Handle("/goodbye", gh)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			logger.Fatal(err)
		}

	}()

	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, os.Interrupt)

	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)

	//graceful shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
