package util

import "unicode"

func ScoreString(str string) int {
	scoreMap := make(map[rune]int)
	scoreMap['E'] = 13
	scoreMap['T'] = 12
	scoreMap['A'] = 11
	scoreMap['O'] = 10
	scoreMap['I'] = 9
	scoreMap['N'] = 8
	scoreMap[' '] = 7
	scoreMap['S'] = 6
	scoreMap['H'] = 5
	scoreMap['R'] = 4
	scoreMap['D'] = 3
	scoreMap['L'] = 2
	scoreMap['U'] = 1

	score := 0

	for _, r := range str {
		score += scoreMap[unicode.ToUpper(r)]
	}

	return score
}
