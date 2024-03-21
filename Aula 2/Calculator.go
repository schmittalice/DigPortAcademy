package main

import "fmt"

func main() {
	var num1, num2, resposta float64

	var 1- soma float64 = num1 + num2
	var 2- subtracao float64 = num1 - num2
	var 3- multiplicacao float64 = num1 * num2
	var 4- divisao float64 = num1 / num2

	fmt.Println("Bem-vindo(a) a calculadora!")

	fmt.Println("Digite o primeiro numero: ")
	fmt.Scanln(&num2)

	fmt.Println("Digite o valor do segundo numero: ")
	fmt.Scanln(&num2)

	fmt.Println("Digite a operacao")

	fmt.Scanln(&resposta)
	switch resposta {
	case 1:
		fmt.Println("1- soma")
	case 2:
		fmt.Println("2- subtracao")
	case 3:
		fmt.Println("3- multiplicacao")
	case 4:
		fmt.Println("4- divisao")
	default:
		fmt.Println("Numero invalido!")
	}

}
