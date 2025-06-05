package data

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/desafio-go-bases/internal/entity"
)

func readCsv(path string) ([][]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("erro ao abrir arquivo: %w", err)

	}

	//Não fecha só em caso de erro, fecha sempre ao final da função
	//É uma prática recomendada ao trabalhar com arquivos em Go
	defer file.Close()

	// Criar um novo leitor para CSV
	// Sabe como interpretar o formato CSV
	reader := csv.NewReader(file)

	// le todas as linhas do CSV
	// retorna uma matriz [][]string
	// fatia de strings, onde cada elementos é uma linha e cada campo um valor

	records, err := reader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo CSV: %w", err)

	}

	return records, nil

}

func convertCsv(records [][]string) ([]entity.Passageiros, error) {

	if len(records) == 0 {
		return nil, errors.New("arquivo está vazio")
	}

	var passageiros []entity.Passageiros

	for _, l := range records {
		if len(l) < 6 {
			continue // pula para a próxima linha do CSV
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
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Arquivo não encontrado!")
		} else {
			fmt.Println("Outro erro:", err)
		}
		return nil, err
	}

	return convertCsv(records)

}
