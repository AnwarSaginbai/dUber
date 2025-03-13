package token

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/AnwarSaginbai/auth-service/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GetSecretKey(secret string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return nil, fmt.Errorf("ошибка при декодировании ключа: %w", err)
	}
	return decoded, nil
}

func GenerateToken(userID int64, email, role string) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id":   userID,
		"email":     email,
		"user_role": role,
		"exp":       expirationTime.Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("my-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(cfg *config.Config, tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(cfg.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ParseToken(cfg *config.Config, tokenString string) (map[string]interface{}, error) {
	token, err := ValidateToken(cfg, tokenString)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
