package cremul

import (
	"fmt"

	"github.com/2xic/nordic-payment-formats/helpers"
	"github.com/2xic/nordic-payment-formats/parser"
)

type Cremul struct {
	helpers.Caller
}

var rollback_char = char_to_byte("'")

func (Cremul) Parse(parser *parser.Parser) ([]helpers.SimpleTransaction, error) {
	var txs []helpers.SimpleTransaction

	// There is some part of the begging of the cremul file i'm not sure what is
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
			fmt.Println(parse_message_header(parser))
			continue
		} else if code == "BGM" {
			fmt.Println(parse_begging_of_message(parser))
			continue
		} else if code == "NAD" {
			fmt.Println(parse_payer(parser))
			continue
		} else if code == "DTM" {
			fmt.Println(parse_timestamp(parser))
			continue
		} else if code == "MOA" {
			fmt.Println(parse_monetary_amount(parser))
			continue
		} else if code == "FII" {
			fmt.Println(parse_business_function(parser))
			continue
		} else if code == "LIN" {
			(parse_line_item(parser))
			continue
		} else if code == "BUS" {
			parse_business_function(parser)
			continue
		} else if code == "RFF" {
			fmt.Println(parse_reference(parser))
			continue
		} else if code == "SEQ" {
			fmt.Println(parse_sequence(parser))
			continue
		} else if code == "PRC" {
			parse_process_identification(parser)
			continue
		} else if code == "FTX" {
			fmt.Println(parse_free_text(parser))
			continue
		} else if code == "CNT" {
			fmt.Println(parse_control_total(parser))
			continue
		} else if code == "UNT" {
			fmt.Println(parse_message_trailer(parser))
			continue
		} else if code == "UNZ" {
			parser.ReadUntil(char_to_byte("'"), rollback_char)
			continue
		}
		fmt.Printf("UNknown code '%s' \n", code)
		break
	}

	/*
		fmt.Println(parse_message_header(parser))
		fmt.Println(parse_begging_of_message(parser))
		fmt.Println(parse_timestamp(parser))
		parse_payer(parser)
		parse_line_item(parser)
		fmt.Println(parse_timestamp(parser))
		fmt.Println(parse_business_function(parser))
		fmt.Println(parse_monetary_amount(parser))
		fmt.Println(parse_reference(parser))
		fmt.Println(parse_timestamp(parser))
	*/
	return txs, nil
}

func parse_message_header(parser *parser.Parser) (message_header, error) {
	//	parser.ReadUntil(char_to_byte("+"), rollback_char)
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

func parse_begging_of_message(parser *parser.Parser) (beginning_of_message, error) {
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

	return payer{
		party_qualifier:              string(party_qualifier),
		party_identification_details: string(party_identification_details),
	}, nil
}

func parse_line_item(parser *parser.Parser) {
	parser.ReadUntil(char_to_byte("'"), rollback_char)
}

func parse_business_function(parser *parser.Parser) (business_function, error) {
	parser.ReadUntil(char_to_byte("'"), rollback_char)
	return business_function{}, nil
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
