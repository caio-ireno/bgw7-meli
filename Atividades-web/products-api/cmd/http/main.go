package main

import (
	"net/http"

	"github.com/caio-ireno/prodcts-api-go/internal/domain"
	"github.com/caio-ireno/prodcts-api-go/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	rt := chi.NewRouter()

	rt.Use(middleware.Logger)

	rt.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	db := make(map[int]*domain.Product)
	productHandler := handler.NewProductHandler(db)

	// -> rotas
	rt.Route("/products", func(r chi.Router) {
		r.Post("/", productHandler.Create())
		r.Get("/", productHandler.GetAll())
		r.Get("/{productId}", productHandler.GetById())
		r.Put("/{productId}", productHandler.UpdateProdcut())
		r.Delete("/{productId}", productHandler.DeleteProdcut())
		r.Patch("/{productId}", productHandler.PatchProdcut())
	})

	// -> executar
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
