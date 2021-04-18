package main

import (
	"testing"
)

func TestHangmanPic(t *testing.T) {
	for i := 0; i < 7; i++ {
		msg := getPic(0)
		if len(msg) < 1 {
			t.Fatalf(`getPic(0) = %q, want non-empty string`, msg)
		}
	}
}

func TestRandomWord(t *testing.T) {
	msg := getRandomWord()
	if len(msg) < 1 {
		t.Fatalf(`getRandomWord() = %q, want non-empty string`, msg)
	}
}

func TestIsValidInput(t *testing.T) {
	boolean := isValidInput("NOT")
	if boolean {
		t.Fatalf("Should have returned false")
	}
	boolean = isValidInput("1232")
	if boolean {
		t.Fatalf("Should have returned false")
	}
	boolean = isValidInput("a")
	if boolean == false {
		t.Fatalf("Should have returned true")
	}
}

func TestPrintUnderlineTestEmptyMap(t *testing.T) {
	var m map[string]bool
	str := printUnderlineWord("otter", m)
	expected := "_ _ _ _ _ "
	if str != expected {
		t.Fatalf("Should have returned %q %v", expected, str)
	}
}

func TestPrintUnderlineTestFullMap(t *testing.T) {
	m := map[string]bool{"o": true, "t": true, "e": true, "r": true}
	str := printUnderlineWord("otter", m)
	expected := "o t t e r "
	if str != expected {
		t.Fatalf("Should have returned %q %v", expected, str)
	}
}

func TestGetWordMap(t *testing.T) {
	word := "world"
	m := getWordMap(word)
	for _, c := range word {
		_key := string(c)
		if m[_key] == 0 {
			t.Fatalf("Missing key %q in %v", _key, m)
		}
	}
}
