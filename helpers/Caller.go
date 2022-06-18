package helpers

import (
	"github.com/2xic/norwegian-payment-formats/parser"
)

type Caller interface {
	Parse(*parser.Parser) ([]SimpleTransaction, error)
}

type SimpleTransaction struct {
	Kid                 string
	From_account_number string
	Amount              string
}
