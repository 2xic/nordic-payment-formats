GO_PATH=~/go
PATH=$PATH:/$GO_PATH/bin

SRC_DIR="./protocol-buffers/"
DST_DIR="./"
protoc  -I=$SRC_DIR/ \
        --go_out=$DST_DIR \
        --go-grpc_out=$DST_DIR \
        $SRC_DIR/service.proto $SRC_DIR/transaction.proto

go vet ./cmd/server
go build ./cmd/server
go build ./cmd/client
go build ./cmd/cli

go test github.com/2xic/nordic-payment-formats/parser
go test github.com/2xic/nordic-payment-formats/ocr
go test github.com/2xic/nordic-payment-formats/bgmax
go test github.com/2xic/nordic-payment-formats/cremul
