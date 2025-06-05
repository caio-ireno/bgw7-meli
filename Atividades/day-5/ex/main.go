package main

import (
	"errors"
	"fmt"
)

type Persona struct {
	Salary float64
	Name   string
}

func main() {

	p1 := Persona{Salary: 4000, Name: "Caio"}

	_, err := validSalary(float64(p1.Salary))
	if err != nil {
		fmt.Println(fmt.Errorf("erro: %w", err))
		return
	}
	fmt.Println("Must pay tax")

}

func validSalary(salary float64) (float64, error) {
	if salary < 150000 {
		return salary, errors.New("error: the salary entered does not reach the taxable minimum")
	}
	return salary, nil

}
