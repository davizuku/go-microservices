package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// The id of this product
	//
	// required: true
	// min: 1
	ID int `json:"id"`

	// The name for this product
	//
	// required: true
	// min length: 1
	Name string `json:"name" validate:"required"`

	// The description of the product
	//
	// required: false
	// max length: 1000
	Description string `json:"description"`

	// The price of the product
	//
	// required: true
	// min: 0.01
	// max: 10.00
	Price float32 `json:"price" validate:"gt=0"`

	// the unique stock keeping unit (SKU) of the product
	//
	// required: true
	// pattern: [a-z0-9]+\-[a-z0-9]+\-[a-z0-9]+
	// example: abc-123-b4d
	SKU       string `json:"sku" validate:"required,sku"` // Internal id
	CreatedOn string `json:"-"`                           // Ommit these fields in JSON representation
	UpdatedOn string `json:"-"`                           // @see https://golang.org/pkg/encoding/json/#Marshal
	DeletedOn string `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-abcd-abcde
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	return len(matches) == 1
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

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

func DeleteProduct(id int) error {
	_, i, err := findProduct(id)
	if err != nil {
		return err
	}
	productList = append(productList[:i], productList[i+1])
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
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
