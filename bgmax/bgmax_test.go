package bgmax

import (
	"testing"

	"github.com/2xic/nordic-payment-formats/parser"
)

func TestParse(t *testing.T) {
	var files [4]string
	files[0] = "bgmax/test-files/bangiroinbetalningar_exempelfil_avtalutantillval_kontroller_sv.txt"
	files[1] = "bgmax/test-files/bankgiroinbetalningar_bildfil_avtal-ocr-kontroll_utokad-blankettregistrering-ocr-kontroll_exempelfil.txt"
	files[2] = "bgmax/test-files/bankgiroinbetalningar_avtal-ocr-kontroll_utokad-blankettregistrering-ocr-kontroll_exempelfil.txt"
	files[3] = "bgmax/test-files/bankgiroinbetalningar_exempelfil_avtal-om-ocr-kontroll_sv.txt"

	for _, file := range files {
		data := parser.Read_file(parser.Get_path(file))
		BgMax{}.Parse(&parser.Parser{
			Data: data,
		})
	}
}
