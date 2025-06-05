package main

import "fmt"

func main() {

	// Memoria (Endereço) <----- a <----- 10

	// a := 10
	// fmt.Println(a)
	// fmt.Println(&a)

	// var ponteiro *int = &a

	// fmt.Println(ponteiro)
	// fmt.Println(*ponteiro)

	// *ponteiro = 20
	// fmt.Println(a)
	// fmt.Println(*ponteiro)

	// b := &a

	// *b = 30
	// fmt.Println(b)
	// fmt.Println(a)

	variravel := 10

	abc(&variravel)

	fmt.Println(variravel)
}

// Muda o endereço de memória, ele recebe o endereço de memória da "variavel"
func abc(a *int) {
	*a = 200
}
