package main

import (
	"log"
	"os"

	"github.com/Lgdev07/ecommerce_app/apps/ecommerce/handler"
	"github.com/Lgdev07/ecommerce_app/apps/ecommerce/repository"
	"github.com/Lgdev07/ecommerce_app/config"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.LoadDotEnv()
	db, err := repository.LoadDatabase()
	if err != nil {
		log.Fatal("error when creating database")
	}

	InitRoutes(e, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}

func InitRoutes(e *echo.Echo, db repository.ProductRepository) {
	checkoutHandler := handler.NewCheckoutHandler(db)

	e.POST("/checkout", checkoutHandler.Products)
}
