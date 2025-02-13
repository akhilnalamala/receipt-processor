package memory_store

import (
	"receipt-processor/models"
	"sync"
)

var (
	mu       sync.Mutex
	Receipts = make(map[string]models.Receipt)
)

func StoreReceipt(id string, receipt models.Receipt) {
	mu.Lock()
	defer mu.Unlock()

	Receipts[id] = receipt
}

func GetReceipt(id string) (models.Receipt, bool) {
	mu.Lock()
	defer mu.Unlock()

	receipt, exists := Receipts[id]
	return receipt, exists
}
