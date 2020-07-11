package noAuthRoutes

import (
	"github.com/gofiber/fiber"
)

func GetIndex(c *fiber.Ctx){
	c.Render("index",fiber.Map{"Title": "hello world"}, "layouts/guest")
}

func GetLogin(c *fiber.Ctx){
	c.Render("login",fiber.Map{"Title": "login page"},"layouts/guest")
}