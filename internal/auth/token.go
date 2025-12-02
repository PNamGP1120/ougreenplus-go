package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/PNamGP1120/ougreenplus-go/internal/config"
	"github.com/PNamGP1120/ougreenplus-go/internal/user"
)

func generateTokens(u *user.User, cfg *config.Config) (string, string, error) {
	now := time.Now()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  u.ID,
		"role": u.Role,
		"exp":  now.Add(15 * time.Minute).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.ID,
		"exp": now.Add(7 * 24 * time.Hour).Unix(),
	})

	at, err := accessToken.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", "", err
	}

	rt, err := refreshToken.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", "", err
	}

	return at, rt, nil
}
