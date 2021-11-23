package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCheckoutSuccess(t *testing.T) {
	checkout := Checkout{
		TotalAmount:             1,
		TotalDiscount:           2,
		TotalAmountWithDiscount: 1,
		Products: []Product{
			{
				Id:            1,
				Quantity:      1,
				UnitAmount:    1,
				TotalAmount:   2,
				DiscountGiven: 12,
				IsGift:        false,
			},
			{
				Id:            1,
				Quantity:      1,
				UnitAmount:    1,
				TotalAmount:   2,
				DiscountGiven: 12,
				IsGift:        false,
			},
		},
	}

	err := checkout.IsValid()
	assert.Nil(t, err)
}

func TestNewCheckoutError(t *testing.T) {
	checkout := Checkout{
		TotalAmount:             1,
		TotalDiscount:           2,
		TotalAmountWithDiscount: 1,
		Products:                nil,
	}

	err := checkout.IsValid()
	assert.Error(t, err)
}
