package parser

import "fmt"

type Parser struct {
	Data   []byte
	index  int
	Tokens int
}

func (parser *Parser) Read_and_increment(length int) []byte {
	data, change, read := parser.read(length)
	parser.increment(
		max(change, length),
	)
	parser.Tokens += read
	return data
}

func (parser *Parser) Peek(length int) []byte {
	bytes, _, _ := parser.read(length)
	return bytes
}

func (parser *Parser) Len() int {
	return len(parser.Data)
}

func (parser *Parser) ReadUntil(char byte, rollback_char byte) []byte {
	var response []byte
	for true {
		if parser.Done() {
			break
		}
		read_char := parser.Read_and_increment(1)
		if read_char[0] == char {
			break
		}
		if read_char[0] == rollback_char {
			parser.increment(-1)
			break
		}
		response = append(response, read_char...)
	}
	return response
}

func (parser *Parser) IsPrevByte(char byte) bool {
	return parser.Data[parser.index - 2] == char;
}

func (parser *Parser) Done() bool {
	return len(parser.Data) <= parser.index+1
}

func (parser *Parser) Index() int {
	return parser.index
}

func (parser *Parser) read(length int) ([]byte, int, int) {
	var parsed []byte
	change := 0
	read := 0
	for read < length {
		index := parser.index + change
		if !(index <= len(parser.Data)-1) {
			break
		}
		read_byte := parser.Data[index]
		if string(read_byte) != "\n" {
			parsed = append(parsed, read_byte)
			read++
		}
		change++
	}
	return parsed, change, read
}

func (parser *Parser) increment(length int) {
	parser.index += length
}

func (parser *Parser) Validate(expectedDelta int) {
	if parser.Tokens%expectedDelta != 0 {
		panic(
			fmt.Sprintf(
				"Something is wrong with the parser (index %d)",
				parser.Tokens,
			),
		)
	}
}

func (parser *Parser) AutoAddPadding(lineLength int) *Parser {
	var parsed []byte
	lengthTracker := 0
	funcPadding := func() {
		for i := 0; i < lengthTracker%lineLength; i++ {
			parsed = append(parsed, byte(" "[0]))
			lengthTracker++
		}
	}
	for !parser.Done() {
		token := parser.Data[parser.index]
		if string(token) != "\n" {
			parsed = append(parsed, token)
			lengthTracker++
		} else {
			funcPadding()
		}
		parser.index++
	}
	funcPadding()

	return &Parser{
		Data:   parsed,
		index:  0,
		Tokens: 0,
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
