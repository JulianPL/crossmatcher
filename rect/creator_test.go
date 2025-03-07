package rect

import (
	"crossmatcher/collection"
	"regexp"
	"strings"
	"testing"
)

func TestCreator_MakeCrosswordRandomTrivial(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	crossword := MakeCrosswordRandomTrivial(alphabet, 4, 5)
	if !crossword.HasUniqueSolution() {
		t.Errorf("MakeCrosswordRandomTrivial should return a crossword with unique_solution")
	}
	rules := strings.Join(crossword.Vertical, "")
	rowRule := "^(a|b)+$"
	matched, _ := regexp.MatchString(rowRule, rules)
	if !matched {
		t.Errorf("MakeCrosswordRandomTrivial should return a crossword with only alphabet characters in the rules, got %s", rules)
	}
	rules = strings.Join(crossword.Horizontal, "")
	rowRule = "^(a|b)+$"
	matched, _ = regexp.MatchString(rowRule, rules)
	if !matched {
		t.Errorf("MakeCrosswordRandomTrivial should return a crossword with only alphabet characters in the rules, got %s", rules)
	}
}

func TestCreator_MakeCrosswordTreeRandomTrivial(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	crossword := MakeCrosswordTreeRandomTrivial(alphabet, 4, 5).ToCrossword()
	if !crossword.HasUniqueSolution() {
		t.Errorf("MakeCrosswordTreeRandomTrivial should return a crossword with unique_solution")
	}
	rules := strings.Join(crossword.Vertical, "")
	rowRule := "^(a|b)+$"
	matched, _ := regexp.MatchString(rowRule, rules)
	if !matched {
		t.Errorf("MakeCrosswordTreeRandomTrivial should return a crossword with only alphabet characters in the rules, got %s", rules)
	}
	rules = strings.Join(crossword.Horizontal, "")
	rowRule = "^(a|b)+$"
	matched, _ = regexp.MatchString(rowRule, rules)
	if !matched {
		t.Errorf("MakeCrosswordTreeRandomTrivial should return a crossword with only alphabet characters in the rules, got %s", rules)
	}
}

func TestCreator_initialSeparationTransformations(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	crosswordTrivial := MakeCrosswordTreeRandomTrivial(alphabet, 4, 5)
	crossword := crosswordTrivial.initialSeparationTransformations().ToCrossword()
	if !crossword.HasUniqueSolution() {
		t.Errorf("initialSeparationTransformations should return a crossword with unique_solution")
	}
	rules := strings.Join(crossword.Vertical, "")
	rowRule := "^(\\((a|b)+\\)\\+)+$"
	matched, _ := regexp.MatchString(rowRule, rules)
	if !matched {
		t.Errorf("initialSeparationTransformations should return a crossword with repetitions of only alphabet characters in the rules, got %s", rules)
	}
	rules = strings.Join(crossword.Horizontal, "")
	rowRule = "^(\\((a|b)+\\)\\+)+$"
	matched, _ = regexp.MatchString(rowRule, rules)
	if !matched {
		t.Errorf("initialSeparationTransformations should return a crossword with repetitions of only alphabet characters in the rules, got %s", rules)
	}
}

func TestCreator_MakeCrosswordRandom(t *testing.T) {
	alphabet := collection.MakeAlphabet("ab")
	crossword := MakeCrosswordRandom(alphabet, 4, 5)
	if !crossword.HasUniqueSolution() {
		t.Errorf("MakeCrosswordRandom should return a crossword with unique_solution")
	}
	rules := strings.Join(crossword.Vertical, "")
	rowRule := "^(\\((a|b|\\|)+\\)\\+)+$"
	matched, _ := regexp.MatchString(rowRule, rules)
	if !matched {
		t.Errorf("MakeCrosswordRandom should return a crossword with repetitions of alternations of only alphabet characters in the rules, got %s", rules)
	}
	rules = strings.Join(crossword.Horizontal, "")
	rowRule = "^(\\((a|b|\\|)+\\)\\+)+$"
	matched, _ = regexp.MatchString(rowRule, rules)
	if !matched {
		t.Errorf("MakeCrosswordRandom should return a crossword with repetitions of alternations of only alphabet characters in the rules, got %s", rules)
	}
	if len(crossword.Vertical) != 5 {
		t.Errorf("MakeCrosswordRandom should return a crossword with the correct width. Expected 5, got %d", len(crossword.Vertical))
	}

	if len(crossword.Horizontal) != 4 {
		t.Errorf("MakeCrosswordRandom should return a crossword with the correct height. Expected 4, got %d", len(crossword.Horizontal))
	}
}
