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

func (candidate Candidate) GetRow() (string, bool) {
	rowString := ""
	for i := 0; i < len(candidate.Content); i++ {
		char, ok := candidate.Alphabet.Char(candidate.Content[i])
		if !ok {
			return "", false
		}
		rowString += string(char)
	}
	return rowString, true
}
