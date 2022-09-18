package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"io/ioutil"
	"github.com/2xic/nordic-payment-formats/ocr"
	"github.com/2xic/nordic-payment-formats/parser"

	"encoding/json"
	"strings"

)

func rootEndpoint(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Use OCR path to parse OCR file\n")
}

func ocrEndpoint(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
		io.WriteString(w, "could not read body\n")
	} else {
		ocr := ocr.Ocr{}
		transactions, _ := ocr.Parse(&parser.Parser{
			Data: body,
		})
		encodedJson := new(strings.Builder)
		json.NewEncoder(encodedJson).Encode(transactions)

		io.WriteString(w, encodedJson.String())
	}
}

func main() {
	http.HandleFunc("/", rootEndpoint)
	http.HandleFunc("/ocr", ocrEndpoint)

	fmt.Println("Listening on localhost:3333")

	err := http.ListenAndServe(":3333", nil)

  	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

