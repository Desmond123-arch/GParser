package gParser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//import "fmt"

type Token struct {
	Type  string
	Value string
}

func Tokenize(str string) []Token {
	tokens := make([]Token, 0)
	insideQuotes := false
	temp_str := ""

	for idx := range str {
		val := string(str[idx])
		// fmt.Println(val)

		//skip spaces and newlines
		if val == " " || val == "\n" || len(val) == 0 {
			continue
		}

		//opening braces
		if val == "{" {
			tokens = append(tokens, Token{Type: "START_OBJECT"})
			continue
		}

		//ending braces
		if val == "}" {
			if len(temp_str) > 0 {
				tokens = append(tokens, Token{Type: "VALUE", Value: temp_str})
			}
			tokens = append(tokens, Token{Type: "END_OBJECT"})
			continue
		}
		// comma
		if val == "," {
			if len(temp_str) > 0 {
				tokens = append(tokens, Token{Type: "VALUE", Value: temp_str})
				temp_str = ""
			}
			tokens = append(tokens, Token{Type: "COMMA"})
			continue
		}

		// colons
		if val == ":" && !insideQuotes {
			tokens = append(tokens, Token{Type: "KEY", Value: temp_str})
			temp_str = ""
			tokens = append(tokens, Token{Type: "COLON"})
			continue
		}
		if val == "\"" {
			insideQuotes = !insideQuotes
		}

		if val != "" {
			temp_str += val
		}
	}
	for idx := range tokens {
		fmt.Println(tokens[idx].Type, tokens[idx].Value)
	}
	return tokens
}

func Parse(str string) (map[string]interface{}, int) {
	if len(str) == 0 {
		return nil, 1
	}
	stream := Tokenize(str)
	mainObject := make(map[string]interface{})
	stack := []map[string]interface{}{}
	key_val_stack := make([]string, 0)

	validKey := regexp.MustCompile(`^".*"$`)
	intRegex := regexp.MustCompile(`^-?\d+$`)
	floatRegex := regexp.MustCompile(`^-?\d+(\.\d+)?$`)
	booleanRegex := regexp.MustCompile(`^(true|false)$`)
	nullRegex := regexp.MustCompile(`null`)

	for idx, i := range stream {

		switch i.Type {
		case "START_OBJECT":
			mainObject = make(map[string]interface{})
			stack = append(stack, mainObject)

		case "END_OBJECT":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				return nil, 1
			}

		case "KEY":
			key := i.Value
			fmt.Println(key, validKey.MatchString(key))
			if validKey.MatchString(key) {
				new_val := strings.ReplaceAll(i.Value, "\"", "")
				key_val_stack = append(key_val_stack, new_val)
			} else {
				return nil, 1
			}

		case "VALUE":
			inputValue := strings.ReplaceAll(i.Value, "\"", "")
			fmt.Println(inputValue)
			var val interface{}
			switch {
			//for ints
			case intRegex.MatchString(inputValue):
				value, err := strconv.Atoi(i.Value)
				if err != nil {
					return nil, 1
				}
				val = value
				// booleans
			case booleanRegex.MatchString(inputValue):
				value, err := strconv.ParseBool(inputValue)
				if err != nil {
					return nil, 1
				}
				val = value
			case floatRegex.MatchString(inputValue):
				value, err := strconv.ParseFloat(inputValue, 64)
				if err != nil {
					return nil, 1
				}
				val = value
			case nullRegex.MatchString(inputValue):
				val = nil
			default:
				val = inputValue
			}

			key := key_val_stack[len(key_val_stack)-1]
			mainObject[key] = val
			key_val_stack = key_val_stack[:len(key_val_stack)-1]

		case "COLON":
			if idx+1 < len(stream) {
				if stream[idx+1].Type != "VALUE" {
					return nil, 1
				}
				continue
			} else {

				return nil, 1
			}

		case "COMMA":
			if idx+1 < len(stream) {
				if stream[idx+1].Type != "KEY" {
					return nil, 1
				}
				continue
			} else {
				return nil, 1
			}
		}

	}
	// for key, val := range mainObject {
	// 	fmt.Println(key, val)
	// }
	return mainObject, 0
}
