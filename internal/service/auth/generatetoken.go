package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	tokenTTL     = time.Hour * 12
	signatureKey = "NOi8B64Gtt3E8LoBY7T1bs"
)

func (s *AuthService) GenerateToken(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.GetUser(ctx, username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(tokenTTL).Unix(),
		"iat":     time.Now().Unix(),
		"user_id": user.ID,
	})

	return token.SignedString([]byte(signatureKey))
}
