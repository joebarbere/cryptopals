package cryptopals

import "testing"
import "./util"

func TestSet01Challenge01(t *testing.T) {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	correctValue := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	byteSlice := util.HexStringToByteSlice(hexString)
	base64String := util.ByteSliceToBase64(byteSlice)

	if base64String != correctValue {
		t.Error(base64String, " != ", correctValue)
	}
}
