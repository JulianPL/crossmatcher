package rect

import (
	"crossmatcher/collection"
	"strings"
)

type Model struct {
	crossword Crossword
	candidate Candidate
}

// NewModel creates a model from the given vRules, hRules, alphabet and candidate (as strings)
func NewModel(vRules, hRules []string, alphabetString string, candidate []string) *Model {
	m := &Model{}
	alphabet := collection.MakeAlphabet(alphabetString, '.')

	m.crossword = MakeCrossword(alphabet, hRules, vRules)
	m.candidate = MakeCandidate(candidate, '.')

	return m
}

// NewModelRandom creates a model for a random crossword with given alphabet, height and width
func NewModelRandom(alphabetString string, height, width int) *Model {
	m := &Model{}
	alphabet := collection.MakeAlphabet(alphabetString, '.')

	m.crossword = MakeCrosswordRandom(alphabet, height, width)

	candidate := make([]string, height)
	for i := range height {
		candidate[i] = strings.Repeat(".", width)
	}
	m.candidate = MakeCandidate(candidate, '.')

	return m
}

// Solve returns the solution to the model as string
// Returns Repeat("#") if there is no solution
func (m *Model) Solve() []string {
	candidate, count := m.crossword.SolveLinearReductions(m.candidate)

	width := len(m.crossword.Vertical)
	height := len(m.crossword.Horizontal)

	if count == 0 {
		ret := make([]string, height)
		for i := range height {
			ret[i] = strings.Repeat("#", width)
		}

		return ret
	}

	ret := make([]string, height)
	for i := range height {
		row, _ := candidate.GetRow(i)
		ret[i] = row.String()
	}

	candidate, count = m.crossword.SolveBruteforce(candidate)

	if count == 1 {
		for i := range height {
			row, _ := candidate.GetRow(i)
			ret[i] = row.String()
		}
	}

	return ret
}
