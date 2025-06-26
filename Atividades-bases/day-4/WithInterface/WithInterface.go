package main

import "fmt"

type Store interface {
	Save(string)
	GetAll()
}

type Loja struct {
	produtos []string
}

func (l *Loja) Save(produto string) {
	l.produtos = append(l.produtos, produto)
}

func (l *Loja) GetAll() {
	for _, p := range l.produtos {
		fmt.Println(p)
	}
}

type MockLoja struct{}

func (m *MockLoja) Save(produto string) {
	fmt.Println("Mock: Produto salvo!")
}

func (m *MockLoja) GetAll() {
	fmt.Println("Mock: Listando produtos fictícios")
}

func main() {

	var store Store = &Loja{}
	// var store Store = &MockLoja{}

	store.Save("Bola")
	store.Save("Celular")
	store.GetAll()
}

//   SEM INTERFACE SERIA ASSIM:

// func main() {
//     // Para usar Loja:
//     // loja := &Loja{}
//     // loja.Save("Bola")
//     // loja.Save("Celular")
//     // loja.GetAll()

//     // Para usar MockLoja, precisa trocar o tipo e os nomes:
//     mock := &MockLoja{}
//     mock.Save("Bola")
//     mock.Save("Celular")
//     mock.GetAll()
// }
