package auth

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(secret, userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(30 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(secret, tokenStr string) (string, error) {
	t, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !t.Valid { return "", err }
	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		if sub, ok := claims["sub"].(string); ok { return sub, nil }
	}
	return "", jwt.ErrTokenInvalidClaims
}
