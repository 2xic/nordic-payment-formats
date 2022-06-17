module norwegian-payment-formats

go 1.18

require(
    github.com/2xic/norwegian-payment-formats/ocr v1.0.0
    github.com/2xic/norwegian-payment-formats/parser v1.0.0
)

replace "github.com/2xic/norwegian-payment-formats/ocr" => "./ocr"
replace "github.com/2xic/norwegian-payment-formats/parser" => "./parser"
