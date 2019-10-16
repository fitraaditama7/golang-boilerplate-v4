package authentication

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	key := []byte(viper.GetString(`secret`))
	claims["authorization"] = true
	claims["client"] = "Fitra Aditama"
	claims["exp"] = time.Now().Add(time.Minute * 525600).Unix()

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
