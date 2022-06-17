package main

import (
	"fmt"

	"github.com/2xic/norwegian-payment-formats/ocr"
	"github.com/2xic/norwegian-payment-formats/parser"
)

func main() {
	fmt.Print(parser.Mul(2, 4))
	ocr.Sum(2, 2)
}
