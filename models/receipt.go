package models

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Receipt struct {
	Retailer     string `json:"retailer" validate:"required,regex=^[\\w\\s\\-&]+$"`
	PurchaseDate string `json:"purchaseDate" validate:"required, datetime=2008-02-05"`
	PurchaseTime string `json:"purchaseTime" validate:"required, datetime=13:05"`
	Items        []Item `json:"items" validate:"required,min=1"`
	Total        string `json:"total" validate:"required,regex=^\\d+\\.\\d{2}$"`
}

func (r *Receipt) Validate() error {
	return validate.Struct(r)
}
