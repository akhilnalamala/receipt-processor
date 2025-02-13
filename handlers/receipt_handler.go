package handlers

import (
	"net/http"

	"receipt-processor/memory_store"
	"receipt-processor/models"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt

	if err := render.DecodeJSON(r.Body, &receipt); err != nil {
		render.JSON(w, r, map[string]string{"error": "The receipt is invalid."})
		return
	}

	if err := receipt.Validate(); err != nil {
		render.JSON(w, r, map[string]string{"error": "The receipt is invalid."})
		return
	}

	receiptID := uuid.New().String()

	memory_store.StoreReceipt(receiptID, receipt)

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]string{"id": receiptID})
}
