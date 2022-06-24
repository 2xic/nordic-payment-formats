package helpers

import (
	"github.com/2xic/nordic-payment-formats/parser"
)

type Caller interface {
	Parse(*parser.Parser) ([]Transaction, error)
}
