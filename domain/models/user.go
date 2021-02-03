package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `josn:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"pass"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type UserUpdate struct {
	Id       uuid.UUID `json:"id"`
	Username *string   `json:"username,omitempty"`
	Password *string   `json:"pass,omitempty"`
	Email    *string   `json:"email,omitempty"`
	Role     *string   `json:"role,omitempty"`
}

type UserRegisterationRequest struct {
	Username string `json:"username"`
	Password string `json:"pass"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

// Claimes are infos that being stored in jwt
type Claimes struct {
	Userid    uuid.UUID `json:"userid"`
	UserRole  string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	jwt.StandardClaims
}

type Cridentials struct {
	Username string `json:"username"`
	Password string `json:"pass"`
}

type PasswordModel struct {
	Password string `json:"pass"`
}
