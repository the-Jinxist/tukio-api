package registration

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/the-Jinxist/tukio-api/internal/mailer"
	twofa "github.com/the-Jinxist/tukio-api/internal/two-fa"
)

var _ service = RegistrationService{}

type service interface {
	registerUser(ctx context.Context, req registerUserReq) error
}

type RegistrationService struct {
	repo repo
}

func NewService(r repo) service {
	return RegistrationService{repo: r}
}

// registerUser implements service.
func (r RegistrationService) registerUser(ctx context.Context, req registerUserReq) error {
	if exists := r.repo.emailExists(ctx, req.Email); exists {
		return errors.New("a user with this email already exists")
	}

	if exists := r.repo.phoneExists(ctx, req.PhoneNumber); exists {
		return errors.New("a user with this phone number already exists")
	}

	err := r.registerUser(ctx, req)
	if err != nil {
		return err
	}

	otp, err := twofa.GenerateOTP(req.Email)
	if err != nil {
		return err
	}

	go func() {
		if err = mailer.SendEmail(
			req.Email,
			fmt.Sprintf("OTP code for registration: %s", otp),
			"", nil,
		); err != nil {
			log.Printf("mailer error: %v", err)
		}

	}()

	return nil

}
