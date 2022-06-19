package ocr

import (
	"testing"

	"github.com/2xic/nordic-payment-formats/parser"
)

func TestParse(t *testing.T) {
	Ocr{}.Parse(&parser.Parser{
		Data: parser.Read_file(parser.Get_path("ocr/test-files/ocr_giro_transaction.txt")),
	})
	/*
		Ocr_Parser(&parser.Parser{
			Data: parser.Read_file(parser.Get_path("ocr/test-files/ocr_nets_example.txt")),
		})*/
}
