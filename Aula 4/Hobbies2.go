package main

import "fmt"

func main() {
	contatos := map[string]string{
		"João":    "Futebol",
		"Maria":   "Pintura",
		"Pedro":   "Leitura",
		"Lucas":   "Música",
		"Carla":   "Dança",
		"Mariana": "Culinária",
	}

	for nome, hobby := range contatos {
		fmt.Println("O hobby de", nome, "é", hobby)
	}
}
