package routes

import (
	"receipt-processor/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/receipts/process", handlers.ProcessReceipt)
	r.Get("/receipts/{id}/points", handlers.GetReceiptPointsHandler)

	return r
}
