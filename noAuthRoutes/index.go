package noAuthRoutes

import (
	"github.com/gofiber/fiber"
)

func GetIndex(c *fiber.Ctx){
	c.Render("index",fiber.Map{"Title": "hello world"}, "layouts/main")
}