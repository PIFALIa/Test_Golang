package halpersPakage

import (
	"strconv"
	"strings"
)

func MarksDel(str string) string {

	str = strings.ReplaceAll(str, "\"", "")
	counts := make(map[string]int)
	slStr := strings.Fields(str)

	result := "\""
	runes := []rune(str)
	resultRues := ""

	minus := false
	mno := false
	del := false
	pluse := false

	for _, s := range slStr {

		switch s {
		case "-":

			minus = true
		case "*":
			s = ""
			mno = true
		case "/":
			s = ""
			del = true
		case "+":
			s = ""
			pluse = true
		}
		counts[s]++
	}

	if del {

		runesHelp := []rune(str)
		num, _ := strconv.Atoi(string(runesHelp[len(runes)-1]))
		for i, s := range runesHelp {
			if i < num {
				result += string(s)
			}
		}
	}

	if mno {
		num, _ := strconv.Atoi(slStr[len(slStr)-1])

		if num > 10 {
			result = "\"программа принимает значение до 10 включительно\""
			return result
		}
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
		if len(runesHelp) < 40 {

			result += resultRues
		}
	}

	for i, s := range counts {

		if minus && counts[i] < 2 {

			if i == "-" {
				break
			}
			if slStr[s] != slStr[s+2] {
				result += " "
			}
			result += i
		}

		if pluse {
			result += i
		}
	}

	result += "\""
	//result = strings.ReplaceAll(result, " ", "")

	return result
}
