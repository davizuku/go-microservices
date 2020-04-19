package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"` // Internal id
	CreatedOn   string  `json:"-"`   // Ommit these fields in JSON representation
	UpdatedOn   string  `json:"-"`   // @see https://golang.org/pkg/encoding/json/#Marshal
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	// Using Encode is better than using Marshal since it does not require
	// allocating extra memory for the transformed JSON data. Encoder writes it
	// directly to the Writer.
	// @see https://golang.org/pkg/encoding/json/#Marshal
	// @see https://golang.org/pkg/encoding/json/#Encoder
	return json.NewEncoder(w).Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
