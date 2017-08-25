package main

import (
	"github.com/lileio/lile"
	"github.com/mattn/lile-example/gene9go"
	"github.com/mattn/lile-example/gene9go/gene9go/cmd"
	"github.com/mattn/lile-example/gene9go/server"
	"google.golang.org/grpc"
)

func main() {
	s := &server.Gene9goServer{}

	lile.Name("gene9go")
	lile.Server(func(g *grpc.Server) {
		gene9go.RegisterGene9goServer(g, s)
	})

	cmd.Execute()
}
