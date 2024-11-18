package halpersPakage

import (
	"strconv"
	"strings"
)

func MarksDel(str string) string {

	str = strings.ReplaceAll(str, "\"", "")
	counts := make(map[string]int)
	words := strings.Fields(str)
	var result string
	var g int
	var mno int
	var dot bool

	for i, word := range words {

		if g == 1 {
			g = 2
			mno, _ := strconv.Atoi(words[i])
			if mno > 10 {
				mno = 10
				dot = true
			}
			result = "\""
			for j := 0; j < mno; j++ {
				result += words[0]
			}
			break
		}
		if word == "*" {
			g = 1
		}
		if word == "/" {
			runes := []rune(words[0])
			mno, _ = strconv.Atoi(words[i+1])
			for k := 0; k < mno; k++ {
				result += string(runes[k])
				counts[word]++
			}
			break
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

		if i == len(counts) || i == 1 && g == 0 {
			result += "\""
		}
		if i == 1 && counts["/"] < 1 && g == 0 {

			result += line
		}
	}
	result = strings.ReplaceAll(result, " ", "")
	result = strings.ReplaceAll(result, "+", "")
	result = strings.ReplaceAll(result, "*", "")
	if dot {
		result += "..."
	}
	result += "\""
	return result
}
