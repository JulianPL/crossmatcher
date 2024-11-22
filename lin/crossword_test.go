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

func TestCrossword_SolveBruteforce(t *testing.T) {
	rule := "(ab)*(ba)*"
	alphabet := collection.MakeAlphabet("ab")
	numA, _ := alphabet.Number('a')
	numB, _ := alphabet.Number('b')
	crossword := MakeCrossword(rule, alphabet)
	content := []int{numA, -1, numB, -1, -1, numA}
	candidate := MakeCandidate(content, alphabet)
	solution, count := crossword.SolveBruteforce(candidate)
	if count != 1 {
		t.Errorf("SolveBruteforce did not find the correct number of solutions. Expected %d, got %d", 1, count)
	}
	row, _ := solution.GetRowWithWildcard('.')
	if row != "abbaba" {
		t.Errorf("SolveBruteforce did not find the correct first row. Expected %s, got %s", "abbaba", row)
	}
	content = []int{numA, -1, -1, -1, -1, numA}
	candidate = MakeCandidate(content, alphabet)
	solution, count = crossword.SolveBruteforce(candidate)
	if count != 2 {
		t.Errorf("SolveBruteforce did not find the correct number of solutions. Expected %d, got %d", 1, count)
	}
	row, _ = solution.GetRowWithWildcard('.')
	if row != "ab..ba" {
		t.Errorf("SolveBruteforce did not find the correct first row. Expected %s, got %s", "abbaba", row)
	}
}
