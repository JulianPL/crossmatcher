package collection

import "slices"

type Alphabet struct {
	number map[rune]int
	char   map[int]rune
}

const WildcardNumber = -1

// MakeAlphabet makes an alphabet that contains all non-wildcard characters from a given string.
func MakeAlphabet(characters string, wildcards ...rune) Alphabet {
	alphabet := Alphabet{make(map[rune]int), make(map[int]rune)}
	for _, char := range characters {
		if !slices.Contains(wildcards, char) {
			alphabet.Insert(char)
		}
	}
	return alphabet
}

func (a Alphabet) String() string {
	charString := ""
	for _, char := range a.char {
		charString += string(char)
	}
	return charString
}

func (a Alphabet) Copy() Alphabet {
	newAlphabet := Alphabet{make(map[rune]int), make(map[int]rune)}
	for key, value := range a.number {
		newAlphabet.number[key] = value
	}
	for key, value := range a.char {
		newAlphabet.char[key] = value
	}
	return newAlphabet
}

// Merge returns an alphabet with all characters from alphabet and from insert.
func (a Alphabet) Merge(insert Alphabet) Alphabet {
	newAlphabet := a.Copy()
	for _, char := range insert.char {
		newAlphabet.Insert(char)
	}
	return newAlphabet
}

func (a Alphabet) Insert(char rune) {
	if a.Contains(char) {
		return
	}
	num := a.Len()

	a.number[char] = num
	a.char[num] = char
}

func (a Alphabet) Contains(char rune) bool {
	_, ok := a.number[char]
	return ok
}

func (a Alphabet) Len() int {
	return len(a.number)
}

func (a Alphabet) Char(num int) (rune, bool) {
	char, ok := a.char[num]
	return char, ok
}

func (a Alphabet) Number(char rune) (int, bool) {
	num, ok := a.number[char]
	return num, ok
}
