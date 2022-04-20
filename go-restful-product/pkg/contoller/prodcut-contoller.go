package contoller

import (
	"github.com/gin-gonic/gin"
	"go-restful-product-svc/pkg/models"
	"net/http"
)

type AddProductRequestBody struct {
	Id                             int64                          `json:"id" gorm:"primaryKey"`
	Name                           string                         `json:"name"`
	Stock                          int64                          `json:"stock"`
	Price                          int64                          `json:"price"`
	AddStockDecreaseLogRequestBody AddStockDecreaseLogRequestBody `gorm:"foreignKey:ProductRefer"`
}

type AddStockDecreaseLogRequestBody struct {
	Id           int64 `json:"id" gorm:"primaryKey"`
	OrderId      int64 `json:"order_id"`
	ProductRefer int64 `json:"product_id"`
}

func (h handler) CreateProduct(c *gin.Context) {
	body := AddProductRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var product models.Product

	product.Name = body.Name
	product.Stock = body.Stock
	product.Id = body.Id
	product.Price = body.Price

	if result := h.DB.Create(&product); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusCreated, &product)

}

func (h handler) FindOne(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if result := h.DB.First(&product, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}
	c.JSON(http.StatusOK, &product)

}

func (h handler) DecreaseStock(c *gin.Context) {

	body := AddProductRequestBody{}

	var product models.Product

	if result := h.DB.First(&product, c.Param("id")); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if product.Stock <= 0 {
		c.JSON(http.StatusConflict, "Stock already decreased")
		return
	}
	product.Stock = body.Stock - 1

	h.DB.Save(&product)

	c.JSON(http.StatusOK, "")

}
