package models

import (
	"errors"
	"regexp"
)

var (
	descriptionRegex = regexp.MustCompile(`^[\w\s\-]+$`)
	priceRegex       = regexp.MustCompile(`^\d+\.\d{2}$`)
)

type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required"`
}

func (i *Item) Validate() error {

	// this validates that the required fields are present
	if err := validate.Struct(i); err != nil {
		return err
	}

	// regex validation for the shortDescription field
	if !descriptionRegex.MatchString(i.ShortDescription) {
		return errors.New("invalid description format")
	}

	// regex validation for the price field
	if !priceRegex.MatchString(i.Price) {
		return errors.New("invalid price format")
	}

	return nil
}
