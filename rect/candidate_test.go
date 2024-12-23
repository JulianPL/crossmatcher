package rect

import (
	"crossmatcher/collection"
	"crossmatcher/lin"
	"strconv"
	"testing"
)

func TestCandidate_MakeCandidateFirst(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	candidate, ok := MakeCandidateFirst(alphabet, 2, 3)
	expected := "aaa\naaa"
	actual := candidate.String()
	if !ok {
		t.Errorf("MakeCandidateFirst incorectly reports fail.")
	}
	if expected != actual {
		t.Errorf("MakeCandidateFirst is incorrect expected %s, actual %s", strconv.Quote(expected), strconv.Quote(actual))
	}

	alphabet = collection.MakeAlphabet("")
	candidate, ok = MakeCandidateFirst(alphabet, 2, 3)
	expected = ""
	actual = candidate.String()
	if ok {
		t.Errorf("MakeCandidateFirst incorectly accepts empty alphabet.")
	}
	if expected != actual {
		t.Errorf("MakeCandidateFirst is incorrect expected %s, actual %s", strconv.Quote(expected), strconv.Quote(actual))
	}
}

func TestCandidate_MakeCandidateEmpty(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	candidate := MakeCandidateEmpty(alphabet, 2, 3)
	expected := "...\n..."
	actual := candidate.String()
	if expected != actual {
		t.Errorf("MakeCandidateEmpty is incorrect expected %s, actual %s", strconv.Quote(expected), strconv.Quote(actual))
	}

	alphabet = collection.MakeAlphabet("")
	candidate = MakeCandidateEmpty(alphabet, 2, 3)
	expected = "...\n..."
	actual = candidate.String()
	if expected != actual {
		t.Errorf("MakeCandidateEmpty is incorrect expected %s, actual %s", strconv.Quote(expected), strconv.Quote(actual))
	}
}

func TestCandidate_MakeCandidate(t *testing.T) {
	candidate := MakeCandidate([]string{"a€", "ab", "2a"})
	expected := "€ba"
	actual, _ := candidate.GetCol(1)
	if expected != actual.String() {
		t.Errorf("MakeCandidate is incorrect expected second column %s, actual %s", expected, actual.String())
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
	if actual.String() != expected {
		t.Errorf("Incorrect row at index 0: expected %s, got %s", expected, actual.String())
	}
	actual, success = candidate.GetRow(1)
	if success {
		t.Errorf("Incorrect row at index 1: incorrect success reported %s", actual.String())
	}
}

func TestCandidate_UpdateRow(t *testing.T) {
	candidate := MakeCandidate([]string{"a€", "ab", "2a"})
	candidate, ok := candidate.UpdateRow(lin.MakeCandidate("c.", '.'), 1)
	expected := "€_a"
	actual, _ := candidate.GetCol(1)
	if !ok {
		t.Errorf("UpdateRow not successful.")
	}
	if expected != actual.String('_') {
		t.Errorf("UpdateRow is incorrect expected second column %s, actual %s", expected, actual.String())
	}
	_, ok = candidate.UpdateRow(lin.MakeCandidate("ccb"), 3)
	if ok {
		t.Errorf("UpdateRow accepts colNumber which is too large.")
	}
	_, ok = candidate.UpdateRow(lin.MakeCandidate("c"), 1)
	if ok {
		t.Errorf("UpdateRow accepts column which is too short.")
	}
	_, ok = candidate.UpdateRow(lin.MakeCandidate("ccc"), 1)
	if ok {
		t.Errorf("UpdateRow accepts column which is too long.")
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
	if actual.String() != expected {
		t.Errorf("Incorrect col at index 0: expected %s, got %s", expected, actual.String())
	}
	_, success = candidate.GetCol(1)
	if success {
		t.Errorf("Incorrect col at index 1: incorrect success reported")
	}
}

func TestCandidate_UpdateCol(t *testing.T) {
	candidate := MakeCandidate([]string{"a€", "ab", "2a"})
	candidate, ok := candidate.UpdateCol(lin.MakeCandidate("cc.", '.'), 1)
	expected := "cc_"
	actual, _ := candidate.GetCol(1)
	if !ok {
		t.Errorf("UpdateCol not successful.")
	}
	if expected != actual.String('_') {
		t.Errorf("UpdateCol is incorrect expected second column %s, actual %s", expected, actual.String())
	}
	_, ok = candidate.UpdateCol(lin.MakeCandidate("ccb"), 2)
	if ok {
		t.Errorf("UpdateCol accepts colNumber which is too large.")
	}
	_, ok = candidate.UpdateCol(lin.MakeCandidate("cc"), 1)
	if ok {
		t.Errorf("UpdateCol accepts column which is too short.")
	}
	_, ok = candidate.UpdateCol(lin.MakeCandidate("cccc"), 1)
	if ok {
		t.Errorf("UpdateCol accepts column which is too long.")
	}
}
