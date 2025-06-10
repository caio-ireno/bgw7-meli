package main

import (
	"log"
	"net/http"

	"github.com/caio-ireno/employee/cmd/http/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	employeeHandler := handler.NewHandlerEmployee()

	employeeHandler.SetStore(map[string]string{
		"1": "Maria",
		"2": "João",
	})
	router.Get("/employees/{id}", employeeHandler.GetById())

	log.Fatal(http.ListenAndServe(":8000", router))

}
