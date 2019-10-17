package routes

import (
	"golang-websocket/api/controllers/auth"

	"github.com/gin-gonic/gin"
)

func RouteAuth(route *gin.RouterGroup) {
	handlerAuth := auth.NewAuthHandler()
	router := route.Group("/auth")
	{
		router.POST("/login", handlerAuth.Login)
		router.POST("/register", handlerAuth.Register)
	}
}
