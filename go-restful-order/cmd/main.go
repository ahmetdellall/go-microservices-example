package main

import (
	"github.com/gin-gonic/gin"
	"go-restful-order-svc/pkg/config"
	contoller "go-restful-order-svc/pkg/controller"
	"go-restful-order-svc/pkg/db"
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
