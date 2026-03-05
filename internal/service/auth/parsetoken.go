package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *AuthService) ParseToken(token string) (int, error) {
	userToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signatureKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := userToken.Claims.(*Claims)
	if !ok || !userToken.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	return claims.UserID, nil
}
