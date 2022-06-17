module norwegian-payment-formats

go 1.18

require (
	github.com/2xic/norwegian-payment-formats/ocr v1.0.0
	github.com/2xic/norwegian-payment-formats/parser v1.0.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.7.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/2xic/norwegian-payment-formats/ocr => ./ocr

replace github.com/2xic/norwegian-payment-formats/parser => ./parser