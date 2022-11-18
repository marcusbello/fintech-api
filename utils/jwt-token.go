package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const JWTKEY = "YOUR_JWT_KEY_HERE"

type MyCustomClaims struct {
	user string `json:"user"`
	jwt.RegisteredClaims
}

func GenerateToken(user string) (string, error) {
	// Create claims while leaving out some of the optional fields
	claims := MyCustomClaims{
		user,
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(JWTKEY)
	fmt.Printf("%v %v", ss, err)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTKEY, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//fmt.Printf("%v %v", claims.user, claims.RegisteredClaims.Issuer)
		return claims.user, nil
	} else {
		return "", err
	}

}
