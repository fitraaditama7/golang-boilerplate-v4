package middleware

import (
	"fmt"
	"golang-websocket/api/helper"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var allowClientAccessAPI = []string{"X-MAHASISWA-API"}

func requiredAuthorization(r *http.Request) string {
	auth := r.Header["Authorization"][0]
	if r.Header["Authorization"] == nil {
		return "Unauthorized, need access token to access this API route"
	}
	isAuth := helper.Include(allowClientAccessAPI, auth)

	if isAuth == false {
		return "Unauthorized, authorization value is invalid/expired"
	}

	return "success"
}

func requiredToken(r *http.Request) string {
	if r.Header["Token"] == nil {
		return "Unauthorized, need access token to access this API route"
	}

	token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		key := []byte(viper.GetString(`secret`))
		return key, nil
	})
	if err != nil {
		return err.Error()
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		return "success"
	}
	return "Unauthorized, need access token to access this API route"
}

func MiddlewareAuthentication(c *gin.Context) {
	var req = c.Request
	var res = c.Writer
	var isAuth = requiredAuthorization(req)
	var isToken = requiredToken(req)
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
