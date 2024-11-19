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
	// var num int
	var dot bool

	for i, word := range words {

		if g == 1 {
			g++
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

	result += "\""
	for line, i := range counts {

		// if i == len(counts) || i == 1 && g == 0 && num == 0 {

		// 	num++
		// }

		if i == 1 && counts["/"] < 1 && g == 0 {

			result += line
		}
	}
	result = strings.ReplaceAll(result, " ", "")
	result = strings.ReplaceAll(result, "+", "")
	result = strings.ReplaceAll(result, "*", "")

	run := []rune(result)
	if len(run) >= 40 {
		bo := true
		if bo {
			result = ""
			for inum, _ := range run {
				if inum <= 40 {
					result += string(run[inum])
				}
			}
			bo = false
		}
		result += "..."
		result += "\""
		return result
	}
	if dot {
		result += "..."
	}
	result += "\""
	return result
}
