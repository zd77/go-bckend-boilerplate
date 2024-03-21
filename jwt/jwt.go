package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken() (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := jwt.MapClaims{
		"Username":  "Dereck",
		"Useremail": "zavala@gmail.com",
		"ExpiresAt": expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_SECRET"))
	ss, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	fmt.Println("ss: ", ss)
	return ss, nil
}

func VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}
	if !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}
