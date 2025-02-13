package models

import (
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	retailerRegex = regexp.MustCompile(`^[\w\s\-&]+$`)
	totalRegex    = regexp.MustCompile(`^\d+\.\d{2}$`)
)

var validate = validator.New()

type Receipt struct {
	Retailer     string `json:"retailer" validate:"required"`
	PurchaseDate string `json:"purchaseDate" validate:"required,datetime=2006-01-02"`
	PurchaseTime string `json:"purchaseTime" validate:"required,datetime=15:04"`
	Items        []Item `json:"items" validate:"required,min=1"`
	Total        string `json:"total" validate:"required"`
}

func (r *Receipt) Validate() error {

	// this validates that the required fields are present
	// this also validates the purchaseDate and purchaseTime fields
	if err := validate.Struct(r); err != nil {
		return err
	}

	// regex validation for retailer field
	if !retailerRegex.MatchString(r.Retailer) {
		return errors.New("invalid retailer format")
	}

	// regex validation for total field
	if !totalRegex.MatchString(r.Total) {
		return errors.New("invalid total format")
	}

	// this validates each item using the validation defined in the item model
	for i := range r.Items {
		if err := r.Items[i].Validate(); err != nil {
			return errors.New("invalid item in receipt")
		}
	}

	return nil
}
