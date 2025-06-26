package main

import (
	"fmt"
	"strings"
)

func main() {
	palavra("caio")
	palavraWithRuna("Eduardo")

}

func palavra(palavra string) {
	r := strings.Split(palavra, "")
	var count int
	for _, r := range r {
		count += 1
		fmt.Println("Letra: ", r)
	}

	fmt.Println("Tamanho da palavra:", count)
}

func palavraWithRuna(palavra string) {

	runas := []rune(palavra)
	fmt.Println("Tamanho da palavra", len(runas))

	for i := range runas {
		fmt.Printf("letras:%c\n", runas[i])
	}
}
