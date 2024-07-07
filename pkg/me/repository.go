package me

import (
	"context"

	"github.com/jmoiron/sqlx"
)

var _ repo = ProfileRepo{}

type repo interface {
	get(ctx context.Context, uid string) (Profile, error)
}

type ProfileRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) repo {
	return ProfileRepo{
		db: db,
	}
}

// get implements repo.
func (p ProfileRepo) get(ctx context.Context, uid string) (Profile, error) {
	var profile Profile
	err := p.db.GetContext(ctx, &profile, "select * from profiles where user_id = $1", uid)
	return profile, err

}
