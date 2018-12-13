package cryptopals

import "testing"
import "./util"

func TestSet01Challenge03(t *testing.T) {
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	correctValue := "Cooking MC's like a pound of bacon"
	cryptopalsResult := util.BruteXorOneByteString(hexString)

	if cryptopalsResult.DecryptedString != correctValue {
		t.Error(cryptopalsResult.DecryptedString, " != ", correctValue)
	}
}
