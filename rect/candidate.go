package rect

import (
	"crossmatcher/collection"
	"strings"
)

type Candidate struct {
	Content  Content
	Alphabet collection.Alphabet
}

type Content [][]int

func MakeCandidateFirst(alphabet collection.Alphabet, horizontalSize int, verticalSize int) Candidate {
	content := make(Content, horizontalSize)
	for i := 0; i < horizontalSize; i++ {
		content[i] = make([]int, verticalSize)
		for j := 0; j < verticalSize; j++ {
			content[i][j] = 0
		}
	}
	return Candidate{content, alphabet}
}

func MakeCandidate(rows []string) Candidate {
	alphabet := collection.MakeAlphabet(strings.Join(rows, ""))
	var content Content
	for _, rowString := range rows {
		var row []int
		for _, char := range rowString {
			num, _ := alphabet.Number(char)
			row = append(row, num)
		}
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
