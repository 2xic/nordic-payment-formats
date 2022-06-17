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

func (parser *Parser) read(length int) []byte {
	if !(parser.index+length < len(parser.Data)) {
		panic("Out of range")
	}
	return parser.Data[parser.index : parser.index+length]
}

func (parser *Parser) increment(length int) {
	parser.index += length
}
