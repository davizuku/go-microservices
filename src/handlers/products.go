package handlers

import (
	"data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		p.getProducts(res, req)
		return
	}
	if req.Method == http.MethodPost {
		p.addProduct(res, req)
		return
	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle GET Product")
	lp := data.GetProducts()
	err := lp.ToJSON(res)
	if err != nil {
		http.Error(res, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POST Product")
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
