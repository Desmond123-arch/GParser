package gParser

import (
	// "fmt"

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

		// Skip spaces and newlines
		if (val == " " && !insideQuotes) || val == "\n" {
			continue
		}

		if val == "\"" {
			insideQuotes = !insideQuotes
			temp_str += val
			continue
		}

		if insideQuotes {
			temp_str += val
			continue
		}
		switch val {
		case "{":
			tokens = append(tokens, Token{Type: "START_OBJECT"})

		case "}":
			if len(temp_str) > 0 {
				tokens = append(tokens, Token{Type: "VALUE", Value: temp_str})
				temp_str = ""
			}
			tokens = append(tokens, Token{Type: "END_OBJECT"})

		case "[":
			tokens = append(tokens, Token{Type: "START_ARR"})
			continue
		case "]":
			if len(temp_str) > 0 {
				tokens = append(tokens, Token{Type: "VALUE", Value: temp_str})
				temp_str = ""
			}
			tokens = append(tokens, Token{Type: "END_ARR"})
			continue
		case ":":
			//replace strings with  empty spaces
			temp_str = strings.ReplaceAll(temp_str, " ", "-")
			tokens = append(tokens, Token{Type: "KEY", Value: temp_str})
			temp_str = ""
			tokens = append(tokens, Token{Type: "COLON"})
		case ",":
			if len(temp_str) > 0 {
				tokens = append(tokens, Token{Type: "VALUE", Value: temp_str})
				temp_str = ""
			}
			tokens = append(tokens, Token{Type: "COMMA"})

		default:
			temp_str += val
		}
	}

	// Handle leftover temp_str
	if len(temp_str) > 0 {
		tokens = append(tokens, Token{Type: "VALUE", Value: temp_str})
	}
	return tokens
}

func Parse(str string) (map[string]interface{}, int) {
	if len(str) == 0 {
		return nil, 1
	}
	stream := Tokenize(str)
	fmt.Println(stream)
	mainObject := make(map[string]interface{})
	var currentObject map[string]interface{}
	stack := []map[string]interface{}{}
	key_val_stack := []string{}
	insideArr := false
	var arr []interface{}
	validString := regexp.MustCompile(`^".*"$`)
	intRegex := regexp.MustCompile(`^-?\d+$`)
	floatRegex := regexp.MustCompile(`^-?\d+(\.\d+)?$`)
	booleanRegex := regexp.MustCompile(`^(true|false)$`)
	nullRegex := regexp.MustCompile(`null`)

	for idx, token := range stream {
		switch token.Type {
		case "START_OBJECT":
			newObject := make(map[string]interface{})
			if len(stack) > 0 {
				// Assign new object to the key in the parent object
				parent := stack[len(stack)-1]
				key := key_val_stack[len(key_val_stack)-1]
				parent[key] = newObject
				key_val_stack = key_val_stack[:len(key_val_stack)-1]

			} else {
				mainObject = newObject
			}
			stack = append(stack, newObject)
			currentObject = newObject

		case "END_OBJECT":
			if len(stack) == 0 {
				return nil, 1
			}
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				currentObject = stack[len(stack)-1]
			} else {
				currentObject = nil
			}
		case "KEY":
			if validString.MatchString(token.Value) {
				key := strings.Trim(token.Value, "\"")
				key_val_stack = append(key_val_stack, key)
			} else {
				return nil, 1
			}
		case "START_ARR":
			insideArr = true

		case "END_ARR":
			insideArr = false
			key := key_val_stack[len(key_val_stack)-1]
			currentObject[key] = arr
			key_val_stack = key_val_stack[:len(key_val_stack)-1]
			arr = make([]interface{}, 0)

		case "VALUE":
			var value interface{}
			rawValue := strings.Trim(token.Value, "\"")

			switch {
			case validString.MatchString(token.Value):
				value = rawValue
			case intRegex.MatchString(rawValue):
				parsed, err := strconv.Atoi(rawValue)
				if err != nil {
					return nil, 1
				}
				value = parsed
			case floatRegex.MatchString(rawValue):
				parsed, err := strconv.ParseFloat(rawValue, 64)
				if err != nil {
					return nil, 1
				}
				value = parsed
			case booleanRegex.MatchString(rawValue):
				parsed, err := strconv.ParseBool(rawValue)
				if err != nil {
					return nil, 1
				}
				value = parsed
			case nullRegex.MatchString(rawValue):
				value = nil
			default:
				return nil, 1
			}

			if len(key_val_stack) == 0 || len(stack) == 0 {
				return nil, 1
			}
			if insideArr {
				arr = append(arr, value)
			} else {
				key := key_val_stack[len(key_val_stack)-1]
				currentObject[key] = value
				key_val_stack = key_val_stack[:len(key_val_stack)-1]
			}
		case "COLON":
			if idx+1 >= len(stream) || (stream[idx+1].Type != "VALUE" && stream[idx+1].Type != "START_OBJECT" && stream[idx+1].Type != "START_ARR") {
				return nil, 1
			}

		case "COMMA":
			if idx+1 >= len(stream) || (stream[idx+1].Type != "KEY" && stream[idx+1].Type != "VALUE") {
				return nil, 1
			}
		}
	}

	if len(stack) > 0 || len(key_val_stack) > 0 {
		return nil, 1
	}

	return mainObject, 0
}
