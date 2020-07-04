package main

import (
	"github.com/gofiber/fiber"
	"github.com/nielsGal/restaurant-api/views"
)


func setupRoutes(app *fiber.App){
	app.Get("/api/v1/getmenu",views.GetMenu)
	app.Post("/api/v1/createmenu",views.CreateMenu)
	app.Delete("/api/v1/deletemenu",views.DeleteMenu)

	app.Get("/api/v1/getproduct",views.GetProduct)
	app.Get("/api/v1/getproducts",views.GetProducts)
	app.Post("/api/v1/createproduct",views.CreateProduct)
	app.Post("/api/v1/createproducts",views.CreateProducts)
	app.Delete("/api/v1/deleteproduct",views.DeleteProduct)
	app.Delete("/api/v1/deleteproducts",views.DeleteProducts)

	app.Get("/api/v1/getcategory",views.GetCategory)
	app.Get("/api/v1/getcategories",views.GetCategories)
	app.Post("/api/v1/createcategory",views.CreateCategory)
	app.Post("/api/v1/createcategories",views.CreateCategories)
	app.Delete("/api/v1/deletecategory",views.DeleteCategory)
	app.Delete("/api/v1/deletecategories",views.DeleteCategories)
}


func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen("3000")
}