package main

import "fmt"

type Pessoa struct {
	Idade        int
	Nome         string
	Empregado    bool
	TempoEmprego int
	Salario      int
}

func main() {
	pessoa := make(map[string]Pessoa)

	pessoa["Caio"] = Pessoa{Idade: 22, Nome: "Caio", Empregado: true, TempoEmprego: 2, Salario: 105000}
	pessoa["Nathalia"] = Pessoa{Idade: 20, Nome: "Nathalia", Empregado: true, TempoEmprego: 2, Salario: 200000}
	pessoa["Felipe"] = Pessoa{Idade: 22, Nome: "Felipe", Empregado: false, TempoEmprego: 2, Salario: 5000000}
	pessoa["Michele"] = Pessoa{Idade: 48, Nome: "Michele", Empregado: true, TempoEmprego: 10, Salario: 10000}

	for chave, valor := range pessoa {

		if valor.Idade >= 22 && valor.Empregado && valor.TempoEmprego >= 1 {
			fmt.Printf("%s pode pedir um emprestimo\n", chave)
			if valor.Salario > 100000 {
				fmt.Printf("%s não será cobrado juros devido ao salario ser maior do que U$100.000\n", chave)
			} else {
				fmt.Printf("%s  será cobrado juros devido ao salario ser menor do que U$100.000\n", chave)
			}

		} else {
			fmt.Printf("%s näo pode pedir um emprestimo\n", chave)
		}

	}
}
