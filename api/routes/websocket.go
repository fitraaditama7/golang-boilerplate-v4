package routes

import (
	"golang-websocket/api/websocket"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func RouteWebsocket(route *gin.RouterGroup, app *firebase.App) {
	handlerWebsocket := websocket.WebsocketHandler{App: app}
	router := route.Group("/websocket")
	{
		router.POST("/push", handlerWebsocket.SendAndroidMessage)
	}
}
