package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/lileio/lile"
	"github.com/mattn/lile-example/gene9go"
)

var s = Gene9goServer{}
var cli gene9go.Gene9goClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		gene9go.RegisterGene9goServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = gene9go.NewGene9goClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
