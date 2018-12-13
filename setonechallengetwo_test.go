package cryptopals

import (
	"testing"

	"./util"
)

func TestSet01Challenge02(t *testing.T) {
	hexString1 := "1c0111001f010100061a024b53535009181c"
	hexString2 := "686974207468652062756c6c277320657965"
	correctValue := "746865206b696420646f6e277420706c6179"
	hexEncodedString := util.XorHexStrings(hexString1, hexString2)

	if hexEncodedString != correctValue {
		t.Error(hexEncodedString, " != ", correctValue)
	}
}
