package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"admin-pro/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword 检查提供的密码是否与哈希匹配
// 兼容旧系统：如果 hash 长度为 64（SHA256），则使用旧算法验证
func CheckPassword(password, hash string) bool {
	// 兼容旧系统的 SHA256 哈希（长度为 64）
	if len(hash) == 64 {
		return checkPasswordLegacy(password, hash)
	}

	// 使用 bcrypt 验证
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// checkPasswordLegacy 兼容旧系统的 SHA256 验证
func checkPasswordLegacy(password, hash string) bool {
	h := sha256.Sum256([]byte(password))
	return hex.EncodeToString(h[:]) == hash
}
