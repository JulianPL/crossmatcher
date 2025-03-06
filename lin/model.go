package lin

import (
	"crossmatcher/collection"
	"strings"
)

type Model struct {
	crossword Crossword
	candidate Candidate
}

// NewModel creates a model from the given rule, alphabet and candidate (as strings)
func NewModel(rule string, alphabetString string, candidate string) *Model {
	alphabet := collection.MakeAlphabet(alphabetString, '.')
	return &Model{MakeCrossword(rule, alphabet), MakeCandidate(candidate, '.')}
}

// Solve returns the solution to the model as string
// Returns Repeat("#") if there is no solution
func (m *Model) Solve() string {
	candidate, ok := m.crossword.SolveBruteforce(m.candidate)
	if ok == 0 {
		return strings.Repeat("#", m.candidate.Len())
	}
	return candidate.String()
}
