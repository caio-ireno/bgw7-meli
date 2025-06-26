package main

import "fmt"

func main() {
	pessoa1 := salario("A", 1)
	fmt.Println(pessoa1)

}

func salario(categoria string, horas float64) float64 {

	switch categoria {
	case "A":
		salario := horas * 3000
		return salario + salario*0.5
	case "B":
		salario := horas * 1500
		return salario + salario*0.2
	case "C":
		return horas * 1000
	default:
		return 0
	}

}
