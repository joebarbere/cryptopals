package util

import (
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"math/bits"
	"net/http"
	"sort"
	"unicode"
)

const RUNES = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789"

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
	var cryptopalsResultSlice []CryptopalsResult

	hexBytes, _ := hex.DecodeString(hexString)

	for _, r := range RUNES {
		cryptopalsResult := CryptopalsResult{}
		key := []byte(string(r))[0]
		decryptedString := string(XorOneByte(key, hexBytes))
		cryptopalsResult.Key = []rune{r}
		cryptopalsResult.DecryptedString = decryptedString
		cryptopalsResult.Score = ScoreString(decryptedString)
		cryptopalsResultSlice = append(cryptopalsResultSlice, cryptopalsResult)
	}

	sort.Slice(cryptopalsResultSlice, func(i, j int) bool {
		return cryptopalsResultSlice[i].Score > cryptopalsResultSlice[j].Score
	})

	cryptopalsResultSlice[0].EncryptedString = hexString

	return cryptopalsResultSlice[0]
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

func GetCryptopalsData(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	responseString := string(responseData)
	return responseString
}
