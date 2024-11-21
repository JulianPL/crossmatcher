package collection

import (
	"testing"
)

func TestAlphabet_Make(t *testing.T) {
	alphabet := MakeAlphabet("t€st")
	alphabet.Len()
	if alphabet.Len() != 3 {
		t.Errorf("alphabet.Len given by \"t€st\" expected %d, got %d.", 3, alphabet.Len())
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
