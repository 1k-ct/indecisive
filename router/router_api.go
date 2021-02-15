package router

import (
	"net/http"

	"github.com/1k-ct/indecisive/api"
	"github.com/gin-gonic/gin"
)

func APIRouter(r *gin.RouterGroup) (R gin.IRoutes) {
	router := r.Group("/api")
	{
		router.POST("/select", api.SelectOne)
	}
	return r
}
func APIRouterV2(r *gin.RouterGroup) (R gin.IRoutes) {
	router := r.Group("/api")
	{
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "pong"})
		})
	}
	return r
}
