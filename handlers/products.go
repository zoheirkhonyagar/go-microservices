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
	if request.Method == http.MethodPost {
		products.addProduct(responseWriter, request)
		return
	}

	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(responseWriter http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handle GET Products")

	listOfProducts := data.GetProducts()

	err := listOfProducts.ToJSON(responseWriter)
	// data, err := json.Marshal(listOfProducts)

	if err != nil {
		http.Error(responseWriter, "unable to handle marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(responseWriter http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handle POST Products")

	prod := &data.Product{}

	err := prod.FromJSON(request.Body)

	if err != nil {
		http.Error(responseWriter, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.logger.Printf("Prod: %#v", prod)

	data.AddProduct(prod)
}
