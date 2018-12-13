package cryptopals

import (
	"encoding/base64"
	"strings"
	"testing"

	"./util"
)

func TestSet01Challenge07(t *testing.T) {
	key := "YELLOW SUBMARINE"
	url := "https://cryptopals.com/static/challenge-data/7.txt"
	responseString := util.GetCryptopalsData(url)
	cipherTextBytes, _ := base64.StdEncoding.DecodeString(responseString)
	plainTextBytes := util.DecryptAes128Ecb(cipherTextBytes, []byte(key))
	correctValueStartsWith := "I'm back and I'm ringin' the bell"

	if !strings.HasPrefix(string(plainTextBytes), correctValueStartsWith) {
		t.Error(string(plainTextBytes)[0:25], " does not start with ", correctValueStartsWith)
	}
}
