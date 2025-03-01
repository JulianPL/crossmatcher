package rect

import (
	"crossmatcher/collection"
	"crossmatcher/lin"
	"math/rand"
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

// MakeCrosswordRandomTrivial makes a random trivial crossword over an underlying alphabet with given size.
func MakeCrosswordRandomTrivial(alphabet collection.Alphabet, height, width int) Crossword {
	runes := []rune(alphabet.String())
	solution := make([][]rune, height)
	for i := range height {
		row := make([]rune, width)
		for j := range width {
			row[j] = runes[rand.Intn(len(runes))]
		}
		solution[i] = row
	}
	horizontal := make([]string, height)
	vertical := make([]string, width)
	for i := range height {
		for j := range width {
			horizontal[i] += string(solution[i][j])
			vertical[j] += string(solution[i][j])
		}
	}
	return MakeCrossword(alphabet, horizontal, vertical)
}

func SeparateRuleArrayIntoBlocks(rules []string, alphabet collection.Alphabet) []string {
	ret := make([]string, len(rules))
	for ruleNumber := range rules {
		row := lin.MakeCrossword(rules[ruleNumber], alphabet)
		row = row.SeparateIntoBlocks()
		ret[ruleNumber] = row.Rule
	}
	return ret
}

func (c Crossword) hasUniqueSolution() bool {
	candidate := MakeCandidateEmpty(c.Alphabet, len(c.Vertical), len(c.Horizontal))
	solution, _ := c.SolveLinearReductions(candidate)
	return c.CheckSolution(solution)
}

func (c Crossword) tryRuleChange(ruleRef *string, newRule string) Crossword {
	if newRule == "" {
		return c
	}
	oldRule := *ruleRef
	*ruleRef = newRule
	if !c.hasUniqueSolution() {
		*ruleRef = oldRule
	}
	return c
}

func (c Crossword) MergeBlocks(ruleRef *string) Crossword {
	oldRule := *ruleRef
	row := lin.MakeCrossword(oldRule, c.Alphabet)
	row = row.MergeRandomBlocks()
	newRule := row.Rule
	return c.tryRuleChange(ruleRef, newRule)
}

func (c Crossword) getRandomRuleRef() *string {
	dimSum := len(c.Horizontal) + len(c.Vertical)
	rule := rand.Intn(dimSum)
	if rule < len(c.Horizontal) {
		return &c.Horizontal[rule]

	} else {
		rule -= len(c.Horizontal)
		return &c.Vertical[rule]
	}
}

func (c Crossword) transformSingleRule() Crossword {
	ruleRef := c.getRandomRuleRef()
	//Todo randomly select different transformation methods
	return c.MergeBlocks(ruleRef)
}

func MakeCrosswordRandomGrouped(alphabet collection.Alphabet, height, width int) Crossword {
	ret := MakeCrosswordRandomTrivial(alphabet, height, width)
	ret.Horizontal = SeparateRuleArrayIntoBlocks(ret.Horizontal, alphabet)
	ret.Vertical = SeparateRuleArrayIntoBlocks(ret.Vertical, alphabet)
	for range 100 {
		ret = ret.transformSingleRule()
	}
	return ret
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
