package rest 

import (
	"amz1-main/controller"
	"github.com/gin-gonic/gin"
)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine{
	r:=gin.Default()
	r.Use(CORSMiddleware())
	v1:=r.Group("/v1")
	{
		v1.GET("product",controller.GetProducts)
		v1.POST("product",controller.CreateProduct)
v1.GET("product/:id",controller.GetProduct)
v1.PUT("product/:id", controller.UpdateProduct)
		v1.DELETE("product/:id", controller.DeleteProduct)

	}
		return r
}