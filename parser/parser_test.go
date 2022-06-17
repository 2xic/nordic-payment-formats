package parser

import "testing"

func test_read_and_increment(t *testing.T) {
	fake_array := []byte{0, 2, 3}
	parser := Parser{
		Data: fake_array,
	}
	value := parser.Read_and_increment(1)
	if len(value) != 1 {
		t.Errorf("Wrong read")
	} else if value[0] != 0 {
		t.Errorf("Wrong read")
	} else if parser.index != 1 {
		t.Errorf("Wrong read")
	}
}
