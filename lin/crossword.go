package lin

import (
	"crossmatcher/collection"
	"regexp"
)

type Crossword struct {
	Rule     string
	Alphabet collection.Alphabet
}

func MakeCrossword(rule string, alphabet collection.Alphabet) Crossword {
	return Crossword{rule, alphabet}
}

func (crossword Crossword) CheckSolution(candidate Candidate) bool {
	row, success := candidate.GetRow()
	if !success {
		return false
	}
	rowRule := "^(" + crossword.Rule + ")$"
	matched, _ := regexp.MatchString(rowRule, row)
	return matched
}
