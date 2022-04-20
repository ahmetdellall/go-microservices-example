package main

import (
	"github.com/gin-gonic/gin"
	"go-grpc-api-gateway/pkg/auth"
	"go-grpc-api-gateway/pkg/config"
	"go-grpc-api-gateway/pkg/order"
	"go-grpc-api-gateway/pkg/product"
	"log"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	authSvc := *auth.RegisterRoutes(r, &c)

	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)
	r.Run(c.Port)
}
