package input

import "github.com/caio-ireno/prodcts-api-go/internal/domain"

func (r RequestBodyProduct) ToDomain() domain.Product {
	return domain.Product{
		Name:     r.Name,
		Type:     r.Type,
		Quantity: r.Quantity,
		Price:    r.Price,
	}
}

type RequestBodyProduct struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func (r UpdateBodyProduct) ToDomain() domain.Product {
	return domain.Product{
		Name:     r.Name,
		Type:     r.Type,
		Quantity: r.Quantity,
		Price:    r.Price,
	}
}

type UpdateBodyProduct struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
