package routes

import (
	"golang-websocket/api/controllers/user"

	"github.com/gin-gonic/gin"
)

func RouterUser(route *gin.RouterGroup) {
	handlerUser := user.NewUserHandler()
	router := route.Group("/user")
	{
		router.GET("/list", handlerUser.List)
		router.GET("/detail/:id", handlerUser.Detail)
		router.POST("/insert", handlerUser.Insert)
		router.PUT("/update/:id", handlerUser.Update)
		router.DELETE("/delete/:id", handlerUser.Delete)
	}
}
