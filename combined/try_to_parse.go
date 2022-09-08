package combined

import (
	"log"

	"github.com/2xic/nordic-payment-formats/bgmax"
	"github.com/2xic/nordic-payment-formats/cremul"
	"github.com/2xic/nordic-payment-formats/generated"
	"github.com/2xic/nordic-payment-formats/helpers"
	"github.com/2xic/nordic-payment-formats/ocr"
	"github.com/2xic/nordic-payment-formats/parser"
)

func Try_to_parse(file_bytes []byte) []generated.Transaction {
	var parsers [3]helpers.Caller
	parsers[0] = bgmax.BgMax{}
	parsers[1] = ocr.Ocr{}
	parsers[2] = cremul.Cremul{}

	for _, local_parser := range parsers {
		transactions, _ := try_parser(file_bytes, local_parser)

		if len(transactions) > 0 {
			return transactions
		}
	}
	return []generated.Transaction{}
}

func try_parser(file []byte, local_parser helpers.Caller) ([]generated.Transaction, error) {
	// TODO : Add proper error handling inside each parser.
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err, local_parser.Name())
		}
	}()

	transactions, failure := local_parser.Parse(&parser.Parser{
		Data: file,
	})

	return transactions, failure
}
