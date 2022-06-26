package encoder

import (
	"container/heap"
	"strings"

	"github.com/obbap1/huffman/common"
)

type Huffman struct {
	Item string
	PQ   *PriorityQueue
}

func (h Huffman) getItemsAndFrequency() map[string]int {
	if len(h.Item) == 0 {
		panic("invalid input")
	}
	items := strings.Split(strings.TrimSpace(h.Item), "")
	result := make(map[string]int)
	for _, v := range items {
		if count, ok := result[v]; ok {
			result[v] = count + 1
		} else {
			result[v] = 1
		}
	}
	return result
}

func (h Huffman) heapify() int {
	itemsFrequency := h.getItemsAndFrequency()
	for item, frequency := range itemsFrequency {
		heap.Push(h.PQ, &common.HuffManNode{
			Value:    item,
			Priority: frequency,
			NodeType: common.Leaf,
		})
	}
	return len(itemsFrequency)
}

func (h Huffman) Do() (string, *common.HuffManNode, *common.Stats) {
	// heapify inputs
	frequencyLength := h.heapify()
	// implement the huffman algorithm
	for h.PQ.Len() > 1 {
		// pop the 2 smallest items
		first, second := heap.Pop(h.PQ).(*common.HuffManNode), heap.Pop(h.PQ).(*common.HuffManNode)
		// merge the nodes
		// To be strict, some versions of huffman use ASCII for ordering if the weights are the same
		// but this just merges left to right
		newNode := first.Merge(second)
		// push new node back to the queue
		heap.Push(h.PQ, newNode)
	}
	// the final item should be the result
	c := h.PQ.Pop().(*common.HuffManNode)
	// prepare results
	result := c.PrepareResults("", frequencyLength, h.Item)
	compressedByteSize := common.GetCompressedByteSize(len(result))
	originalByteSize := len([]byte(h.Item))
	return result, c, &common.Stats{
		OriginalByteSize:   originalByteSize,
		CompressedByteSize: compressedByteSize,
		PercentageSaved:    common.GetPercentageSaved(originalByteSize, compressedByteSize),
	}

}

// Init initializes a new huffman object
func Init(input string) *Huffman {
	h := &Huffman{}
	h.Item = input
	h.PQ = &PriorityQueue{}
	heap.Init(h.PQ)
	return h
}
