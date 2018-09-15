package setone

import (
	"encoding/hex"
	"fmt"

	"../util"
)

func DoChallengeFive() {
	fmt.Println("\nSet 1. Challenge 5.")
	runes := "ICE"
	message := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	f := util.NextRune([]byte(runes))
	encrypted := make([]byte, len(message))
	for i, r := range message {
		x := f()
		//fmt.Println(string(r), " ^ ", string(x))
		encrypted[i] = byte(r) ^ x
	}
	fmt.Println(hex.EncodeToString(encrypted))
}
