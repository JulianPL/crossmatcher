package rect

import (
	"crossmatcher/collection"
	"testing"
)

func TestCrossword_CheckSolution(t *testing.T) {
	horizontal := []string{"ab|ba", "aa|bb"}
	vertical := []string{"ba", ".."}
	alphabet := collection.MakeAlphabet("ab")
	crossword := MakeCrossword(alphabet, horizontal, vertical)
	rows := []string{"ba", "aa"}
	candidate := MakeCandidate(rows)
	solved := crossword.CheckSolution(candidate)
	if !solved {
		t.Errorf("Solution was not verified.")
	}
	rows = []string{"ba", "ab"}
	candidate = MakeCandidate(rows)
	solved = crossword.CheckSolution(candidate)
	if solved {
		t.Errorf("Non-solution was mistakenly verified.")
	}
	rows = []string{"ba", "bb"}
	candidate = MakeCandidate(rows)
	solved = crossword.CheckSolution(candidate)
	if solved {
		t.Errorf("Non-solution was mistakenly verified.")
	}
	//TODO Test non-successful row/column retrieval
}

func TestCrossword_SolveBruteforce(t *testing.T) {
	horizontal := []string{"ab|ba", "aa|bb"}
	vertical := []string{"ba", ".."}
	alphabet := collection.MakeAlphabet("ab")
	crossword := MakeCrossword(alphabet, horizontal, vertical)
	solution, count := crossword.SolveBruteforce()
	if count != 1 {
		t.Errorf("SolveBruteforce did not find the correct number of solutions. Expected %d, got %d", 1, count)
	}
	row, _ := solution.GetRow(0)
	if row.String() != "ba" {
		t.Errorf("SolveBruteforce did not find the correct first row. Expected %s, got %s", "ba", row.String())
	}
}
