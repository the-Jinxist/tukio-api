package registration

type registerUserReq struct {
	Email       string `json:"email" validate:"required,email"`
	FirstName   string `json:"first_name" validate:"required,min=2"`
	LastName    string `json:"last_name" validate:"required,min=2"`
	Password    string `json:"password" validate:"required,min=10"`
	PhoneNumber string `json:"phonenumber" validate:"required,numeric"`
}

type verifyCodeReq struct {
	Code string `json:"code" validate:"required,min=6,max=6"`
}
