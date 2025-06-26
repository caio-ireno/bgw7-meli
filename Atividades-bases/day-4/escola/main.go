package main

import (
	"fmt"

	"github.com/caio-ireno/bgw7-meli/Atividades/day-4/escola/entities"
)

func main() {
	p1 := entities.Aluno{
		Name: "Caio",
		Professor: entities.Professor{
			Name:    "Nicolas",
			Materia: "Excel",
		},
	}

	fmt.Println(p1)

}
