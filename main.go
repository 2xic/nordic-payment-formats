package main

import (
	"fmt"

	"github.com/2xic/norwegian-payment-formats/ocr"
	"github.com/2xic/norwegian-payment-formats/parser"
)

func main() {
	path := parser.Get_path(
		"ocr_nets_example.txt",
	)
	ocr_file := parser.Read_file(path)

	fmt.Println(ocr_file)
	transaction := ocr.Ocr_Parser(&parser.Parser{
		Data: ocr_file,
	})
	fmt.Println(transaction)

	//	fmt.Print(parser.Mul(2, 4))
	//	ocr.Sum(2, 2)
}
