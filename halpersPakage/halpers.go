package halpersPakage

import (
	"strconv"
	"strings"
)

func MarksDel(str string) string {

	str = strings.ReplaceAll(str, "\"", "")
	counts := make(map[string]int)
	slStr := strings.Fields(str)

	var result string
	runes := []rune(str)
	resultRues := ""

	minus := false
	mno := false
	// countRunes := 0
	// del := false

	for _, s := range slStr {

		switch s {
		case "-":
			s = ""
			minus = true
		case "*":
			s = ""
			mno = true
			// case "/":
			// 	del = true
			// }
		}
		counts[s]++
	}

	if mno {
		num, _ := strconv.Atoi(string(runes[len(runes)-1]))
		for range num {
			resultRues += slStr[0]
		}
		runesHelp := []rune(resultRues)

		if len(runesHelp) > 40 {
			for i := range runesHelp {
				result += string(runesHelp[i])
				if i > 38 {
					result += "..."
					break
				}
			}

		}
	}

	for i := range counts {

		if minus && counts[i] < 2 {

			result += i
		}
	}

	result = strings.ReplaceAll(result, " ", "")

	return result
}
