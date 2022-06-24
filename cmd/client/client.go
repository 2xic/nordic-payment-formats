package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/2xic/nordic-payment-formats/generated"
	"github.com/2xic/nordic-payment-formats/parser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	file := flag.String("input", "", "file to be parsed")
	flag.Parse()

	if len(*file) == 0 {
		panic("Bad input, use --input field to provide a file")
	}
	path := parser.Get_path(
		*file,
	)
	file_bytes := parser.Read_file(path)

	conn, err := grpc.Dial("localhost:4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := generated.NewPaymentParserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetTransactions(ctx, &generated.FileRequest{
		Content: string(file_bytes),
	})
	if err != nil {
		log.Fatalf("could not parse file: %v", err)
	}
	log.Printf("Transactions found: %s", r.Transactions)
}
