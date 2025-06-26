package main

import "testing"

func TestMedia(t *testing.T) {
	arr := []float64{4, 2, 3, 4}

	ExpectedResult := 3.25

	result := calculaMedia(arr...)

	if result != ExpectedResult {
		t.Errorf("O resultado deveria ser %f", ExpectedResult)
	}
}
