package cremul

import (
	"testing"

	"github.com/2xic/norwegian-payment-formats/parser"
)

func TestParse(t *testing.T) {
	data := parser.Read_file(parser.Get_path("cremul/test-files/CREMUL0001-utf-8.txt"))
	Cremul{}.Parse(&parser.Parser{
		Data: data,
	})
}
