package main

import "fmt"

func main() {

	a := NewHandlerEmployee()
	b := NewHandlerEmployee()

	fmt.Println(a, b)
	// a e b são instâncias diferentes, cada uma no seu espaço de memória.

	p := NewHandlerEmployee2()
	q := NewHandlerEmployee2()

	fmt.Println(p, q)

	// p e q são dois ponteiros diferentes que apontam para structs diferentes.
}

type handlerEmployee struct {
	dados map[string]string
}

// Constructor de valor
func NewHandlerEmployee() handlerEmployee {
	return handlerEmployee{}
}

// Constructor de ponteiros
func NewHandlerEmployee2() *handlerEmployee {
	return &handlerEmployee{}
}
