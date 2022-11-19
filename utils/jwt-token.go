package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

var JWT_KEY = []byte("SECRET_JWT_KEY_HERE")

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	// Create claims while leaving out some optional fields
	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(JWT_KEY)
	log.Printf("GenerateToken: %v %v %v", ss, err, username)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWT_KEY, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//fmt.Printf("%v %v", claims.user, claims.RegisteredClaims.Issuer)
		return claims.Username, nil
	} else {
		return "", err
	}

}
