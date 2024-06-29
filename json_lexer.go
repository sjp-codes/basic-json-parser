package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type keywords struct {
	JSON_COMMA        string
	JSON_COLON        string
	JSON_LEFTBRACKET  string
	JSON_RIGHTBRACKET string
	JSON_LEFTBRACE    string
	JSON_RIGHTBRACE   string
	JSON_QUOTE        string
}

const (
	JSON_COMMA        = ","
	JSON_COLON        = ":"
	JSON_LEFTBRACKET  = "["
	JSON_RIGHTBRACKET = "]"
	JSON_LEFTBRACE    = "{"
	JSON_RIGHTBRACE   = "}"
	JSON_QUOTE        = "\""
)

type bool struct {
	FALSE string
	TRUE  string
	NULL  string
}

const (
	FALSE = "false"
	TRUE  = "true"
	NULL  = "null"
)

func lex_string(input string) (string, string) {
	var jsonString strings.Builder

	if strings.HasPrefix(input, JSON_QUOTE) {
		input = input[1:]
	} else {
		return "", input
	}

	for i, c := range input {
		if string(c) == JSON_QUOTE {
			return jsonString.String(), input[i+1:]
		}
		jsonString.WriteRune(c)
	}
	panic("Expected end-of-string quote")

}
func lex_number(input string) (interface{}, string) {
	var jsonNumber strings.Builder //incomplete
	number_characters := "0123456789+-e."

	for _, c := range input {
		if strings.Contains(number_characters, string(c)) {
			jsonNumber.WriteRune(c)
		} else {
			break
		}
	}

	rest := input[len(jsonNumber.String()):]

	if jsonNumber.Len() == 0 {
		return nil, input
	}

	if strings.Contains(jsonNumber.String(), ".") {
		floatVal, err := strconv.ParseFloat(jsonNumber.String(), 64)
		if err != nil {
			panic("Invalid float number")
		}
		return floatVal, rest
	}

	intVal, err := strconv.Atoi(jsonNumber.String())
	if err != nil {
		panic("Invalid integer number")
	}
	return intVal, rest
}

func lex_bool(input string) (bool, string) {
	if strings.HasPrefix(input, TRUE) {
		return true, input[len(TRUE):]
	} else if strings.HasPrefix(input, FALSE) {
		return false, input[len(FALSE):]
	}
	return false, input
}

func lex_null(input string) (interface{}, string) {
	if strings.HasPrefix(input, NULL) {
		return nil, input[len(NULL):]
	}
	return nil, input
}

func lex(input string) []interface{} {
	var tokens []interface{}

	for len(input) > 0 {
		input = strings.TrimSpace(input)
		jsonString, rest := lex_string(input)
		if jsonString != "" {
			tokens = append(tokens, jsonString)
			input = rest
			continue
		}

		jsonNumber, rest := lex_number(input)
		if jsonNumber != nil {
			tokens = append(tokens, jsonNumber)
			input = rest
			continue
		}

		jsonBool, rest := lex_bool(input)
		if jsonBool != false || jsonBool != true {
			tokens = append(tokens, jsonBool)
			input = rest
			continue
		}

		jsonNull, rest := lex_null(input)
		if jsonNull != nil {
			tokens = append(tokens, jsonNull)
			input = rest
			continue
		}

		c := input[0]
		if c == ' ' || c == '\t' || c == '\b' || c == '\n' || c == '\r' {
			input = input[1:]
		} else if c == JSON_COMMA[0] || c == JSON_COLON[0] || c == JSON_LEFTBRACKET[0] || c == JSON_RIGHTBRACKET[0] || c == JSON_LEFTBRACE[0] || c == JSON_RIGHTBRACE[0] {
			tokens = append(tokens, string(c))
			input = input[1:]
		} else {
			panic(fmt.Sprintf("Unexpected character: %c", c))
		}

	}
	return tokens

}

func main() {
	// Read the contents of trial.json
	jsonData, err := os.Open("C:\\Users\\Saijyoti\\Desktop\\CODING\\AIEP\\JSON PARSER\\trial.json")
	if err != nil {
		fmt.Println("Error reading trial.json:", err)
		return
	}

	// Convert the byte slice to a string
	jsonString := string(jsonData)

	// Tokenize the JSON string
	tokens := lex(jsonString)

	// Print the tokens
	fmt.Println(tokens)
}

// c==JSON_COMMA || c == JSON_COLON || c == JSON_LEFTBRACKET || c == JSON_RIGHTBRACKET || c == JSON_LEFTBRACE || c == JSON_RIGHTBRACE
