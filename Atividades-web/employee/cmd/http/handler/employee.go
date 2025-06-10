package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handlerEmployee struct {
	st map[string]string
}

type ResponseGetByIdEmployee struct {
	Message string `json:"message"`
	Data    *struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
	Error bool `json:"error"`
}

// Constructor
func NewHandlerEmployee() *handlerEmployee {
	return &handlerEmployee{}
}

func (c *handlerEmployee) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathId := chi.URLParam(r, "id")
		queryParansId := r.URL.Query().Get("id")

		fmt.Println("PathId", pathId)
		fmt.Println("Query parans", queryParansId)

		employee, ok := c.st[pathId]

		if !ok {
			code := http.StatusNotFound
			body := &ResponseGetByIdEmployee{Message: "Employee not found", Data: nil, Error: true}

			w.WriteHeader(code)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(body)
			return
		}

		code := http.StatusOK
		body := &ResponseGetByIdEmployee{
			Message: "Employee Found",
			Data: &struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			}{Id: pathId, Name: employee},
			Error: false}

		w.WriteHeader(code)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)

	}

}

func (h *handlerEmployee) SetStore(store map[string]string) {
	h.st = store
}
