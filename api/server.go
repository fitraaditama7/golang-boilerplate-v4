package api

import (
	"golang-websocket/api/database"
	"golang-websocket/api/routes"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	db := database.Load()
	router := gin.Default()
	timeout := time.Duration(viper.GetInt(`context.timeout`)) * time.Second
	var app *firebase.App

	v1 := router.Group("/v1")
	{
		routes.RouteMahasiswa(v1, db, timeout)
		routes.RouteWebsocket(v1, app)
		router.Run(":3000")
	}

}
