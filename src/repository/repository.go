package repository

import "fmt"

type ProductRepository interface {
    GetStock() int
    UpdateStock(amount int)
}

type productRepository struct {
    stock int
}

func NewProductRepository(initialStock int) ProductRepository {
    return &productRepository{stock: initialStock}
}

func (r *productRepository) GetStock() int {
    return r.stock
}

func (r *productRepository) UpdateStock(amount int) {
    r.stock += amount
    fmt.Printf("Updated stock: %d\n", r.stock)
}
