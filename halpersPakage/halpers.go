package halpersPakage

import (
	"strconv"
	"strings"
)

func MarksDel(str string) string {

	runesE := []rune(str)
	if string(runesE[0]) != "\"" {
		panic("Первый аргумент болжен быть строкой!")
	}
	str = strings.ReplaceAll(str, "\"", "")
	counts := make(map[string]int)
	slStr := strings.Fields(str)

	result := "\""
	//runes := []rune(str)
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
		num, err := strconv.Atoi(slStr[len(slStr)-1])
		if err != nil {
			panic("Ошибка преобразования:" + slStr[len(slStr)-1] + "не является целым числом ") // Паника при ошибке
		}
		for i, s := range runesHelp {
			if i < num {
				result += string(s)
			}
		}
	}

	if mno {
		num, err := strconv.Atoi(slStr[len(slStr)-1])
		if err != nil {
			panic("Ошибка преобразования: не является целым числом " + slStr[len(slStr)-1]) // Паника при ошибке
		}
		if num < 1 {
			panic("\"программа принимает значение от 1 до 10\"")

		}

		if num > 10 {
			panic("\"программа принимает значение до 10 включительно\"")

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

	for i := range counts {

		if minus && counts[i] < 2 {

			if i == "-" {
				break
			}
			if slStr[0] == slStr[len(slStr)-1] {
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
