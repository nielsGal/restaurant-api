package main

import (
	"github.com/gofiber/fiber"
	"github.com/nielsGal/restaurant-api/views"
)


func setupRoutes(app *fiber.App){
	app.Get("/api/v1/getmenu",views.GetMenu)
	app.Post("/api/v1/createmenu",views.CreateMenu)
	app.Delete("/api/v1/deletemenu",views.DeleteMenu)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen("3000")
}