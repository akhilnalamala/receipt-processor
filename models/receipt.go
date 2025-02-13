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
	if err := validate.Struct(r); err != nil {
		return err
	}

	if !retailerRegex.MatchString(r.Retailer) {
		return errors.New("invalid retailer format")
	}

	if !totalRegex.MatchString(r.Total) {
		return errors.New("invalid total format")
	}

	for i := range r.Items {
		if err := r.Items[i].Validate(); err != nil {
			return errors.New("invalid item in receipt")
		}
	}

	return nil
}
