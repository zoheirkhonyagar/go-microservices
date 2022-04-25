package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if request.Method == http.MethodPut {
		products.logger.Println("PUT update method")
		// products.update()
		path := request.URL.Path
		regex := regexp.MustCompile(`/([0-9]+)`)
		group := regex.FindAllStringSubmatch(path, -1)

		if len(group) != 1 {
			products.logger.Println("Invalid URI more than one id")
			http.Error(responseWriter, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			products.logger.Println("Invalid URI more than one capture group")
			http.Error(responseWriter, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := group[0][1]

		id, _ := strconv.Atoi(idString)

		products.logger.Println(id)

		products.updateProducts(id, responseWriter, request)
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

func (p *Products) updateProducts(id int, responseWriter http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handle PUT Product")

	prod := &data.Product{}

	err := prod.FromJSON(request.Body)

	if err != nil {
		http.Error(responseWriter, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(responseWriter, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(responseWriter, "Product not found", http.StatusInternalServerError)
		return
	}

}
