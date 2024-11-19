package rect

import (
	"crossmatcher/collection"
)

type Candidate struct {
	Content  [][]int
	Alphabet collection.Alphabet
}

func MakeCandidateFirst(alphabet collection.Alphabet, horizontalSize int, verticalSize int) Candidate {
	content := make([][]int, horizontalSize)
	for i := 0; i < horizontalSize; i++ {
		content[i] = make([]int, verticalSize)
		for j := 0; j < verticalSize; j++ {
			content[i][j] = 0
		}
	}
	return Candidate{content, alphabet}
}

func (candidate Candidate) IncrementCandidate() (Candidate, bool) {
	success := false
	increment := Candidate{candidate.Content, candidate.Alphabet}
	for i := 0; i < len(candidate.Content); i++ {
		for j := 0; j < len(candidate.Content[i]); j++ {
			if increment.Content[i][j] < increment.Alphabet.Len()-1 {
				increment.Content[i][j] += 1
				success = true
				break
			}
			increment.Content[i][j] = 0
		}
		if success {
			break
		}
	}
	return increment, success
}

func (candidate Candidate) GetRow(row int) (string, bool) {
	if len(candidate.Content) <= row {
		return "", false
	}
	rowString := ""
	for i := 0; i < len(candidate.Content[row]); i++ {
		// TODO test for non-existing numbers
		rowString += string(candidate.Alphabet.Char(candidate.Content[row][i]))
	}
	return rowString, true
}

func (candidate Candidate) GetCol(col int) (string, bool) {
	colString := ""
	for i := 0; i < len(candidate.Content); i++ {
		if len(candidate.Content[i]) <= col {
			return "", false
		}
		// TODO test for non-existing numbers
		colString += string(candidate.Alphabet.Char(candidate.Content[i][col]))
	}
	return colString, true
}
