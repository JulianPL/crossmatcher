package rect

import (
	"crossmatcher/collection"
	"crossmatcher/lin"
	"strings"
)

type Candidate struct {
	Content  Content
	Alphabet collection.Alphabet
}

type Content [][]int

// MakeCandidateFirst makes starting candidate for Incrementation.
// Fails on empty alphabet.
func MakeCandidateFirst(alphabet collection.Alphabet, verticalSize int, horizontalSize int) (Candidate, bool) {
	if alphabet.Len() == 0 {
		return Candidate{}, false
	}

	content := make(Content, verticalSize)
	for i := range content {
		content[i] = make(lin.Content, horizontalSize)
		for j := range content[i] {
			content[i][j] = 0
		}
	}
	return Candidate{content, alphabet}, true
}

// MakeCandidate makes a candidate representing a string. All wildcards are mapped to alphabet-number -1.
func MakeCandidate(rows []string, wildcards ...rune) Candidate {
	alphabet := collection.MakeAlphabet(strings.Join(rows, ""), wildcards...)
	var content Content
	for _, rowString := range rows {
		row, _ := lin.MakeContent(rowString, alphabet, wildcards...)
		content = append(content, row)
	}
	return Candidate{content, alphabet}
}

func (candidate Candidate) IncrementCandidate() (Candidate, bool) {
	success := false
	increment := Candidate{candidate.Content.Copy(), candidate.Alphabet}
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

func (content Content) Copy() Content {
	var newContent Content
	for _, row := range content {
		var newRow []int
		for _, char := range row {
			newRow = append(newRow, char)
		}
		newContent = append(newContent, newRow)
	}
	return newContent
}

func (candidate Candidate) GetRow(row int) (string, bool) {
	if len(candidate.Content) <= row {
		return "", false
	}
	rowString := ""
	for i := 0; i < len(candidate.Content[row]); i++ {
		char, ok := candidate.Alphabet.Char(candidate.Content[row][i])
		if !ok {
			return "", false
		}
		rowString += string(char)
	}
	return rowString, true
}

func (candidate Candidate) GetCol(col int) (string, bool) {
	colString := ""
	for i := 0; i < len(candidate.Content); i++ {
		if len(candidate.Content[i]) <= col {
			return "", false
		}
		char, ok := candidate.Alphabet.Char(candidate.Content[i][col])
		if !ok {
			return "", false
		}
		colString += string(char)
	}
	return colString, true
}
