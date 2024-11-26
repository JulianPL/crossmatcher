package lin

import (
	"crossmatcher/collection"
	"testing"
)

func TestCandidate_MakeCandidateFirst(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	candidate, ok := MakeCandidateFirst(alphabet, 2)
	expected := "aa"
	actual := candidate.String()
	if !ok {
		t.Errorf("MakeCandidateFirst incorectly reports fail.")
	}
	if expected != actual {
		t.Errorf("MakeCandidateFirst is incorrect expected %s, actual %s", expected, actual)
	}

	alphabet = collection.MakeAlphabet("")
	candidate, ok = MakeCandidateFirst(alphabet, 2)
	expected = ""
	actual = candidate.String()
	if ok {
		t.Errorf("MakeCandidateFirst incorectly accepts empty alphabet.")
	}
	if expected != actual {
		t.Errorf("MakeCandidateFirst is incorrect expected %s, actual %s", expected, actual)
	}
}

func TestCandidate_MakeCandidateEmpty(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	candidate := MakeCandidateEmpty(alphabet, 2)
	expected := ".."
	actual := candidate.String()
	if expected != actual {
		t.Errorf("MakeCandidateEmpty is incorrect expected %s, actual %s", expected, actual)
	}
}

func TestCandidate_MakeCandidateManual(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	numA, _ := alphabet.Number('a')
	numB, _ := alphabet.Number('b')
	content := Content{numA, numB, -1}
	actual, ok := MakeCandidateManual(content, alphabet)
	if !ok {
		t.Errorf("MakeCandidateManual incorectly reports fail.")
	}
	if actual.String() != "ab." {
		t.Errorf("MakeCandidateManual is incorrect expected %s, actual %s", "ab.", actual.String())
	}
	content = Content{numA, numB, 24}
	_, ok = MakeCandidateManual(content, alphabet)
	if ok {
		t.Errorf("MakeCandidateManual is incorrect accepts content not defined by alphabet.")
	}
}

func TestCandidate_MakeCandidate(t *testing.T) {
	candidate := MakeCandidate("€ba.", '.')
	expected := "€ba_"
	actual := candidate.String('_')
	if expected != actual {
		t.Errorf("MakeCandidateManual is incorrect expected %s, actual %s", expected, actual)
	}
}

func TestCandidate_MakeContent(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	row := "abc"
	_, ok := MakeContent(row, alphabet, '.')
	if ok {
		t.Errorf("MakeContent is incorrectly accepts alphabet without all relevant characters.")
	}
}

func TestCandidate_Len(t *testing.T) {
	candidate := MakeCandidate("a.a..", '.')
	if candidate.Len() != 5 {
		t.Errorf("Len is incorrect expected %d, actual %d", 5, candidate.Len())
	}
}

func TestCandidate_CountWildcards(t *testing.T) {
	candidate := MakeCandidate("a.a..", '.')
	if candidate.CountWildcards() != 3 {
		t.Errorf("CountWildcards is incorrect expected %d, actual %d", 3, candidate.Len())
	}
}

func TestCandidate_IncrementCandidate(t *testing.T) {
	alphabet := collection.MakeAlphabet("0€1")
	candidate, _ := MakeCandidateFirst(alphabet, 6)
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

func TestCandidate_Merge(t *testing.T) {
	candidate1 := MakeCandidate("aa..a", '.')
	candidate2 := MakeCandidate("bb")
	merge, ok := candidate1.Merge(candidate2)
	if !ok {
		t.Errorf("Merge is incorrect, expected success")
	}
	if merge.String('.') != "aabba" {
		t.Errorf("Merge is incorrect, expected \"aabba\", got %s", merge.String('.'))
	}
	candidate1 = MakeCandidate("aa..a", '.')
	candidate2 = MakeCandidate("b")
	merge, ok = candidate1.Merge(candidate2)
	if ok {
		t.Errorf("Merge is incorrect, merged rows with filler of wrong size")
	}
	candidate1 = MakeCandidate("aa..a", '.')
	candidate2 = MakeCandidate("b.", '.')
	merge, ok = candidate1.Merge(candidate2)
	if !ok {
		t.Errorf("Merge is incorrect, expected success")
	}
	if merge.String('.') != "aab.a" {
		t.Errorf("Merge is incorrect, expected \"aabba\", got %s", merge.String('.'))
	}
}

func TestCrossword_GreatestCommonPattern(t *testing.T) {
	candidate1 := MakeCandidate("aa..a", '.')
	candidate2 := MakeCandidate("ab,aa", ',')
	expected := "a...a"
	candidateGCP, ok := candidate1.GreatestCommonPattern(candidate2)
	actual := candidateGCP.String()
	if !ok {
		t.Errorf("GreatestCommonPattern is incorrect, expected success")
	}
	if actual != expected {
		t.Errorf("GreatestCommonPattern is incorrect expected %s, actual %s", expected, actual)
	}
	candidate1 = MakeCandidate("aa..a", '.')
	candidate2 = MakeCandidate("ab,aab", ',')
	expected = ""
	candidateGCP, ok = candidate1.GreatestCommonPattern(candidate2)
	actual = candidateGCP.String()
	if ok {
		t.Errorf("GreatestCommonPattern is incorrect, expected failure on candidates with different lengths")
	}
	if actual != expected {
		t.Errorf("GreatestCommonPattern is incorrect expected %s, actual %s", expected, actual)
	}
	candidate1 = MakeCandidate("", '.')
	candidate2 = MakeCandidate("ab,aab", ',')
	expected = "ab.aab"
	candidateGCP, ok = candidate1.GreatestCommonPattern(candidate2)
	actual = candidateGCP.String()
	if !ok {
		t.Errorf("GreatestCommonPattern is incorrect, expected success on empty candidate")
	}
	if actual != expected {
		t.Errorf("GreatestCommonPattern is incorrect expected %s, actual %s", expected, actual)
	}
}
