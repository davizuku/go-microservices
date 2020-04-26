package main

import (
	"fmt"
	"testing"

	"github.com/davizuku/go-microservices/cmd/client"
	"github.com/davizuku/go-microservices/cmd/client/products"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:3000")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v", prod.GetPayload())
	t.Fail()
}
