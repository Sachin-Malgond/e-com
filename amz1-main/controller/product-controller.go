package controller

import (
	"log"
	"net/http"
	"amz1-main/model"
	"amz1-main/service"
	"github.com/gin-gonic/gin"
)



func GetProducts(c *gin.Context) {
	var products []model.Product
	err := service.GetAllProducts(&products)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

func CreateProduct(c *gin.Context) {
	var product model.Product

	c.BindJSON(&product)
	log.Printf(" Prod to be Created %v",product)
	err := service.CreateProduct(&product)
	if err != nil {
		log.Println("Error", err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		log.Printf("Product Created %v", product)
		c.JSON(http.StatusOK, product)
		
	}
}

func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product model.Product
	err := service.GetProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(c *gin.Context) {
	var product model.Product
	id := c.Params.ByName("id")
	err := service.GetProduct(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = service.UpdateProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}


func DeleteProduct(c *gin.Context) {
	var product model.Product
	id := c.Params.ByName("id")
	err := service.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
