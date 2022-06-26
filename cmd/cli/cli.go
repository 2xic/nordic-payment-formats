package main

import (
	"flag"
	"fmt"

	"github.com/2xic/nordic-payment-formats/combined"
	"github.com/2xic/nordic-payment-formats/generated"
	"github.com/2xic/nordic-payment-formats/parser"
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
	transactions := (combined.Try_to_parse(file_bytes))
	fmt.Println("======")
	print_transactions(transactions)
}

func print_transactions(transactions []generated.Transaction) {
	for index := range transactions {
		transaction := &transactions[index]
		fmt.Println(transaction)
	}
}
