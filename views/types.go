package views

import(
	"time"
)

type IDList struct{
	Ids []int `json:"Ids"`
}

type Product struct{
	Name string `json:"Name"`
	Price float32 `json:"Price"`
	Description string `json:"Description"`
	ImageUrl string `json:"ImageUrl"`
}

type ProductList struct{
	Items []Product `json:"Items"` 
}

type Category struct{
	Name string `json:"Name"`
	From *time.Time `json:"From"`
	To *time.Time `json:"To"`
	Products []Product `json:"Products"`
}  

type CategoryList struct{
	Categories []Category `json:"Categories"`
}

type Menu struct{
	Name string `json:"Name"`
	From *time.Time `json:"From"`
	To *time.Time `json:"To"`
	Categories []Category `json:"Categories"`
}

type MenuList struct{
	Menus []Menu `json:"Menus"`
}