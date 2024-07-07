package me

import (
	"context"

	"github.com/the-Jinxist/tukio-api/middleware"
)

var _ service = ProfileService{}

type service interface {
	get(ctx context.Context) (Profile, error)
}

type ProfileService struct {
	repo repo
}

func NewService(r repo) service {
	return ProfileService{repo: r}
}

// get implements service.
func (p ProfileService) get(ctx context.Context) (Profile, error) {

	uid := middleware.GetUserID(ctx)
	profile, err := p.repo.get(ctx, uid)
	return profile, err
}
