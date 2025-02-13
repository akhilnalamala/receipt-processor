package models

type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required,regex=^[\\w\\s\\-]+$"`
	Price            string `json:"price" validate:"required,regex=^\\d+\\.\\d{2}$"`
}
