package parser

import (
	"io/ioutil"
)

func Read_file(path string) []byte {
	results, response := ioutil.ReadFile(path)
	if response != nil {
		panic("Empty content")
	} else {
		return results
	}
}
