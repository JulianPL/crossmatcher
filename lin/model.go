package lin

import (
	"crossmatcher/collection"
	"strings"
)

type Model struct {
	crossword Crossword
	candidate Candidate
}

func NewModel(rule string, alphabetString string, candidate string) *Model {
	alphabet := collection.MakeAlphabet(alphabetString, '.')
	return &Model{MakeCrossword(rule, alphabet), MakeCandidate(candidate, '.')}
}

func (m *Model) Solve() string {
	candidate, ok := m.crossword.SolveBruteforce(m.candidate)
	if ok == 0 {
		return strings.Repeat("#", m.candidate.Len())
	}
	return candidate.String()
}
