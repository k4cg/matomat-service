package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

type AuthJWT struct {
	tokenSigningSecret          string
	defaultTokenValiditySeconds uint32
	maximumTokenValiditySeconds uint32
	tokenIssuer                 string
}

func NewAuthJWT(tokenIssuer string, tokenSigningSecret string, defaultTokenValiditySeconds uint32, maximumTokenValiditySeconds uint32) *AuthJWT {
	return &AuthJWT{tokenIssuer: tokenIssuer, tokenSigningSecret: tokenSigningSecret, defaultTokenValiditySeconds: defaultTokenValiditySeconds, maximumTokenValiditySeconds: maximumTokenValiditySeconds}
}

func (a *AuthJWT) NewAuthTokenClaims(userID uint32, userName string, requestedTokenValiditySeconds uint32) AuthTokenClaims {
	tokenValiditySeconds := a.defaultTokenValiditySeconds
	if requestedTokenValiditySeconds > 0 && requestedTokenValiditySeconds <= a.maximumTokenValiditySeconds {
		tokenValiditySeconds = requestedTokenValiditySeconds
	}

	claims := AuthTokenClaims{
		userID,
		userName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(tokenValiditySeconds)).Unix(),
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
