package pessoa

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func ListaPessoa() []Pessoa {
	pessoas := []Pessoa{
		{Nome: "Cate", Idade: 32},
		{Nome: "Carol", Idade: 31},
	}
	return pessoas
}
