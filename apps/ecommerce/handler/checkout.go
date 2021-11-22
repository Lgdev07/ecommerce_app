package handler

import (
	"log"
	"net/http"

	"github.com/Lgdev07/ecommerce_app/apps/discount"
	"github.com/Lgdev07/ecommerce_app/apps/ecommerce/repository"
	"github.com/Lgdev07/ecommerce_app/apps/ecommerce/schema"
	"github.com/Lgdev07/ecommerce_app/apps/ecommerce/usecase"
	"github.com/labstack/echo/v4"
)

type CheckoutHandler struct {
	CheckoutUseCase usecase.CheckoutUseCase
}

func (f *CheckoutHandler) Products(c echo.Context) error {
	checkout := &schema.CheckoutRequest{}
	if err := c.Bind(checkout); err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusBadRequest, "an error occurred when decoding the request body")
	}

	result, err := f.CheckoutUseCase.CalculateProducts(*checkout)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, "an error occurred during execution")
	}

	return c.JSON(http.StatusOK, result)
}

func NewCheckoutHandler(db repository.ProductRepository) *CheckoutHandler {
	fareUseCase := usecase.CheckoutUseCase{
		Repo:     db,
		Discount: &discount.DiscountService{},
	}
	fareHandler := &CheckoutHandler{
		CheckoutUseCase: fareUseCase,
	}

	return fareHandler
}
