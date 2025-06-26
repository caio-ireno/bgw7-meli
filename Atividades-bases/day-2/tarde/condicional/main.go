package main

import (
	"fmt"
)

func main() {
	fmt.Println("Vamos começar...")

	fmt.Println(validaIdade(15))
	fmt.Println(validaIdade(60))
	fmt.Println(validaIdade(65))

	fmt.Println(validaCase(1))
	fmt.Println(validaCase(2))
	fmt.Println(validaCase(3))

}

func validaIdade(idade int) string {
	if idade < 18 {
		return "menor de idade"
	} else if idade >= 18 && idade <= 60 {
		return "Pessoa Adulta"
	} else {
		return "Pessoa Idosa"
	}
}

func validaCase(teste int) string {

	switch teste {

	case 1:
		return "Numero 1"

	case 2:
		return "Numero 2"

	case 3:
		fallthrough
	default:
		return "Outro numero"
	}

}
