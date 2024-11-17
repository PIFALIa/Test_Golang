package halpersPakage

import (
	"strconv"
	"strings"
)

func MarksDel(str string) string {

	counts := make(map[string]int)
	words := strings.Fields(str)
	var result string
	var g int

	for i, word := range words {

		if word == "/" {
			runes := []rune(words[0])
			mno, _ := strconv.Atoi(words[i+1])
			for k := 0; k < mno; k++ {
				result += string(runes[k])
				counts[word]++
			}
			break
		}

		if g == 1 {
			g = 0
			mno, _ := strconv.Atoi(words[i])
			for j := 1; j < mno; j++ {
				result += words[0]
			}
			break
		}
		if word == "*" {
			g = 1
		}

		if words[i] == "-" && words[i-1] == words[i+1] {
			word = strings.ReplaceAll(word, "-", "")
			counts[words[i-1]]++
		}
		if word == "-" {
			word = strings.ReplaceAll(word, "-", "")
			words[i+1] = ""
		}
		counts[word]++
	}

	for line, i := range counts {

		if i == 1 && counts["/"] < 1 {
			result += line
			result += " "
		}
	}
	result = strings.ReplaceAll(result, " ", "")
	result = strings.ReplaceAll(result, "+", "")
	result = strings.ReplaceAll(result, "*", "")
	return result
}
