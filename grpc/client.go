package main

import (
	"context"
	"fmt"

	protos "github.com/davizuku/go-microservices/grpc/protos/currency"
	"google.golang.org/grpc"
)

func main() {
	// gRPC client @see: https://grpc.io/docs/tutorials/basic/go/#client
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close() // deferred until the end of the execution of this function
	client := protos.NewCurrencyClient(conn)
	req := &protos.RateRequest{
		Base:        protos.Currencies_EUR,
		Destination: protos.Currencies_GBP,
	}
	res, err := client.GetRate(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Rate: ", res.Rate)
}
