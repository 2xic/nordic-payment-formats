package helpers

import "fmt"

func Require(actual string, expected string) {
	if actual != expected {
		panic(fmt.Sprintf("Bad value got %s, but expected %s", actual, expected))
	}
}
