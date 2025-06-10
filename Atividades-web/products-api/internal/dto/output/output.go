package output

import "github.com/caio-ireno/prodcts-api-go/internal/domain"

type ResponseBodyProduct struct {
	Message string          `json:"message"`
	Data    *domain.Product `json:"data"`
	Error   bool            `json:"error"`
}

type ResponseBodyProductAll struct {
	Message string            `json:"message"`
	Data    []*domain.Product `json:"data"`
	Error   bool              `json:"error"`
}
