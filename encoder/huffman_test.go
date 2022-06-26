package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoderHuffman(t *testing.T) {
	assert := assert.New(t)
	input := "good day sir"
	encoderHuffman := Init(input)
	compressedStr, huffmanNode, stats := encoderHuffman.Do()
	assert.NotNil(huffmanNode)
	assert.Equal(stats.OriginalByteSize, len(input))
	assert.NotEmpty(compressedStr)
}
