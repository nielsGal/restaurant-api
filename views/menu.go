package views

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/nielsGal/restaurant-api/database"
)

func GetMenu(c *fiber.Ctx){
	db := database.DBConn
	id := c.Params("id")
	menu := new(Menu)
	db.First(menu,id)
	c.JSON(menu)
}

func GetMenus(c *fiber.Ctx){
	db:= database.DBConn
	var menus []Menu
	db.Find(&menus)
	c.JSON(menus)
}

func CreateMenu(c *fiber.Ctx){
	db := database.DBConn
	menu := new(Menu)
	if err := c.BodyParser(menu); err != nil {
		c.Status(422).Send("there was a problem with the request")
	}
	if result := db.Create(menu); result.Error != nil {
		fmt.Println(result.Error)
		c.Status(500).Send("there was an issue creating the menu")
	}
	c.JSON(menu)
}

func CreateMenus(c* fiber.Ctx){
	db := database.DBConn
	menus := new(MenuList)
	if err := c.BodyParser(menus); err != nil{
		fmt.Println(err)
		c.Status(422).Send("could not process request")
	}
	for _ ,menu := range menus.Menus {
		if result := db.Create(menu); result.Error != nil {
			fmt.Println(result.Error)
			c.Status(500).Send("there was an issue creating the menus")
		}
	}
	c.JSON(menus)
}

func DeleteMenu(c *fiber.Ctx){
	db := database.DBConn
	id := c.Params("id")
	menu := new(Menu)
	db.First(menu,id)
	if (menu.Name == ""){
		c.Status(404).Send("there is no menu like this")
	}
	if result := db.Delete(menu); result.Error !=  nil {
		c.Status(500).Send("there is an error deleting this menu")
	}
	c.JSON(menu)
}

func DeleteMenus(c *fiber.Ctx){
	db := database.DBConn
	IdList := new(IDList)
	if err := c.BodyParser(IdList); err !=nil{
		c.Status(422).Send("there was a problem with the request")
	}
	if result := db.Where("ID IN (?)",IdList.Ids).Delete(&Menu{}); result.Error != nil {
		c.Status(500).Send("there was some issue with deleting these ids")
	}
	c.JSON(IdList.Ids)
}