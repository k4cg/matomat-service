package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

type AuthJWT struct {
	tokenSigningSecret   string
	tokenValiditySeconds uint32
	tokenIssuer          string
}

func NewAuthJWT(tokenIssuer string, tokenSigningSecret string, tokenValiditySeconds uint32) *AuthJWT {
	return &AuthJWT{tokenIssuer: tokenIssuer, tokenSigningSecret: tokenSigningSecret, tokenValiditySeconds: tokenValiditySeconds}
}

func (a *AuthJWT) NewAuthTokenClaims(userID uint32, userName string) AuthTokenClaims {
	claims := AuthTokenClaims{
		userID,
		userName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(a.tokenValiditySeconds)).Unix(),
			Issuer:    a.tokenIssuer,
		},
	}
	return claims
}

func (a *AuthJWT) GetToken(claims AuthTokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.tokenSigningSecret))

	return tokenString, err
}

func (a *AuthJWT) VerifyToken(tokenString string) (AuthTokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.tokenSigningSecret), nil
	})
	var authTokenClaims AuthTokenClaims
	if token.Valid {
		mapstructure.Decode(token.Claims.(jwt.MapClaims), &authTokenClaims)
	}
	return authTokenClaims, err
}
