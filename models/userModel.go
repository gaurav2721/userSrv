package models

type User struct {
	Id         string `json:"id" validate:"required"`
	Name       string `json:"name" validate:"required,min=2"`
	SignupTime int64  `json:"signupTime" validate:"required,MinTimeAllowed=1850"` //1850 is the minimum year allowed
}
