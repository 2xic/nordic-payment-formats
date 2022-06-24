module norwegian-payment-formats

go 1.18

require (
	github.com/2xic/nordic-payment-formats/bgmax v1.0.0
	github.com/2xic/nordic-payment-formats/cremul v1.0.0
	github.com/2xic/nordic-payment-formats/generated v1.0.0
	github.com/2xic/nordic-payment-formats/helpers v1.0.0
	github.com/2xic/nordic-payment-formats/ocr v1.0.0
	github.com/2xic/nordic-payment-formats/parser v1.0.0
	github.com/2xic/nordic-payment-formats/combined v1.0.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.7.2 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect

)

replace github.com/2xic/nordic-payment-formats/ocr => ./ocr

replace github.com/2xic/nordic-payment-formats/parser => ./parser

replace github.com/2xic/nordic-payment-formats/bgmax => ./bgmax

replace github.com/2xic/nordic-payment-formats/helpers => ./helpers

replace github.com/2xic/nordic-payment-formats/cremul => ./cremul

replace github.com/2xic/nordic-payment-formats/generated => ./generated

replace github.com/2xic/nordic-payment-formats/combined => ./combined
