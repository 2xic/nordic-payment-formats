package ocr

import (
	"testing"

	"github.com/2xic/nordic-payment-formats/parser"
)

func TestParse(t *testing.T) {
	var files [4]string
	files[0] = "ocr/test-files/ocr_giro_transaction.txt"
	files[1] = "ocr/test-files/ocr_1.txt"
	files[2] = "ocr/test-files/ocr_3.txt"
	files[3] = "ocr/test-files/ocr_giro_transaction-with-new-line.txt"

	for _, file := range files {
		data := parser.Read_file(parser.Get_path(file))
		Ocr{}.Parse(&parser.Parser{
			Data: data,
		})
	}
}
