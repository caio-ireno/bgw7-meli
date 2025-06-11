package repository

import (
	"fmt"

	"github.com/caio-ireno/prodcts-api-go/internal/domain"
	"github.com/caio-ireno/prodcts-api-go/pkg/apperrors"
)

type repository struct {
	storage map[int]domain.ProductAttributes
	lastId  int
}

func NewProductRepository(db map[int]domain.ProductAttributes) *repository {
	return &repository{
		storage: db,
	}
}

func (r repository) Get() ([]domain.Product, error) {

	products := make([]domain.Product, 0)

	for produtId, productAttr := range r.storage {
		product := productAttr.ToDomain()
		product.Id = produtId
		products = append(products, *product)

	}

	return products, nil
}

func (r *repository) Save(product *domain.Product) error {
	attr := domain.ProductAttributes{
		Name:     product.Name,
		Type:     product.Type,
		Quantity: product.Quantity,
		Price:    product.Price,
	}

	r.lastId++
	r.storage[r.lastId] = attr
	product.Id = r.lastId
	return nil
}

func (r repository) GetById(productId int) (*domain.Product, error) {
	p, ok := r.storage[productId]

	if !ok {
		return nil, apperrors.ErrResourceNotFound
	}

	product := p.ToDomain()
	product.Id = productId

	return product, nil
}

func (r *repository) Update(product *domain.Product) error {
	attr := domain.ProductAttributes{
		Name:     product.Name,
		Type:     product.Type,
		Quantity: product.Quantity,
		Price:    product.Price,
	}

	fmt.Println("rp:", attr)

	fmt.Println("rp:", *product)
	_, ok := r.storage[product.Id]

	if !ok {
		return apperrors.ErrResourceNotFound
	}

	r.storage[product.Id] = attr

	return nil
}

func (r *repository) Delete(productId int) error {
	_, ok := r.storage[productId]

	if !ok {
		return apperrors.ErrResourceNotFound
	}

	delete(r.storage, productId)

	return nil
}
