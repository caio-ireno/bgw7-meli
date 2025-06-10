package domain

// Representa um produto completo.
// Esse struct geralmente é a sua ENTIDADE de DOMÍNIO,
// ou seja, equivale ao que existe “no seu sistema”.
type Product struct {
	Id       int
	Name     string
	Type     string
	Quantity int
	Price    float64
}

//Representa apenas os atributos que podem ser enviados (input) ou retornados (output) sem o ID.

type ProductAttributes struct {
	Name     string
	Type     string
	Quantity int
	Price    float64
}

type ProductRepository interface {
	Get() ([]Product, error)

	GetById(productId int) (*Product, error)

	Save(product *Product) error

	Update(product *Product) error

	Delete(productId int) error
}
