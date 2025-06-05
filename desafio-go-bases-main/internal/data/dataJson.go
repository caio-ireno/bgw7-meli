package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/desafio-go-bases/internal/entity"
)

func readCsv(path string) ([][]string, error) {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	//Não fecha só em caso de erro, fecha sempre ao final da função
	//É uma prática recomendada ao trabalhar com arquivos em Go
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	return records, nil

}

func convertCsv(records [][]string) ([]entity.Passageiros, error) {

	var passageiros []entity.Passageiros

	for _, l := range records {
		if len(l) < 6 {
			continue
		}

		//convert string -> int
		id, err := strconv.Atoi(l[0])

		if err != nil {
			fmt.Printf("Erro ao converter id: %v\n", err)
			continue
		}

		//convert string -> int
		preco, err := strconv.ParseFloat(l[5], 64)

		if err != nil {
			fmt.Printf("Erro ao converter preco: %v\n", err)
			continue
		}

		passageiros = append(passageiros, entity.Passageiros{
			Id:      id,
			Nome:    l[1],
			Email:   l[2],
			Pais:    l[3],
			Horario: l[4],
			Preco:   preco,
		})
	}

	return passageiros, nil

}

func Data() ([]entity.Passageiros, error) {

	records, err := readCsv("tickets.csv")

	if err != nil {
		return nil, err
	}

	return convertCsv(records)

}
