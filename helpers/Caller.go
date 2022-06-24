package helpers

import (
	"github.com/2xic/nordic-payment-formats/generated"
	"github.com/2xic/nordic-payment-formats/parser"
)

type Caller interface {
	Parse(*parser.Parser) ([]generated.Transaction, error)
}
