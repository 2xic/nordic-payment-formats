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
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)

require (
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.0 // indirect

)

replace github.com/2xic/nordic-payment-formats/ocr => ./ocr

replace github.com/2xic/nordic-payment-formats/parser => ./parser

replace github.com/2xic/nordic-payment-formats/bgmax => ./bgmax

replace github.com/2xic/nordic-payment-formats/helpers => ./helpers

replace github.com/2xic/nordic-payment-formats/cremul => ./cremul

replace github.com/2xic/nordic-payment-formats/generated => ./generated

replace github.com/2xic/nordic-payment-formats/combined => ./combined
