package main

import "fmt"

func main() {

	salario1, imposto1 := imposto(25000)
	salario2, imposto2 := imposto(50000)
	salario3, imposto3 := imposto(100000)
	salario4, imposto4 := imposto(150000)

	fmt.Println("O salario 1 é", salario1, "O imposto sobre é:", imposto1)
	fmt.Println("O salario 2 é", salario2, "O imposto sobre é:", imposto2)
	fmt.Println("O salario 3 é", salario3, "O imposto sobre é:", imposto3)
	fmt.Println("O salario 4 é", salario4, "O imposto sobre é:", imposto4)

}

func imposto(num float64) (float64, float64) {
	if num >= 50000 {
		return num * 0.83, num * 0.17
	} else if num >= 150000 {
		return num * 0.73, num * 0.27
	} else {
		return num, 0
	}

}
