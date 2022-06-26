package common

import "math"

func GetPercentageSaved(originalByteSize, compressedByteSize int) float64 {
	return float64(originalByteSize-compressedByteSize) / float64(originalByteSize) * 100
}

func GetCompressedByteSize(resultLength int) int {
	compressedByteSize := resultLength / 8
	if math.Remainder(float64(resultLength), float64(8)) != 0 {
		compressedByteSize += 1
	}
	return compressedByteSize
}
