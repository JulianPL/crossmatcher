package rect

import (
	"crossmatcher/collection"
	"crossmatcher/lin"
	"math/rand"
)

type CrosswordTree struct {
	Horizontal []lin.RegexNode
	Vertical   []lin.RegexNode
	Alphabet   collection.Alphabet
}

func MakeRandomCrossword(alphabet collection.Alphabet, height, width int) Crossword {
	trivial := MakeCrosswordRandomTrivial(alphabet, height, width)
	horizontal := make([]lin.RegexNode, height)
	vertical := make([]lin.RegexNode, width)
	for i, rule := range trivial.Horizontal {
		horizontal[i] = lin.MakeRegexNode(rule)
	}
	for i, rule := range trivial.Vertical {
		vertical[i] = lin.MakeRegexNode(rule)
	}
	ret := CrosswordTree{Horizontal: horizontal, Vertical: vertical, Alphabet: alphabet}
	ret = ret.SeparateRulesIntoBlocks()

	for range 100 {
		ret.transformSingleRule()
	}

	return ret.ToCrossword()
}

// MakeCrosswordRandomTrivial makes a random trivial crossword over an underlying alphabet with given size.
func MakeCrosswordRandomTrivial(alphabet collection.Alphabet, height, width int) Crossword {
	runes := []rune(alphabet.String())
	solution := make([][]rune, height)
	for i := range height {
		row := make([]rune, width)
		for j := range width {
			row[j] = runes[rand.Intn(len(runes))]
		}
		solution[i] = row
	}
	horizontal := make([]string, height)
	vertical := make([]string, width)
	for i := range height {
		for j := range width {
			horizontal[i] += string(solution[i][j])
			vertical[j] += string(solution[i][j])
		}
	}
	return MakeCrossword(alphabet, horizontal, vertical)
}

func (c CrosswordTree) DeepCopy() CrosswordTree {
	ret := CrosswordTree{}
	horizontal := make([]lin.RegexNode, len(c.Horizontal))
	vertical := make([]lin.RegexNode, len(c.Vertical))
	for i, node := range c.Horizontal {
		horizontal[i] = node.DeepCopy()
	}
	for i, node := range c.Vertical {
		vertical[i] = node.DeepCopy()
	}
	ret.Horizontal = horizontal
	ret.Vertical = vertical
	ret.Alphabet = c.Alphabet
	return ret
}

func (c CrosswordTree) ToCrossword() Crossword {
	horizontal := make([]string, len(c.Horizontal))
	vertical := make([]string, len(c.Vertical))
	for i, rule := range c.Horizontal {
		horizontal[i] = rule.String()
	}
	for i, rule := range c.Vertical {
		vertical[i] = rule.String()
	}
	return MakeCrossword(c.Alphabet, horizontal, vertical)
}

func (c CrosswordTree) getRandomRuleRef() *lin.RegexNode {
	dimSum := len(c.Horizontal) + len(c.Vertical)
	rule := rand.Intn(dimSum)
	if rule < len(c.Horizontal) {
		return &c.Horizontal[rule]

	} else {
		rule -= len(c.Horizontal)
		return &c.Vertical[rule]
	}
}

func (c CrosswordTree) transformSingleRule() CrosswordTree {
	ruleRef := c.getRandomRuleRef()
	//Todo randomly select different transformation methods
	return c.MergeBlocks(ruleRef)
}

func (c CrosswordTree) MergeBlocks(ruleRef *lin.RegexNode) CrosswordTree {
	rule := *ruleRef
	rule = rule.MergeRandomBlocks()
	return c.tryRuleChange(ruleRef, rule)
}

func (c CrosswordTree) tryRuleChange(ruleRef *lin.RegexNode, newRule lin.RegexNode) CrosswordTree {
	oldRule := *ruleRef
	*ruleRef = newRule
	if !c.ToCrossword().hasUniqueSolution() {
		*ruleRef = oldRule
	}
	return c
}

// applyInitialSeparationTransformations applies a sequence of transformations to a slice of RegexNodes
func applyInitialSeparationTransformations(rules []lin.RegexNode) []lin.RegexNode {
	newRules := make([]lin.RegexNode, len(rules))
	for i, rule := range rules {
		newRules[i] = rule.
			SeparateIntoBlocks().
			WithAlternationSubgroups().
			WithRepetitionSubgroups()
	}
	return newRules
}

func (c CrosswordTree) SeparateRulesIntoBlocks() CrosswordTree {
	ret := c.DeepCopy()
	ret.Horizontal = applyInitialSeparationTransformations(ret.Horizontal)
	ret.Vertical = applyInitialSeparationTransformations(ret.Vertical)
	return ret
}
