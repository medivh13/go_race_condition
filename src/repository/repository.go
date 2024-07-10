package repository

import "fmt"

// ProductRepository is an interface that defines methods for managing product stock.
type ProductRepository interface {
    // GetStock retrieves the current stock count.
    GetStock() int
    // UpdateStock updates the stock count by the specified amount.
    UpdateStock(amount int)
}

type productRepository struct {
    stock int
}

// NewProductRepository creates a new instance of ProductRepository with the given initial stock.
func NewProductRepository(initialStock int) ProductRepository {
    return &productRepository{stock: initialStock}
}

// GetStock retrieves the current stock count.
func (r *productRepository) GetStock() int {
    return r.stock
}

// UpdateStock updates the stock count by the specified amount and prints the updated stock.
func (r *productRepository) UpdateStock(amount int) {
    r.stock += amount
    fmt.Printf("Updated stock: %d\n", r.stock)
}
