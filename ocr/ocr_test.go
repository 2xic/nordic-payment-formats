package ocr

import (
	"testing"

	"github.com/2xic/norwegian-payment-formats/parser"
)

func TestParse(t *testing.T) {
	Ocr_Parser(&parser.Parser{
		Data: parser.Read_file(Get_path("./test-files/ocr_giro_transaction.txt")),
	})
	Ocr_Parser(&parser.Parser{
		Data: parser.Read_file(Get_path("./test-files/ocr_nets_example.txt")),
	})
}
