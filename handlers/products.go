package handlers

import (
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
	if request.Method == http.MethodGet {
		products.getProducts(responseWriter, request)
		return
	}

	// handle update
	if request.Method == http.MethodPut {

	}

	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(responseWriter http.ResponseWriter, request *http.Request) {
	listOfProducts := data.GetProducts()

	err := listOfProducts.ToJSON(responseWriter)
	// data, err := json.Marshal(listOfProducts)

	if err != nil {
		http.Error(responseWriter, "unable to handle marshal json", http.StatusInternalServerError)
	}
}
