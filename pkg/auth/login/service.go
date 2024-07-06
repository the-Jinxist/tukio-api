package login

import (
	"context"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/the-Jinxist/tukio-api/internal/token"
)

var _ service = LoginService{}

type service interface {
	login(ctx context.Context, req loginrReq) (string, error)
}

type LoginService struct {
	repo repo
}

func NewService(r repo) service {
	return LoginService{repo: r}
}

// login implements service.
func (l LoginService) login(ctx context.Context, req loginrReq) (string, error) {
	user, err := l.repo.login(ctx, req)
	if err != nil {
		return "", err
	}

	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID.String(),
		"token_id": "auth",
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
	})

	str, err := tokenStr.SignedString(token.AUTH_TOKEN_SECRET)
	if err != nil {
		log.Printf("mailer error: %v", err)
	}

	return str, nil
}
