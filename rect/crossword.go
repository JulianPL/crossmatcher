package rect

import (
	"crossmatcher/collection"
	"regexp"
)

type Crossword struct {
	Horizontal []string
	Vertical   []string
	Alphabet   collection.Alphabet
}

func MakeCrossword(alphabet collection.Alphabet, horizontal []string, vertical []string) Crossword {
	return Crossword{horizontal, vertical, alphabet}
}

func (rules Crossword) CheckSolution(candidate Candidate) bool {
	for rowNumber := 0; rowNumber < len(candidate.Content); rowNumber++ {
		row, success := candidate.GetRow(rowNumber)
		if !success {
			return false
		}
		rowRule := "^(" + rules.Horizontal[rowNumber] + ")$"
		matched, _ := regexp.MatchString(rowRule, row)
		if !matched {
			return false
		}
	}
	for colNumber := 0; colNumber < len(candidate.Content[0]); colNumber++ {
		col, success := candidate.GetCol(colNumber)
		if !success {
			return false
		}
		colRule := "^(" + rules.Vertical[colNumber] + ")$"
		matched, _ := regexp.MatchString(colRule, col)
		if !matched {
			return false
		}
	}
	return true
}