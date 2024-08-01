package exercicios

import (
	"fmt"

	"github.com/Nhanderu/brdoc"
)

func Cpf() {
	//ler do teclado um texto
	var texto string
	//verificar se texto é um cpf
	s := brdoc.IsCPF(texto)

	//caso sim
	fmt.Printf("%s é um CPF", texto)
	//caso não
	fmt.Printf("%s não é um CPF", texto)

}
