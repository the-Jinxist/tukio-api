package me

import (
	"context"

	"github.com/jmoiron/sqlx"
)

var _ repo = ProfileRepo{}

type repo interface {
	get(ctx context.Context, uid string) (Profile, error)
	getUserProfile(ctx context.Context, uid string) (Profile, error)

	update(ctx context.Context, uid string, req updateProfileReq) error
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

// update implements repo.
func (p ProfileRepo) update(ctx context.Context, uid string, req updateProfileReq) error {
	_, err := p.db.ExecContext(ctx, "update profiles set first_name = $1, last_name = $2, phone_number = $3 where user_id = $4",
		req.FirstName, req.LastName, req.PhoneNumber)
	return err
}

// getUserProfile implements repo.
func (p ProfileRepo) getUserProfile(ctx context.Context, uid string) (Profile, error) {
	var profile Profile
	err := p.db.GetContext(ctx, &profile, "select * from profiles where user_id = $1", uid)
	return profile, err
}
