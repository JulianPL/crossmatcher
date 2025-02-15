package rect

import (
	"crossmatcher/collection"
	"testing"
)

func TestCrossword_GetRow(t *testing.T) {
	horizontal := []string{"ab|ba", "aa|bb"}
	vertical := []string{"ba", ".."}
	alphabet := collection.MakeAlphabet("ab")
	crossword := MakeCrossword(alphabet, horizontal, vertical)
	actual, ok := crossword.GetRow(1)
	expected := "aa|bb"
	if !ok {
		t.Errorf("GetRow did not successfully get the valid row")
	}
	if actual.Rule != expected {
		t.Errorf("GetRow returns wrong row. Expected %s, got %s", expected, actual.Rule)
	}
	_, ok = crossword.GetRow(2)
	if ok {
		t.Errorf("GetRow did return success for non-existing row")
	}
}

func TestCrossword_GetCol(t *testing.T) {
	horizontal := []string{"ab|ba", "aa|bb"}
	vertical := []string{"ba", ".."}
	alphabet := collection.MakeAlphabet("ab")
	crossword := MakeCrossword(alphabet, horizontal, vertical)
	actual, ok := crossword.GetCol(1)
	expected := ".."
	if !ok {
		t.Errorf("GetCol did not successfully get the valid row")
	}
	if actual.Rule != expected {
		t.Errorf("GetCol returns wrong row. Expected %s, got %s", expected, actual.Rule)
	}
	_, ok = crossword.GetCol(2)
	if ok {
		t.Errorf("GetCol did return success for non-existing row")
	}
}

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
	rows = []string{"ba", "a."}
	candidate = MakeCandidate(rows, '.')
	solved = crossword.CheckSolution(candidate)
	if solved {
		t.Errorf("Solution with wildcards was mistakenly verified.")
	}
}

func TestCrossword_SolveBruteforce(t *testing.T) {
	horizontal := []string{"ab|ba", "aa|bb"}
	vertical := []string{"b.", ".."}
	alphabet := collection.MakeAlphabet("ab")
	crossword := MakeCrossword(alphabet, horizontal, vertical)
	constraint := MakeCandidateEmpty(alphabet, 2, 2)
	solution, count := crossword.SolveBruteforce(constraint)
	if count != 2 {
		t.Errorf("SolveBruteforce did not find the correct number of solutions. Expected %d, got %d", 1, count)
	}
	constraint = MakeCandidate([]string{"..", "a."}, '.')
	solution, count = crossword.SolveBruteforce(constraint)
	if count != 1 {
		t.Errorf("SolveBruteforce did not find the correct number of solutions. Expected %d, got %d", 1, count)
	}
	row, _ := solution.GetRow(0)
	if row.String() != "ba" {
		t.Errorf("SolveBruteforce did not find the correct first row. Expected %s, got %s", "ba", row.String())
	}
	/* Performance Test: about 250,000 candidates in about 5 seconds
	horizontal = []string{".*", ".*", "a.*"}
	vertical = []string{".*", ".*", ".*"}
	alphabet = collection.MakeAlphabet("abcd")
	crossword = MakeCrossword(alphabet, horizontal, vertical)
	constraint = MakeCandidate([]string{"...", "...", "..."}, '.')
	solution, _ = crossword.SolveBruteforce(constraint)
	row, _ = solution.GetRow(2)
	if row.String() != "abb" {
		t.Errorf("SolveBruteforce did not find the correct first row. Expected %s, got %s", "a..", row.String())
	} */
}

func TestCrossword_SolveLinearReductions(t *testing.T) {
	horizontal := []string{"a.", "ab||ba"}
	vertical := []string{"(aa)|(bb)", ".."}
	alphabet := collection.MakeAlphabet("ab")
	crossword := MakeCrossword(alphabet, horizontal, vertical)
	constraint := MakeCandidateEmpty(alphabet, 2, 2)
	solution, _ := crossword.SolveLinearReductions(constraint)
	expected := "a.\nab"
	if solution.String() != expected {
		t.Errorf("SolveLinearReductions did not find the solution. Expected %s, got %s", expected, solution.String())
	}
	horizontal = []string{"0*1{3}0*", "((01)|(10))*", "(00.)*", "0*1{4}0*", "(01)*(10)*", "0*1{3}0*"}
	vertical = []string{"(0.)*", "((01)|(10))*", "1*01*", "(.1)*", "(01)*(11)*", "(.00)*"}
	alphabet = collection.MakeAlphabet("01")
	crossword = MakeCrossword(alphabet, horizontal, vertical)
	constraint = MakeCandidateEmpty(alphabet, 6, 6)
	solution, _ = crossword.SolveLinearReductions(constraint)
	expected = "011100\n100110\n001000\n011110\n011010\n001110"
	if solution.String() != expected {
		t.Errorf("SolveLinearReductions did not find the solution. Expected %s, got %s", expected, solution.String())
	}
}
