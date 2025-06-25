package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateJWT(userID uint, email string) (string, error) {
	jwtSecret := []byte(viper.GetString("auth_jwt_secret"))
	expiry := viper.GetDuration("auth_jwt_expiry")
	if expiry == 0 {
		expiry = 15 * time.Minute
	}
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(expiry).Unix(), // expired 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
