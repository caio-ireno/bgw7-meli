package handler

import "net/http"

type ControllerProducts struct {
	st map[string]string
}

type ResponseGetByIdProducts struct {
	Message string
	Data    *struct {
		Id           int
		Name         string
		Quantity     int
		Code_value   string
		Is_published bool
		Expiration   string
		Price        float64
	}
}

func (c *ControllerProducts) GetById() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		return
	}
}
