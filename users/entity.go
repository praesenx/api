package users

import "time"

type User struct {
	ID         uint       `json:"id"`
	UUID       string     `json:"uuid"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Token      string     `json:"token"`
	VerifiedAt *time.Time `json:"verified_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
