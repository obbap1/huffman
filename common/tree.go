package common

import (
	"fmt"
)

type NodeType int

const (
	Leaf NodeType = iota
	Internal
)

type Stats struct {
	OriginalByteSize   int
	CompressedByteSize int
	PercentageSaved    float64
}

type HuffManNode struct {
	Parent     *HuffManNode
	LeftChild  *HuffManNode
	RightChild *HuffManNode
	NodeType   NodeType
	Value      string
	Priority   int
	Index      int
}

// Merge joins two trees together and gives them a common parent
func (h *HuffManNode) Merge(node *HuffManNode) *HuffManNode {
	parent := &HuffManNode{
		LeftChild:  h,
		RightChild: node,
		NodeType:   Internal,
		Priority:   h.Priority + node.Priority,
	}
	h.Parent = parent
	node.Parent = parent
	return parent
}

func (h HuffManNode) IsLeaf() bool { return h.NodeType == Leaf }

// RecursivePrint prints the tree from root to leaf
func (h HuffManNode) RecursivePrint() {
	if h.IsLeaf() {
		fmt.Printf("leaf value is %s .", h.Value)
		return
	}
	fmt.Printf("%+v, %+v", h.LeftChild, h.RightChild)
	h.LeftChild.RecursivePrint()
	h.RightChild.RecursivePrint()
}

var Result map[string]string = map[string]string{}

func (h HuffManNode) recursiveWalk(value string, itemLength int) map[string]string {
	if h.IsLeaf() {
		Result[h.Value] = value
	} else {
		h.LeftChild.recursiveWalk(value+"0", itemLength)  // left is a zero
		h.RightChild.recursiveWalk(value+"1", itemLength) // right is a one
	}

	if len(Result) == itemLength {
		return Result
	}
	return nil
}

func (h HuffManNode) PrepareResults(value string, itemLength int, item string) string {
	resultsMap := h.recursiveWalk(value, itemLength)
	result := ""
	for _, v := range item {
		val := resultsMap[string(v)]
		result += val
	}
	return result
}
