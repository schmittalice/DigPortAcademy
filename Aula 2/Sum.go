package main

import "fmt"

func main() {
	var Num1 int = 10
	var Num2 int

	fmt.Println("Diga um número")
	fmt.Scanf("%d", &Num2)
	soma := Num1 + Num2
	fmt.Println("Resultado:", soma)
}
