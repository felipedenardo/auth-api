package user

import "github.com/felipedenardo/chameleon-common/pkg/base"

type User struct {
	base.Model
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"`
	Status       string `json:"status"`
}
