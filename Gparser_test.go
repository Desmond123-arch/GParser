package gParser

import (
	"fmt"
	"os"
	"testing"
)

// step 1
func TestParseEmptyValid(t *testing.T) {
	data, err := os.ReadFile("./step1/valid.json")
	if err != nil {
		t.Errorf("File could not be opened")
	}
	_, got := Parse(string(data))
	want := 0
	if got != want {
		t.Errorf("Got %q wanted %q for file data %s", got, want, string(data))
	}
}

func TestParseEmptyInvalid(t *testing.T) {
	data, err := os.ReadFile("./step1/invalid.json")
	if err != nil {
		t.Errorf("File could not be opened")
	}
	_, got := Parse(string(data))
	want := 0
	if got != want {
		t.Errorf("Got %q wanted %q for file data %s", got, want, string(data))
	}
}

func TestParseSimpleKeyValue(t *testing.T) {
	data, err := os.ReadFile("./step2/valid.json")
	if err != nil {
		t.Errorf("File could not be opened")
	}
	fmt.Println(string(data))
	obj, _ := Parse(string(data))
	fmt.Println(obj["key"])
}
