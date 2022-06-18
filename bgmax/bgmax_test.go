package bgmax

import (
	"testing"

	"github.com/2xic/norwegian-payment-formats/parser"
)

func TestParse(t *testing.T) {
	BgMax_Parser(&parser.Parser{
		Data: parser.Read_file(parser.Get_path("bgmax/test-files/example_1.txt")),
	})
}
