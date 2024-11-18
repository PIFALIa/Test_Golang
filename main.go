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
	// var a string
	// fmt.Scanf("%s\n", &a)

	fmt.Print("Ваш ответ:\t| ")
	fmt.Println(hel.MarksDel(text))
}
