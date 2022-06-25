package helpers

import "fmt"

func Require[T comparable](actual T, expected T) {
	if actual != expected {
		panic(fmt.Sprintf("Bad value got %+v, but expected %+v", actual, expected))
	}
}
