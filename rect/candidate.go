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

// MakeCandidateEmpty makes a candidate consisting of only wildcards.
func MakeCandidateEmpty(alphabet collection.Alphabet, verticalSize int, horizontalSize int) Candidate {
	content := make(Content, verticalSize)
	for i := range content {
		content[i] = make(lin.Content, horizontalSize)
		for j := range content[i] {
			content[i][j] = -1
		}
	}
	return Candidate{content, alphabet}
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

// String returns the candidate.
// Wildcards are presented by the passed rune (default = '.').
func (c Candidate) String(wildcard ...rune) string {
	ret := ""
	for _, rowContent := range c.Content {
		row, _ := lin.MakeCandidateManual(rowContent, c.Alphabet)
		ret += row.String(wildcard...) + "\n"
	}
	return strings.Trim(ret, "\n")
}

// CountWildcards returns the number of wildcards in the content of a candidate.
func (c Candidate) CountWildcards() int {
	count := 0
	for _, rowContent := range c.Content {
		row, _ := lin.MakeCandidateManual(rowContent, c.Alphabet)
		count += row.CountWildcards()
	}
	return count
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

func (c Candidate) Merge(cFill lin.Candidate) (Candidate, bool) {
	if c.CountWildcards() != cFill.Len() {
		return Candidate{}, false
	}

	alphabetMerge := c.Alphabet.Merge(cFill.Alphabet)
	var contentMerge Content
	currentFill := 0

	for _, row := range c.Content {
		var rowContent lin.Content
		for _, num := range row {
			var newNum int
			if num != -1 {
				char, _ := c.Alphabet.Char(num)
				newNum, _ = alphabetMerge.Number(char)
			} else {
				if cFill.Content[currentFill] == -1 {
					newNum = -1
				} else {
					char, _ := cFill.Alphabet.Char(cFill.Content[currentFill])
					newNum, _ = alphabetMerge.Number(char)
				}
				currentFill++
			}

			rowContent = append(rowContent, newNum)
		}
		contentMerge = append(contentMerge, rowContent)
	}
	return Candidate{contentMerge, alphabetMerge}, true
}

// GetRow restrict a candidate to the given row (which leaves a linear candidate)
func (c Candidate) GetRow(rowNumber int) (lin.Candidate, bool) {
	if len(c.Content) <= rowNumber {
		return lin.MakeCandidate(""), false
	}
	row, _ := lin.MakeCandidateManual(c.Content[rowNumber], c.Alphabet)
	return row, true
}

func (c Candidate) UpdateRow(rowInsert lin.Candidate, rowNumber int) (Candidate, bool) {
	if len(c.Content) <= rowNumber {
		return c.Copy(), false
	}
	if len(c.Content[rowNumber]) != len(rowInsert.Content) {
		return c.Copy(), false
	}

	alphabet := c.Alphabet.Merge(rowInsert.Alphabet)
	context := c.Content.Copy()
	for colNumber := range rowInsert.Content {
		num := rowInsert.Content[colNumber]
		if num == -1 {
			context[rowNumber][colNumber] = -1
		} else {
			char, _ := rowInsert.Alphabet.Char(num)
			newNum, _ := alphabet.Number(char)
			context[rowNumber][colNumber] = newNum
		}
	}
	return Candidate{context, alphabet}, true
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

func (c Candidate) UpdateCol(colInsert lin.Candidate, colNumber int) (Candidate, bool) {
	if len(c.Content) != len(colInsert.Content) {
		return c.Copy(), false
	}

	alphabet := c.Alphabet.Merge(colInsert.Alphabet)
	context := c.Content.Copy()

	for rowNumber := range colInsert.Content {
		if len(c.Content[rowNumber]) <= colNumber {
			return c.Copy(), false
		}

		num := colInsert.Content[rowNumber]
		if num == -1 {
			context[rowNumber][colNumber] = -1
		} else {
			char, _ := colInsert.Alphabet.Char(num)
			newNum, _ := alphabet.Number(char)
			context[rowNumber][colNumber] = newNum
		}
	}
	return Candidate{context, alphabet}, true
}
