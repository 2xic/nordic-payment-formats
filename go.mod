module norwegian-payment-formats

go 1.18

require (
	github.com/2xic/nordic-payment-formats/ocr v1.0.0
	github.com/2xic/nordic-payment-formats/parser v1.0.0
	github.com/2xic/nordic-payment-formats/bgmax v1.0.0
	github.com/2xic/nordic-payment-formats/helpers v1.0.0
	github.com/2xic/nordic-payment-formats/cremul v1.0.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.7.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/2xic/nordic-payment-formats/ocr => ./ocr
replace github.com/2xic/nordic-payment-formats/parser => ./parser
replace github.com/2xic/nordic-payment-formats/bgmax => ./bgmax
replace github.com/2xic/nordic-payment-formats/helpers => ./helpers
replace github.com/2xic/nordic-payment-formats/cremul => ./cremul
