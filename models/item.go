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

	if err := validate.Struct(i); err != nil {
		return err
	}

	if !descriptionRegex.MatchString(i.ShortDescription) {
		return errors.New("invalid description format")
	}

	if !priceRegex.MatchString(i.Price) {
		return errors.New("invalid price format")
	}

	return nil
}
