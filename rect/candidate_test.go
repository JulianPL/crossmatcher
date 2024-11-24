package rect

import (
	"crossmatcher/collection"
	"testing"
)

func TestCandidate_MakeCandidate(t *testing.T) {
	candidate := MakeCandidate([]string{"a€", "ab", "2a"})
	expected := "€ba"
	actual, _ := candidate.GetCol(1)
	if expected != actual {
		t.Errorf("MakeCandidate is incorrect expected second column %s, actual %s", expected, actual)
	}
}

func TestCandidate_IncrementCandidate(t *testing.T) {
	alphabet := collection.MakeAlphabet("0€1")
	candidate, _ := MakeCandidateFirst(alphabet, 2, 3)
	success := true
	count := 0
	for success {
		candidate, success = candidate.IncrementCandidate()
		count++
	}
	if count != 729 {
		t.Errorf("Incorrect number of distinct candidates: expected 729, got %d", count)
	}
}

func TestCandidate_GetRow(t *testing.T) {
	alphabet := collection.MakeAlphabet("0€1")
	candidate, _ := MakeCandidateFirst(alphabet, 1, 3)
	for i := 0; i < 5; i++ {
		candidate, _ = candidate.IncrementCandidate()
	}
	expected := "1€0"
	actual, success := candidate.GetRow(0)
	if !success {
		t.Errorf("Incorrect row at index 0: no success reported")
	}
	if actual != expected {
		t.Errorf("Incorrect row at index 0: expected %s, got %s", expected, actual)
	}
	actual, success = candidate.GetRow(1)
	if success {
		t.Errorf("Incorrect row at index 1: incorrect success reported %s", actual)
	}
}

func TestCandidate_GetCol(t *testing.T) {
	alphabet := collection.MakeAlphabet("0€1")
	candidate, _ := MakeCandidateFirst(alphabet, 3, 1)
	for i := 0; i < 5; i++ {
		candidate, _ = candidate.IncrementCandidate()
	}
	expected := "1€0"
	actual, success := candidate.GetCol(0)
	if !success {
		t.Errorf("Incorrect col at index 0: no success reported")
	}
	if actual != expected {
		t.Errorf("Incorrect col at index 0: expected %s, got %s", expected, actual)
	}
	_, success = candidate.GetCol(1)
	if success {
		t.Errorf("Incorrect col at index 1: incorrect success reported")
	}
}
