package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numero := rand.Intn(15)

	mes := validaMes(numero)

	fmt.Println(numero, mes)

}

func validaMes(numero int) string {

	switch {
	case numero == 1:
		return "Janeiro"
	case numero == 2:
		return "Fevereiro"
	case numero == 3:
		return "Março"
	case numero == 4:
		return "Abril"
	case numero == 5:
		return "Maio"
	case numero == 6:
		return "Junho"
	case numero == 7:
		return "Julho"
	case numero == 7:
		return "Julho"
	case numero == 8:
		return "Agosto"
	case numero == 9:
		return "Setembro"
	case numero == 10:
		return "Outubro"
	case numero == 11:
		return "Novembro"
	case numero == 12:
		return "Dezembro"
	default:
		return "Digite um numero valido"

	}

}
