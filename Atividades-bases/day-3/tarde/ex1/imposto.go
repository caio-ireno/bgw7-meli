package main

import "fmt"

func main() {
	fmt.Println(calcImposto(150000))

}

func calcImposto(salario float64) float64 {
	if salario < 50000 {
		return 0
	} else if salario >= 50000 && salario < 150000 {
		return salario * 0.17
	} else {
		return salario * 0.27
	}
}
