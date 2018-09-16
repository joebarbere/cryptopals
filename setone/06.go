package setone

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(strings.NewReader(responseString))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
