package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/service"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	productRepository := repository.NewProductRepository(dbConnection)

	productService := service.NewProductUseCase(productRepository)

	ProductController := controller.NewProductController(productService)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.POST("/products", ProductController.CreateProduct)

	server.GET("/products/:id", ProductController.GetProductById)

	server.Run(":8000")

}
