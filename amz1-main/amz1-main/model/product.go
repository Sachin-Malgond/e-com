package model 

import (
	
)

type Product struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`

}

func (product *Product) TableName() string {
	return "product"
}