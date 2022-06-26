package decoder

import (
	"github.com/obbap1/huffman/common"
)

type Huffman struct {
	Item string
	Tree *common.HuffManNode
}

func Init(str string, tree *common.HuffManNode) *Huffman {
	return &Huffman{
		Item: str,
		Tree: tree,
	}
}

func (h *Huffman) Do() (string, *common.Stats) {
	node := &common.HuffManNode{}
	answer := ""
	for _, b := range h.Item {
		bit := string(b)
		// if node is empty, initialize to root
		if node.Priority == 0 {
			node = h.Tree
		}
		if bit == "0" {
			// Go left
			node = node.LeftChild

		} else if bit == "1" {
			// Go right
			node = node.RightChild
		}
		if node.IsLeaf() {
			answer += node.Value
			node = h.Tree
		}
	}
	originalByteSize := len(answer)
	return answer, &common.Stats{
		OriginalByteSize: originalByteSize,
	}
}
