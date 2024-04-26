package presenter

import (
	"net/mail"

	"github.com/golang-jwt/jwt"
)

type Registration struct {
	Name            string `json:"name" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Auth
}

type Auth struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type SSJWTClaim struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	*jwt.StandardClaims
}

func (u *Registration) Validate() (string, bool) {
	u.Auth.Validate()

	if u.ConfirmPassword != u.Password {
		return "Kata sandi tidak cocok", true
	}
	return "", false
}

func (u *Auth) Validate() (string, bool) {
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return "Email tidak valid", true
	}

	return "", false
}
