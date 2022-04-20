package main

import (
	"github.com/gin-gonic/gin"
	"go-restful-product-svc/pkg/config"
	"go-restful-product-svc/pkg/contoller"
	"go-restful-product-svc/pkg/db"
)

func main() {

	c, err := config.LoadConfig()

	if err != nil {
		return
	}

	r := gin.Default()
	h := db.Init(c.DBUrl)

	contoller.RegisterRoutes(r, h.DB)
	// register more routes here

	r.Run(c.Port)
}
