package lin

import (
	"crossmatcher/collection"
	"testing"
)

func TestCrossword_CheckSolution(t *testing.T) {
	rule := "(ab)*(ba)*"
	alphabet := collection.MakeAlphabet("ab")
	numA, _ := alphabet.Number('a')
	numB, _ := alphabet.Number('b')
	crossword := MakeCrossword(rule, alphabet)
	candidate := MakeCandidateFromString("ababba")
	solved := crossword.CheckSolution(candidate)
	if !solved {
		t.Errorf("Solution was not verified.")
	}
	candidate = MakeCandidateFromString("ababbb")
	solved = crossword.CheckSolution(candidate)
	if solved {
		t.Errorf("Non-solution was mistakenly verified.")
	}
	content := []int{numA, numB, numA, numB, numB, numA}
	candidate = MakeCandidate(content, alphabet)
	solved = crossword.CheckSolution(candidate)
	if !solved {
		t.Errorf("Solution was not verified.")
	}
	content = []int{numA, numB, numA, numB, numB, -1}
	candidate = MakeCandidate(content, alphabet)
	solved = crossword.CheckSolution(candidate)
	if solved {
		t.Errorf("Non-valid candidate was mistakenly verified.")
	}
}
