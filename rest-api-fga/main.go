package main

import (
	"rest-api-fga/database"
	"rest-api-fga/handler"
	productHandler "rest-api-fga/handler/product"
	"rest-api-fga/repo/product"
	productSvc "rest-api-fga/service/product"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}

	// Setup Repo
	productRepo := product.NewRepository(db)

	// Setup Service
	productService := productSvc.NewService(productRepo)

	// Setup Handler
	productHandler := productHandler.NewHandler(productService)

	handler.NewHttpServer(productHandler)
}
