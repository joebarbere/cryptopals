package setone

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"sort"
	"strings"
	"testing"

	"../util"
)

func TestChallengeOne(t *testing.T) {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	correctValue := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	t.Log(hexString)
	byteSlice := util.HexStringToByteSlice(hexString)
	t.Log(string(byteSlice))
	base64String := util.ByteSliceToBase64(byteSlice)
	t.Log(base64String)

	if base64String != correctValue {
		t.Error(base64String, " != ", correctValue)
	}
}

func TestChallengeTwo(t *testing.T) {
	hexString1 := "1c0111001f010100061a024b53535009181c"
	hexString2 := "686974207468652062756c6c277320657965"
	hexEncodedString := util.XorHexStrings(hexString1, hexString2)
	t.Log(hexString1)
	t.Log(string(util.HexStringToByteSlice(hexString1)))
	t.Log(hexString2)
	t.Log(string(util.HexStringToByteSlice(hexString2)))
	t.Log(hexEncodedString)
}

func TestChallengeThree(t *testing.T) {
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	cryptopalsResult := util.BruteXorOneByteString(hexString)
	t.Log("key: ", cryptopalsResult.Key, " ", string(cryptopalsResult.Key))
	t.Log("score: ", cryptopalsResult.Score)
	t.Log("encrypted: ", cryptopalsResult.EncryptedString)
	t.Log("decrypted: ", cryptopalsResult.DecryptedString)
}

func TestChallengeFour(t *testing.T) {
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
	topScore := cryptopalsResultSlice[0].Score
	decryptedString := cryptopalsResultSlice[0].DecryptedString
	encryptedString := cryptopalsResultSlice[0].EncryptedString
	t.Log("key: ", key, " ", string(key))
	t.Log("score: ", topScore)
	t.Log("encrypted: ", encryptedString)
	t.Log("decrypted: ", decryptedString)
}

func TestChallengeFive(t *testing.T) {
	key := "ICE"
	message := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	hexEncodedCiphertext := util.RepeatingKeyXorString(key, message)
	t.Log(hexEncodedCiphertext)
}

func TestChallengeSix(t *testing.T) {
	str1 := "this is a test"
	str2 := "wokka wokka!!!"
	bytes1 := []byte(str1)
	bytes2 := []byte(str2)
	distance := util.HammingDistanceBytes(bytes1, bytes2)
	t.Log("Distance: ", distance)

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
		//t.Log(line)
	}
	cipherTextBase64Encoded := buffer.String()
	//t.Log(cipherTextBase64Encoded)
	cipherTextBytes, _ := base64.StdEncoding.DecodeString(cipherTextBase64Encoded)

	//test := "ABCDEABCDEABCDEABCDEABCDEABCDEABCDEABCDEABCDEABCDEABCDEABCDEABCDEABCDEABCDEABCDE"

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
		//t.Log("keysize: ", keysize, "distance: ", distance, "normalized distance: ", normalizedDistance)
		keySizeHammingDistanceSlice = append(keySizeHammingDistanceSlice, util.KeySizeHammingDistance{KeySize: keysize, HammingDistance: normalizedDistance})
	}

	sort.Slice(keySizeHammingDistanceSlice, func(i, j int) bool {
		return keySizeHammingDistanceSlice[i].HammingDistance < keySizeHammingDistanceSlice[j].HammingDistance
	})

	//testBytes := []byte{'1', '2', '3', '4', '5', '6', '7', '8'}
	//testLength := 3
	//testTransposedBytes := util.Transpose(testBytes, testLength)
	//for t := range testTransposedBytes {
	//	t.Log(testTransposedBytes[t])
	//}

	/*
		for ks_i, ks := range keySizeHammingDistanceSlice {
			theKeySize := ks.KeySize
			transposedBytes := util.Transpose(cipherTextBytes, theKeySize)
			var theKey []byte
			for i := 0; i < theKeySize; i++ {
				cryptopalsResult := util.BruteXorOneByte(transposedBytes[i])
				keyByte := cryptopalsResult.Key[0]
				theKey = append(theKey, byte(keyByte))
				//t.Log(i, ": ", string(keyByte))
			}

			//foo := hex.EncodeToString(cipherTextBytes)
			t.Log(ks_i, ": ", string(theKey))
		}
	*/

	theKeySize := keySizeHammingDistanceSlice[0].KeySize
	transposedBytes := util.Transpose(cipherTextBytes, theKeySize)
	var theKey []byte
	for i := 0; i < theKeySize; i++ {
		cryptopalsResult := util.BruteXorOneByte(transposedBytes[i])
		keyByte := cryptopalsResult.Key[0]
		theKey = append(theKey, byte(keyByte))
	}
	t.Log(string(theKey))
	plainText := util.RepeatingKeyXor(string(theKey), string(cipherTextBytes))
	t.Log(string(plainText))
}
