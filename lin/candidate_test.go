package lin

import (
	"crossmatcher/collection"
	"testing"
)

func TestCandidate_MakeCandidateFirst(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	candidate := MakeCandidateFirst(alphabet, 2)
	expected := "aa"
	actual, ok := candidate.GetRow()
	if expected != actual {
		t.Errorf("MakeCandidateFirst is incorrect expected %s, actual %s", expected, actual)
	}
	if !ok {
		t.Errorf("MakeCandidateFirst incorectly reports fail on row-retrieval.")
	}
}

func TestCandidate_MakeCandidateEmpty(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	candidate := MakeCandidateEmpty(alphabet, 2)
	expected := ""
	actual, ok := candidate.GetRow()
	if expected != actual {
		t.Errorf("MakeCandidateEmpty is incorrect expected %s, actual %s", expected, actual)
	}
	if ok {
		t.Errorf("MakeCandidateEmpty is incorrect expected non-retrievable Row.")
	}
}

func TestCandidate_MakeCandidate(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	numA, _ := alphabet.Number('a')
	numB, _ := alphabet.Number('b')
	content := []int{numA, numB, numA, numB, numB, numA}
	candidate := MakeCandidate(content, alphabet)
	expected := "ababba"
	actual, ok := candidate.GetRow()
	if expected != actual {
		t.Errorf("MakeCandidate is incorrect expected %s, actual %s", expected, actual)
	}
	if !ok {
		t.Errorf("MakeCandidate incorectly reports fail on row-retrieval.")
	}
}

func TestCandidate_MakeCandidateFromString(t *testing.T) {
	candidate := MakeCandidateFromString("€ba")
	expected := "€ba"
	actual, ok := candidate.GetRow()
	if expected != actual {
		t.Errorf("MakeCandidate is incorrect expected %s, actual %s", expected, actual)
	}
	if !ok {
		t.Errorf("MakeCandidateFromString incorectly reports fail on row-retrieval.")
	}
}
