package lin

import (
	"strings"
	"testing"
)

func TestRuleTree_MakeRegexNode(t *testing.T) {
	actual := MakeRegexNode("abcda").String()
	expected := "abcda"
	if expected != actual {
		t.Errorf("MakeRegexNod or String is incorrect. Expected:%s, actual:%s", expected, actual)
	}

	actual = (&RegexNode{-1, "abcda", nil}).String()
	expected = ""
	if expected != actual {
		t.Errorf("RegexNode or String is incorrect. Expected:%s, actual:%s", expected, actual)
	}
}

func TestRuleTree_Stringer(t *testing.T) {
	child := RegexNode{Literal, "abcda", nil}
	parent := RegexNode{Repetition, "+", []RegexNode{child}}
	actual := parent.String()
	expected := "(abcda)+"
	if actual != expected {
		t.Errorf("String is incorrect. Expected:%s, actual:%s", expected, actual)
	}

	child1 := RegexNode{Literal, "abc", nil}
	child2 := RegexNode{Literal, "da", nil}
	parent = RegexNode{Alternation, "", []RegexNode{child1, child2}}
	actual = parent.String()
	expected = "abc|da"
	if actual != expected {
		t.Errorf("String is incorrect. Expected:%s, actual:%s", expected, actual)
	}
}

func TestRuleTree_WithRepetitionSubgroups(t *testing.T) {
	base := MakeRegexNode("RegularExpression")
	actual := base.SeparateIntoBlocks().WithAlternationSubgroups().WithRepetitionSubgroups().String()
	prefix := "(R"
	suffix := "n)+"
	if prefix != actual[:len(prefix)] {
		t.Errorf("WithRepetitionSubgroups is incorrect. Expected prefix:%s, actual:%s", prefix, actual)
	}
	if suffix != actual[len(actual)-len(suffix):] {
		t.Errorf("WithRepetitionSubgroups is incorrect. Expected suffix:%s, actual:%s", suffix, actual)
	}
}

func TestRuleTree_MergeRandomBlocks(t *testing.T) {
	base := MakeRegexNode("RegularExpression").SeparateIntoBlocks().WithAlternationSubgroups().WithRepetitionSubgroups()
	goal := base.MergeRandomBlocks()
	actual := goal.String()
	prefix := "(R"
	suffix := "n)+"
	infix := "|"
	if prefix != actual[:len(prefix)] {
		t.Errorf("WithRepetitionSubgroups is incorrect. Expected prefix:%s, actual:%s", prefix, actual)
	}
	if suffix != actual[len(actual)-len(suffix):] {
		t.Errorf("WithRepetitionSubgroups is incorrect. Expected suffix:%s, actual:%s", suffix, actual)
	}
	if strings.Count(actual, infix) != 1 {
		t.Errorf("WithRepetitionSubgroups is incorrect. Expected one occurrence of :%s, actual:%s", infix, actual)
	}
	if len(base.Children) != len(goal.Children)+1 {
		t.Errorf("WithRepetitionSubgroups is incorrect. Expected removal of one group :%s, actual:%s", base, goal)
	}
	// Small case should do nothing
	base = MakeRegexNode("R").SeparateIntoBlocks().WithAlternationSubgroups().WithRepetitionSubgroups()
	goal = base.MergeRandomBlocks()
	if base.String() != goal.String() {
		t.Errorf("WithRepetitionSubgroups is incorrect. Expected no change on short group :%s, actual:%s", base, goal)
	}
}

func TestRuleTree_SimplifyAlternations(t *testing.T) {
	base := strings.Repeat("a", 100)
	grouped := MakeRegexNode(base).SeparateIntoBlocks().WithAlternationSubgroups().WithRepetitionSubgroups()
	merged := grouped.MergeRandomBlocks()
	for range len(base) {
		merged = merged.MergeRandomBlocks()
	}
	simplified := merged.SimplifyAlternations()
	if len(simplified.Children) != 1 {
		t.Errorf("MergeRandomBlocks is incorrect. Expected only one child after many exections, actual:%s", merged.String())
	}
	if len(simplified.Children[0].Children[0].Children) > 6 {
		t.Errorf("SimplifyAlternations is incorrect. Expected no duplicates, actual:%s", simplified.String())
	}
}
