package main

import (
	"context"
	"fmt"
	"net"

	"github.com/2xic/nordic-payment-formats/combined"
	"github.com/2xic/nordic-payment-formats/generated"

	"google.golang.org/grpc"
)

type Server struct {
	generated.PaymentParserServer
}

func RemoveIndex(s []generated.Transaction, index int) []generated.Transaction {
	return append(s[:index], s[index+1:]...)
}

func (s *Server) GetTransactions(ctx context.Context, in *generated.FileRequest) (*generated.Response, error) {
	transactions := combined.Try_to_parse([]byte(in.Content))
	var transactionConverted []*generated.Transaction
	for 0 < len(transactions) {
		transactionConverted = append(transactionConverted, &transactions[0])
		transactions = RemoveIndex(transactions, 0)
	}

	return &generated.Response{
		Transactions: transactionConverted,
	}, nil
}

func main() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	generated.RegisterPaymentParserServer(grpcServer, &Server{})
	listener, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		panic(err)
	}

	fmt.Println(listener.Addr())
	fmt.Println("Ready to serve :)")

	grpcServer.Serve(listener)
	fmt.Println(listener.Addr())
}
