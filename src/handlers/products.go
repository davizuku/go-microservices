package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/davizuku/go-microservices/data"
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
	if req.Method == http.MethodPut {
		p.l.Println("Parse PUT Product URL")
		// expect the id to be in the URI
		regex := regexp.MustCompile(`/([0-9]+)`)
		groups := regex.FindAllStringSubmatch(req.URL.Path, -1)
		if len(groups) != 1 || len(groups[0]) != 2 {
			http.Error(res, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := groups[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(res, "Could not convert string ID to int", http.StatusBadRequest)
			return
		}
		p.updateProducts(id, res, req)
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

func (p *Products) updateProducts(id int, res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle PUT Product")
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Unable to unmarshal json", http.StatusBadRequest)
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
