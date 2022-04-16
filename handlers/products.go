package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zoheirkhonyagar/go-microservices/product-api/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (products *Products) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	listOfProducts := data.GetProducts()

	data, err := json.Marshal(listOfProducts)

	if err != nil {
		http.Error(responseWriter, "unable to handle marshal json", http.StatusInternalServerError)
	}

	responseWriter.Write(data)

}
