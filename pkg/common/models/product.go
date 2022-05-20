package models

// swagger:model
type Product struct {
	// the id of the product
	Id int `json:"id" gorm:"primaryKey"`
	// the name of the product
	Name string `json:"name"`
	// the stock of the product
	Stock int32 `json:"stock"`
	// the price of the product
	Price int32 `json:"price"`
}
