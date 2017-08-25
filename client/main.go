package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mattn/lile-example/gene9go"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:11111", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := gene9go.NewGene9goClient(conn)
	req := &gene9go.Request{
		Word: os.Args[1],
	}
	resp, err := client.Translate(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Text)
}
