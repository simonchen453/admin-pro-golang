package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"admin-pro/internal/config"
)

type Claims struct {
	UserID     string `json:"userId"`
	UserDomain string `json:"userDomain"`
	LoginName  string `json:"loginName"`
	jwt.RegisteredClaims
}

func GenerateToken(userID, userDomain, loginName string, cfg *config.Config) (string, error) {
	expirationTime := time.Now().Add(time.Duration(cfg.JWT.Expire) * time.Hour)
	claims := &Claims{
		UserID:     userID,
		UserDomain: userDomain,
		LoginName:  loginName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "admin-pro-golang",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}

func ParseToken(tokenString string, cfg *config.Config) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func EncryptPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// CheckPassword checks if the provided password matches the hash
// In this legacy system, simple SHA256 is used (based on sample data)
func CheckPassword(password, hash string) bool {
	return EncryptPassword(password) == hash
}
