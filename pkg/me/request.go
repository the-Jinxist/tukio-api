package me

type updateProfileReq struct {
	PhoneNumber string `json:"phone_number" db:"phone_number" validate:"required,numeric"`
	FirstName   string `json:"first_name" db:"first_name" validate:"required,min=2"`
	LastName    string `json:"last_name" db:"last_name" validate:"required,min=2"`
}
