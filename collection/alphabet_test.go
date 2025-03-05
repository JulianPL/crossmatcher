package collection

import (
	"strings"
	"testing"
)

func TestAlphabet_Make(t *testing.T) {
	alphabet := MakeAlphabet("t€st")
	actual := alphabet.Len()
	expected := 3
	if actual != expected {
		t.Errorf("alphabet.Len given by \"t€st\" expected %d, got %d.", expected, actual)
	}
	if !alphabet.Contains('t') {
		t.Errorf("alphabet given by \"t€st\" expected to contain 't'.")
	}
	if !alphabet.Contains('€') {
		t.Errorf("alphabet given by \"t€st\" expected to contain '€'.")
	}
	if alphabet.Contains('T') {
		t.Errorf("alphabet given by \"t€st\" expected not to contain 'T'.")
	}
}

func TestAlphabet_String(t *testing.T) {
	alphabetString := MakeAlphabet("abd", 'b').String()
	if !strings.Contains(alphabetString, "a") {
		t.Errorf("alphabet given by \"abd\" with wildcard 'b' expected to contain 'a'.")
	}
	if strings.Contains(alphabetString, "b") {
		t.Errorf("alphabet given by \"abd\" with wildcard 'b' expected not to contain 'b'.")
	}
	if strings.Contains(alphabetString, "c") {
		t.Errorf("alphabet given by \"abd\" with wildcard 'b' expected not to contain 'c'.")
	}
}

func TestAlphabet_Bijective(t *testing.T) {
	alphabet := MakeAlphabet("abcA€")
	num, _ := alphabet.Number('a')
	char, _ := alphabet.Char(num)
	if char != 'a' {
		t.Errorf("alphabet Char/Number are not inverse on 'a'")
	}
	num, _ = alphabet.Number('A')
	char, _ = alphabet.Char(num)
	if char != 'A' {
		t.Errorf("alphabet Char/Number are not inverse on 'A'")
	}
	num, _ = alphabet.Number('b')
	char, _ = alphabet.Char(num)
	if char != 'b' {
		t.Errorf("alphabet Char/Number are not inverse on 'b'")
	}
	num, _ = alphabet.Number('€')
	char, _ = alphabet.Char(num)
	if char != '€' {
		t.Errorf("alphabet Char/Number are not inverse on '€'")
	}
}

func TestAlphabet_Merge(t *testing.T) {
	alphabet1 := MakeAlphabet("abcd..", '.')
	alphabet2 := MakeAlphabet("a.bcdee", '.')
	expected := 5
	actual := alphabet1.Merge(alphabet2).Len()
	if actual != expected {
		t.Errorf("Merge is incorrect. Expected length %d, got %d.", expected, actual)
	}
}
