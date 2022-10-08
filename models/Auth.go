package models

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Username string `json:"username"`
	ID       uint   `json:"id"`
	Time     string `json:"time"`
	jwt.StandardClaims
}