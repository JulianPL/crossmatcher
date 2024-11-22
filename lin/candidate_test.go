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

func TestCandidate_Len(t *testing.T) {
	alphabet := collection.MakeAlphabet("a")
	numA, _ := alphabet.Number('a')
	candidate := MakeCandidate([]int{numA, -1, numA, -1, -1}, alphabet)
	if candidate.Len() != 5 {
		t.Errorf("Len is incorrect expected %d, actual %d", 5, candidate.Len())
	}
}

func TestCandidate_CountWildcards(t *testing.T) {
	alphabet := collection.MakeAlphabet("a")
	numA, _ := alphabet.Number('a')
	candidate := MakeCandidate([]int{numA, -1, numA, -1, -1}, alphabet)
	if candidate.CountWildcards() != 3 {
		t.Errorf("CountWildcards is incorrect expected %d, actual %d", 3, candidate.Len())
	}
}

func TestCandidate_IncrementCandidate(t *testing.T) {
	alphabet := collection.MakeAlphabet("0€1")
	candidate := MakeCandidateFirst(alphabet, 6)
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

func TestCandidate_MergeRow(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	numA, _ := alphabet.Number('a')
	numB, _ := alphabet.Number('b')
	candidate1 := MakeCandidate([]int{numA, numA, -1, -1, numA}, alphabet)
	candidate2 := MakeCandidate([]int{numB, numB}, alphabet)
	merge, ok := candidate1.MergeRow(candidate2)
	if !ok {
		t.Errorf("MergeRow is incorrect, expected success")
	}
	if merge != "aabba" {
		t.Errorf("MergeRow is incorrect, expected \"aabba\", got %s", merge)
	}
	candidate1 = MakeCandidate([]int{numA, numA, -1, -1, numA}, alphabet)
	candidate2 = MakeCandidate([]int{numB}, alphabet)
	merge, ok = candidate1.MergeRow(candidate2)
	if ok {
		t.Errorf("MergeRow is incorrect, merged rows with filler of wrong size")
	}
	candidate1 = MakeCandidate([]int{numA, numA, -1, -1, numA}, alphabet)
	candidate2 = MakeCandidate([]int{numB, -1}, alphabet)
	merge, ok = candidate1.MergeRow(candidate2)
	if ok {
		t.Errorf("MergeRow is incorrect, merged rows with wildcards")
	}
	candidate1 = MakeCandidate([]int{numA, numA, -1, 10, numA}, alphabet)
	candidate2 = MakeCandidate([]int{numB}, alphabet)
	merge, ok = candidate1.MergeRow(candidate2)
	if ok {
		t.Errorf("MergeRow is incorrect, merged rows with undefined charactes")
	}
}
