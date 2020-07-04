package views

import "github.com/gofiber/fiber"

func getCategory(c *fiber.Ctx){
	c.Send("get a specific category")
}
func getCategories(c *fiber.Ctx){
	c.Send("get all categories for a specific menu")
}

func createCategory(c *fiber.Ctx){
	c.Send("create a specific category")
}

func createCategories(c *fiber.Ctx){
	c.Send("create a batch of categories")
}

func deleteCategory(c *fiber.Ctx){
	c.Send("delete a specific category")
}

func deleteCategories(c *fiber.Ctx){
	c.Send("delete a batch of categories")
}