package setone

import (
	"fmt"

	"../util"
)

// DoChallengeOne completes the Cryptopals Set 1 Challenge 1
func DoChallengeOne() {
	fmt.Println("\nSet 1. Challenge 1.")
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println(hexString)
	fmt.Println(util.HexToBase64(hexString))
}
