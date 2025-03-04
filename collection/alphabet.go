package collection

import "slices"

type Alphabet struct {
	number map[rune]int
	char   map[int]rune
}

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

func (alphabet Alphabet) String() string {
	charString := ""
	for _, char := range alphabet.char {
		charString += string(char)
	}
	return charString
}

func (alphabet Alphabet) Copy() Alphabet {
	newAlphabet := Alphabet{make(map[rune]int), make(map[int]rune)}
	for key, value := range alphabet.number {
		newAlphabet.number[key] = value
	}
	for key, value := range alphabet.char {
		newAlphabet.char[key] = value
	}
	return newAlphabet
}

// Merge returns an alphabet with all characters from alphabet and from insert.
func (alphabet Alphabet) Merge(insert Alphabet) Alphabet {
	newAlphabet := alphabet.Copy()
	for _, char := range insert.char {
		newAlphabet.Insert(char)
	}
	return newAlphabet
}

func (alphabet Alphabet) Insert(char rune) {
	if alphabet.Contains(char) {
		return
	}
	num := alphabet.Len()

	alphabet.number[char] = num
	alphabet.char[num] = char
}

func (alphabet Alphabet) Contains(char rune) bool {
	_, ok := alphabet.number[char]
	return ok
}

func (alphabet Alphabet) Len() int {
	return len(alphabet.number)
}

func (alphabet Alphabet) Char(num int) (rune, bool) {
	char, ok := alphabet.char[num]
	return char, ok
}

func (alphabet Alphabet) Number(char rune) (int, bool) {
	num, ok := alphabet.number[char]
	return num, ok
}
