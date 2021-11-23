package repository

import (
	"errors"
)

type ProductDAO struct {
	Id          int
	Title       string
	Description string
	Amount      int
	IsGift      bool `json:"is_gift"`
}

func (m ProductRepository) Get(id int) (ProductDAO, error) {
	for _, p := range m.Products {
		if p.Id == id {
			return p, nil
		}
	}
	return ProductDAO{}, errors.New("product not found")
}

func (m ProductRepository) GetGift() (ProductDAO, error) {
	for _, p := range m.Products {
		if p.IsGift {
			return p, nil
		}
	}

	return ProductDAO{}, errors.New("gift not found")
}
