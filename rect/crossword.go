package rect

import (
	"crossmatcher/collection"
	"crossmatcher/lin"
	"fmt"
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

func (c Crossword) String() string {
	return fmt.Sprintf("%v\n%v\n%v", c.Horizontal, c.Vertical, c.Alphabet)
}

func (c Crossword) hasUniqueSolution() bool {
	candidate := MakeCandidateEmpty(c.Alphabet, len(c.Horizontal), len(c.Vertical))
	solution, _ := c.SolveLinearReductions(candidate)
	return c.CheckSolution(solution)
}

func (c Crossword) GetRow(rowNumber int) (lin.Crossword, bool) {
	if len(c.Horizontal) <= rowNumber {
		return lin.MakeCrossword("", collection.MakeAlphabet("")), false
	}
	row := lin.MakeCrossword(c.Horizontal[rowNumber], c.Alphabet)
	return row, true
}

func (c Crossword) GetCol(colNumber int) (lin.Crossword, bool) {
	if len(c.Vertical) <= colNumber {
		return lin.MakeCrossword("", collection.MakeAlphabet("")), false
	}
	col := lin.MakeCrossword(c.Vertical[colNumber], c.Alphabet)
	return col, true
}

// CheckSolution checks, whether a candidate without wildcards satisfies a crossword.
func (c Crossword) CheckSolution(candidate Candidate) bool {
	if candidate.CountWildcards() > 0 {
		return false
	}

	for rowNumber := range candidate.Content {
		rowContent, _ := candidate.GetRow(rowNumber)
		rowCrossword, _ := c.GetRow(rowNumber)
		if !rowCrossword.CheckSolution(rowContent) {
			return false
		}
	}

	for colNumber := range candidate.Content[0] {
		colContent, _ := candidate.GetCol(colNumber)
		colCrossword, _ := c.GetCol(colNumber)
		if !colCrossword.CheckSolution(colContent) {
			return false
		}
	}
	return true
}

// SolveBruteforce checks all candidates that fill the wildcards given by the constraint.
func (c Crossword) SolveBruteforce(constraint Candidate) (Candidate, int) {
	candidateFill, _ := lin.MakeCandidateFirst(c.Alphabet, constraint.CountWildcards())

	candidateIsValid := true
	solutionNum := 0
	var solution Candidate
	for candidateIsValid {
		candidateMerge, _ := constraint.Merge(candidateFill)
		if c.CheckSolution(candidateMerge) {
			solutionNum++
			solution = candidateMerge
		}
		candidateFill, candidateIsValid = candidateFill.IncrementCandidate()
	}
	return solution, solutionNum
}

func (c Crossword) SolveLinearReductions(constraint Candidate) (Candidate, int) {
	next := constraint.Copy()
	for rowNumber := range c.Horizontal {
		rowRule := c.Horizontal[rowNumber]
		rowCrossword := lin.MakeCrossword(rowRule, c.Alphabet)
		rowConstraint, _ := next.GetRow(rowNumber)
		rowSolved, rowNumSolutions := rowCrossword.SolveBruteforce(rowConstraint)
		if rowNumSolutions == 0 {
			return Candidate{}, 0
		}
		next, _ = next.UpdateRow(rowSolved, rowNumber)
	}
	for colNumber := range c.Vertical {
		colRule := c.Vertical[colNumber]
		colCrossword := lin.MakeCrossword(colRule, c.Alphabet)
		colConstraint, _ := next.GetCol(colNumber)
		colSolved, colNumSolutions := colCrossword.SolveBruteforce(colConstraint)
		if colNumSolutions == 0 {
			return Candidate{}, 0
		}
		next, _ = next.UpdateCol(colSolved, colNumber)
	}
	if next.CountWildcards() == constraint.CountWildcards() {
		return next, 1
	}
	result, depth := c.SolveLinearReductions(next)
	return result, depth + 1
}
