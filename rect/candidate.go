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

type Content []lin.Content

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

// IncrementCandidate makes the lexicographically next candidate.
// The order uses the reversed content and the given alphabet.
// Fails on last candidate.
func (c Candidate) IncrementCandidate() (Candidate, bool) {
	success := false
	increment := c.Copy()
	for i := range increment.Content {
		row, _ := increment.GetRow(i)
		var ok bool
		row, ok = row.IncrementCandidate()
		increment.Content[i] = row.Content

		if ok {
			success = true
			break
		}
	}
	return increment, success
}

// Copy creates an exact copy of the content.
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

// Copy creates an exact copy of the candidate
func (c Candidate) Copy() Candidate {
	return Candidate{c.Content.Copy(), c.Alphabet.Copy()}
}

// GetRow restrict a candidate to the given row (which leaves a linear candidate)
func (c Candidate) GetRow(rowNumber int) (lin.Candidate, bool) {
	if len(c.Content) <= rowNumber {
		return lin.MakeCandidate(""), false
	}
	row, _ := lin.MakeCandidateManual(c.Content[rowNumber], c.Alphabet)
	return row, true
}

// GetCol restrict a candidate to the given col (which leaves a linear candidate)
func (c Candidate) GetCol(colNumber int) (lin.Candidate, bool) {
	var content lin.Content
	for _, row := range c.Content {
		if len(row) <= colNumber {
			return lin.MakeCandidate(""), false
		}
		content = append(content, row[colNumber])
	}
	col, _ := lin.MakeCandidateManual(content, c.Alphabet)
	return col, true
}
