package main

import (
	"bufio"
	"fmt"
	hel "main/halpersPakage"
	"os"
)

func main() {

	fmt.Print("Введите букавы:\t| ")
	riders := bufio.NewReader(os.Stdin)
	text, _ := riders.ReadString('\n')

	fmt.Print("Ваш ответ:\t| ")
	fmt.Println(hel.MarksDel(text))
}
