package parser

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
)

func Get_path(file string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Something went wrong")
	}
	caller_path := filepath.Dir(filename)
	path := path.Join(
		caller_path, "../", file,
	)
	fmt.Print(path)
	return path
}
