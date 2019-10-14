package api

import (
	"golang-websocket/api/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := SetupRouter()
	router.Run(":3000")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		routes.RouteMahasiswa(v1)
		routes.RouterUser(v1)
	}
	return router
}
