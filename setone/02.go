package setone

import (
	"encoding/hex"
	"fmt"

	"../util"
)

// DoChallengeTwo completes the Cryptopals Set 1 Challenge 2
func DoChallengeTwo() {
	fmt.Println("\nSet 1. Challenge 2.")
	hexString1 := "1c0111001f010100061a024b53535009181c"
	hexString2 := "686974207468652062756c6c277320657965"
	hexBytes1, _ := hex.DecodeString(hexString1)
	hexBytes2, _ := hex.DecodeString(hexString2)
	destBytes := util.Xor(hexBytes1, hexBytes2)
	fmt.Println(hexString1)
	fmt.Println(hexString2)
	hexEncodedString := hex.EncodeToString(destBytes)
	fmt.Println(hexEncodedString)
}
