package exercicios

import "fmt"

var helloWorldPrivada = "Hello World"
var HelloWorldPublica = "Hello World"

func Funcao1() string {
	return helloWorldPrivada
}

func funcao() {
	fmt.Println("Hello World")
}
