package cremul

import (
	"testing"

	"github.com/2xic/nordic-payment-formats/parser"
)

func TestParse(t *testing.T) {
	var files [4]string
	files[0] = "cremul/test-files/CREMUL0001-utf-8.txt"
	files[1] = "cremul/test-files/CREMUL0002-utf-8.txt"
	files[2] = "cremul/test-files/CREMUL0003-utf-8.txt"
	files[3] = "cremul/test-files/CREMUL0001-utf-8-with-new-line.txt"

	for _, file := range files {
		data := parser.Read_file(parser.Get_path(file))
		Cremul{}.Parse(&parser.Parser{
			Data: data,
		})
	}
}
