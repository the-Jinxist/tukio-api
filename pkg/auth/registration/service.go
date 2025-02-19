package registration

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/the-Jinxist/tukio-api/internal/mailer"
	"github.com/the-Jinxist/tukio-api/internal/token"
	twofa "github.com/the-Jinxist/tukio-api/internal/two-fa"
	"github.com/the-Jinxist/tukio-api/middleware"
)

var _ service = RegistrationService{}

type service interface {
	registerUser(ctx context.Context, req registerUserReq) (string, error)
	verifyCode(ctx context.Context, req verifyCodeReq) (string, error)
}

type RegistrationService struct {
	repo repo
}

func NewService(r repo) service {
	return RegistrationService{repo: r}
}

// registerUser implements service.
func (r RegistrationService) registerUser(ctx context.Context, req registerUserReq) (string, error) {
	if exists := r.repo.emailExists(ctx, req.Email); exists {
		return "", errors.New("a user with this email already exists")
	}

	if exists := r.repo.phoneExists(ctx, req.PhoneNumber); exists {
		return "", errors.New("a user with this phone number already exists")
	}

	userID, err := r.repo.registerUser(ctx, req)
	if err != nil {
		return "", err
	}

	str, err := token.GenerateJwt(userID, time.Now().Add(time.Minute*5), token.VERIFY_TOKEN_SECRET, token.VERIFY_TOKEN_TYPE)
	if err != nil {
		log.Printf("mailer error: %v", err)
	}

	go func() {
		otp, err := twofa.GenerateOTP(req.Email)
		if err != nil {
			log.Printf("totp error: %v", err)
			return
		}
		log.Printf("OTP code for registration: %s", otp)

		_ = mailer.SendEmail(
			req.Email,
			fmt.Sprintf("OTP code for registration: %s", otp),
			"", nil,
		)

	}()

	return str, nil

}

// verifyCode implements service.
func (r RegistrationService) verifyCode(ctx context.Context, req verifyCodeReq) (string, error) {

	verify := twofa.ValdateOTP(req.Code)
	if !verify {
		return "", errors.New("invalid otp code")
	}

	userID := middleware.GetUserID(ctx)
	err := r.repo.verifyUser(ctx, userID)
	if err != nil {
		return "", err
	}

	// create user and profile model

	// add code to get profile and user

	// load values into jwt claims and send to user

	str, err := token.GenerateJwt(userID, time.Now().Add(time.Hour*72), token.AUTH_TOKEN_SECRET, token.AUTH_TOKEN_TYPE)
	if err != nil {
		log.Printf("mailer error: %v", err)
	}

	return str, nil

}
