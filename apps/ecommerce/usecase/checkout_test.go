package usecase

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Lgdev07/ecommerce_app/apps/ecommerce/repository"
	"github.com/Lgdev07/ecommerce_app/apps/ecommerce/schema"
	"github.com/stretchr/testify/assert"
)

type MockDiscount struct{}

func (s MockDiscount) GetProductDiscount(productID int) (float64, error) {
	return 0.1, nil
}

func MockProducts() repository.ProductRepository {
	return repository.ProductRepository{
		Products: []repository.ProductDAO{
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

func TestCalculateProductsSuccess(t *testing.T) {
	checkout := CheckoutUseCase{
		Repo:     MockProducts(),
		Discount: &MockDiscount{},
	}

	checkoutRequest := schema.CheckoutRequest{
		Products: []schema.ProductRequest{
			{
				Id:       1,
				Quantity: 1,
			},
			{
				Id:       2,
				Quantity: 1,
			},
		},
	}

	response, _ := checkout.CalculateProducts(checkoutRequest)

	assert.Len(t, response.Products, 1)
	assert.Equal(t, response.TotalAmount, 123)
	assert.Equal(t, response.TotalDiscount, 12)
	assert.Equal(t, response.TotalAmountWithDiscount, 111)

	assert.Equal(t, response.Products[0].Id, 1)
	assert.Equal(t, response.Products[0].DiscountGiven, 12)
	assert.Equal(t, response.Products[0].IsGift, false)
	assert.Equal(t, response.Products[0].Quantity, 1)
	assert.Equal(t, response.Products[0].TotalAmount, 123)
	assert.Equal(t, response.Products[0].UnitAmount, 123)
}

func TestCalculateProductsErrorProductDoesntExist(t *testing.T) {
	checkout := CheckoutUseCase{
		Repo:     MockProducts(),
		Discount: &MockDiscount{},
	}

	checkoutRequest := schema.CheckoutRequest{
		Products: []schema.ProductRequest{
			{
				Id:       3,
				Quantity: 1,
			},
		},
	}

	_, err := checkout.CalculateProducts(checkoutRequest)

	assert.Error(t, err)
}

func TestCalculateProductsWithGiftSuccess(t *testing.T) {
	today := time.Now()

	closer := envSetter(map[string]string{
		"BLACK_FRIDAY_MONTH": fmt.Sprint(int(today.Month())),
		"BLACK_FRIDAY_DAY":   fmt.Sprint(today.Day()),
	})
	defer closer()

	checkout := CheckoutUseCase{
		Repo:     MockProducts(),
		Discount: &MockDiscount{},
	}

	checkoutRequest := schema.CheckoutRequest{
		Products: []schema.ProductRequest{
			{
				Id:       1,
				Quantity: 1,
			},
			{
				Id:       2,
				Quantity: 1,
			},
		},
	}

	response, _ := checkout.CalculateProducts(checkoutRequest)

	assert.Len(t, response.Products, 2)
	assert.Equal(t, response.TotalAmount, 123)
	assert.Equal(t, response.TotalDiscount, 12)
	assert.Equal(t, response.TotalAmountWithDiscount, 111)

	assert.Equal(t, response.Products[0].Id, 1)
	assert.Equal(t, response.Products[0].DiscountGiven, 12)
	assert.Equal(t, response.Products[0].IsGift, false)
	assert.Equal(t, response.Products[0].Quantity, 1)
	assert.Equal(t, response.Products[0].TotalAmount, 123)
	assert.Equal(t, response.Products[0].UnitAmount, 123)

	assert.Equal(t, response.Products[1].Id, 2)
	assert.Equal(t, response.Products[1].DiscountGiven, 0)
	assert.Equal(t, response.Products[1].IsGift, true)
	assert.Equal(t, response.Products[1].Quantity, 1)
	assert.Equal(t, response.Products[1].TotalAmount, 0)
	assert.Equal(t, response.Products[1].UnitAmount, 0)
}

func TestItsBlackFridayFalse(t *testing.T) {
	result := ItsBlackFriday()

	assert.False(t, result)
}

func TestItsBlackFridayTrue(t *testing.T) {
	today := time.Now()

	closer := envSetter(map[string]string{
		"BLACK_FRIDAY_MONTH": fmt.Sprint(int(today.Month())),
		"BLACK_FRIDAY_DAY":   fmt.Sprint(today.Day()),
	})
	defer closer()

	result := ItsBlackFriday()

	assert.True(t, result)
}

func envSetter(envs map[string]string) (closer func()) {
	originalEnvs := map[string]string{}

	for name, value := range envs {
		if originalValue, ok := os.LookupEnv(name); ok {
			originalEnvs[name] = originalValue
		}
		_ = os.Setenv(name, value)
	}

	return func() {
		for name := range envs {
			origValue, has := originalEnvs[name]
			if has {
				_ = os.Setenv(name, origValue)
			} else {
				_ = os.Unsetenv(name)
			}
		}
	}
}
