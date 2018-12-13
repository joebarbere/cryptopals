package cryptopals

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"sort"
	"strings"
	"testing"

	"./util"
)

func TestSet01Challenge06(t *testing.T) {
	str1 := "this is a test"
	str2 := "wokka wokka!!!"
	bytes1 := []byte(str1)
	bytes2 := []byte(str2)
	distance := util.HammingDistanceBytes(bytes1, bytes2)

	correctDistance := 37
	if distance != correctDistance {
		t.Error(distance, " != ", correctDistance)
	}

	url := "https://cryptopals.com/static/challenge-data/6.txt"

	responseString := util.GetCryptopalsData(url)

	scanner := bufio.NewScanner(strings.NewReader(responseString))
	scanner.Split(bufio.ScanLines)
	var buffer bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()
		buffer.WriteString(line)
	}
	cipherTextBase64Encoded := buffer.String()
	cipherTextBytes, _ := base64.StdEncoding.DecodeString(cipherTextBase64Encoded)

	var keySizeHammingDistanceSlice []util.KeySizeHammingDistance
	for keysize := 2; keysize <= 40; keysize++ {
		bytes1 := cipherTextBytes[:keysize]
		bytes2 := cipherTextBytes[keysize : keysize*2]
		bytes3 := cipherTextBytes[keysize*2 : keysize*3]
		bytes4 := cipherTextBytes[keysize*3 : keysize*4]
		bytes5 := cipherTextBytes[keysize*4 : keysize*5]
		bytes6 := cipherTextBytes[keysize*5 : keysize*6]
		bytes7 := cipherTextBytes[keysize*6 : keysize*7]
		bytes8 := cipherTextBytes[keysize*7 : keysize*8]
		bytes9 := cipherTextBytes[keysize*8 : keysize*9]
		bytes10 := cipherTextBytes[keysize*9 : keysize*10]
		bytes11 := cipherTextBytes[keysize*10 : keysize*11]
		bytes12 := cipherTextBytes[keysize*11 : keysize*12]

		distance1 := util.HammingDistanceBytes(bytes1, bytes2)
		distance2 := util.HammingDistanceBytes(bytes3, bytes4)
		distance3 := util.HammingDistanceBytes(bytes5, bytes6)
		distance4 := util.HammingDistanceBytes(bytes7, bytes8)
		distance5 := util.HammingDistanceBytes(bytes9, bytes10)
		distance6 := util.HammingDistanceBytes(bytes11, bytes12)
		distance := float64(distance1+distance2+distance3+distance4+distance5+distance6) / float64(6)
		normalizedDistance := float64(distance) / float64(keysize)
		keySizeHammingDistanceSlice = append(keySizeHammingDistanceSlice, util.KeySizeHammingDistance{KeySize: keysize, HammingDistance: normalizedDistance})
	}

	sort.Slice(keySizeHammingDistanceSlice, func(i, j int) bool {
		return keySizeHammingDistanceSlice[i].HammingDistance < keySizeHammingDistanceSlice[j].HammingDistance
	})

	theKeySize := keySizeHammingDistanceSlice[0].KeySize
	transposedBytes := util.Transpose(cipherTextBytes, theKeySize)
	var theKey []byte
	for i := 0; i < theKeySize; i++ {
		cryptopalsResult := util.BruteXorOneByte(transposedBytes[i])
		keyByte := cryptopalsResult.Key[0]
		theKey = append(theKey, byte(keyByte))
	}

	correctValue := "Terminator X: Bring the noise"

	if string(theKey) != correctValue {
		t.Error(string(theKey), " != ", correctValue)
	}

	//plainText := util.RepeatingKeyXor(string(theKey), string(cipherTextBytes))
	//t.Log(string(plainText))
}
