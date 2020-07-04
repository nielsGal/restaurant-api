package views

import "github.com/gofiber/fiber"

func GetMenu(c *fiber.Ctx){
	c.Send("get a specific menu")
}

func CreateMenu(c *fiber.Ctx){
	c.Send("create a specific menu")
}

func DeleteMenu(c *fiber.Ctx){
	c.Send("delete a specific menu")
}