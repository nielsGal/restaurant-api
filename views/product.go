package views

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/nielsGal/restaurant-api/database"
)

type Product struct{
	gorm.Model
	Name string `json:"Name"`
	Price float32 `json:"Price"`
	Description string `json:"Description"`
	ImageUrl string `json:"ImageUrl"`
}

type ProductList struct{
	Items []Product `json:"Items"` 
}

func GetProduct(c * fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn
	var product Product
	db.Find(&product,id)
	c.JSON(product)
}

func GetProducts(c *fiber.Ctx){
	db := database.DBConn
	var product []Product
	db.Find(&product)
	c.JSON(product)
}

func CreateProduct(c* fiber.Ctx){
	db := database.DBConn
	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		c.Status(422).Send(err)
	}
	db.Create(&product)
	c.JSON(product)
}

func CreateProducts(c* fiber.Ctx){
	products := new(ProductList)
	if err := c.BodyParser(products); err != nil{
		c.Status(422).Send("could not process request")
		fmt.Println(err)
	}
	c.JSON(products)
}

func DeleteProduct(c* fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn
	product := new(Product)
	db.First(product,id)
	if (product.Name == ""){
		c.Status(404).Send("no book with that ID")
		return
	}
	db.Delete(&product)
	c.Send("deleted product")
}

func DeleteProducts(c* fiber.Ctx){
	type IDList struct{
		Ids []int `json:"Ids"`
	}
	IdList := new(IDList)
	if err := c.BodyParser(IdList); err !=nil{
		c.Status(422).Send("could not process requet")
	}
	c.JSON(IdList)
}