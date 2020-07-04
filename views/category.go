package views

import "github.com/gofiber/fiber"

func GetCategory(c *fiber.Ctx){
	c.Send("get a specific category")
}
func GetCategories(c *fiber.Ctx){
	c.Send("get all categories for a specific menu")
}

func CreateCategory(c *fiber.Ctx){
	c.Send("create a specific category")
}

func CreateCategories(c *fiber.Ctx){
	c.Send("create a batch of categories")
}

func DeleteCategory(c *fiber.Ctx){
	c.Send("delete a specific category")
}

func DeleteCategories(c *fiber.Ctx){
	c.Send("delete a batch of categories")
}