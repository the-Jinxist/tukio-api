package login

import (
	"context"
	"time"

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

	str, err := token.GenerateJwt(
		user.ID.String(), time.Now().Add(time.Hour*72), token.AUTH_TOKEN_SECRET, token.AUTH_TOKEN_TYPE)
	if err != nil {
		return "", err
	}

	return str, nil
}
