package parser

type Parser struct {
	Data  []byte
	index int
}

func (parser *Parser) Read_and_increment(length int) []byte {
	data := parser.read(length)
	parser.increment(length)
	return data
}

func (parser *Parser) Peek(length int) []byte {
	return parser.read(3)
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

func (parser *Parser) Done() bool {
	return len(parser.Data) == parser.index
}

func (parser *Parser) Index() int {
	return parser.index
}

func (parser *Parser) read(length int) []byte {
	if !(parser.index+length <= len(parser.Data)) {
		panic("Out of range")
	}
	return parser.Data[parser.index : parser.index+length]
}

func (parser *Parser) increment(length int) {
	parser.index += length
}
