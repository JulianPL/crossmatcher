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

func (crossword Crossword) CheckSolutionString(row string) bool {
	rowRule := "^(" + crossword.Rule + ")$"
	matched, _ := regexp.MatchString(rowRule, row)
	return matched
}

func (crossword Crossword) SolveBruteforce(constraint Candidate) (Candidate, int) {
	numWildcards := constraint.CountWildcards()
	candidate := MakeCandidateFirst(crossword.Alphabet, numWildcards)
	candidateIsValid := true
	solutionNum := 0
	var solution Candidate
	for candidateIsValid {
		row, _ := constraint.MergeRow(candidate)
		if crossword.CheckSolutionString(row) {
			solutionNum++
			solution, _ = solution.GreatestCommonPattern(MakeCandidateFromString(row))
		}
		candidate, candidateIsValid = candidate.IncrementCandidate()
	}
	return solution, solutionNum
}
