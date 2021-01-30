package api

import (
	"database/sql"
)

type User struct {
	Login     string       `json:"login"`
	FirstName *string      `json:"first_name"`
	LastName  *string      `json:"last_name"`
	BirthDate sql.NullTime `json:"birth_date"`
	Gender    *string      `json:"gender"`
	City      *string      `json:"city"`
	Token     string       `json:"-"`
}

type UserProfile struct {
	User
	Hobbies []string `json:"hobbies"`
}
