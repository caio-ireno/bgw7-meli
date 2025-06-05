package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/data"
	"github.com/bootcamp-go/desafio-go-bases/internal/entity"
)

func main() {
	jsonData, err := data.Data()

	if err != nil {
		panic(err)
	}

	// for _, p := range jsonData {
	// 	if p.Pais == "Indonesia" {
	// 		fmt.Println(p)
	// 	}
	// }

	service := entity.PassageirosService{List: jsonData}

	total, err := service.GetTotalPassagens()

	if err != nil {
		fmt.Println("Erro ao trazer a quantidade de passageiros")
	}

	local, err := service.CustoPorDestino("China")
	if err != nil {
		fmt.Println("Erro ao trazer o custo por destino")
	}
	fmt.Println(total)
	fmt.Println(local)

}
