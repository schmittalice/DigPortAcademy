package exercicios

import "fmt"

var Contacts string

func main() {
	fmt.Println("Contatos existentes:")
	for _, contact := range Contacts {
		fmt.Println("Nome:", contact.Name)
		fmt.Println("Telefone:", contact.Phone)
		fmt.Println("CPF", contact.CPF)
	}

	//Adicionando novos contatos

}
