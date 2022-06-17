package ocr

import "testing"
import "github.com/2xic/norwegian-payment-formats/parser"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestMul(t *testing.T) {
	total := parser.Mul(5, 5)
	if total != 2 {
		t.Errorf("Mul was incorrect, got: %d, want: %d.", total, 25)
	}
}
