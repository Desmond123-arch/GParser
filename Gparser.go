package gParser

import (
	"fmt"
)

//import "fmt"

type Token struct {
	Type  string
	Value string
}

func tokenize(str string) []Token {
	tokens := make([]Token, 0)
	insideQuotes := false
	temp_str := ""
	isKey := true
	for idx := range str {
		val := string(str[idx])
		// fmt.Println(val)

		//skip spaces and newlines
		if val == " " || val == "\n" {
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
	stream := tokenize(str)
	mainObject := make(map[string]interface{})
	stack := []map[string]interface{}{}
	key_val_stack := make([]string, 0)
	for _, i := range stream {
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
		}
	}
	if len(stack) > 0 {
		fmt.Println(stack, len(stack))
		return nil, 1
	}
	return mainObject, 0
}
