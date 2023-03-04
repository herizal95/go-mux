package utils

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("adk1l23lkjfa093kjadf")

type JWTClaims struct {
	Username string
	jwt.RegisteredClaims
}
