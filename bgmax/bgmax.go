package bgmax

import (
	"fmt"
	"math/big"

	"github.com/2xic/nordic-payment-formats/generated"
	"github.com/2xic/nordic-payment-formats/helpers"
	"github.com/2xic/nordic-payment-formats/parser"
)

type BgMax struct {
	helpers.Caller
}

func (BgMax) Parse(inputParser *parser.Parser) ([]generated.Transaction, error) {
	var txs []generated.Transaction
	parser := inputParser.AutoAddPadding(80)
	currency := ""
	for !parser.Done() {
		transaction_code := string(parser.Read_and_increment(2))

		if transaction_code == "01" {
			header := parse_header(parser)
			helpers.Require(header.version, "01")
		} else if transaction_code == "05" {
			section := parse_section_header(parser)
			currency = section.currency
		} else if transaction_code == "20" || transaction_code == "21" || transaction_code == "22" || transaction_code == "23" {
			transaction := parse_payment_section(parser, transaction_code)
			if transaction.transaction_code == "22" && transaction.payment_amount.String() == "0" {
				txs[len(txs)-1].Reference = transaction.reference
			} else {
				txs = append(txs, generated.Transaction{
					FromAccountNumber: transaction.from_bank_giro_number,
					Amount:            transaction.payment_amount.String(),
					Kid:               "",
					Reference:         helpers.Trim(transaction.reference),
					Payer:             &generated.Payer{},
					Currency:          currency,
					TransactionCode:   transaction.transaction_code,
				})
			}
		} else if transaction_code == "25" {
			parse_information_post(parser)
		} else if transaction_code == "26" {
			payer := parse_name_post(parser)
			txs[len(txs)-1].Payer.Name = helpers.Trim(payer.name)
		} else if transaction_code == "27" {
			payer := parse_address_1(parser)
			txs[len(txs)-1].Payer.Address1 = helpers.Trim(payer.address1)
		} else if transaction_code == "28" {
			payer := parse_address_2(parser)
			txs[len(txs)-1].Payer.Address2 = helpers.Trim(payer.local_address)
			txs[len(txs)-1].Payer.CountryCode = helpers.Trim(payer.country_code)
		} else if transaction_code == "29" {
			parse_organizations_number(parser)
		} else if transaction_code == "15" {
			parse_deposit_section(parser)
		} else if transaction_code == "70" {
			parse_tail_header(parser)
		} else {
			next_bytes := parser.Read_and_increment(8)
			panic(
				fmt.Sprintf(
					"Unknown state (tx code '%s', next '%s')", transaction_code, next_bytes,
				),
			)
		}
		parser.Validate(80)
	}

	parser.Validate(80)

	return txs, nil
}

func parse_header(parser *parser.Parser) StartHeader {
	layout_name := parser.Read_and_increment(20)
	version := parser.Read_and_increment(2)
	timestamp := parser.Read_and_increment(20)
	is_production := parser.Read_and_increment(1)

	// filler
	parser.Read_and_increment(35)

	return StartHeader{
		layout_name:   string(layout_name),
		version:       string(version),
		timestamp:     string(timestamp),
		is_production: string(is_production) == "P",
	}
}

func parse_section_header(parser *parser.Parser) SectionHeader {
	to_giro_number := parser.Read_and_increment(10)
	to_giro_plus_number := parser.Read_and_increment(10)
	currency := parser.Read_and_increment(3)

	// filler
	parser.Read_and_increment(55)

	return SectionHeader{
		to_giro_number:      string(to_giro_number),
		to_plus_giro_number: string(to_giro_plus_number),
		currency:            string(currency),
	}
}

func parse_payment_section(parser *parser.Parser, transaction_code string) SectionPayment {
	from_bank_giro_number := parser.Read_and_increment(10)
	reference := parser.Read_and_increment(25)
	payment_amount := parser.Read_and_increment(18)

	// TODO : Add metadata from table 5
	reference_code := parser.Read_and_increment(1)

	// TODO : Add metadata from table 6
	payment_channel_code := parser.Read_and_increment(1)

	bgc_identifier := parser.Read_and_increment(12)

	// "Avibildmarkering" -> brevgiro ?
	has_photo := parser.Read_and_increment(1)

	// filler
	if string(transaction_code) == "21" {
		// deduction code
		parser.Read_and_increment(1)

		parser.Read_and_increment(9)
	} else {
		// case 20, 22, 23
		parser.Read_and_increment(10)
	}

	return SectionPayment{
		transaction_code:      string(transaction_code),
		from_bank_giro_number: string(from_bank_giro_number),
		reference:             string(reference),
		payment_amount:        helpers.ConvertToBigInt(string(payment_amount)),
		reference_code:        string(reference_code),
		payment_channel_code:  string(payment_channel_code),
		bgc_identifier:        string(bgc_identifier),
		has_photo:             string(has_photo) == "1",
	}
}

func parse_information_post(parser *parser.Parser) InformationPost {
	text := string(parser.Read_and_increment(50))
	// filler
	parser.Read_and_increment(28)

	return InformationPost{
		text: text,
	}
}

func parse_name_post(parser *parser.Parser) NamePost {
	name := string(parser.Read_and_increment(35))
	extra_name := string(parser.Read_and_increment(35))
	// filler
	parser.Read_and_increment(8)

	return NamePost{
		name:       name,
		extra_name: extra_name,
	}
}

func parse_address_1(parser *parser.Parser) Address1 {
	address1 := string(parser.Read_and_increment(35))
	zip_code := string(parser.Read_and_increment(9))
	// filler
	parser.Read_and_increment(34)

	return Address1{
		address1: address1,
		zip_code: zip_code,
	}
}

func parse_address_2(parser *parser.Parser) Address2 {
	address := string(parser.Read_and_increment(35))
	payer_country := string(parser.Read_and_increment(35))
	payer_country_code := string(parser.Read_and_increment(2))
	// filler
	parser.Read_and_increment(6)

	return Address2{
		local_address: address,
		country:       payer_country,
		country_code:  payer_country_code,
	}
}

func parse_organizations_number(parser *parser.Parser) OrganizationNumber {
	organization_number := string(parser.Read_and_increment(12))
	// filler
	parser.Read_and_increment(66)

	return OrganizationNumber{
		organization_number: organization_number,
	}
}

func parse_deposit_section(parser *parser.Parser) Deposit {
	to_bank_account_number := string(parser.Read_and_increment(35))
	payment_date := string(parser.Read_and_increment(8))
	(parser.Read_and_increment(5))

	amount := string(parser.Read_and_increment(18))
	currency := string(parser.Read_and_increment(3))

	count_payments := string(parser.Read_and_increment(8))

	deposit_type := string(parser.Read_and_increment(1))

	return Deposit{
		to_bank_account_number: to_bank_account_number,
		date:                   payment_date,
		deposit:                amount,
		currency:               currency,
		count_payments:         count_payments,
		deposit_type:           deposit_type,
	}
}

func parse_tail_header(parser *parser.Parser) EndHeader {
	count_payment_post := string(parser.Read_and_increment(8))
	count_deduction_post := string(parser.Read_and_increment(8))
	count_additional_reference_post := string(parser.Read_and_increment(8))
	count_deposit_post := string(parser.Read_and_increment(8))

	// filler
	parser.Read_and_increment(46)

	return EndHeader{
		count_payment_post:              count_payment_post,
		count_deduction_post:            count_deduction_post,
		count_deposit_post:              count_deposit_post,
		count_additional_reference_post: count_additional_reference_post,
	}
}

type StartHeader struct {
	layout_name   string
	version       string
	timestamp     string
	is_production bool
}

type SectionHeader struct {
	to_giro_number      string
	to_plus_giro_number string
	currency            string
}

type SectionPayment struct {
	transaction_code      string
	from_bank_giro_number string
	reference             string
	payment_amount        *big.Int
	reference_code        string
	payment_channel_code  string
	bgc_identifier        string
	has_photo             bool
}

type InformationPost struct {
	text string
}

type NamePost struct {
	name       string
	extra_name string
}

type Address1 struct {
	address1 string
	zip_code string
}

type Address2 struct {
	local_address string
	country       string
	country_code  string
}

type OrganizationNumber struct {
	organization_number string
}

type Deposit struct {
	to_bank_account_number string
	date                   string
	deposit                string
	currency               string
	count_payments         string
	deposit_type           string
}

type EndHeader struct {
	count_payment_post              string
	count_deduction_post            string
	count_deposit_post              string
	count_additional_reference_post string
}
