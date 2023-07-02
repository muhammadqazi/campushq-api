package utils

import (
	"fmt"

	extract "github.com/golang-jwt/jwt"
)

func Extractor(tokenString string) extract.MapClaims {
	token, _, err := new(extract.Parser).ParseUnverified(tokenString, extract.MapClaims{})
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	if claims, ok := token.Claims.(extract.MapClaims); ok {

		return claims
	}

	return nil
}
