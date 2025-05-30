package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	validaIdade(employees)
	employees["Frederico"] = 25
	validaIdade(employees)
	delete(employees, "Pedro")
	validaIdade(employees)

}

func validaIdade(m map[string]int) {
	var count int
	for _, valor := range m {
		if valor > 21 {
			count += 1
		}
	}
	fmt.Println("Temos:", count, "Funcionarios maiores de 21 anos")
}
