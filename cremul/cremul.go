package cremul

import (
	"fmt"

	"github.com/2xic/nordic-payment-formats/generated"
	"github.com/2xic/nordic-payment-formats/helpers"
	"github.com/2xic/nordic-payment-formats/parser"
)

type Cremul struct {
	helpers.Caller
}

var rollback_char = char_to_byte("'")

func (Cremul) Parse(parser *parser.Parser) ([]generated.Transaction, error) {
	var txs []generated.Transaction

	// There is some part of the begining of the cremul file i'm not sure what is
	// skipping over it for now - it's not clearly mentioned in the docs
	for true {
		current := string(parser.Read_and_increment(1))
		if current == "'" {
			keyword := string(parser.Peek(3))
			if keyword == "UNH" {
				break
			}
		}
	}

	for !parser.Done() {
		code := string(parser.ReadUntil(char_to_byte("+"), rollback_char))

		if code == "UNH" {
			parse_message_header(parser)
		} else if code == "BGM" {
			parse_beginning_of_message(parser)
		} else if code == "NAD" {
			payer, _ := parse_payer(parser)
			// TODO : This is incorrect.
			txs = append(txs, generated.Transaction{Payer: &generated.Payer{
				Name: payer.party_identification_details,
			}})
		} else if code == "DTM" {
			parse_timestamp(parser)
		} else if code == "MOA" {
			monetary_amount, _ := parse_monetary_amount(parser)

			txs[len(txs)-1].Amount = monetary_amount.amount
			txs[len(txs)-1].Currency = monetary_amount.currency
		} else if code == "FII" {
			skip_parsing(parser)
		} else if code == "LIN" {
			// Todo : implement this
			parser.ReadUntil(char_to_byte("'"), rollback_char)
		} else if code == "BUS" {
			business_function, _ := parse_business_function(parser)
			txs[len(txs)-1].IsInternational = business_function.environment == "IN"
		} else if code == "RFF" {
			parse_reference(parser)
		} else if code == "SEQ" {
			parse_sequence(parser)
		} else if code == "PRC" {
			parse_process_identification(parser)
		} else if code == "FTX" {
			parse_free_text(parser)
		} else if code == "CNT" {
			parse_control_total(parser)
		} else if code == "UNT" {
			parse_message_trailer(parser)
		} else if code == "UNZ" {
			// This does not seem to be used by nets.
			skip_parsing(parser)
		} else if code == "DOC" {
			// This does not seem to be used by nets.
			skip_parsing(parser)
		} else if code == "GIS" {
			// Todo : implement this
			skip_parsing(parser)
		} else if code == "INP" {
			// Todo : implement this
			skip_parsing(parser)
		} else {
			panic(fmt.Sprintf("UNknown code '%s' \n", code))
		}
	}

	return txs, nil
}

func parse_message_header(parser *parser.Parser) (message_header, error) {
	reference_number := string(parser.ReadUntil(char_to_byte("+"), rollback_char))
	message_type_identifier := string(parser.ReadUntil(char_to_byte(":"), rollback_char))
	message_version_number := string(parser.ReadUntil(char_to_byte(":"), rollback_char))
	message_release_number := string(parser.ReadUntil(char_to_byte(":"), rollback_char))
	control_agency := string(parser.ReadUntil(char_to_byte("'"), rollback_char))
	// other codes should not be used

	return message_header{
		reference_number:        reference_number,
		message_type_identifier: message_type_identifier,
		message_version_number:  message_version_number,
		control_agency:          control_agency,
		message_release_number:  message_release_number,
	}, nil
}

func parse_beginning_of_message(parser *parser.Parser) (beginning_of_message, error) {
	message_name_coded := string(parser.ReadUntil(char_to_byte("+"), rollback_char))
	unique_identifier := string(parser.ReadUntil(char_to_byte("'"), rollback_char))

	// other codes should not be used

	return beginning_of_message{
		message_name_coded: message_name_coded,
		unique_identifier:  unique_identifier,
	}, nil
}

func parse_timestamp(parser *parser.Parser) (timestamp, error) {
	period_qualifier := string(parser.ReadUntil(char_to_byte(":"), rollback_char))
	date := string(parser.ReadUntil(char_to_byte(":"), rollback_char))
	format_qualifier := string(parser.ReadUntil(char_to_byte("'"), rollback_char))

	return timestamp{
		period_qualifier: period_qualifier,
		date:             date,
		format_qualifier: format_qualifier,
	}, nil
}

func parse_payer(parser *parser.Parser) (payer, error) {
	party_qualifier := parser.ReadUntil(char_to_byte("+"), rollback_char)
	party_identification_details := parser.ReadUntil(char_to_byte("'"), rollback_char)

	// TODO, this should probably be handled in a better way, It's how cremul does an char escape.
	if parser.IsPrevByte(char_to_byte("?")) {
		additional := parser.ReadUntil(char_to_byte("'"), rollback_char)
		party_identification_details = append(
			party_identification_details,
			additional...,
		)
	}

	return payer{
		party_qualifier:              string(party_qualifier),
		party_identification_details: string(party_identification_details),
	}, nil
}

func parse_business_function(parser *parser.Parser) (business_function, error) {
	parser.ReadUntil(char_to_byte("+"), rollback_char)
	parser.ReadUntil(char_to_byte("+"), rollback_char)

	environment := parser.ReadUntil(char_to_byte("+"), rollback_char)

	parser.ReadUntil(char_to_byte("'"), rollback_char)

	return business_function{
		environment: string(environment),
	}, nil
}

func skip_parsing(parser *parser.Parser) {
	parser.ReadUntil(char_to_byte("'"), rollback_char)
}

func parse_monetary_amount(parser *parser.Parser) (monetary_amount, error) {
	amount_type := parser.ReadUntil(char_to_byte(":"), rollback_char)
	amount := parser.ReadUntil(char_to_byte(":"), rollback_char)
	currency := parser.ReadUntil(char_to_byte("'"), rollback_char)

	return monetary_amount{
		amount_type: string(amount_type),
		amount:      string(amount),
		currency:    string(currency),
	}, nil
}

func parse_reference(parser *parser.Parser) (reference, error) {
	reference_qualifier := parser.ReadUntil(char_to_byte(":"), rollback_char)
	reference_number := parser.ReadUntil(char_to_byte("'"), rollback_char)

	return reference{
		reference_qualifier: string(reference_qualifier),
		reference_number:    string(reference_number),
	}, nil
}

func parse_sequence(parser *parser.Parser) (sequence_details, error) {
	sequence_information := parser.ReadUntil(char_to_byte("+"), rollback_char)
	sequence_number := parser.ReadUntil(char_to_byte("'"), rollback_char)
	return sequence_details{
		sequence_information: string(sequence_information),
		sequence_number:      string(sequence_number),
	}, nil
}

func parse_process_identification(parser *parser.Parser) {
	parser.ReadUntil(char_to_byte("'"), rollback_char)
}

func parse_free_text(parser *parser.Parser) (free_text, error) {
	text_subject_qualifier := parser.ReadUntil(char_to_byte("+"), rollback_char)
	text_function_coded := parser.ReadUntil(char_to_byte("+"), rollback_char)
	text_reference := parser.ReadUntil(char_to_byte("'"), rollback_char)

	return free_text{
		text_subject_qualifier: string(text_subject_qualifier),
		text_function_coded:    string(text_function_coded),
		text_reference:         string(text_reference),
	}, nil
}

func parse_control_total(parser *parser.Parser) (control_total, error) {
	control_qualifier := parser.ReadUntil(char_to_byte(":"), rollback_char)
	control_value := parser.ReadUntil(char_to_byte(":"), rollback_char)
	parser.ReadUntil(char_to_byte("'"), rollback_char)

	return control_total{
		control_qualifier: string(control_qualifier),
		control_value:     string(control_value),
	}, nil
}

func parse_message_trailer(parser *parser.Parser) (message_trailer, error) {
	number_of_segments := parser.ReadUntil(char_to_byte(":"), rollback_char)
	message_reference_number := parser.ReadUntil(char_to_byte(":"), rollback_char)
	parser.ReadUntil(char_to_byte("'"), rollback_char)

	return message_trailer{
		number_of_segments:       string(number_of_segments),
		message_reference_number: string(message_reference_number),
	}, nil
}

func char_to_byte(input string) byte {
	if len(input) > 1 {
		panic("Single char only")
	}
	bytes := []byte(input)
	return bytes[0]
}

type message_header struct {
	reference_number        string
	message_type_identifier string
	message_version_number  string
	message_release_number  string
	control_agency          string
}

type beginning_of_message struct {
	message_name_coded string
	unique_identifier  string
}

type timestamp struct {
	message_name     string
	period_qualifier string
	date             string
	format_qualifier string
}

type payer struct {
	party_qualifier              string
	party_identification_details string
}

type business_function struct {
	environment string
}

type monetary_amount struct {
	amount_type string
	amount      string
	currency    string
}

type reference struct {
	reference_qualifier string
	reference_number    string
}

type sequence_details struct {
	sequence_information string
	sequence_number      string
}

type free_text struct {
	text_subject_qualifier string
	text_function_coded    string
	text_reference         string
}

type control_total struct {
	control_qualifier string
	control_value     string
}

type message_trailer struct {
	number_of_segments       string
	message_reference_number string
}
