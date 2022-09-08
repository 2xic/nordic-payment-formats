module norwegian-payment-formats

go 1.18

require (
	github.com/2xic/nordic-payment-formats/combined v1.0.0
	github.com/2xic/nordic-payment-formats/generated v1.0.0
	github.com/2xic/nordic-payment-formats/parser v1.0.0
)

require (
	github.com/2xic/nordic-payment-formats/bgmax v1.0.0 // indirect
	github.com/2xic/nordic-payment-formats/cremul v1.0.0 // indirect
	github.com/2xic/nordic-payment-formats/helpers v1.0.0 // indirect
	github.com/2xic/nordic-payment-formats/ocr v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220907135653-1e95f45603a7 // indirect
	golang.org/x/sys v0.0.0-20220907062415-87db552b00fd // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220902135211-223410557253 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

require google.golang.org/grpc v1.49.0

replace github.com/2xic/nordic-payment-formats/ocr => ./ocr

replace github.com/2xic/nordic-payment-formats/parser => ./parser

replace github.com/2xic/nordic-payment-formats/bgmax => ./bgmax

replace github.com/2xic/nordic-payment-formats/helpers => ./helpers

replace github.com/2xic/nordic-payment-formats/cremul => ./cremul

replace github.com/2xic/nordic-payment-formats/generated => ./generated

replace github.com/2xic/nordic-payment-formats/combined => ./combined
