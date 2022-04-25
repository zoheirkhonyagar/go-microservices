package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreateOn    string  `json:"-"`
	UpdateOn    string  `json:"-"`
	DeleteOn    string  `json:"-"`
}

type Products []*Product

func (prodcuts *Products) ToJSON(writer io.Writer) error {
	e := json.NewEncoder(writer)

	return e.Encode(prodcuts)
}

func (products *Product) FromJSON(reader io.Reader) error {
	d := json.NewDecoder(reader)

	return d.Decode(products)
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc 45",
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fj33",
		CreateOn:    time.Now().String(),
		UpdateOn:    time.Now().String(),
	},
}

func GetProducts() Products {
	return productList
}

func AddProduct(product *Product) {
	product.ID = getNextID()

	productList = append(productList, product)
}

func getNextID() int {

	productsListLength := len(productList)

	if productsListLength == 0 {
		return 1
	}

	p := productList[productsListLength-1]

	return p.ID + 1
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for index, product := range productList {
		if product.ID == id {
			return product, index, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func UpdateProduct(id int, product *Product) error {
	_, index, err := findProduct(id)
	if err != nil {
		return err
	}

	product.ID = id
	productList[index] = product
	return nil
}
