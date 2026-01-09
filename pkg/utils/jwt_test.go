package utils

import (
	"testing"

	"admin-pro/internal/config"
)

func TestGenerateToken(t *testing.T) {
	cfg := &config.Config{
		JWT: config.JWTConfig{
			Secret: "test_secret_key_for_unit_testing",
			Expire: 24,
		},
	}

	token, err := GenerateToken("user123", "domain1", "testuser", cfg)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	if token == "" {
		t.Fatal("Generated token is empty")
	}
}

func TestParseToken(t *testing.T) {
	cfg := &config.Config{
		JWT: config.JWTConfig{
			Secret: "test_secret_key_for_unit_testing",
			Expire: 24,
		},
	}

	// Generate a token first
	token, err := GenerateToken("user123", "domain1", "testuser", cfg)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Parse the token
	claims, err := ParseToken(token, cfg)
	if err != nil {
		t.Fatalf("ParseToken failed: %v", err)
	}

	// Verify claims
	if claims.UserID != "user123" {
		t.Errorf("Expected UserID to be 'user123', got '%s'", claims.UserID)
	}
	if claims.UserDomain != "domain1" {
		t.Errorf("Expected UserDomain to be 'domain1', got '%s'", claims.UserDomain)
	}
	if claims.LoginName != "testuser" {
		t.Errorf("Expected LoginName to be 'testuser', got '%s'", claims.LoginName)
	}
}

func TestParseTokenInvalid(t *testing.T) {
	cfg := &config.Config{
		JWT: config.JWTConfig{
			Secret: "test_secret_key_for_unit_testing",
			Expire: 24,
		},
	}

	// Try to parse an invalid token
	_, err := ParseToken("invalid.token.here", cfg)
	if err == nil {
		t.Fatal("Expected error when parsing invalid token, got nil")
	}
}

func TestEncryptPassword(t *testing.T) {
	password := "testpassword123"

	hash, err := EncryptPassword(password)
	if err != nil {
		t.Fatalf("EncryptPassword failed: %v", err)
	}

	if hash == "" {
		t.Fatal("Generated hash is empty")
	}

	// Bcrypt hash should start with $2a$, $2b$, or $2y$
	if hash[0:4] != "$2a$" && hash[0:4] != "$2b$" && hash[0:4] != "$2y$" {
		t.Errorf("Hash doesn't look like a bcrypt hash: %s", hash[0:4])
	}
}

func TestCheckPassword(t *testing.T) {
	password := "testpassword123"

	// Test with bcrypt hash
	hash, err := EncryptPassword(password)
	if err != nil {
		t.Fatalf("EncryptPassword failed: %v", err)
	}

	if !CheckPassword(password, hash) {
		t.Error("CheckPassword failed to verify correct password")
	}

	if CheckPassword("wrongpassword", hash) {
		t.Error("CheckPassword incorrectly verified wrong password")
	}
}

func TestCheckPasswordLegacy(t *testing.T) {
	// Test legacy SHA256 hash for backward compatibility
	password := "testpassword123"
	hash := "b55c8792d1ce458e279308835f8a97b580263503e76e1998e279703e35ad0c2e" // SHA256 of "testpassword123"

	if !CheckPassword(password, hash) {
		t.Error("CheckPassword failed to verify legacy SHA256 hash")
	}
}
