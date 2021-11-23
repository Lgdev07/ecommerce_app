package usecase

import (
	"os"
	"strconv"
	"time"

	"github.com/Lgdev07/ecommerce_app/apps/discount"
	"github.com/Lgdev07/ecommerce_app/apps/ecommerce/model"
	"github.com/Lgdev07/ecommerce_app/apps/ecommerce/repository"
	"github.com/Lgdev07/ecommerce_app/apps/ecommerce/schema"
)

type CheckoutUseCase struct {
	Repo     repository.ProductRepository
	Discount discount.DiscountServiceInt
}

func (c *CheckoutUseCase) CalculateProducts(checkoutRequest schema.CheckoutRequest) (model.Checkout, error) {
	checkout := model.Checkout{}

	for _, product := range checkoutRequest.Products {
		productResult, err := c.Repo.Get(product.Id)
		if err != nil {
			return checkout, err
		}

		if productResult.IsGift {
			continue
		}

		discountResult, err := c.Discount.GetProductDiscount(productResult.Id)
		if err != nil {
			return checkout, err
		}

		product := &model.Product{
			Id:            productResult.Id,
			Quantity:      product.Quantity,
			UnitAmount:    productResult.Amount,
			TotalAmount:   productResult.Amount * product.Quantity,
			DiscountGiven: int(float64(productResult.Amount*product.Quantity) * discountResult),
			IsGift:        productResult.IsGift,
		}

		if err := product.IsValid(); err != nil {
			return checkout, err
		}

		checkout.Products = append(checkout.Products, *product)
		checkout.TotalAmount += product.TotalAmount
		checkout.TotalDiscount += product.DiscountGiven
	}

	if ItsBlackFriday() {
		gift, err := c.Repo.GetGift()
		if err != nil {
			return checkout, err
		}

		giftResponse := model.Product{
			Id:            gift.Id,
			Quantity:      1,
			UnitAmount:    0,
			TotalAmount:   0,
			DiscountGiven: 0,
			IsGift:        gift.IsGift,
		}

		checkout.Products = append(checkout.Products, giftResponse)
	}

	checkout.TotalAmountWithDiscount = checkout.TotalAmount - checkout.TotalDiscount

	return checkout, nil
}

func ItsBlackFriday() bool {
	blackFridayMonth, _ := strconv.Atoi(os.Getenv("BLACK_FRIDAY_MONTH"))
	blackFridayDay, _ := strconv.Atoi(os.Getenv("BLACK_FRIDAY_DAY"))

	today := time.Now()
	return blackFridayMonth == int(today.Month()) && blackFridayDay == today.Day()
}
