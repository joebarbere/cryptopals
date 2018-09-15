package util

import (
	"encoding/base64"
	"encoding/hex"
	"unicode"
)

type ScoreKeyValue struct {
	Key   rune
	Value int
}

type ChallengeFourStruct struct {
	Rune            rune
	Score           int
	EncryptedString string
	DecryptedString string
}

func HexToBase64(hexString string) string {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(hexBytes)
}

func Xor(a, b []byte) []byte {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	dst := make([]byte, n)
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return dst
}

func XorOneRune(runeByte byte, srcBytes []byte) []byte {
	n := len(srcBytes)
	destBytes := make([]byte, n)
	for i := 0; i < n; i++ {
		destBytes[i] = srcBytes[i] ^ runeByte
	}
	return destBytes
}

func NextRune(runeBytes []byte) func() byte {
	var pos int
	return func() byte {
		endPos := len(runeBytes) - 1
		curPos := pos
		var returnPos int
		if pos == endPos {
			pos = 0
			returnPos = endPos
		} else {
			pos++
			returnPos = curPos
		}
		return runeBytes[returnPos]
	}
}

func ScoreString(str string) int {
	scoreMap := make(map[rune]int)
	scoreMap['E'] = 13
	scoreMap['T'] = 12
	scoreMap['A'] = 11
	scoreMap['O'] = 10
	scoreMap['I'] = 9
	scoreMap['N'] = 8
	scoreMap[' '] = 7
	scoreMap['S'] = 6
	scoreMap['H'] = 5
	scoreMap['R'] = 4
	scoreMap['D'] = 3
	scoreMap['L'] = 2
	scoreMap['U'] = 1

	score := 0

	for _, r := range str {
		score += scoreMap[unicode.ToUpper(r)]
	}

	return score
}
