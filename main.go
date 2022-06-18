package main

import (
	"flag"
	"fmt"

	"github.com/2xic/norwegian-payment-formats/bgmax"
	"github.com/2xic/norwegian-payment-formats/helpers"
	"github.com/2xic/norwegian-payment-formats/ocr"
	"github.com/2xic/norwegian-payment-formats/parser"
)

func main() {
	file := flag.String("input", "", "file to be parsed")
	flag.Parse()

	if len(*file) == 0 {
		panic("Bad input")
	}
	path := parser.Get_path(
		*file,
	)
	ocr_file := parser.Read_file(path)

	var parsers [2]helpers.Caller
	parsers[0] = bgmax.BgMax{}
	parsers[1] = ocr.Ocr{}

	for _, local_parser := range parsers {
		transactions, _ := try_parser(ocr_file, local_parser)

		if len(transactions) > 0 {
			fmt.Println(transactions)
			break
		}
	}
}

func try_parser(file []byte, local_parser helpers.Caller) ([]helpers.SimpleTransaction, error) {
	// TODO : Add proper error handling inside each parser.
	defer func() {
		if err := recover(); err != nil {
			//log.Println("panic occurred:", err)
		}
	}()

	transactions, failure := local_parser.Parse(&parser.Parser{
		Data: file,
	})

	return transactions, failure
}
