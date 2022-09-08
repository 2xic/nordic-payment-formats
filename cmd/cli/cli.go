package main

import (
	"flag"
	"fmt"

	"encoding/json"
	"strings"

	"github.com/2xic/nordic-payment-formats/combined"
	"github.com/2xic/nordic-payment-formats/generated"
	"github.com/2xic/nordic-payment-formats/parser"
)

func main() {
	file := flag.String("input", "", "file to be parsed")
	output := flag.String("output", "print", "output format")
	filterAccountNumber := flag.String("filter-account-number", "", "filter account number")
	flag.Parse()

	if len(*file) == 0 {
		panic("Bad input, use --input field to provide a file")
	}
	path := parser.Get_path(
		*file,
	)
	file_bytes := parser.Read_file(path)
	transactions := filter_transactions(combined.Try_to_parse(file_bytes), filterAccountNumber)

	if *output == "print" {
		fmt.Println("======")
		print_transactions(transactions)
	} else if *output == "json" {
		encodedJson := new(strings.Builder)
		json.NewEncoder(encodedJson).Encode(transactions)
		print(encodedJson.String())
	}
}

func print_transactions(transactions []generated.Transaction) {
	for index := range transactions {
		transaction := &transactions[index]
		fmt.Println(transaction)
	}
}

func filter_transactions(transactions []generated.Transaction, filterAccountNumber *string) []generated.Transaction {
	var filtered []generated.Transaction
	for index := range transactions {
		transaction := &transactions[index]
		if 0 < len(*filterAccountNumber) && transaction.ToBankAccountNumber == *filterAccountNumber {
			filtered = append(filtered, *transaction)
		}
	}

	return filtered
}
