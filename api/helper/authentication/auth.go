package authentication

import (
	"fmt"
	"golang-websocket/api/helper"
	"golang-websocket/api/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var allowClientAccessAPI = []string{"X-MAHASISWA-API"}

func GenerateToken(payload interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	key := []byte(viper.GetString(`secret`))
	data := payload.(*models.User)

	claims["authorization"] = true
	claims["client"] = data.Nama
	claims["iss"] = data.ID
	claims["exp"] = time.Now().Add(time.Minute * 525600).Unix()
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func RequiredAuthorization(r *http.Request) string {

	if r.Header["Authorization"] == nil {
		return "Unauthorized, need access token to access this API route"
	}
	auth := r.Header["Authorization"][0]
	isAuth := helper.Include(allowClientAccessAPI, auth)

	if isAuth == false {
		return "Unauthorized, authorization value is invalid/expired"
	}

	return "success"
}

func RequiredToken(r *http.Request) string {
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

func SetPassword(s string) (string, error) {
	password := []byte(s)

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashPassword string, plainPassword string) bool {
	byteHash := []byte(hashPassword)
	bytePass := []byte(plainPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)
	if err != nil {
		return false
	}
	return true
}
