package login

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var _ repo = LoginRepo{}

type repo interface {
	login(ctx context.Context, req loginrReq) (User, error)
}

type LoginRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) repo {
	return LoginRepo{
		db: db,
	}
}

// registerUser implements repo.
func (l LoginRepo) login(ctx context.Context, req loginrReq) (User, error) {
	var exists bool
	res, err := l.db.QueryContext(ctx, "select exists (select * from users where email = $1)", req.Email)
	if err != nil {
		return User{}, err
	}

	res.Next()
	defer res.Close()
	err = res.Scan(&exists)
	if err != nil {
		return User{}, err
	}

	if !exists {
		return User{}, errors.New("no user with this email exists")
	}

	var user User
	err = l.db.GetContext(ctx, &user, "select * from users where email = $1", req.Email)
	if err != nil {
		return User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return User{}, errors.New("email or password is incorrect")
	}

	return user, nil
}
