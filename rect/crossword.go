package rect

import (
	"crossmatcher/collection"
	"crossmatcher/lin"
)

type Crossword struct {
	Horizontal []string
	Vertical   []string
	Alphabet   collection.Alphabet
}

// MakeCrossword makes a crossword from two given sets of strings and an underlying alphabet.
func MakeCrossword(alphabet collection.Alphabet, horizontal []string, vertical []string) Crossword {
	return Crossword{horizontal, vertical, alphabet}
}

func (crossword Crossword) GetRow(rowNumber int) (lin.Crossword, bool) {
	if len(crossword.Horizontal) <= rowNumber {
		return lin.MakeCrossword("", collection.MakeAlphabet("")), false
	}
	row := lin.MakeCrossword(crossword.Horizontal[rowNumber], crossword.Alphabet)
	return row, true
}

func (crossword Crossword) GetCol(colNumber int) (lin.Crossword, bool) {
	if len(crossword.Vertical) <= colNumber {
		return lin.MakeCrossword("", collection.MakeAlphabet("")), false
	}
	col := lin.MakeCrossword(crossword.Vertical[colNumber], crossword.Alphabet)
	return col, true
}

// CheckSolution checks, whether a candidate without wildcards satisfies a crossword.
func (crossword Crossword) CheckSolution(candidate Candidate) bool {
	if candidate.CountWildcards() > 0 {
		return false
	}

	for rowNumber := range candidate.Content {
		rowContent, _ := candidate.GetRow(rowNumber)
		rowCrossword, _ := crossword.GetRow(rowNumber)
		if !rowCrossword.CheckSolution(rowContent) {
			return false
		}
	}

	for colNumber := range candidate.Content[0] {
		colContent, _ := candidate.GetCol(colNumber)
		colCrossword, _ := crossword.GetCol(colNumber)
		if !colCrossword.CheckSolution(colContent) {
			return false
		}
	}
	return true
}

// SolveBruteforce checks all candidates that fill the wildcards given by the constraint.
func (crossword Crossword) SolveBruteforce(constraint Candidate) (Candidate, int) {
	candidateFill, _ := lin.MakeCandidateFirst(crossword.Alphabet, constraint.CountWildcards())

	candidateIsValid := true
	solutionNum := 0
	var solution Candidate
	for candidateIsValid {
		candidateMerge, _ := constraint.Merge(candidateFill)
		if crossword.CheckSolution(candidateMerge) {
			solutionNum++
			solution = candidateMerge
		}
		candidateFill, candidateIsValid = candidateFill.IncrementCandidate()
	}
	return solution, solutionNum
}
