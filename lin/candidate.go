package lin

import (
	"crossmatcher/collection"
)

type Candidate struct {
	Content  Content
	Alphabet collection.Alphabet
}

type Content []int

func MakeCandidateFirst(alphabet collection.Alphabet, size int) Candidate {
	content := make(Content, size)
	for i := 0; i < size; i++ {
		content[i] = 0
	}
	return Candidate{content, alphabet}
}

func MakeCandidateEmpty(alphabet collection.Alphabet, size int) Candidate {
	content := make(Content, size)
	for i := 0; i < size; i++ {
		content[i] = -1
	}
	return Candidate{content, alphabet}
}

func MakeCandidate(content Content, alphabet collection.Alphabet) Candidate {
	return Candidate{content, alphabet}
}

func MakeCandidateFromString(contentString string) Candidate {
	alphabet := collection.MakeAlphabet(contentString)
	var content Content
	for _, char := range contentString {
		num, _ := alphabet.Number(char)
		content = append(content, num)
	}
	return Candidate{content, alphabet}
}

func (c Candidate) Len() int {
	return len(c.Content)
}

func (c Candidate) CountWildcards() int {
	count := 0
	for _, char := range c.Content {
		if char == -1 {
			count++
		}
	}
	return count
}

func (c Candidate) IncrementCandidate() (Candidate, bool) {
	success := false
	increment := Candidate{c.Content.Copy(), c.Alphabet}
	for i := 0; i < len(c.Content); i++ {
		if increment.Content[i] < increment.Alphabet.Len()-1 {
			increment.Content[i] += 1
			success = true
			break
		}
		increment.Content[i] = 0
	}
	return increment, success
}

func (c Content) Copy() Content {
	var newContent Content
	for _, char := range c {
		newContent = append(newContent, char)
	}
	return newContent
}

func (c Candidate) GetRow() (string, bool) {
	rowString := ""
	for i := 0; i < len(c.Content); i++ {
		char, ok := c.Alphabet.Char(c.Content[i])
		if !ok {
			return "", false
		}
		rowString += string(char)
	}
	return rowString, true
}

func (c Candidate) MergeRow(cFill Candidate) (string, bool) {
	if c.CountWildcards() != cFill.Len() {
		return "", false
	}
	rowFill, ok := cFill.GetRow()
	if !ok {
		return "", false
	}
	runesFill := []rune(rowFill)
	currentFill := 0
	var runesMerge []rune
	for _, num := range c.Content {
		if num != -1 {
			char, ok := c.Alphabet.Char(num)
			if !ok {
				return "", false
			}
			runesMerge = append(runesMerge, char)
		} else {
			runesMerge = append(runesMerge, runesFill[currentFill])
			currentFill++
		}
	}
	return string(runesMerge), true
}
