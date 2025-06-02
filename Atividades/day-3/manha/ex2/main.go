package main

import "fmt"

func main() {
	m1 := calculaMedia(1, 2, 3, 4, 5)
	fmt.Println(m1)
}

func calculaMedia(num ...float64) float64 {
	var result float64
	var i int
	for _, num := range num {
		result += num
		i += 1
	}
	return result / float64(i)
}
