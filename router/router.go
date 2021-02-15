package router

import (
	"github.com/1k-ct/indecisive/api"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()

	r.GET("/", api.Introduction)

	apiGroup := r.Group("/v1")

	{
		APIRouter(apiGroup)
	}
	return r
}
