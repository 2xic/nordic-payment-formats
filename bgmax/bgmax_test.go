package bgmax

import (
	"fmt"
	"testing"

	"github.com/2xic/nordic-payment-formats/parser"
)

func TestParse(t *testing.T) {
	data := parser.Read_file(parser.Get_path("bgmax/test-files/example_1.txt"))
	if len(data) != 2240 {
		panic(
			fmt.Sprintf(
				"Wrong encoding %d", len(data),
			),
		)
	}
	BgMax{}.Parse(&parser.Parser{
		Data: data,
	})
}
