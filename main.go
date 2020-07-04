package main

import (
	"github.com/gofiber/fiber"
	"github.com/nielsGal/restaurant-api/views"
)


func setupRoutes(app *fiber.App){
	app.Get("/api/v1/get-menu",views.GetMenu)
	app.Post("/api/v1/create-menu",views.CreateMenu)
	app.Delete("/api/v1/delete-menu",views.DeleteMenu)

	app.Get("/api/v1/get-product",views.GetProduct)
	app.Get("/api/v1/get-products",views.GetProducts)
	app.Post("/api/v1/create-product",views.CreateProduct)
	app.Post("/api/v1/create-products",views.CreateProducts)
	app.Delete("/api/v1/delete-product",views.DeleteProduct)
	app.Delete("/api/v1/delete-products",views.DeleteProducts)

	app.Get("/api/v1/get-category",views.GetCategory)
	app.Get("/api/v1/get-categories",views.GetCategories)
	app.Post("/api/v1/create-category",views.CreateCategory)
	app.Post("/api/v1/create-categories",views.CreateCategories)
	app.Delete("/api/v1/delete-category",views.DeleteCategory)
	app.Delete("/api/v1/delete-categories",views.DeleteCategories)
}


func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen("3000")
}