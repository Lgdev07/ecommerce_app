package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProductSuccess(t *testing.T) {
	product := Product{
		Id:            1,
		Quantity:      1,
		UnitAmount:    1,
		TotalAmount:   2,
		DiscountGiven: 12,
		IsGift:        false,
	}

	err := product.IsValid()
	assert.Nil(t, err)
}

func TestNewProductError(t *testing.T) {
	product := Product{
		Id:            1,
		Quantity:      0,
		UnitAmount:    1,
		TotalAmount:   2,
		DiscountGiven: 12,
		IsGift:        false,
	}

	err := product.IsValid()
	assert.Error(t, err)
}
