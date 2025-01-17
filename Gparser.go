package gParser

import (
	"fmt"
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
	isKey := true

	for idx := range str {
		val := string(str[idx])
		// fmt.Println(val)

		//skip spaces and newlines
		if val == " " || val == "\n"{
			continue
		}

		//opening braces
		if val == "{" {
			tokens = append(tokens, Token{Type: "START_OBJECT"})
			continue
		}

		//ending braces
		if val == "}" {
			tokens = append(tokens, Token{Type: "END_OBJECT"})
			continue
		}
		if val == "," {
			tokens = append(tokens, Token{Type: "COMMA"})
			continue
		}
		//handle quotes
		if val == "\"" {
			insideQuotes = !insideQuotes
			if val != " " && !insideQuotes {
				if temp_str != "" {
					if isKey {
						tokens = append(tokens, Token{Type: "KEY", Value: temp_str})
					} else {
						tokens = append(tokens, Token{Type: "VALUE", Value: temp_str})
					}
				}
				temp_str = ""
				isKey = !isKey
				continue
			}
			continue
		}
		if val == ":" && !insideQuotes {
			tokens = append(tokens, Token{Type: "COLON"})
			continue
		}
		temp_str += val
	}
	if temp_str != "" {
		if isKey {
			tokens = append(tokens, Token{Type: "KEY", Value: temp_str})
		} else {
			tokens = append(tokens, Token{Type: "VALUE", Value: temp_str})
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
			key_val_stack = append(key_val_stack, key)

		case "VALUE":
			value := i.Value
			key := key_val_stack[len(key_val_stack)-1]
			mainObject[key] = value
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
