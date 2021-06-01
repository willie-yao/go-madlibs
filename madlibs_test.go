package main

import (
	"encoding/json"
	"testing"
)

// Tests expected behavior for reading a file
func TestReadFile(t *testing.T) {
	filename := "inventory.json"
	expected := Inventory{"wood", 16}
	inv, err := ReadFile(filename)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}
	parsed_inv, err := json.Marshal(inv)
	if err != nil {
		t.Errorf("Error parsing inv to json: %v", err)
	}
	parsed_expected, err := json.Marshal(expected)
	if err != nil {
		t.Errorf("Error parsing expected to json: %v", err)
	}
	if inv.Count != expected.Count || inv.Material != expected.Material || err != nil {
		t.Fatalf(`ReadFile("inventory.json") output %s does not match expected output %s`, string(parsed_inv), string(parsed_expected))
	}
}

// Makes sure an error is returned when given an invalid filename
func TestReadInvalidFileName(t *testing.T) {
	filename := "bad"
	empty := Inventory{}
	file, err := ReadFile(filename)
	if err == nil || file != empty {
		t.Errorf("ReadFile should return error")
	}
}
