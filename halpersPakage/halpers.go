package halpersPakage

import (
	"strings"
)

func MarksDel(str string) string {

	str = strings.ReplaceAll(str, "\"", "")
	// counts := make(map[string]int)
	// words := strings.Fields(str)
	return str
}
