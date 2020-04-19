package handlers

import (
	"data"
	"encoding/json"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(res http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(res, "Unable to marshal json", http.StatusInternalServerError)
	}
	res.Write(d)
}
