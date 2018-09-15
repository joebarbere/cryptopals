package setone

import (
	"fmt"

	"../util"
)

// DoChallengeFive completes the Cryptopals Set 1 Challenge 5
func DoChallengeFive() {
	fmt.Println("\nSet 1. Challenge 5.")
	key := "ICE"
	message := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	hexEncodedCiphertext := util.RepeatingKeyXorString(key, message)
	fmt.Println(hexEncodedCiphertext)
}
