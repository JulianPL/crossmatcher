package lin

import (
	"crossmatcher/collection"
	"slices"
)

type Candidate struct {
	Content  Content
	Alphabet collection.Alphabet
}

type Content []int

// MakeCandidateFirst makes starting candidate for Incrementation.
// Fails on empty alphabet
func MakeCandidateFirst(alphabet collection.Alphabet, size int) (Candidate, bool) {
	if alphabet.Len() == 0 {
		return Candidate{}, false
	}

	content := make(Content, size)
	for i := 0; i < size; i++ {
		content[i] = 0
	}
	return Candidate{content, alphabet}, true
}

// MakeCandidateEmpty makes a candidate consisting of only wildcards.
func MakeCandidateEmpty(alphabet collection.Alphabet, size int) Candidate {
	content := make(Content, size)
	for i := 0; i < size; i++ {
		content[i] = -1
	}
	return Candidate{content, alphabet}
}

// MakeCandidate makes a candidate representing a string. All wildcards are mapped to alphabet-number -1.
func MakeCandidate(contentString string, wildcards ...rune) Candidate {
	alphabet := collection.MakeAlphabet(contentString, wildcards...)
	var content Content
	for _, char := range contentString {
		var num int
		if slices.Contains(wildcards, char) {
			num = -1
		} else {
			num, _ = alphabet.Number(char)
		}
		content = append(content, num)
	}
	return Candidate{content, alphabet}
}

// String returns the candidate. Wildcards are presented by the passed rune (default = '.')
func (c Candidate) String(wildcard ...rune) string {
	wildRune := '.'
	if wildcard != nil {
		wildRune = wildcard[0]
	}

	rowString := ""
	for _, num := range c.Content {
		if num == -1 {
			rowString += string(wildRune)
		} else {
			char, _ := c.Alphabet.Char(num)
			rowString += string(char)
		}
	}

	return rowString
}

// Len returns the length of content of a candidate.
func (c Candidate) Len() int {
	return len(c.Content)
}

// CountWildcards returns the number of wildcards in the content of a candidate.
func (c Candidate) CountWildcards() int {
	count := 0
	for _, char := range c.Content {
		if char == -1 {
			count++
		}
	}
	return count
}

// IncrementCandidate makes the lexicographically next candidate.
// The order uses the reversed content and the given alphabet.
// Fails on last candidate.
func (c Candidate) IncrementCandidate() (Candidate, bool) {
	success := false
	increment := c.Copy()
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

// Copy creates an exact copy of the content
func (c Content) Copy() Content {
	var newContent Content
	for _, char := range c {
		newContent = append(newContent, char)
	}
	return newContent
}

// Copy creates an exact copy of the candidate
func (c Candidate) Copy() Candidate {
	return Candidate{c.Content.Copy(), c.Alphabet.Copy()}
}

// Merge fills the candidate cFill into the wildcards of candidate c.
// Fails if the length of cFill does not match the number of wildcards of c.
func (c Candidate) Merge(cFill Candidate) (Candidate, bool) {
	if c.CountWildcards() != cFill.Len() {
		return Candidate{}, false
	}

	alphabetMerge := c.Alphabet.Merge(cFill.Alphabet)
	var contentMerge Content
	currentFill := 0

	for _, num := range c.Content {
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

		contentMerge = append(contentMerge, newNum)
	}

	return Candidate{contentMerge, alphabetMerge}, true
}

func (c Candidate) GreatestCommonPattern(cFill Candidate) (Candidate, bool) {
	if c.Len() == 0 {
		return cFill.Copy(), true
	}
	if c.Len() != cFill.Len() {
		return Candidate{}, false
	}
	alphabet := c.Alphabet.Merge(cFill.Alphabet)
	var content Content
	for i := range c.Content {
		num1 := c.Content[i]
		num2 := cFill.Content[i]
		if (num1 == -1) || (num2 == -1) {
			content = append(content, -1)
			continue
		}
		char1, _ := c.Alphabet.Char(num1)
		char2, _ := cFill.Alphabet.Char(num2)
		if char1 == char2 {
			num, _ := alphabet.Number(char1)
			content = append(content, num)
		} else {
			content = append(content, -1)
		}
	}
	return Candidate{content, alphabet}, true
}
