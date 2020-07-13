package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/basicauth"
	"github.com/gofiber/template/pug"

	"github.com/nielsGal/restaurant-api/noAuthRoutes"
	"github.com/nielsGal/restaurant-api/authRoutes"
	"github.com/nielsGal/restaurant-api/database"
	"github.com/nielsGal/restaurant-api/types"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupAPI(app *fiber.App){
	app.Get("/api/v1/get-menu",authRoutes.GetMenu)
	app.Get("/api/v1/get-menus",authRoutes.GetMenus)
	app.Post("/api/v1/create-menu",authRoutes.CreateMenu)
	app.Delete("/api/v1/delete-menu",authRoutes.DeleteMenu)
	app.Post("/api/v1/delete-menus",authRoutes.DeleteMenus)

	app.Get("/api/v1/get-product",authRoutes.GetProduct)
	app.Get("/api/v1/get-products",authRoutes.GetProducts)
	app.Post("/api/v1/create-product",authRoutes.CreateProduct)
	app.Post("/api/v1/create-products",authRoutes.CreateProducts)
	app.Delete("/api/v1/delete-product",authRoutes.DeleteProduct)
	app.Post("/api/v1/delete-products",authRoutes.DeleteProducts)

	app.Get("/api/v1/get-category",authRoutes.GetCategory)
	app.Get("/api/v1/get-categories",authRoutes.GetCategories)
	app.Post("/api/v1/create-category",authRoutes.CreateCategory)
	app.Post("/api/v1/create-categories",authRoutes.CreateCategories)
	app.Delete("/api/v1/delete-category",authRoutes.DeleteCategory)
	app.Delete("/api/v1/delete-categories",authRoutes.DeleteCategories)
}

func setupAuthRoutes(app *fiber.App){
	app.Get("/admin",authRoutes.GetAdminPage)
	app.Get("/admin/restaurants",authRoutes.GetAdminRestaurantPage)
}

func setupNonAuthRoutes(app *fiber.App){
	app.Get("/",noAuthRoutes.GetIndex)
	app.Get("/login",noAuthRoutes.GetLogin)
	app.Get("/register",noAuthRoutes.GetRegister)
	app.Get("/contact",noAuthRoutes.GetContact)
	app.Post("/contact",noAuthRoutes.PostContact)
}

func InitDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3","rest.db")
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("database connection established")
	database.DBConn.AutoMigrate(&types.Product{})
	database.DBConn.AutoMigrate(&types.Category{})
	database.DBConn.AutoMigrate(&types.Menu{})
	fmt.Println("database migrated")

}

func setupAuth(app *fiber.App){
	cfg := basicauth.Config{
		Users: map[string]string{
			"john": "doe",
		},
	}
	
	app.Use(basicauth.New(cfg))
}


func main() {
	engine := pug.New("./views",".pug")
	app := fiber.New(&fiber.Settings{
		Views: engine,
	})
	app.Static("/public","./public")
	InitDatabase()
	defer database.DBConn.Close()
	setupNonAuthRoutes(app)
	setupAuth(app)
	setupAuthRoutes(app)
	app.Listen("3000")
}