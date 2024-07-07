package token

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenSecret []byte
type TokenType string

var (
	VERIFY_TOKEN_SECRET = []byte("verify_token")
	AUTH_TOKEN_SECRET   = []byte("auth_token_secret")
)

var (
	VERIFY_TOKEN_TYPE = TokenType("verify_otp")
	AUTH_TOKEN_TYPE   = TokenType("auth")
)

type User struct {
}

func GenerateJwt(userID string, exp time.Time, tokenSecret []byte, tokenType TokenType) (string, error) {
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"token_id": string(tokenType),
		"exp":      exp.Unix(),
	})

	str, err := tokenStr.SignedString([]byte(tokenSecret))
	if err != nil {
		log.Printf("mailer error: %v", err)
	}

	return str, err
}
