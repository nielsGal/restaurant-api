package views

import "github.com/gofiber/fiber"

func GetProduct(c * fiber.Ctx){
	c.Send("get a specific product")
}

func GetProducts(c *fiber.Ctx){
	c.Send("get a batch of products")
}

func CreateProduct(c* fiber.Ctx){
	c.Send("create a specific product")
}

func CreateProducts(c* fiber.Ctx){
	c.Send("create a batch of products")
}

func DeleteProduct(c* fiber.Ctx){
	c.Send("delete a specifc product")
}

func DeleteProducts(c* fiber.Ctx){
	c.Send("delete a batch of products")
}