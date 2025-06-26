package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/caio-ireno/prodcts-api-go/internal/domain"
	"github.com/caio-ireno/prodcts-api-go/internal/dto/input"
	"github.com/caio-ireno/prodcts-api-go/internal/dto/output"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	repository domain.ProductRepository
}

// Implementa a interface criada no domain
func NewProductHandler(rp domain.ProductRepository) *ProductHandler {
	return &ProductHandler{repository: rp}
}

func (c *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		if token != "123456" {
			code := http.StatusUnauthorized // 401
			body := &output.ResponseBodyProduct{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		var reqBody input.RequestBodyProduct // do tipo RequestBodyProduct
		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			code := http.StatusBadRequest
			body := &output.ResponseBodyProduct{
				Message: "Bad Request",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		// pr := input.RequestBodyProduct.ToDomain(reqBody)
		pr := reqBody.ToDomain()

		err = c.repository.Save(&pr)

		if err != nil {
			code := http.StatusBadRequest
			body := &output.ResponseBodyProduct{
				Message: "Erro ao salvar no banco de dados",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		code := http.StatusCreated
		body := &output.CreateBodyProductOutput{
			Message: "Product created",
			Data:    &pr,
			Error:   false,
		}

		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)

	}
}

func (c ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		products, err := c.repository.Get()

		if err != nil {
			code := http.StatusBadRequest
			body := &output.ResponseBodyProduct{
				Message: "Bad Request",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		body := &output.GetBodyProductOutput{
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
			code := http.StatusBadRequest
			body := &output.ResponseBodyProduct{
				Message: "Bad Request",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		product, ok := c.repository.GetById(productsId)

		if ok != nil {
			code := http.StatusNotFound
			body := &output.ResponseBodyProduct{
				Message: "Product not found",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		code := http.StatusOK
		body := &output.ResponseBodyProduct{
			Message: "Product found",
			Data:    product,
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
			body := &output.UpdateBodyProductOutput{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		productId, err := strconv.Atoi(chi.URLParam(r, "productId"))

		if err != nil {
			code := http.StatusBadRequest
			body := &output.UpdateBodyProductOutput{Message: "Bad Request to get id", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		_, ok := c.repository.GetById(productId)

		if ok != nil {
			code := http.StatusNotFound
			body := &output.UpdateBodyProductOutput{
				Message: "Product not found",
				Data:    nil,
				Error:   true,
			}
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		var reqBody input.UpdateBodyProduct

		err = json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			code := http.StatusBadRequest
			body := &output.UpdateBodyProductOutput{
				Message: "Erro ao converter dados",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		pr := reqBody.ToDomain()
		pr.Id = productId

		fmt.Println(pr)

		err = c.repository.Update(&pr)

		if err != nil {
			code := http.StatusBadRequest
			body := &output.UpdateBodyProductOutput{
				Message: "Erro ao atualizar o produto no Storage",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		code := http.StatusOK
		body := &output.UpdateBodyProductOutput{
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
			body := &output.ResponseBodyProduct{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		productId, err := strconv.Atoi(chi.URLParam(r, "productId"))

		if err != nil {
			code := http.StatusBadRequest
			body := &output.ResponseBodyProduct{Message: "Bad Request", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		// Validação: verifica se o produto existe
		_, ok := c.repository.GetById(productId)

		if ok != nil {
			code := http.StatusNotFound
			body := &output.ResponseBodyProduct{
				Message: "Product not found",
				Data:    nil,
				Error:   true,
			}
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		err = c.repository.Delete(productId)

		if err != nil {
			code := http.StatusBadRequest
			body := &output.ResponseBodyProduct{
				Message: "Erro ao deletar produto",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		code := http.StatusOK
		body := &output.ResponseBodyProduct{
			Message: "Deleted Product",
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
			body := &output.ResponseBodyProduct{Message: "Unauthorized", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		productId, err := strconv.Atoi(chi.URLParam(r, "productId"))

		if err != nil {
			code := http.StatusBadRequest
			body := &output.ResponseBodyProduct{Message: "Bad Request", Data: nil, Error: true}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		product, ok := c.repository.GetById(productId)

		if ok != nil {
			code := http.StatusNotFound
			body := &output.ResponseBodyProduct{
				Message: "Product not found",
				Data:    nil,
				Error:   true,
			}
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}
		/// ------ok -----

		// Cria um reqBody preenchido com os dados atuais:
		reqBody := input.UpdateBodyProduct{
			Name:     product.Name,
			Type:     product.Type,
			Quantity: product.Quantity,
			Price:    product.Price,
		}

		err = json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			code := http.StatusUnprocessableEntity
			body := &output.ResponseBodyProduct{
				Message: "Erro no formato dos dados",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		pr := reqBody.ToDomain()
		pr.Id = productId
		// O update ja recebe um produto com ID, por isso nao passamos no método
		err = c.repository.Update(&pr)

		if err != nil {
			code := http.StatusBadRequest
			body := &output.ResponseBodyProduct{
				Message: "Erro ao atualizar parcialmente o produto",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		code := http.StatusOK
		body := &output.ResponseBodyProduct{
			Message: "Product Update",
			Data:    &pr,
			Error:   false,
		}

		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)

	}
}

