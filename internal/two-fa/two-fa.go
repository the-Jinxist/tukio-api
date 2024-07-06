package twofa

import (
	"encoding/base32"

	"github.com/pquerna/otp/hotp"
)

var otpSecret = "pwmfimwpofmowefwmfw"

func GenerateOTP(email string) (string, error) {
	// key, err := totp.Generate(totp.GenerateOpts{
	// 	AccountName: email,
	// 	Issuer:      "Tukio",
	// 	Secret:      []byte(otpSecret),

	// })

	// if key != nil {
	// 	digits := key.Digits().String()
	// 	_ = digits
	// 	return key.String(), err
	// }

	// return "", err
	secretStr := base32.StdEncoding.EncodeToString([]byte(otpSecret))
	return hotp.GenerateCode(secretStr, uint64(6))
}

func ValdateOTP(code string) bool {
	// totp := gotp.NewDefaultTOTP(otpSecret)
	secretStr := base32.StdEncoding.EncodeToString([]byte(otpSecret))
	return hotp.Validate(code, uint64(6), secretStr)

	// return totp.VerifyTime(code, time.Now())
}
