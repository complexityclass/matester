package api

import (
	"database/sql"
)

type User struct {
	Login     string       `json:"login"`
	FirstName *string      `json:"first_name"`
	LastName  *string      `json:"last_name"`
	BirthDate sql.NullTime `json:"birth_date"`
	JobTitle  *string      `json:"job_title"`
	City      *string      `json:"city"`
	Token     string       `json:"-"`
}
