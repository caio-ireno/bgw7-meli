package main

import "testing"

func TestSum(t *testing.T) {
	total := soma(2, 3)
	if total != 5 {
		t.Errorf("Soma incorreta, esperado 5, obteve %d", total)
	}
}
