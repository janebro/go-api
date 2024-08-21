package controller

import (
	"go-api/model"
	"go-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *service.ProductService
}

func NewProductController(usecase *service.ProductService) *ProductController {
	return &ProductController{
		ProductService: usecase,
	}
}

func (pc *ProductController) GetProducts(ctx *gin.Context) {
	products, err := pc.ProductService.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.ShouldBindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	createdProduct, err := pc.ProductService.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdProduct)
}

func (pc *ProductController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		response := model.Response{
			ErrorCode: http.StatusBadRequest,
			Message:   "Product ID could not be null or empty",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			ErrorCode: http.StatusBadRequest,
			Message:   "Product ID has to be an integer number",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := pc.ProductService.GetProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			ErrorCode: http.StatusNotFound,
			Message:   "Product with ID " + id + " not found",
		}

		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
