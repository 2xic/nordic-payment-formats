package parser

type Parser struct {
	data []byte
	read func(uint8) []byte
}

func Mul(x int, y int) int {
	return x * y
}
