package setone

import (
	"fmt"

	"../util"
)

// DoChallengeTwo completes the Cryptopals Set 1 Challenge 2
func DoChallengeTwo() {
	fmt.Println("\nSet 1. Challenge 2.")
	hexString1 := "1c0111001f010100061a024b53535009181c"
	hexString2 := "686974207468652062756c6c277320657965"
	hexEncodedString := util.XorHexStrings(hexString1, hexString2)
	fmt.Println(hexString1)
	fmt.Println(hexString2)
	fmt.Println(hexEncodedString)
}
