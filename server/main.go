package main

import (
	"log"
	"net"

	"github.com/mattn/lile-example/gene9go"
	"github.com/mattn/lile-example/gene9go/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":11111")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()

	gene9go.RegisterGene9goServer(srv, &server.Gene9goServer{})
	srv.Serve(lis)
}
