FROM golang:1.18-bullseye

WORKDIR /app

COPY . ./

RUn apt-get update
RUN apt-get install -y protobuf-compiler 
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
RUN export PATH="$PATH:$(go env GOPATH)/bin"

RUN go mod download

RUN chmod +x ./pipeline.sh

RUN ./pipeline.sh

ENTRYPOINT ["./cli"]

# docker build --tag docker-payments-format-parser .
# docker run -v=$(pwd):/host/root docker-payments-format-parser --input ./host/root/ocr-file

