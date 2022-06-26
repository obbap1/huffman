package encoder

import (
	"container/heap"
	"testing"

	"github.com/obbap1/huffman/common"
	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	assert := assert.New(t)
	items := map[string]int{
		"f": 4,
		"e": 100,
		"g": 2,
		"l": 15,
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	for k, v := range items {
		heap.Push(pq, &common.HuffManNode{
			Value:    k,
			Priority: v,
		})
	}
	// the first item should be "g", since it has the least priority
	item := heap.Pop(pq).(*common.HuffManNode)
	assert.Equal(item.Value, "g")

	// Add item inbetween
	heap.Push(pq, &common.HuffManNode{
		Value:    "w",
		Priority: 50,
	})
	// next should be "f"
	item = heap.Pop(pq).(*common.HuffManNode)
	assert.Equal(item.Value, "f")
	// next should be "l"
	item = heap.Pop(pq).(*common.HuffManNode)
	assert.Equal(item.Value, "l")
	// next should be "w"
	item = heap.Pop(pq).(*common.HuffManNode)
	assert.Equal(item.Value, "w")
	// next should be "e"
	item = heap.Pop(pq).(*common.HuffManNode)
	assert.Equal(item.Value, "e")
}
