package util

import (
	"encoding/hex"
	"sort"
)

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

func BruteXorOneByteString(hexString string) CryptopalsResult {
	hexBytes, _ := hex.DecodeString(hexString)
	cryptopalsResult := BruteXorOneByte(hexBytes)
	cryptopalsResult.EncryptedString = hexString
	return cryptopalsResult
}

func BruteXorOneByte(hexBytes []byte) CryptopalsResult {
	var cryptopalsResultSlice []CryptopalsResult

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

func RepeatingKeyXorString(key, message string) string {
	encryptedBytes := RepeatingKeyXor(key, message)
	hexEncodedCiphertext := hex.EncodeToString(encryptedBytes)
	return hexEncodedCiphertext
}

func RepeatingKeyXor(key, message string) []byte {
	fNextByteOfRepeatingKey := NextByteOfRepeatingKey([]byte(key))
	encryptedBytes := make([]byte, len(message))
	for i, r := range message {
		nextByte := fNextByteOfRepeatingKey()
		encryptedBytes[i] = byte(r) ^ nextByte
	}
	return encryptedBytes
}
