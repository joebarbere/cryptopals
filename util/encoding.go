package util

import (
	"encoding/base64"
	"encoding/hex"
)

func HexStringToByteSlice(hexString string) []byte {
	byteSlice, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}
	return byteSlice
}

func ByteSliceToBase64(byteSlice []byte) string {
	return base64.StdEncoding.EncodeToString(byteSlice)
}
