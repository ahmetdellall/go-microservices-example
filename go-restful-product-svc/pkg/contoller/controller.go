package contoller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/product")
	routes.POST("/", h.CreateProduct)
	routes.GET("/:id", h.FindOne)
	routes.POST("/:id", h.DecreaseStock)
}
