package handlers

import (
	"net/http"

	"receipt-processor/memory_store"
	"receipt-processor/models"
	"receipt-processor/points"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// return: a newly generated uuid for a given receipt as part of the response
// stores the receipt in a map, with the new id acting as the key
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

// return: the points tally of a receipt as part of the response
// looks up the receipt by the provided receiptID in the map
func GetReceiptPointsHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	receipt, exists := memory_store.GetReceipt(id)
	if !exists {
		http.Error(w, `{"error": "No receipt found for that ID."}`, http.StatusNotFound)
		return
	}

	points := points.CalculatePoints(receipt)

	render.JSON(w, r, map[string]int{"points": points})
}
