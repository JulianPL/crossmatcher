package lin

import (
	"crossmatcher/collection"
	"testing"
)

func TestCrossword_CheckSolution(t *testing.T) {
	rule := "(ab)*(ba)*"
	alphabet := collection.MakeAlphabet("ab")
	crossword := MakeCrossword(rule, alphabet)
	candidate := MakeCandidate("ababba")
	solved := crossword.CheckSolution(candidate)
	if !solved {
		t.Errorf("Solution was not verified.")
	}
	candidate = MakeCandidate("ababbb")
	solved = crossword.CheckSolution(candidate)
	if solved {
		t.Errorf("Non-solution was mistakenly verified.")
	}
	candidate = MakeCandidate("ababba", '.')
	solved = crossword.CheckSolution(candidate)
	if !solved {
		t.Errorf("Solution was not verified.")
	}
	candidate = MakeCandidate("ababb.", '.')
	solved = crossword.CheckSolution(candidate)
	if solved {
		t.Errorf("Non-valid candidate was mistakenly verified.")
	}
}

func TestCrossword_SolveBruteforce(t *testing.T) {
	rule := "(ab)*(ba)*"
	alphabet := collection.MakeAlphabet("ab")
	crossword := MakeCrossword(rule, alphabet)
	candidate := MakeCandidate("a.b..a", '.')
	solution, count := crossword.SolveBruteforce(candidate)
	if count != 1 {
		t.Errorf("SolveBruteforce did not find the correct number of solutions. Expected %d, got %d", 1, count)
	}
	row := solution.String()
	if row != "abbaba" {
		t.Errorf("SolveBruteforce did not find the correct first row. Expected %s, got %s", "abbaba", row)
	}
	candidate = MakeCandidate("a....a", '.')
	solution, count = crossword.SolveBruteforce(candidate)
	if count != 2 {
		t.Errorf("SolveBruteforce did not find the correct number of solutions. Expected %d, got %d", 1, count)
	}
	row = solution.String()
	if row != "ab..ba" {
		t.Errorf("SolveBruteforce did not find the correct first row. Expected %s, got %s", "ab..ba", row)
	}
	candidate = MakeCandidate("a.....", '.')
	solution, count = crossword.SolveBruteforce(candidate)
	if count != 3 {
		t.Errorf("SolveBruteforce did not find the correct number of solutions. Expected %d, got %d", 3, count)
	}
	row = solution.String()
	if row != "ab...." {
		t.Errorf("SolveBruteforce did not find the correct first row. Expected %s, got %s", "ab....", row)
	}
}
