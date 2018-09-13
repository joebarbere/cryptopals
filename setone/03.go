package setone

import (
	"encoding/hex"
	"fmt"
	"sort"

	"../util"
)

// DoChallengeThree completes the Cryptopals Set 1 Challenge 3
func DoChallengeThree() {
	fmt.Println("\nSet 1. Challenge 3.")
	hexString3 := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	hexBytes3, _ := hex.DecodeString(hexString3)
	const runes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789"
	bruteForceMap := make(map[rune]string)
	scoresMap := make(map[rune]int)
	for _, r := range runes {
		runeByte := []byte(string(r))[0]
		decrypted := string(util.XorOneRune(runeByte, hexBytes3))
		bruteForceMap[r] = decrypted
		scoresMap[r] = util.ScoreString(decrypted)
	}

	var scoresSlice []util.ScoreKeyValue
	for k, v := range scoresMap {
		scoresSlice = append(scoresSlice, util.ScoreKeyValue{Key: k, Value: v})
	}

	sort.Slice(scoresSlice, func(i, j int) bool {
		return scoresSlice[i].Value > scoresSlice[j].Value
	})

	topScoredRune := scoresSlice[0].Key
	topScore := scoresSlice[0].Value
	fmt.Println("key: ", string(topScoredRune))
	fmt.Println("decrypted: ", bruteForceMap[topScoredRune])
	fmt.Println("score: ", topScore)
}
