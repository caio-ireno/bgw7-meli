package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

type Funcionario struct {
	Pessoa
	Salario float64
	Cargo   string
}

func (f Funcionario) pegaSalario() float64 {
	return f.Salario

}

func (p Pessoa) pegaIdade() int {
	return p.Idade
}

func main() {

	f := Funcionario{Pessoa: Pessoa{Nome: "João", Idade: 30}, Salario: 2000.0, Cargo: "Desenvolvedor"}

	fmt.Println(f.pegaSalario())
	fmt.Println(f.pegaIdade())
}
