package setone

import (
	"fmt"

	"../util"
)

// DoChallengeThree completes the Cryptopals Set 1 Challenge 3
func DoChallengeThree() {
	fmt.Println("\nSet 1. Challenge 3.")
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	cryptopalsResult := util.BruteXorOneByte(hexString)
	fmt.Println("key: ", cryptopalsResult.Key, " ", string(cryptopalsResult.Key))
	fmt.Println("score: ", cryptopalsResult.Score)
	fmt.Println("encrypted: ", cryptopalsResult.EncryptedString)
	fmt.Println("decrypted: ", cryptopalsResult.DecryptedString)
}
