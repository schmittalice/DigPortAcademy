package main

import (
	"fmt"
)

type employee struct {
	name      string
	idade     int
	profissao string
	salario   int
}

func main() {

	a := employee{"Livia", 27, "desenvolvedora", 8000}
	b := employee{"Maria", 33, "manager", 10000}

	fmt.Printf("%+v/n", a)
	fmt.Scanf("%s", &b.idade)

}
