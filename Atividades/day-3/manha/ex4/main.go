package main

import (
	"errors"
	"fmt"
)

func main() {

	const (
		minimum = "minimum"

		average = "average"

		maximum = "maximum"
	)

	minFunc, err := operation(minimum)

	if err != nil {
		fmt.Println(err)
	}

	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
	}

	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)

	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)

	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("Mínimo: %.2f\n", minValue)
	fmt.Printf("Média: %.2f\n", averageValue)
	fmt.Printf("Máximo: %.2f\n", maxValue)

	// Testando operação inválida
	_, err = operation("mediana")
	if err != nil {
		fmt.Println("Erro:", err)
	}
}

func operation(operation string) (func(value ...int) float64, error) {
	switch operation {
	case "minimum":
		return func(values ...int) float64 {
			if len(values) == 0 {
				return 0
			}

			min := values[0]
			for _, i := range values {
				if i < min {
					min = i
				}
			}

			return float64(min)
		}, nil

	case "average":
		return func(values ...int) float64 {
			if len(values) == 0 {
				return 0
			}
			sum := 0
			for _, i := range values {
				sum += i

			}
			return float64(sum) / float64(len(values))
		}, nil

	case "maximum":
		return func(values ...int) float64 {
			if len(values) == 0 {
				return 0
			}

			max := 0
			for _, i := range values {
				if i > max {
					max = i
				}
			}
			return float64(max)
		}, nil

	default:
		return nil, errors.New("Operation invalid")

	}

}
