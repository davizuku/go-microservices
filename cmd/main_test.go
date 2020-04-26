package main

import (
	"fmt"
	"testing"

	"github.com/davizuku/go-microservices/client/client"
)

func TestOurClient(t *testing.T) {
	c := client.Default
	prods := c.Products.ListProducts()
	fmt.Println("Number of products found: ", len(prods))
}
