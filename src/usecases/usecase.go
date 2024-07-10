package usecase

import (
    "go_race_condition/src/repository"
    "sync"
)

type ProductUsecase interface {
    GetStock() int
    UpdateStockWithWaitGroup(amount int)
    UpdateStockWithMutex(amount int)
}

type productUsecase struct {
    productRepo repository.ProductRepository
    mu          sync.Mutex
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
    return &productUsecase{productRepo: productRepo}
}

func (uc *productUsecase) GetStock() int {
    return uc.productRepo.GetStock()
}

func (uc *productUsecase) UpdateStockWithWaitGroup(amount int) {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        uc.productRepo.UpdateStock(amount)
    }()
    wg.Wait()
}

func (uc *productUsecase) UpdateStockWithMutex(amount int) {
    uc.mu.Lock()
    defer uc.mu.Unlock()
    uc.productRepo.UpdateStock(amount)
}
