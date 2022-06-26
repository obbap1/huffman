package decoder

import (
	"testing"

	encoder "github.com/obbap1/huffman/encoder"
	"github.com/stretchr/testify/assert"
)

func TestDecoderHuffman(t *testing.T) {
	assert := assert.New(t)
	input := "good day sir"
	encoderHuffman := encoder.Init(input)
	compressedStr, huffmanNode, encoderStats := encoderHuffman.Do()
	decoderHuffman := Init(compressedStr, huffmanNode)
	originalStr, decodedStats := decoderHuffman.Do()
	assert.Equal(originalStr, input)
	assert.Equal(encoderStats.OriginalByteSize, decodedStats.OriginalByteSize)
}
