package registration

import (
	"context"
	"database/sql"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var _ repo = RegistrationRepo{}

type repo interface {
	registerUser(ctx context.Context, req registerUserReq) (string, error)
	verifyUser(ctx context.Context, uid string) error
	emailExists(ctx context.Context, email string) bool
	phoneExists(ctx context.Context, phoneNumber string) bool
}

type RegistrationRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) repo {
	return RegistrationRepo{
		db: db,
	}
}

// registerUser implements service.
func (r RegistrationRepo) registerUser(ctx context.Context, req registerUserReq) (string, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return "", err
	}

	defer tx.Rollback()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return "", err
	}

	userID := uuid.Must(uuid.NewV7())
	err = tx.QueryRow(`insert into users (id, email, password, verified, created_at) values ($1, $2, $3, false, now()) returning id`,
		userID, req.Email, hashedPassword).Scan(&userID)
	if err != nil {
		return "", err
	}

	pID := uuid.Must(uuid.NewV7())
	_, err = tx.Exec(`insert into profiles (id, user_id, first_name, last_name, phone_number, created_at, updated_at)
		values ($1, $2, $3, $4, $5, now(), now())`,
		pID, userID, req.FirstName, req.LastName, req.PhoneNumber)
	if err != nil {
		return "", err
	}

	return userID.String(), tx.Commit()

}

// emailExists implements repo.
func (r RegistrationRepo) emailExists(ctx context.Context, email string) bool {
	var exists bool
	res, err := r.db.QueryContext(ctx, "select exists (select * from users where email = $1)", email)
	if err != nil {
		return true
	}

	res.Next()
	defer res.Close()
	err = res.Scan(&exists)
	if err != nil {
		return true
	}

	return exists

}

func (r RegistrationRepo) phoneExists(ctx context.Context, phoneNumber string) bool {
	var exists bool
	res, err := r.db.QueryContext(ctx, "select exists (select * from profiles where phone_number = $1)", phoneNumber)

	if err != nil {
		return true
	}

	res.Next()
	defer res.Close()

	err = res.Scan(&exists)
	if err != nil {
		return true
	}

	return exists

}

// verifyUser implements repo.
func (r RegistrationRepo) verifyUser(ctx context.Context, uid string) error {
	_, err := r.db.ExecContext(ctx, "update users set verified = $1 where id = $2", true, uid)
	return err
}
