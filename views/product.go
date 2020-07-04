package views

import "github.com/gofiber/fiber"

func getProduct(c * fiber.Ctx){
	c.Send("get a specific product")
}

func getProducts(c *fiber.Ctx){
	c.Send("get a batch of products")
}

func createProduct(c* fiber.Ctx){
	c.Send("create a specific product")
}

func createProducts(c* fiber.Ctx){
	c.Send("create a batch of products")
}

func deleteProduct(c* fiber.Ctx){
	c.Send("delete a specifc product")
}

func deleteProducts(c* fiber.Ctx){
	c.Send("detelete a batch of products")
}