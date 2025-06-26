package main

import "testing"

func TestImpostoMenor50000(t *testing.T) {
	num := 30000.0

	expectedValue := 0

	result := calcImposto((num))

	if result != float64(expectedValue) {
		t.Errorf("Valor esperado era %f", float64(expectedValue))
	}
}

func TestImpostoMaior50000(t *testing.T) {
	num := 50000.0

	expectedValue := 8500.0

	result := calcImposto((num))

	if result != float64(expectedValue) {
		t.Errorf("Valor esperado era %f", float64(expectedValue))
	}

}

func TestImpostoMaior150000(t *testing.T) {
	num := 150000.0

	expectedValue := 40500.0

	result := calcImposto((num))

	if result != float64(expectedValue) {
		t.Errorf("Valor esperado era %f", float64(expectedValue))
	}

}
