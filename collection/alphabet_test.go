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
	if alphabet.Char(alphabet.Number('a')) != 'a' {
		t.Errorf("alphabet Char/Number are not inverse on 'a'")
	}
	if alphabet.Char(alphabet.Number('A')) != 'A' {
		t.Errorf("alphabet Char/Number are not inverse on 'A'")
	}
	if alphabet.Char(alphabet.Number('b')) != 'b' {
		t.Errorf("alphabet Char/Number are not inverse on 'b'")
	}
	if alphabet.Char(alphabet.Number('€')) != '€' {
		t.Errorf("alphabet Char/Number are not inverse on '€'")
	}
}
