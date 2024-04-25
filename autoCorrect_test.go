package main

import (
	"testing"
)

func TestReplaceHex(t *testing.T) {
	input := "1E (hex) files were added"
	expected := "30 files were added"
	result := ReplaceHex(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestReplaceBin(t *testing.T) {
	input := "It has been 10 (bin) years"
	expected := "It has been 2 years"
	result := ReplaceBin(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestReplaceWithAn(t *testing.T) {
	input := "There it was. A amazing rock!"
	expected := "There it was. An amazing rock!"
	result := ReplaceWithAn(input)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestFormatQuotes(t *testing.T) {
	input := "As Elton John said: ' I am the most well-known homosexual in the world '"
	expected := "As Elton John said: 'I am the most well-known homosexual in the world'"
	result := FormatQuotes(input)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestFormatPunctuation(t *testing.T) {
	input := "I was sitting over there ,and then BAMM !!"
	expected := "I was sitting over there, and then BAMM!!"
	result := FormatPunctuation(input)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestReplaceCase(t *testing.T) {
	input := "This is so exciting (up, 2)"
	expected := "This is SO EXCITING"
	result := ReplaceCase(input)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
