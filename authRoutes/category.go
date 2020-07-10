package authRoutes

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/nielsGal/restaurant-api/database"
	"github.com/nielsGal/restaurant-api/types"
)




func GetCategory(c *fiber.Ctx){
	db := database.DBConn
	id := c.Params("id")
	var category types.Category
	db.Find(&category,id)
	c.JSON(category)
}
func GetCategories(c *fiber.Ctx){
	db := database.DBConn
	var category []types.Category
	db.Find(&category)
	c.JSON(category)
}

func CreateCategory(c *fiber.Ctx){
	db := database.DBConn
	category := new(types.Category)
	if err := c.BodyParser(category); err != nil{
		c.Status(422).Send("there is some error in the request")
	}

	if result := db.Create(&category); result.Error != nil {
		fmt.Println(result.Error)
		c.Status(500).Send("there was some error creating the category")
	}
	c.JSON(category)
}

func CreateCategories(c *fiber.Ctx){
	db := database.DBConn
	categories := new(types.CategoryList)
	if err := c.BodyParser(categories); err != nil {
		c.Status(422).Send("there is some error in the request")
	}	
	for _ , category := range categories.Categories {
		if result := db.Create(&category); result.Error != nil {
			c.Status(500).Send("there was an issue creating these categories")
		}
	}
	c.JSON(categories)
}

func DeleteCategory(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn
	category := new(types.Category)
	db.First(category,id)
	if (category.Name == "") {
		c.Status(404).Send("no such category exists")
	}
	if result := db.Delete(&category); result.Error != nil{
		c.Status(500).Send("error in deleting the category")
	}
	c.Send(category)
}

func DeleteCategories(c *fiber.Ctx){
	db := database.DBConn
	IdList := new(types.IDList)
	if err := c.BodyParser(IdList); err !=nil{
		c.Status(422).Send("there was some error in the request")
	}
	db.Where("ID IN (?)",IdList.Ids).Delete(&types.Category{})
}