package main

import (
	"fmt"
	hel "main/halpersPakage"
)

func main() {
	var a string
	fmt.Println("Введите букавы")
	fmt.Scanf("%s\n", &a)

	fmt.Print("Ваш ответ: ")
	fmt.Println(hel.MarksDel(a))
}
