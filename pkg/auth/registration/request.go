package registration

type registerUserReq struct {
	Email       string `json:"email" validate:"required,email"`
	FirstName   string `json:"first_name" validate:"required,min=2"`
	LastName    string `json:"last_name" validate:"required,min=2"`
	PhoneNumber string `json:"phonenumber" validate:"required,numeric"`
}
