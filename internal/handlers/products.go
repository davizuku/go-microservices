// Package handlers Product API
//
// Documentation for Product API
//
// 	Schemes: http
// 	BasePath: /
// 	Version: 1.0.0
//
// 	Consumes:
// 	- application/json
//
// 	Produces:
// 	- application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/davizuku/go-microservices/internal/data"
	"github.com/gorilla/mux"
)

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct { // This is only used for swagger documentation
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the data store
	// in: path
	// required: true
	ID int `"json:id"`
}

// swagger:response noContent
type productsNoContent struct{}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 	200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle GET Product")
	lp := data.GetProducts()
	err := lp.ToJSON(res)
	if err != nil {
		http.Error(res, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p Products) AddProduct(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POST Product")
	prod := req.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(&prod)
}

func (p Products) UpdateProducts(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Unable to convert ID to int", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT Product", id)
	prod := req.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(res, "Could not update the product", http.StatusInternalServerError)
		return
	}
	p.l.Printf("Prod: %#v", prod)
}

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// responses:
// 	201: noContent

// DeleteProduct removes a product from the data store
func (p Products) DeleteProduct(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Unable to convert ID to int", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle DELETE Product", id)
	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(res, "Could not delete the product", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(req.Body)
		if err != nil {
			http.Error(res, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}
		err = prod.Validate()
		if err != nil {
			http.Error(
				res,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		ctx := context.WithValue(req.Context(), KeyProduct{}, prod)
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	})
}
