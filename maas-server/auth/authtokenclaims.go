package auth

import jwt "github.com/dgrijalva/jwt-go"

type AuthTokenClaims struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
