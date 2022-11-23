package utils

import (
	"fmt"
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

func ValidateToken(tokenString, userName string) (error, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWT_KEY, nil
	})

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		//nil secret key
		return fmt.Errorf("unexpected signing method: %v", token.Header["alg"]), false
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		log.Printf("%v %v", claims.Username, claims.RegisteredClaims.Issuer)
	}
	if err != nil {
		return err, false
	} else if claims.Username == userName {
		return nil, true
	}
	return err, false
}
