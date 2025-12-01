package user

import "context"

type IService interface {
	Register(ctx context.Context, name, email, password, role string) (*User, error)
	Login(ctx context.Context, email, password string) (string, *User, error)
}
