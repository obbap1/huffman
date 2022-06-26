package decoder

import (
	"testing"

	"github.com/obbap1/huffman/common"
	"github.com/stretchr/testify/assert"
)

func TestDecoderHuffman(t *testing.T) {
	// useless test, i know
	assert := assert.New(t)
	input := "good day sir"
	decoderHuffman := Init(input, &common.HuffManNode{})
	originalStr, _ := decoderHuffman.Do()
	assert.Equal(originalStr, "") // since the Huffman node is empty. The only way to test this is to hardcode the node
}
