// Ter uma estrutura chamada Product com os campos ID, Name, Price, Description e Category.

// Ter uma fatia global de Produto chamada Produtos instanciada com valores. 2 métodos associados à estrutura Produto: Save(), GetAll().
//  O método Save() deve pegar a fatia de Products e adicionar o produto a partir do qual o método é chamado.
//  O método GetAll() imprime todos os produtos salvos na fatia Products.

// Uma função getById() para a qual um INT deve ser passado como parâmetro e retorna o produto correspondente ao parâmetro passado.

// Execute pelo menos uma vez cada método e função definidos em main().

// Loja possui produtos

package main

import (
	"errors"
	"fmt"
)

type Loja struct {
	produtos []Produto
}

type Produto struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (l *Loja) Save(myProduct Produto) {
	p := myProduct
	l.produtos = append(l.produtos, p)

}

func (l Loja) GetAll() {
	for _, produto := range l.produtos {
		fmt.Printf("%+v\n", produto)
	}

}

func (l Loja) GetById(id int) (Produto, error) {
	for _, produto := range l.produtos {
		if produto.Id == id {
			return produto, nil
		}
	}
	return Produto{}, errors.New("id não existe")
}

func main() {

	p1 := Produto{
		Id:          1,
		Name:        "Bola",
		Price:       20.50,
		Description: "Bola de couro",
		Category:    "Esporte",
	}

	p2 := Produto{
		Id:          2,
		Name:        "Celular",
		Price:       1500.50,
		Description: "Iphone",
		Category:    "Eletronico",
	}

	myLoja := Loja{} // Instanciando a struct Loja

	myLoja.Save(p1)
	myLoja.Save(p2)
	myLoja.GetAll()

	produtoId, myError := myLoja.GetById(2)
	if myError != nil {
		fmt.Println(myError)
	} else {
		fmt.Println(produtoId.Name)
	}
}
