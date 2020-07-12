package noAuthRoutes

import (
	"github.com/gofiber/fiber"
)

func GetIndex(c *fiber.Ctx){
	c.Render("index",nil, "layouts/guest")
}

func GetLogin(c *fiber.Ctx){
	c.Render("login",nil,"layouts/guest")
}

func GetRegister(c *fiber.Ctx){
	c.Render("register",nil,"layouts/user")
}

func GetContact(c *fiber.Ctx){
	c.Render("contact",nil,"layouts/guest")
}

func PostContact(c *fiber.Ctx){
	c.Render("contact",nil,"layouts/guest")
}