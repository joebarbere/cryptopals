package setone

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"../util"
)

// DoChallengeFour completes the Cryptopals Set 1 Challenge 4
func DoChallengeFour() {
	fmt.Println("\nSet 1. Challenge 4.")

	const runes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789"
	encryptedStringsFileURL := "https://cryptopals.com/static/challenge-data/4.txt"

	response, err := http.Get(encryptedStringsFileURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	responseString := string(responseData)

	var challengeFourStructSlice []util.ChallengeFourStruct
	scanner := bufio.NewScanner(strings.NewReader(responseString))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		hexBytes, _ := hex.DecodeString(line)
		bruteForceMap := make(map[rune]string)
		scoresMap := make(map[rune]int)
		for _, r := range runes {
			runeByte := []byte(string(r))[0]
			decrypted := string(util.XorOneRune(runeByte, hexBytes))
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
		decrypted := bruteForceMap[topScoredRune]

		challengeFourStructSlice = append(challengeFourStructSlice, util.ChallengeFourStruct{Rune: topScoredRune, Score: topScore, EncryptedString: line, DecryptedString: decrypted})
	}

	sort.Slice(challengeFourStructSlice, func(i, j int) bool {
		return challengeFourStructSlice[i].Score > challengeFourStructSlice[j].Score
	})

	pTopScoredRune := challengeFourStructSlice[0].Rune
	pTopScore := challengeFourStructSlice[0].Score
	pTopDecrypted := challengeFourStructSlice[0].DecryptedString
	pTopEncryptedString := challengeFourStructSlice[0].EncryptedString
	fmt.Println("TOP encrypted: ", pTopEncryptedString)
	fmt.Println("TOP key: ", string(pTopScoredRune))
	fmt.Println("TOP decrypted: ", pTopDecrypted)
	fmt.Println("TOP score: ", pTopScore)

}
