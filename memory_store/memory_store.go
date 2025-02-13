package memory_store

import (
	"receipt-processor/models"
	"sync"
)

// memory store implemented using a map + a sync.Mutex
var (
	mu       sync.Mutex
	Receipts = make(map[string]models.Receipt)
)

// stores the receipt in the map with the receiptID as the key
func StoreReceipt(id string, receipt models.Receipt) {
	mu.Lock()
	defer mu.Unlock()

	Receipts[id] = receipt
}

// looks up a receiptID and returns the receipt if it exists
func GetReceipt(id string) (models.Receipt, bool) {
	mu.Lock()
	defer mu.Unlock()

	receipt, exists := Receipts[id]
	return receipt, exists
}
