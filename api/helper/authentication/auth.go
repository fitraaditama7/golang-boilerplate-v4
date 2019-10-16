package authentication

import (
	"golang-websocket/api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

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
