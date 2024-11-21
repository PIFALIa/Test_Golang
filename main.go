package main

import (
	"bufio"
	"fmt"
	hel "main/halpersPakage"
	"os"
	"strings"
)

func main() {

	fmt.Print("Введите букавы:\t| ")
	riders := bufio.NewReader(os.Stdin)
	text, _ := riders.ReadString('\n')

	text = strings.TrimSpace(text)

	fmt.Print("Ваш ответ:\t| ")
	fmt.Println(hel.MarksDel(text))

}
