package data

import (
	"encoding/csv"
	"strings"
	"testing"
)

func TestReadCsv(t *testing.T) {

	csvData := `1, Caio, caio@gmail.com, Brasil, 10:10, 400.50`

	readFile := strings.NewReader(csvData)

	reader := csv.NewReader(readFile)

	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Erro ao ler o CSV: %v", err)
	}

	if len(records) != 1 {
		t.Errorf("Esperado 1 registo, obido %d", len(records))
	}

}

func TestConvertCsv(t *testing.T) {
	records := [][]string{
		{"1", "Caio", "caio@email.com", "Brasil", "10:00", "400.5"},
	}

	passageirosTeste, err := convertCsv(records)

	if err != nil {
		t.Fatalf("Erro na conversão: %v", err)
	}

	if len(passageirosTeste) != 1 {
		t.Errorf("Esperado 1 passageiro, obtido %d", len(passageirosTeste))
	}

	if passageirosTeste[0].Nome != "Caio" {
		t.Errorf("Esperado Caio, obtido %s", passageirosTeste[0].Nome)
	}
}
