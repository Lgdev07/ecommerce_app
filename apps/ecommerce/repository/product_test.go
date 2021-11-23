package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func MockProducts() ProductRepository {
	return ProductRepository{
		Products: []ProductDAO{
			{
				Id:          1,
				Title:       "Test Product 1",
				Description: "Description Product 1",
				Amount:      123,
				IsGift:      false,
			},
			{
				Id:          2,
				Title:       "Test Product 2",
				Description: "Description Product 2",
				Amount:      1234,
				IsGift:      true,
			},
		},
	}
}

func TestGetProductSucess(t *testing.T) {
	productsRepo := MockProducts()

	productDAO, _ := productsRepo.Get(1)

	assert.Equal(t, productDAO.Id, 1)
	assert.Equal(t, productDAO.Title, "Test Product 1")
	assert.Equal(t, productDAO.Description, "Description Product 1")
	assert.Equal(t, productDAO.Amount, 123)
	assert.Equal(t, productDAO.IsGift, false)
}

func TestGetProductError(t *testing.T) {
	productsRepo := ProductRepository{
		Products: []ProductDAO{},
	}

	_, err := productsRepo.Get(1)

	assert.Error(t, err)
}

func TestGetGiftSucess(t *testing.T) {
	productsRepo := MockProducts()

	productDAO, _ := productsRepo.GetGift()

	assert.Equal(t, productDAO.Id, 2)
	assert.Equal(t, productDAO.Title, "Test Product 2")
	assert.Equal(t, productDAO.Description, "Description Product 2")
	assert.Equal(t, productDAO.Amount, 1234)
	assert.Equal(t, productDAO.IsGift, true)
}
