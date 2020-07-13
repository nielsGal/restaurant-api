package authRoutes

import (
	"github.com/gofiber/fiber"
)

func GetAdminPage(c *fiber.Ctx){
	c.Render("admin",nil,"layouts/admin")
}

func GetAdminRestaurantPage(c *fiber.Ctx){
	c.Render("admin",nil,"layouts/admin")
}