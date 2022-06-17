package ocr

import (
	"path/filepath"
	"runtime"
)

func Get_path(file string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Something went wrong")
	}
	exPath := filepath.Dir(filename)
	return exPath + "/test-files/ocr_giro_transaction.txt"
}
