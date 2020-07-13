package types

import(
	"time"
)

type IDList struct{
	Ids []int `json:"Ids"`
}

type Product struct{
	ID uint `json:"ID" form:"ID" gorm:"primary_key"`
	Name string `json:"Name" form:"Name"`
	Price float32 `json:"Price" form:"Price"`
	Description string `json:"Description" form:"Description"`
	ImageUrl string `json:"ImageUrl" form:"ImageUrl"`
}

type ProductList struct{
	Items []Product `json:"Items" form:"Items"` 
}

type Category struct{
	ID uint `json:"ID" form:"ID" gorm:"primary_key"`
	Name string `json:"Name" form:"Name"`
	From *time.Time `json:"From" form:"From"`
	To *time.Time `json:"To" form:"To"`
	Products []Product `json:"Products" form:"Products"`
}  

type CategoryList struct{
	Categories []Category `json:"Categories" form:"Categories"`
}

type Menu struct{
	ID uint `json:"ID" form:"ID" gorm:"primary_key"`
	Name string `json:"Name" form:"Name"`
	From *time.Time `json:"From" form:"From"`
	To *time.Time `json:"To" form:"To"`
	Categories []Category `json:"Categories" form:"Categories"`
}

type MenuList struct{
	Menus []Menu `json:"Menus" form:"Menus"`
}

type Restaurant struct {
	Adress string `json:"Adress"`
	Categories []Category `json:"Categories" form:"Categories"`
	Products []Product `json:"Products" form:"Products"`
	Menus []Menu `json:"Menus" form:"Menus"`
	ActiveMenu Menu `json:"ActiveMenu" form:"ActiveMenu"`
	Name string `json:"Name" form:"Name"`
	ImageUrl string `json:"ImageUrl" form:"ImageUrl"`
}

type ClientProfile struct {
	Restaurant Restaurant  `json:"Restaurant"`
	Adress string `json:"Adress"`
}

type PrintJob  struct {
	MacSdress string `json:"MacAdress" form:"MacAdress"`
	PrintJob string `json:"PrintJob" form:"PrintJob"`
}

type User struct{
	UserName string `grom:"primary_key"`
	Salt string 
	Hash string
	Email string 
	IsAdmin bool
	Profile ClientProfile
}

type ContactRequest struct {
	ID uint `json:"ID" form:"ID" gorm:"primary_key"`
	Name string `json:"Name" form:"Name"`
	Email string `json:"Email" form:"Email"`
	Message string	`json:"Message" form:"Message"`
}