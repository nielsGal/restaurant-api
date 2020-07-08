package views

import (
	"github.com/gofiber/fiber"
	"github.com/nielsGal/restaurant-api/database"
)

func GetMenu(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn
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
		c.Status(500).Send("there was an issue creating the menu")
	}
	c.JSON(menu)
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