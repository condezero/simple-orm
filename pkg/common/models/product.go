package models

type Product struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Stock int32  `json:"stock"`
	PRice int32  `json:"price"`
}
