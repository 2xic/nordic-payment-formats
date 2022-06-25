package ocr

import (
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/2xic/nordic-payment-formats/generated"
	"github.com/2xic/nordic-payment-formats/helpers"
	"github.com/2xic/nordic-payment-formats/parser"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Ocr struct {
	helpers.Caller
}

func (Ocr) Parse(parser *parser.Parser) ([]generated.Transaction, error) {
	var txs []generated.Transaction
	sum_summary_amount := new(big.Int)

	for !parser.Done() {
		header := parse_header(parser)

		if header.record_type == "10" {
			parse_head_transmission(parser)
		} else if header.record_type == "20" {
			parse_start_header_assignment(parser)
		} else if header.record_type == "30" {
			tx := parse_transaction(parser, header.transaction_type)
			converted_date, error := time.Parse("020106", tx.date)
			if error != nil {
				panic(
					fmt.Sprintf("Error in parsing date %s, input %s", error, tx.date),
				)
			}
			txs = append(txs, generated.Transaction{
				Kid:               helpers.Trim(tx.kid),
				Amount:            tx.amount.String(),
				FromAccountNumber: tx.FromAccountNumber,
				Date: timestamppb.New(
					converted_date,
				),
			})
		} else if header.record_type == "88" {
			summary_assignment := parse_end_header_assignment(parser)
			sum_summary_amount = sum_summary_amount.Add(sum_summary_amount, &summary_assignment.total_amount)
		} else if header.record_type == "89" {
			summary := parse_tail_transmission(parser)
			helpers.Require(summary.total_amount.Cmp(sum_summary_amount), 0)
		} else {
			panic("Unknown record type")
		}
	}
	return txs, nil
}

func parse_head_transmission(parser *parser.Parser) TransmissionHeader {
	// NETS-ID
	data_transmitter := (parser.Read_and_increment(8))

	// serial number
	transmission_number := (parser.Read_and_increment(7))

	// serial number
	customer_unit_id := (parser.Read_and_increment(8))

	// serial number
	filler := (parser.Read_and_increment(49))

	return TransmissionHeader{
		data_transmitter:    string(data_transmitter),
		transmission_number: string(transmission_number),
		customer_unit_id:    string(customer_unit_id),
		filler:              string(filler),
	}
}

func parse_start_header_assignment(parser *parser.Parser) HeadAssignment {
	agreement_id := string(parser.Read_and_increment(9))

	assignment_number := string(parser.Read_and_increment(7))

	assignment_account := string(parser.Read_and_increment(11))

	// filler
	parser.Read_and_increment(45)

	return HeadAssignment{
		agreement_id:       agreement_id,
		assignment_number:  assignment_number,
		assignment_account: assignment_account,
	}
}

func parse_transaction(parser *parser.Parser, transaction_type int) Transaction {
	transaction_number := string(parser.Read_and_increment(7))

	nets_date := string(parser.Read_and_increment(6))

	//centre_id := string
	(parser.Read_and_increment(2))

	//day_code := string
	(parser.Read_and_increment(2))

	//partial_settlement_number := string
	(parser.Read_and_increment(1))

	if 18 <= transaction_type &&
		transaction_type <= 21 {
		//		helpers.Require(partial_settlement_number, "0")
	}

	//serial_number := string
	(parser.Read_and_increment(5))

	//sign := string
	(parser.Read_and_increment(1))

	amount := string(parser.Read_and_increment(17))

	kid := string(parser.Read_and_increment(25))

	//filler
	(parser.Read_and_increment(6))

	header_amount_2 := parse_header(parser)
	helpers.Require(header_amount_2.record_type, "31")

	transaction_number_amount_2 := string(parser.Read_and_increment(7))
	helpers.Require(transaction_number, transaction_number_amount_2)

	// form_number := string
	(parser.Read_and_increment(10))

	// agreement_id := string
	(parser.Read_and_increment(9))

	//agreement_id := string
	(parser.Read_and_increment(7))

	//date :=
	(parser.Read_and_increment(6))

	FromAccountNumber := string(parser.Read_and_increment(11))

	(parser.Read_and_increment(22))

	reference := ""

	if header_amount_2.transaction_type == 21 || header_amount_2.transaction_type == 2 {
		parse_header(parser)

		transaction_number_amount_3 := string(parser.Read_and_increment(7))
		helpers.Require(transaction_number, transaction_number_amount_3)

		reference = string(parser.Read_and_increment(40))

		// filler
		(parser.Read_and_increment(25))
	}

	return Transaction{
		transaction_number: transaction_number,
		amount:             *helpers.ConvertToBigInt(amount),
		kid:                kid,
		date:               nets_date,
		FromAccountNumber:  FromAccountNumber,
		reference:          reference,
	}
}

func parse_header(parser *parser.Parser) AmountHeader {
	record := string(parser.Read_and_increment(2))
	helpers.Require(record, "NY")

	service_code := string(parser.Read_and_increment(2))
	transaction_type := string(parser.Read_and_increment(2))
	record_type := string(parser.Read_and_increment(2))

	transaction_type_int, err := strconv.Atoi(transaction_type)

	if err != nil {
		panic("Bad value")
	}

	return AmountHeader{
		record:           record,
		transaction_type: transaction_type_int,
		record_type:      record_type,
		service_code:     service_code,
	}
}

func parse_end_header_assignment(parser *parser.Parser) TailAssignment {
	// Number of transactions
	parser.Read_and_increment(8)

	// Number of records
	number_of_records := (parser.Read_and_increment(8))

	// Total amount
	total_amount := (parser.Read_and_increment(17))

	// Nets date
	(parser.Read_and_increment(6))

	// earliest Nets date
	(parser.Read_and_increment(6))

	// latest Nets date
	(parser.Read_and_increment(6))

	// filler
	(parser.Read_and_increment(21))

	return TailAssignment{
		total_amount:      *helpers.ConvertToBigInt(string(total_amount)),
		number_of_records: string(number_of_records),
	}
}

func parse_tail_transmission(parser *parser.Parser) TailTransmission {
	// Number of transactions
	(parser.Read_and_increment(8))

	// Number of records
	(parser.Read_and_increment(8))

	// total amount
	total_amount := (parser.Read_and_increment(17))

	// total date
	(parser.Read_and_increment(6))

	// filler
	(parser.Read_and_increment(33))

	return TailTransmission{
		total_amount: *helpers.ConvertToBigInt(string(total_amount)),
	}
}

type TransmissionHeader struct {
	data_transmitter    string
	transmission_number string
	customer_unit_id    string
	filler              string
}

type AmountHeader struct {
	record           string
	service_code     string
	transaction_type int
	record_type      string
}

type HeadAssignment struct {
	agreement_id       string
	assignment_number  string
	assignment_account string
}

type TailAssignment struct {
	number_of_records string
	total_amount      big.Int
}

type Transaction struct {
	transaction_number string
	kid                string
	amount             big.Int
	date               string
	FromAccountNumber  string
	reference          string
}

type TailTransmission struct {
	total_amount big.Int
}
