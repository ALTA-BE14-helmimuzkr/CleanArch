package helper

import (
	"api/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func ExtractToken(t interface{}) int {
	user := t.(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userID"].(float64)
		return int(userId)
	}
	return -1
}

func GenerateToken(userID uint) string {
	// Create claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	// Create Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, _ := token.SignedString([]byte(config.JWT_KEY))

	return strToken
}

func ValidateToken(strToken string) *jwt.Token {
	// Decode rawToken, parse from rawToken to jwt.Token
	token, _ := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_KEY), nil
	})

	return token
}
