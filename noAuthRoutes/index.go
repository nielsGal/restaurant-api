package noAuthRoutes

import (
	"github.com/gofiber/fiber"
)

func GetIndex(c *fiber.Ctx){
	c.Send("hello")
}