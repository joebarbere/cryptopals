package cryptopals

import (
	"bufio"
	"sort"
	"strings"
	"testing"

	"./util"
)

func TestSet01Challenge04(t *testing.T) {
	url := "https://cryptopals.com/static/challenge-data/4.txt"

	responseString := util.GetCryptopalsData(url)

	var cryptopalsResultSlice []util.CryptopalsResult
	scanner := bufio.NewScanner(strings.NewReader(responseString))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		cryptopalsResult := util.BruteXorOneByteString(line)

		key := cryptopalsResult.Key
		score := cryptopalsResult.Score
		decryptedString := cryptopalsResult.DecryptedString

		cryptopalsResultSlice = append(cryptopalsResultSlice, util.CryptopalsResult{Key: key, Score: score, EncryptedString: line, DecryptedString: decryptedString})
	}

	sort.Slice(cryptopalsResultSlice, func(i, j int) bool {
		return cryptopalsResultSlice[i].Score > cryptopalsResultSlice[j].Score
	})

	key := cryptopalsResultSlice[0].Key

	correctValue := "5"

	if string(key) != correctValue {
		t.Error(string(key), " != ", correctValue)
	}
}
