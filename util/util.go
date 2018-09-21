package util

import (
	"io/ioutil"
	"math"
	"net/http"
)

const RUNES = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789~!@#$%^&*()_+`-={}|[]\\,./<>?:\";' "

type CryptopalsResult struct {
	Key             []rune
	Score           int
	EncryptedString string
	DecryptedString string
}

type KeySizeHammingDistance struct {
	KeySize         int
	HammingDistance float64
}

func NextByteOfRepeatingKey(key []byte) func() byte {
	var pos int
	return func() byte {
		endPos := len(key) - 1
		curPos := pos
		if curPos == endPos {
			pos = 0
		} else {
			pos++
		}
		return key[curPos]
	}
}

func GetCryptopalsData(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	responseString := string(responseData)
	return responseString
}

func Transpose(bytes []byte, length int) [][]byte {
	dy := int(math.Ceil(float64(len(bytes)) / float64(length)))
	//fmt.Println("dy: ", dy)
	transposedBytes := make([][]byte, dy)
	for i := range transposedBytes {
		for j := 0; j < length && i*length+j < len(bytes); j++ {
			//fmt.Println(i, " ", j, " ", i*length+j)
			transposedBytes[i] = append(transposedBytes[i], bytes[i+j*length])
		}
	}
	return transposedBytes
}
