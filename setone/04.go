package setone

import (
	"bufio"
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

	var cryptopalsResultSlice []util.CryptopalsResult
	scanner := bufio.NewScanner(strings.NewReader(responseString))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		cryptopalsResult := util.BruteXorOneByte(line)

		key := cryptopalsResult.Key
		score := cryptopalsResult.Score
		decryptedString := cryptopalsResult.DecryptedString

		cryptopalsResultSlice = append(cryptopalsResultSlice, util.CryptopalsResult{Key: key, Score: score, EncryptedString: line, DecryptedString: decryptedString})
	}

	sort.Slice(cryptopalsResultSlice, func(i, j int) bool {
		return cryptopalsResultSlice[i].Score > cryptopalsResultSlice[j].Score
	})

	key := cryptopalsResultSlice[0].Key
	topScore := cryptopalsResultSlice[0].Score
	decryptedString := cryptopalsResultSlice[0].DecryptedString
	encryptedString := cryptopalsResultSlice[0].EncryptedString
	fmt.Println("key: ", key, " ", string(key))
	fmt.Println("score: ", topScore)
	fmt.Println("encrypted: ", encryptedString)
	fmt.Println("decrypted: ", decryptedString)
}
