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

func (c Candidate) GetRowWithWildcard(wildcard rune) (string, bool) {
	rowString := ""
	for i := 0; i < len(c.Content); i++ {
		if c.Content[i] == -1 {
			rowString += string(wildcard)
			continue
		}
		char, ok := c.Alphabet.Char(c.Content[i])
		if !ok {
			return "", false
		}
		rowString += string(char)
	}
	return rowString, true
}

func (c Candidate) GetNonWildcards() (string, bool) {
	rowString := ""
	for i := 0; i < len(c.Content); i++ {
		if c.Content[i] == -1 {
			continue
		}
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

func (c Candidate) GreatestCommonPattern(cFill Candidate) (Candidate, bool) {
	if c.Len() == 0 {
		alphabet := cFill.Alphabet
		content := cFill.Content.Copy()
		return Candidate{content, alphabet}, true
	}
	if c.Len() != cFill.Len() {
		return Candidate{}, false
	}
	alphabetString1, ok1 := c.GetNonWildcards()
	alphabetString2, ok2 := cFill.GetNonWildcards()
	if !ok1 || !ok2 {
		return Candidate{}, false
	}
	alphabet := collection.MakeAlphabet(alphabetString1 + alphabetString2)
	var content Content
	for i := range c.Content {
		num1 := c.Content[i]
		num2 := cFill.Content[i]
		if (num1 == -1) || (num2 == -1) {
			content = append(content, -1)
			continue
		}
		char1, ok1 := c.Alphabet.Char(num1)
		char2, ok2 := cFill.Alphabet.Char(num2)
		if !ok1 || !ok2 {
			return Candidate{}, false
		}
		if char1 == char2 {
			num, _ := alphabet.Number(char1)
			content = append(content, num)
		} else {
			content = append(content, -1)
		}
	}
	return Candidate{content, alphabet}, true
}
