package routes

import (
	"context"
	"go-grpc-api-gateway/pkg/product/pb"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FineOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
