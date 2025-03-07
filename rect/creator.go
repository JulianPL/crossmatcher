package rect

import (
	"crossmatcher/collection"
	"crossmatcher/lin"
	"math/rand"
)

type TransformationType int

const (
	Merge int = iota
	Extend
	Shorten
)

type CrosswordTree struct {
	Horizontal []lin.RegexNode
	Vertical   []lin.RegexNode
	Alphabet   collection.Alphabet
}

// MakeCrosswordRandom makes a random crossword over an underlying alphabet with given size.
func MakeCrosswordRandom(alphabet collection.Alphabet, height, width int) Crossword {
	ret := MakeCrosswordTreeRandomTrivial(alphabet, height, width)
	ret = ret.initialSeparationTransformations()

	for range 5 * (height + width) {
		ret = ret.transformSingleRule(alphabet)
	}

	ret = ret.finalSeparationTransformations()

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

// MakeCrosswordTreeRandomTrivial makes a random trivial crossword tree over an underlying alphabet with given size.
func MakeCrosswordTreeRandomTrivial(alphabet collection.Alphabet, height, width int) CrosswordTree {
	trivial := MakeCrosswordRandomTrivial(alphabet, height, width)
	horizontal := make([]lin.RegexNode, height)
	vertical := make([]lin.RegexNode, width)
	for i, rule := range trivial.Horizontal {
		horizontal[i] = lin.MakeRegexNode(rule)
	}
	for i, rule := range trivial.Vertical {
		vertical[i] = lin.MakeRegexNode(rule)
	}
	return CrosswordTree{Horizontal: horizontal, Vertical: vertical, Alphabet: alphabet}
}

// DeepCopy creates a copy of a crossword tree which can be altered without side effects to the original tree
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

// ToCrossword converts the crossword tree to a crossword
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

// getRandomRuleRef returns the reference to a random rule (as lin.RegexNode)
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

// transformSingleRule applies a single random rule transformation to a single random rule
func (c CrosswordTree) transformSingleRule(alphabet collection.Alphabet) CrosswordTree {
	ruleRef := c.getRandomRuleRef()
	rule := getTransformationNumber()
	switch rule {
	case Merge:
		return c.MergeBlocks(ruleRef)
	case Extend:
		return c.ExtendAlternationElement(ruleRef, alphabet)
	case Shorten:
		return c.ShortenAlternationElement(ruleRef)
	default:
		return c
	}
}

// getTransformationProbabilityAcc returns the accumulated probabilities of the different rule transformations
func getTransformationProbabilityAcc() []float64 {
	return []float64{0.8, 0.90}
}

// getTransformationNumber gets a random number representing a rule transformation
func getTransformationNumber() int {
	blockProbabilityAcc := getTransformationProbabilityAcc()
	randomVal := rand.Float64()
	for i, acc := range blockProbabilityAcc {
		if randomVal < acc {
			return i
		}
	}
	return len(blockProbabilityAcc)
}

// MergeBlocks applies MergeRandomBlocks rule to a given rule
// Rollback if the new crossword does not have a unique solution anymore
func (c CrosswordTree) MergeBlocks(ruleRef *lin.RegexNode) CrosswordTree {
	rule := *ruleRef
	rule = rule.MergeRandomBlocks()
	return c.tryRuleChange(ruleRef, rule)
}

// ExtendAlternationElement applies ExtendRandomAlternationElement rule to a given rule
// Rollback if the new crossword does not have a unique solution anymore
func (c CrosswordTree) ExtendAlternationElement(ruleRef *lin.RegexNode, alphabet collection.Alphabet) CrosswordTree {
	rule := *ruleRef
	rule = rule.ExtendRandomAlternationElement(alphabet)
	return c.tryRuleChange(ruleRef, rule)
}

// ShortenAlternationElement applies ShortenRandomAlternationElement rule to a given rule
// Rollback if the new crossword does not have a unique solution anymore
func (c CrosswordTree) ShortenAlternationElement(ruleRef *lin.RegexNode) CrosswordTree {
	rule := *ruleRef
	rule = rule.ShortenRandomAlternationElement()
	return c.tryRuleChange(ruleRef, rule)
}

// tryRuleChange tests if a rule change results in a crossword with a unique solution
// if yes: the rule gets applied
// if no: the rule gets rolled back
func (c CrosswordTree) tryRuleChange(ruleRef *lin.RegexNode, newRule lin.RegexNode) CrosswordTree {
	oldRule := *ruleRef
	*ruleRef = newRule
	if !c.ToCrossword().HasUniqueSolution() {
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

// applyFinalSeparationTransformations applies a sequence of transformations to a slice of RegexNodes
func applyFinalSeparationTransformations(rules []lin.RegexNode) []lin.RegexNode {
	newRules := make([]lin.RegexNode, len(rules))
	for i, rule := range rules {
		newRules[i] = rule.SimplifyAlternations().RandomizeAlternations()
	}
	return newRules
}

// initialSeparationTransformations applies the finalSeparationTransformations to each rule.
func (c CrosswordTree) initialSeparationTransformations() CrosswordTree {
	ret := c.DeepCopy()
	ret.Horizontal = applyInitialSeparationTransformations(ret.Horizontal)
	ret.Vertical = applyInitialSeparationTransformations(ret.Vertical)
	return ret
}

// finalSeparationTransformations applies the finalSeparationTransformations to each rule.
func (c CrosswordTree) finalSeparationTransformations() CrosswordTree {
	ret := c.DeepCopy()
	ret.Horizontal = applyFinalSeparationTransformations(ret.Horizontal)
	ret.Vertical = applyFinalSeparationTransformations(ret.Vertical)
	return ret
}
