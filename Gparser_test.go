package gParser

import (
	"fmt"
	"os"
	"reflect"
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
	obj, got := Parse(string(data))
	want := 1
	if got != want {
		t.Errorf("Got %q wanted %q for file data %s", got, want, string(data))
	} else {
		fmt.Printf("Got %q wanted %q for file data %s", got, want, string(data))
		fmt.Print(obj)
	}
}

func TestParseSimpleKeyValue(t *testing.T) {
	data, err := os.ReadFile("./step2/valid.json")
	if err != nil {
		t.Errorf("File could not be opened")
	}
	obj, got := Parse(string(data))
	want := 0
	if got != want {
		t.Errorf("Got %q wanted %q for file data %s", got, want, string(data))
	} else {
		fmt.Printf("Got %q wanted %q for file data %s.\n", got, want, string(data))
		// for key, val := range obj {
		// 	fmt.Println(key, val)
		// }
		fmt.Println(obj)
	}
}

func TestParseSimpleKeyValueInvalid(t *testing.T) {
	data, err := os.ReadFile("./step2/invalid.json")
	if err != nil {
		t.Errorf("File could not be opened")
	}
	obj, got := Parse(string(data))
	want := 1
	if got != want {
		t.Errorf("Got %q wanted %q for file data %s", got, want, string(data))
	} else {
		fmt.Printf("Got %q wanted %q for file data %s", got, want, string(data))
		fmt.Print(obj)
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

func TestJsonWithTypes(t *testing.T) {
	data, err := os.ReadFile("./step3/valid.json")
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
		fmt.Println(key, val, reflect.TypeOf(val))
	}
}
func TestJsonWithTypesInvalid(t *testing.T) {
	data, err := os.ReadFile("./step3/invalid.json")
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
		fmt.Println(key, val, reflect.TypeOf(val))
	}
}