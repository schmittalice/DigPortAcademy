package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/schmittalice/digportacademy/pessoa"
)

func Rotas() *mux.Router {
	rotas := mux.NewRouter()
	rotas.HandleFunc("/lista/pessoas", HandlePessoa).Methods("GET")
	return rotas
}

func HandlePessoa(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pessoa.ListaPessoa())
}
