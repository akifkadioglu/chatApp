package models

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	ID        uint   `json:"id"`
	Image     string `json:"image"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Biography string `json:"biography"`
	Time      string `json:"time"`
	jwt.StandardClaims
}
