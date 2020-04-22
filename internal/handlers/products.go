package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/davizuku/go-microservices/internal/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle GET Product")
	lp := data.GetProducts()
	err := lp.ToJSON(res)
	if err != nil {
		http.Error(res, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POST Product")
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) UpdateProducts(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Unable to convert ID to int", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT Product", id)
	prod := &data.Product{}
	err = prod.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	err = data.UpdateProduct(id, prod)
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
