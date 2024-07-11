package usecase

import (
	"fmt"
	"go_race_condition/src/repository"
	"sync"
)

// ProductUsecase is an interface defining methods for managing product operations.
type ProductUsecase interface {
	GetStock() int
	UpdateStock(amount int)
	UpdateStockWithWaitGroup(amount int)
	UpdateStockWithMutex(amount int)
}

type productUsecase struct {
	productRepo repository.ProductRepository
	mu          sync.Mutex
}

// NewProductUsecase creates a new instance of ProductUsecase with the provided ProductRepository.
func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
	return &productUsecase{productRepo: productRepo}
}

// GetStock retrieves the current stock count from the repository.
func (uc *productUsecase) GetStock() int {
	return uc.productRepo.GetStock()
}

func (uc *productUsecase) UpdateStock(amount int) {
	currentStock := uc.productRepo.GetStock()
	newStock := currentStock + amount

	// Print a message to indicate that race condition is happening
	fmt.Printf("Process Detected -> Current stock: %d, Updating with amount: %d, New stock: %d\n", currentStock, amount, newStock)

	// Update the stock
	uc.productRepo.UpdateStock(amount)
}

// UpdateStockWithWaitGroup updates the stock using WaitGroup to handle concurrency.
// It creates a WaitGroup to coordinate the concurrent update operation.
// A goroutine is spawned to execute the repository's UpdateStock method asynchronously.
// WaitGroup ensures that the main routine waits until the goroutine finishes its execution.
func (uc *productUsecase) UpdateStockWithWaitGroup(amount int) {
	var wg sync.WaitGroup
	wg.Add(1) // Increment the WaitGroup counter by 1
	// Incrementing the WaitGroup counter by 1 (wg.Add(1))
	// means adding the number of goroutines that will be executed
	// concurrently by the WaitGroup. This approach allows us to inform the WaitGroup of how many
	// goroutines need to run concurrently before continuing the main process.
	go func() {
		defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
		currentStock := uc.productRepo.GetStock()
		newStock := currentStock + amount
		// Print a message to indicate that race condition is happening
		fmt.Printf("Process Detected WG-> Current stock: %d, Updating with amount: %d, New stock: %d\n", currentStock, amount, newStock)
		uc.productRepo.UpdateStock(amount) // Execute the repository's UpdateStock method
	}()
	wg.Wait() // Wait for all goroutines in the WaitGroup to finish
}

// UpdateStockWithMutex updates the stock using a Mutex to handle concurrency.
// It locks a Mutex to ensure exclusive access to the shared resource (stock).
// After updating the stock, it unlocks the Mutex to allow other goroutines to access it.
func (uc *productUsecase) UpdateStockWithMutex(amount int) {
	uc.mu.Lock()         // Acquire the lock of the Mutex
	defer uc.mu.Unlock() // Ensure the Mutex is unlocked when the function exits
	currentStock := uc.productRepo.GetStock()
	newStock := currentStock + amount
	// Print a message to indicate that race condition is happening
	fmt.Printf("Process Detected Mutex-> Current stock: %d, Updating with amount: %d, New stock: %d\n", currentStock, amount, newStock)
	uc.productRepo.UpdateStock(amount) // Execute the repository's UpdateStock method under Mutex lock
}
