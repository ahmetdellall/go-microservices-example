package contoller

import "github.com/gin-gonic/gin"

type AddOrderRequestBody struct {
	Id        int64 `json:"id" gorm:"primaryKey"`
	Price     int64 `json:"price"`
	ProductId int64 `json:"product_id"`
	UserId    int64 `json:"user_id"`
}

func (h handler) CreateOrder(c *gin.Context) {

}
