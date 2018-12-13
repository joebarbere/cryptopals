package util

import "math/bits"

func HammingDistanceBytes(b1, b2 []byte) int {
	distance := 0
	for i, b := range b1 {
		distance += HammingDistanceByte(b, b2[i])
	}
	return distance
}

func HammingDistanceByte(b1, b2 byte) int {
	return CountBitsInByte(b1 ^ b2)
}

func CountBitsInByte(b byte) int {
	return bits.OnesCount(uint(b))
}
