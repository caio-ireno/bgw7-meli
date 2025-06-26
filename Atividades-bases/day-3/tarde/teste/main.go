package main

import "fmt"

func main() {
	t, r := teste(2, 3)
	t2, r2 := teste2(2, 3)
	fmt.Println(t, r)
	fmt.Println(t2, r2)
}

func soma(a, b int) int {
	return a + b
}

func teste(x, z int) (a, b int) {
	return 10 - x, 1 - +z
}
func teste2(x, z int) (a, b int) {
	a = 10 - x
	b = 1 - z
	return // Aqui retorna "a, b" automaticamente
}
