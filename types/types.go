package types

import(
	"time"
)

type IDList struct{
	Ids []int `json:"Ids"`
}

type Product struct{
	ID uint `json:"ID" gorm:"primary_key"`
	Name string `json:"Name"`
	Price float32 `json:"Price"`
	Description string `json:"Description"`
	ImageUrl string `json:"ImageUrl"`
}

type ProductList struct{
	Items []Product `json:"Items"` 
}

type Category struct{
	ID uint `json:"ID" gorm:"primary_key"`
	Name string `json:"Name"`
	From *time.Time `json:"From"`
	To *time.Time `json:"To"`
	Products []Product `json:"Products"`
}  

type CategoryList struct{
	Categories []Category `json:"Categories"`
}

type Menu struct{
	ID uint `json:"ID" gorm:"primary_key"`
	Name string `json:"Name"`
	From *time.Time `json:"From"`
	To *time.Time `json:"To"`
	Categories []Category `json:"Categories"`
}

type MenuList struct{
	Menus []Menu `json:"Menus"`
}

type ClientProfile struct {
	Categories []Category `json:"Categories"`
	Products []Product `json:"Products"`
	Menus []Menu `json:"Menus"`
	ActiveMenu Menu `json:"ActiveMenu"`
	Name string `json:"Name"`
	ImageUrl string `json:"ImageUrl"`
}

type User struct{
	UserName string `grom:"primary_key"`
	Salt string 
	Hash string
	Email string 
	Profile ClientProfile
}

type ContactRequest struct {
	ID uint `json:"ID" gorm:"primary_key"`
	Name string `json:"Name"`
	Email string `json:"Email"`
	Message string	`json:"Message"`
}