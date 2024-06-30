package main

import "fmt"

func main() {
minhas Despesas :=[]string{"aluguel", "gas", "mercado", "celular"}
//for i, v := range minhasDespesas {
for _, v := range minhasDespesas {
	fmt.Println(v)
}}

fmt.Println("Qual item vc deseja saber se está na lista de despesas?")
var itemVerificar string
fmt.Scan(&itemVerificar)
fmt.Println("Verificando se há", itemVerificar, "na sua lista...")
for _, v := range minhasDespesas {
	if v == itemVerificar {
		fmt.Pintln("sim")
		return
	}
}

fmt.Println("Não")

fmt.Printf("Minha lista de despesas tem %d itens", len(minhasDespesas))

}