package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {

	pessoas := make(map[string]Pessoa)

	// Adicionando elementos ao map
	pessoas["joao"] = Pessoa{Nome: "João", Idade: 30}
	pessoas["maria"] = Pessoa{Nome: "Maria", Idade: 25}

	// Acessando e imprimindo os dados
	for chave, pessoa := range pessoas {
		fmt.Printf("Chave: %s, Valor: %+v\n", chave, pessoa)
	}
}
