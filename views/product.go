package views

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/nielsGal/restaurant-api/database"
)




func GetProduct(c * fiber.Ctx){
	db := database.DBConn
	id := c.Params("id")
	var product Product
	db.Find(&product,id)
	c.JSON(product)
}

func GetProducts(c *fiber.Ctx){
	db := database.DBConn
	var products []Product
	db.Find(&products)
	c.JSON(products)
}

func CreateProduct(c* fiber.Ctx){
	db := database.DBConn
	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		c.Status(422).Send("there is some error in the request")
	}
	if result := db.Create(&product); result.Error != nil {
		c.Status(500).Send("there was some error creating the product")
	}
	c.JSON(product)
}

func CreateProducts(c* fiber.Ctx){
	db := database.DBConn
	products := new(ProductList)
	if err := c.BodyParser(products); err != nil{
		fmt.Println(err)
		c.Status(422).Send("could not process request")
	}
	for _ , product := range products.Items {
		if result := db.Create(product); result.Error != nil {
			fmt.Println(result.Error)
			c.Status(500).Send("there was an error creathing these products")
		}
	}
	c.JSON(products)
}

func DeleteProduct(c* fiber.Ctx){
	db := database.DBConn
	id := c.Params("id")
	product := new(Product)
	db.First(product,id)
	if (product.Name == ""){
		c.Status(404).Send("no book with that ID")
		return
	}
	if result := db.Delete(&product); result.Error != nil {
		c.Status(500).Send("there was some issue with deleting this product")
	}
	c.JSON(product)
}

func DeleteProducts(c* fiber.Ctx){
	db := database.DBConn
	IdList := new(IDList)
	if err := c.BodyParser(IdList); err !=nil{
		c.Status(422).Send("could not process request")
	}
	if result := db.Where("ID IN (?)",IdList.Ids).Delete(&Product{}); result.Error != nil {
		c.Status(500).Send("there was some issue with deleting these ids")
	}
	c.JSON(IdList.Ids)
}
