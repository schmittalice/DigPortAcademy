package main

import "fmt"

func main() {
	var Numero int

	fmt.Println("Diga um número")
	fmt.Scanf("%d", &Numero)

	if Numero > 0 {
		fmt.Println("Positivo")
	} else if Numero < 0 {
		fmt.Println("Negativo")
	} else {
		fmt.Println("Zero")
	}

}
