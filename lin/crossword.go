package lin

import (
	"crossmatcher/collection"
	"math/rand"
	"regexp"
	"strings"
)

type Crossword struct {
	Rule     string
	Alphabet collection.Alphabet
}

// MakeCrossword makes a crossword from a given string and an underlying alphabet.
func MakeCrossword(rule string, alphabet collection.Alphabet) Crossword {
	return Crossword{rule, alphabet}
}

func (crossword Crossword) MergeRandomBlocks() Crossword {
	rule := strings.Split(crossword.Rule, ")+(")
	if len(rule) == 1 {
		return MakeCrossword("", crossword.Alphabet)
	}
	sepIndex := 1
	if len(rule) > 2 {
		sepIndex = rand.Intn(len(rule)-2) + 1
	}
	left := strings.Join(rule[:sepIndex], ")+(")
	right := strings.Join(rule[sepIndex:], ")+(")
	return MakeCrossword(left+"|"+right, crossword.Alphabet)
}

// CheckSolution checks, whether a candidate without wildcards satisfies a crossword.
func (crossword Crossword) CheckSolution(candidate Candidate) bool {
	if candidate.CountWildcards() > 0 {
		return false
	}
	row := candidate.String()
	rowRule := "^(" + crossword.Rule + ")$"
	matched, _ := regexp.MatchString(rowRule, row)
	return matched
}

// SolveBruteforce checks all candidates that fill the wildcards given by the constraint.
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
