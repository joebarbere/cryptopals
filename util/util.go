package util

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"sort"
	"unicode"
)

const RUNES = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789"

type ScoreKeyValue struct {
	Key   rune
	Value int
}

type CryptopalsResult struct {
	Key             []rune
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

func BruteXorOneByte(hexString string) CryptopalsResult {
	cryptopalsResult := CryptopalsResult{}

	hexBytes, _ := hex.DecodeString(hexString)
	bruteForceMap := make(map[rune]string)
	scoresMap := make(map[rune]int)
	for _, r := range RUNES {
		runeByte := []byte(string(r))[0]
		decrypted := string(XorOneByte(runeByte, hexBytes))
		bruteForceMap[r] = decrypted
		scoresMap[r] = ScoreString(decrypted)
	}

	var scoresSlice []ScoreKeyValue
	for k, v := range scoresMap {
		scoresSlice = append(scoresSlice, ScoreKeyValue{Key: k, Value: v})
	}

	sort.Slice(scoresSlice, func(i, j int) bool {
		return scoresSlice[i].Value > scoresSlice[j].Value
	})

	topScoredRune := scoresSlice[0].Key
	topScore := scoresSlice[0].Value
	cryptopalsResult.Key = []rune{topScoredRune}
	cryptopalsResult.EncryptedString = hexString
	cryptopalsResult.DecryptedString = bruteForceMap[topScoredRune]
	cryptopalsResult.Score = topScore

	return cryptopalsResult
}

func XorHexStrings(hexString1, hexString2 string) string {
	hexBytes1, _ := hex.DecodeString(hexString1)
	hexBytes2, _ := hex.DecodeString(hexString2)
	destBytes := Xor(hexBytes1, hexBytes2)
	hexEncodedString := hex.EncodeToString(destBytes)
	return hexEncodedString
}

func XorOneByte(key byte, srcBytes []byte) []byte {
	n := len(srcBytes)
	destBytes := make([]byte, n)
	for i := 0; i < n; i++ {
		destBytes[i] = srcBytes[i] ^ key
	}
	return destBytes
}

func NextByte(key []byte) func() byte {
	var pos int
	return func() byte {
		endPos := len(key) - 1
		curPos := pos
		var returnPos int
		if pos == endPos {
			pos = 0
			returnPos = endPos
		} else {
			pos++
			returnPos = curPos
		}
		return key[returnPos]
	}
}

func RepeatingKeyXorString(key, message string) string {
	encryptedBytes := RepeatingKeyXor(key, message)
	hexEncodedCiphertext := hex.EncodeToString(encryptedBytes)
	return hexEncodedCiphertext
}

func RepeatingKeyXor(key, message string) []byte {
	f := NextByte([]byte(key))
	encryptedBytes := make([]byte, len(message))
	for i, r := range message {
		x := f()
		encryptedBytes[i] = byte(r) ^ x
	}
	return encryptedBytes
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

func GetHammingDistance(s1, s2 string) (int, error) {
	if len(s1) != len(s2) {
		return 0, errors.New("ERROR:  Strings are of different lengths")
	}
	var differences int
	for i, x := range s1 {
		if string(x) != string(s2[i]) {
			differences++
		}
	}
	return differences, nil
}
