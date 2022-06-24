go vet
go build
SRC_DIR="./helpers/"
DST_DIR="./"
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/transaction.proto

go test github.com/2xic/nordic-payment-formats/parser
go test github.com/2xic/nordic-payment-formats/ocr
go test github.com/2xic/nordic-payment-formats/bgmax
go test github.com/2xic/nordic-payment-formats/cremul

