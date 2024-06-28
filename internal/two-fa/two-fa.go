package twofa

import "github.com/pquerna/otp/totp"

var otpSecret = "pwmfimwpofmowefwmfw"

func GenerateOTP(email string) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		AccountName: email,
		Secret:      []byte(otpSecret),
	})

	if key != nil {
		return key.String(), err
	}

	return "", err
}

func ValdateOTP(code string) bool {
	return totp.Validate(code, otpSecret)
}
