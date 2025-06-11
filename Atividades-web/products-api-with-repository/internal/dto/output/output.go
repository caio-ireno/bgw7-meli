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

// ------- novos outputs ------------
type CreateBodyProductOutput struct {
	Message string          `json:"message"`
	Data    *domain.Product `json:"data"`
	Error   bool            `json:"error"`
}

type UpdateBodyProductOutput struct {
	Message string          `json:"message"`
	Data    *domain.Product `json:"data"`
	Error   bool            `json:"error"`
}

type GetBodyProductOutput struct {
	Message string           `json:"message"`
	Data    []domain.Product `json:"data"`
	Error   bool             `json:"error"`
}

type GetOneBodyProductOutput struct {
	Message string          `json:"message"`
	Data    *domain.Product `json:"data"`
	Error   bool            `json:"error"`
}
