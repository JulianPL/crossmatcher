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
	if candidate.CountWildcards() > 0 {
		return false
	}
	row := candidate.String()
	rowRule := "^(" + crossword.Rule + ")$"
	matched, _ := regexp.MatchString(rowRule, row)
	return matched
}

func (crossword Crossword) SolveBruteforce(constraint Candidate) (Candidate, int) {
	numWildcards := constraint.CountWildcards()
	candidateFill, _ := MakeCandidateFirst(crossword.Alphabet, numWildcards)
	candidateIsValid := true
	solutionNum := 0
	var solution Candidate
	for candidateIsValid {
		candidateMerge, _ := constraint.Merge(candidateFill)
		if crossword.CheckSolution(candidateMerge) {
			solutionNum++
			solution, _ = solution.GreatestCommonPattern(candidateMerge)
		}
		candidateFill, candidateIsValid = candidateFill.IncrementCandidate()
	}
	return solution, solutionNum
}
