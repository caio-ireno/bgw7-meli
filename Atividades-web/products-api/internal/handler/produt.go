package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/caio-ireno/prodcts-api-go/internal/domain"
	"github.com/caio-ireno/prodcts-api-go/internal/dto"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	storage map[int]*domain.Product
}

func NewProductHandler(db map[int]*domain.Product) *ProductHandler {
	return &ProductHandler{storage: db}
}

func (c *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		if token != "123456" {
			code := http.StatusUnauthorized // 401
			body := &dto.ResponseBodyProduct{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		// request
		var reqBody dto.RequestBodyProduct // do tipo RequestBodyProduct
		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			code := http.StatusBadRequest
			body := &dto.ResponseBodyProduct{
				Message: "Bad Request",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		pr := reqBody.ToDomain()
		pr.Id = len(c.storage) + 1

		c.storage[pr.Id] = &pr

		code := http.StatusCreated
		body := &dto.ResponseBodyProduct{
			Message: "Product created",
			Data: &domain.Product{
				Id:       pr.Id,
				Name:     pr.Name,
				Type:     pr.Type,
				Quantity: pr.Quantity,
				Price:    pr.Price,
			},
			Error: false,
		}

		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)

	}
}

func (c ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		products := make([]*domain.Product, 0, len(c.storage))

		for _, product := range c.storage {
			products = append(products, product)
		}

		body := &dto.ResponseBodyProductAll{
			Message: "Products Found",
			Data:    products,
			Error:   false,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)

	}

}

func (c ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathId := chi.URLParam(r, "productId")
		productsId, err := strconv.Atoi(pathId)

		if err != nil {
			panic("Erro ao converter o pathid")
		}

		products, ok := c.storage[productsId]

		if !ok {
			code := http.StatusNotFound
			body := &dto.ResponseBodyProduct{
				Message: "Product not found",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		code := http.StatusOK
		body := &dto.ResponseBodyProduct{
			Message: "Product found",
			Data: &domain.Product{
				Name:     products.Name,
				Type:     products.Type,
				Quantity: products.Quantity,
				Price:    products.Price,
			},
		}

		w.WriteHeader(code)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)

	}
}

func (c *ProductHandler) UpdateProdcut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		if token != "123456" {
			code := http.StatusUnauthorized // 401
			body := &dto.ResponseBodyProduct{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		productId, err := strconv.Atoi(chi.URLParam(r, "productId"))

		if err != nil {
			code := http.StatusBadRequest
			body := &dto.ResponseBodyProduct{Message: "Bad Request", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		// Validação: verifica se o produto existe
		_, ok := c.storage[productId]

		if !ok {
			code := http.StatusNotFound
			body := &dto.ResponseBodyProduct{
				Message: "Product not found",
				Data:    nil,
				Error:   true,
			}
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		// Cria uma var do tipo UpdateBodyProduct
		var reqBody dto.UpdateBodyProduct

		err = json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			code := http.StatusBadRequest
			body := &dto.ResponseBodyProduct{
				Message: "Bad Request",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		pr := reqBody.ToDomain()
		pr.Id = productId

		c.storage[productId] = &pr

		code := http.StatusOK
		body := &dto.ResponseBodyProduct{
			Message: "Product Update",
			Data:    &pr,
			Error:   false,
		}

		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)

	}
}

func (c *ProductHandler) DeleteProdcut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		if token != "123456" {
			code := http.StatusUnauthorized // 401
			body := &dto.ResponseBodyProduct{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		productId, err := strconv.Atoi(chi.URLParam(r, "productId"))

		if err != nil {
			code := http.StatusBadRequest
			body := &dto.ResponseBodyProduct{Message: "Bad Request", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		// Validação: verifica se o produto existe
		_, ok := c.storage[productId]

		if !ok {
			code := http.StatusNotFound
			body := &dto.ResponseBodyProduct{
				Message: "Product not found",
				Data:    nil,
				Error:   true,
			}
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		delete(c.storage, productId)

		code := http.StatusOK
		body := &dto.ResponseBodyProduct{
			Message: "Deleted Product ",
			Data:    nil,
			Error:   false,
		}

		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)

	}
}

func (c *ProductHandler) PatchProdcut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		if token != "123456" {
			code := http.StatusUnauthorized // 401
			body := &dto.ResponseBodyProduct{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		productId, err := strconv.Atoi(chi.URLParam(r, "productId"))

		if err != nil {
			code := http.StatusBadRequest
			body := &dto.ResponseBodyProduct{Message: "Bad Request", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		// Validação: verifica se o produto existe
		product, ok := c.storage[productId]

		if !ok {
			code := http.StatusNotFound
			body := &dto.ResponseBodyProduct{
				Message: "Product not found",
				Data:    nil,
				Error:   true,
			}
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}
		/// ------ok -----

		// Cria uma var do tipo UpdateBodyProduct
		reqBody := dto.UpdateBodyProduct{
			Name:     product.Name,
			Type:     product.Type,
			Quantity: product.Quantity,
			Price:    product.Price,
		}

		err = json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			code := http.StatusUnprocessableEntity
			body := &dto.ResponseBodyProduct{
				Message: "Bad Request",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		pr := reqBody.ToDomain()
		pr.Id = productId

		// ok--------

		c.storage[productId] = &pr

		code := http.StatusOK
		body := &dto.ResponseBodyProduct{
			Message: "Product Update",
			Data:    &pr,
			Error:   false,
		}

		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)

	}
}
