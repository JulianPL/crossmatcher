package lin

import (
	"math/rand"
)

type RegexNodeType int

const (
	Literal RegexNodeType = iota
	Concatenation
	Alternation
	Repetition
)

type RegexNode struct {
	Type     RegexNodeType
	Value    string
	Children []RegexNode
}

func (node RegexNode) String() string {
	switch node.Type {
	case Literal:
		return node.Value
	case Concatenation:
		ret := ""
		for _, child := range node.Children {
			ret += child.String()
		}
		return ret
	case Alternation:
		ret := ""
		for i, child := range node.Children {
			if i > 0 {
				ret += "|"
			}
			ret += child.String()
		}
		return ret
	case Repetition:
		return "(" + node.Children[0].String() + ")" + node.Value
	default:
		return ""
	}
}

func (node RegexNode) DeepCopy() RegexNode {
	if node.Type == Literal {
		return RegexNode{Type: Literal, Value: node.Value}
	}
	newNode := RegexNode{Type: node.Type, Value: node.Value}
	for _, child := range node.Children {
		newNode.Children = append(newNode.Children, child.DeepCopy())
	}
	return newNode
}

// SimplifyAlternations make a deep copy without string-duplicates in Alteration nodes
func (node RegexNode) SimplifyAlternations() RegexNode {
	if node.Type == Literal {
		return RegexNode{Type: Literal, Value: node.Value}
	}
	newNode := RegexNode{Type: node.Type, Value: node.Value}
	if node.Type == Alternation {
		seen := make(map[string]struct{})
		for _, child := range node.Children {
			simplifiedChild := child.SimplifyAlternations()
			key := simplifiedChild.String()

			if _, exists := seen[key]; !exists {
				seen[key] = struct{}{}
				newNode.Children = append(newNode.Children, simplifiedChild)
			}
		}
	} else {
		for _, child := range node.Children {
			newNode.Children = append(newNode.Children, child.SimplifyAlternations())
		}
	}
	return newNode
}

// MakeRegexNode makes a regexNode corresponding to a concatenation of alphabet characters
func MakeRegexNode(value string) RegexNode {
	children := make([]RegexNode, len(value))
	for i, char := range value {
		children[i] = RegexNode{Type: Literal, Value: string(char)}
	}
	return RegexNode{Type: Concatenation, Value: "", Children: children}
}

// SeparateIntoBlocks splits the rule
// consisting of only a concatenation of characters from the alphabet randomly into blocks
// currently biased because the split starts from the front
func (node RegexNode) SeparateIntoBlocks() RegexNode {
	if node.Type != Concatenation {
		return node
	}

	ret := node.DeepCopy()

	childrenOld := ret.Children
	childrenNew := make([]RegexNode, 0)
	index := 0

	for index < len(childrenOld) {
		end := index + getBlockLength()
		if end > len(childrenOld) {
			end = len(childrenOld)
		}
		group := RegexNode{Concatenation, "", nil}
		for _, child := range childrenOld[index:end] {
			group.Children = append(group.Children, child)
		}
		childrenNew = append(childrenNew, group)
		index = end
	}
	ret.Children = childrenNew

	return ret
}

// WithAlternationSubgroups replaces each child with a repetition of itself
func (node RegexNode) WithAlternationSubgroups() RegexNode {
	ret := node.DeepCopy()
	for i, child := range ret.Children {
		repetition := RegexNode{Alternation, "", []RegexNode{child}}
		ret.Children[i] = repetition
	}
	return ret
}

// WithRepetitionSubgroups replaces each child with a repetition of itself
func (node RegexNode) WithRepetitionSubgroups() RegexNode {
	ret := node.DeepCopy()
	for i, child := range ret.Children {
		repetition := RegexNode{Repetition, "+", []RegexNode{child}}
		ret.Children[i] = repetition
	}
	return ret
}

// MergeRandomBlocks Merges the Alternation-Grandchildren of Repetition-Children
func (node RegexNode) MergeRandomBlocks() RegexNode {
	if len(node.Children) <= 1 {
		return node
	}
	ret := node.DeepCopy()
	leftIndex := 0
	if len(ret.Children) > 2 {
		leftIndex = rand.Intn(len(ret.Children) - 2)
	}
	rightIndex := leftIndex + 1
	leftAlternation := &ret.Children[leftIndex].Children[0].Children
	rightAlternation := &ret.Children[rightIndex].Children[0].Children
	*leftAlternation = append(*leftAlternation, *rightAlternation...)

	ret.Children = append(ret.Children[:rightIndex], ret.Children[rightIndex+1:]...)

	return ret
}

// getBlockProbabilityAcc returns the accumulated probabilities of blocks of different sizes
// TODO: At some point in the future, this should depend on the size of the crossword
func getBlockProbabilityAcc() []float64 {
	return []float64{0.0, 0.05, 0.35, 0.8}
}

// getBlockLength gets a random size from the accumulated block lengths
func getBlockLength() int {
	blockProbabilityAcc := getBlockProbabilityAcc()
	randomVal := rand.Float64()
	for i, acc := range blockProbabilityAcc {
		if randomVal < acc {
			return i
		}
	}
	return len(blockProbabilityAcc)
}
