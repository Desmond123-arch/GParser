// package gParser
package main

import "fmt"

type Token struct {
	Type  string
	Value string
}

func tokenize(str string) {
	tokens := make([]Token, 0)
	insideQuotes := false
	temp_str := ""
	for idx := range str {
		val := string(str[idx])

		if val == "\"" {
			insideQuotes = !insideQuotes
			continue
		}
		if val == " " && insideQuotes == false {
			tokens = append(tokens, Token{Type: "Key", Value: temp_str})
			temp_str = ""
		} else {
			temp_str += val
		}
		//fmt.Println(temp_str)
	}
	if temp_str != "" {
		tokens = append(tokens, Token{Type: "Key", Value: temp_str})
	}
	fmt.Print(tokens)
}

func main() {
	tokenize("\"Pa\" \"Franklin\"")
}
