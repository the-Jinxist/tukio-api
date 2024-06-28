package registration

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gofrs/uuid"
)

var _ repo = RegistrationRepo{}

type repo interface {
	registerUser(ctx context.Context, req registerUserReq) error
	emailExists(ctx context.Context, email string) bool
	phoneExists(ctx context.Context, phoneNumber string) bool
}

type RegistrationRepo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) RegistrationRepo {
	return RegistrationRepo{
		db: db,
	}
}

// registerUser implements service.
func (r RegistrationRepo) registerUser(ctx context.Context, req registerUserReq) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	userID := uuid.Must(uuid.NewV7())
	err = tx.QueryRow(`insert into users (id, email, created_at) values ($1, $2, now()) returning id`, userID, req.Email).Scan(&userID)
	if err != nil {
		return err
	}

	pID := uuid.Must(uuid.NewV7())
	_, err = tx.Exec(`insert into profile (id, user_id, first_name, last_name, phone_number, created_at, updated_at)
		values ($1, $2, $3, $4, $5, now(), now())`, pID, userID, req.FirstName, req.LastName, req.PhoneNumber)
	if err != nil {
		return err
	}

	return tx.Commit()

}

// emailExists implements repo.
func (r RegistrationRepo) emailExists(ctx context.Context, email string) bool {
	_, err := r.db.Query("select * from users where email = $1", email)
	return !errors.Is(err, sql.ErrNoRows)

}

func (r RegistrationRepo) phoneExists(ctx context.Context, phoneNumber string) bool {
	_, err := r.db.Query("select * from profiles where phone_number = $1", phoneNumber)
	return !errors.Is(err, sql.ErrNoRows)

}
