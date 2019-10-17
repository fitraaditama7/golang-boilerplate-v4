package middleware

import (
	"golang-websocket/api/helper"
	"golang-websocket/api/helper/authentication"

	"github.com/gin-gonic/gin"
)

func MiddlewareAuthentication(c *gin.Context) {
	var req = c.Request
	var res = c.Writer
	var isAuth = authentication.RequiredAuthorization(req)
	var isToken = authentication.RequiredToken(req)
	if isAuth != "success" {
		helper.ErrorCustomStatus(res, 401, isAuth)
		c.Abort()
		return
	}

	if isToken != "success" {
		helper.ErrorCustomStatus(res, 401, isToken)
		c.Abort()
		return
	}
	c.Next()
}
