package helpers

import (
	"github.com/2xic/nordic-payment-formats/parser"
)

type Caller interface {
	Parse(*parser.Parser) ([]SimpleTransaction, error)
}

type SimpleTransaction struct {
	Kid                 string
	From_account_number string
	Amount              string
}
