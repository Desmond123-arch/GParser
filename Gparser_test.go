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
	} else {
		fmt.Printf("Got %q wanted %q for file data %s", got, want, string(data))
	}
}

func TestParseEmptyInvalid(t *testing.T) {
	data, err := os.ReadFile("./step1/invalid.json")
	if err != nil {
		t.Errorf("File could not be opened")
	}
	_, got := Parse(string(data))
	want := 1
	if got != want {
		t.Errorf("Got %q wanted %q for file data %s", got, want, string(data))
	} else {
		fmt.Printf("Got %q wanted %q for file data %s", got, want, string(data))
	}
}

func TestParseSimpleKeyValue(t *testing.T) {
	data, err := os.ReadFile("./step2/valid.json")
	if err != nil {
		t.Errorf("File could not be opened")
	}
	_, got := Parse(string(data))
	want := 0
	if got != want {
		t.Errorf("Got %q wanted %q for file data %s", got, want, string(data))
	} else {
		fmt.Printf("Got %q wanted %q for file data %s", got, want, string(data))
	}
}
func TestParseSimpleKeyValueInvalid(t *testing.T) {
	data, err := os.ReadFile("./step2/invalid.json")
	if err != nil {
		t.Errorf("File could not be opened")
	}
	_, got := Parse(string(data))
	want := 1
	if got != want {
		t.Errorf("Got %q wanted %q for file data %s", got, want, string(data))
	} else {
		fmt.Printf("Got %q wanted %q for file data %s", got, want, string(data))
	}
}

func TestParseSimple2KeyValueValid(t *testing.T) {
	data, err := os.ReadFile("./step2/valid2.json")
	if err != nil {
		t.Errorf("File could not be opened")
	}
	obj, got := Parse(string(data))
	want := 0
	if got != want {
		t.Errorf("Got %q wanted %q for file data %s", got, want, string(data))
	} else {
		fmt.Printf("Got %q wanted %q for file data %s", got, want, string(data))
	}
	for key, val := range obj {
		fmt.Println(key, val)
	}
}
func TestParseSimple2KeyValueInValid2(t *testing.T) {
	data, err := os.ReadFile("./step2/invalid2.json")
	if err != nil {
		t.Errorf("File could not be opened")
	}
	obj, got := Parse(string(data))
	want := 1
	if got != want {
		t.Errorf("Got %q wanted %q for file data %s", got, want, string(data))
	} else {
		fmt.Printf("Got %q wanted %q for file data %s", got, want, string(data))
	}
	for key, val := range obj {
		fmt.Println(key, val)
	}
}