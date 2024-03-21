package main

import "fmt"

func main() {
	var Nome string

	fmt.Println("Qual seu nome?")
	fmt.Scanf("%s", &Nome)
	fmt.Println("Bem-vinda,", Nome)

}
