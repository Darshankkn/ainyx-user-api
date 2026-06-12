package models
import "time"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	DOB time.Time `json:"dob"`
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob" validate:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob" validate:"required"`
}