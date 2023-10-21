package models

type User struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,min=2"`
	Signup_Time int64  `json:"signup_time" validate:"required,MinTimeAllowed=1850"` //1850 is the minimum year allowed
}
