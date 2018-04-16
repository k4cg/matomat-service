package auth

type AuthInterface interface {
	GetToken(AuthTokenClaims) (string, error)
	VerifyToken(tokenString string) (AuthTokenClaims, error)
	NewAuthTokenClaims(userID uint32, userName string) AuthTokenClaims
}
