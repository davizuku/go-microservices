package main

import (
	"net"
	"os"

	protos "github.com/davizuku/go-microservices/grpc/protos/currency"
	"github.com/davizuku/go-microservices/grpc/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewCurrency(log)
	protos.RegisterCurrencyServer(gs, cs)
	// Enable Reflection API to list the available services in the server
	reflection.Register(gs)

	listener, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	log.Info("GRPC server listening on 3001")
	gs.Serve(listener)
}
