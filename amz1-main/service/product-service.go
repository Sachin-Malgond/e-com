package service

import (
	"fmt"

	"amz1-main/config"
	"amz1-main/model"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllAlbums - fetch all Albums
func GetAllProducts(product *[]model.Product) (err error) {
	if err = config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

//CreateAlbum - creates an album
func CreateProduct(product *model.Product) (err error) {
	if err = config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

//GetAlbum fetch one alb um
func GetProduct(product *model.Product, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

//UpdateAlbum - modifies an album
func UpdateProduct(product *model.Product, id string) (err error) {
	fmt.Println(product)
	config.DB.Save(product)
	return nil
}

//DeleteAlbum deletes a given album name
func DeleteProduct(product *model.Product, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(product)
	return nil
}
