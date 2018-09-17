package setone

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"sort"
	"strings"

	"../util"
)

// DoChallengeSix completes the Cryptopals Set 1 Challenge 6
func DoChallengeSix() {
	fmt.Println("\nSet 1. Challenge 6.")

	str1 := "this is a test"
	str2 := "wokka wokka!!!"
	bytes1 := []byte(str1)
	bytes2 := []byte(str2)
	distance := util.HammingDistanceBytes(bytes1, bytes2)
	fmt.Println("Distance: ", distance)

	url := "https://cryptopals.com/static/challenge-data/6.txt"

	responseString := util.GetCryptopalsData(url)
	//fmt.Print(responseString)

	scanner := bufio.NewScanner(strings.NewReader(responseString))
	scanner.Split(bufio.ScanLines)
	var buffer bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()
		buffer.WriteString(line)
		//fmt.Print("... ")
		//fmt.Println(line)
	}
	cipherTextBase64Encoded := buffer.String()
	//fmt.Println(cipherTextBase64Encoded)
	cipherTextBytes, _ := base64.StdEncoding.DecodeString(cipherTextBase64Encoded)

	var keySizeHammingDistanceSlice []util.KeySizeHammingDistance
	for keysize := 2; keysize <= 40; keysize++ {
		bytes1 := cipherTextBytes[:keysize]
		bytes2 := cipherTextBytes[keysize : keysize*2]
		distance := util.HammingDistanceBytes(bytes1, bytes2)
		normalizedDistance := float64(distance) / float64(keysize)
		fmt.Println("keysize: ", keysize, "distance: ", distance, "normalized distance: ", normalizedDistance)
		keySizeHammingDistanceSlice = append(keySizeHammingDistanceSlice, util.KeySizeHammingDistance{KeySize: keysize, HammingDistance: normalizedDistance})
	}

	sort.Slice(keySizeHammingDistanceSlice, func(i, j int) bool {
		return keySizeHammingDistanceSlice[i].HammingDistance < keySizeHammingDistanceSlice[j].HammingDistance
	})

	fmt.Println("lowest distance: ", keySizeHammingDistanceSlice[0].HammingDistance, "for keysize: ", keySizeHammingDistanceSlice[0].KeySize)

}
