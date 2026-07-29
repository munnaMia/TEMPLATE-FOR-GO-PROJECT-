package domain

import (
	"context"
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository is the PORT defined by the domain
type UserRepository interface {
	Create(ctx context.Context, u *User) (*User, error)
}
