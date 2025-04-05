package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/cyansnbrst/gesture-guru/sso-service/config"
	"github.com/cyansnbrst/gesture-guru/sso-service/internal/models"
)

// New JWT token (HMAC)
func NewToken(user *models.User, cfg *config.Config) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(cfg.App.JWTTokenTTL).Unix()

	tokenString, err := token.SignedString([]byte(cfg.App.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
