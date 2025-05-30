package main

import (
	"fmt"
)

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
}

func ex1() {
	var nome string = "caio"
	nomeInfe := "Caio Eduardo"
	var lastname = "Ireno"

	fmt.Println(nome)
	fmt.Println(nomeInfe)
	fmt.Println(nomeInfe + " " + lastname)
	fmt.Println("--------------------------------")

}

func ex2() {
	fmt.Println("Análise do clima...")

	temp := 25.9
	umidade := 68
	pressao := 934.8

	fmt.Println("Temperatura:", temp, "C")
	fmt.Println("Umidade:", umidade, "%")
	fmt.Println("Pressão:", pressao, "hPa")
	fmt.Println("--------------------------------------")
}

func ex3() {
	/*
		ERRADO -- var 1firstName string - Errado, inicia com numero
		CERTO -- var lastName string - Certo
		ERRADO -- var int age - Errado, tipo vai depois
		ERRADO -- 1lastName := 6 - Errado, inicia com numero
		ERRADO -- var driver_license = true - Certo
		ERRADO -- var person height int - Errado, possui duas declaraCões
		CERTO --- childsNumber := 2 - certo
	*/

	var firstName string
	var lastName string
	// var int age
	var age int

	firstName = "caio"
	lastName = "Ireno"
	age = 20

	//    1lastName := 6

	var driver_license = true
	//var person height int
	childsNumber := 2

	fmt.Println("Testando variaveis...")
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(driver_license)
	fmt.Println(childsNumber)
	fmt.Println("--------------------------------------")

}

func ex4() {

	var lastName string = "Smith"
	var age int = 35              // estava como string mas a declaracao é int
	boolean := false              // estava como string mas o tipo declarado é bool
	var salary float64 = 45853.91 //Tipo estava como string mas deveria ser float.
	var firstName string = "Marry"

	fmt.Println(firstName, lastName, "Idade:", age, "Salario", salary, "Casada?", boolean)

}
