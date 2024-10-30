package service

import (
	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

type UpdateUserRequest struct {
	Name string
	ID   int64
	Role desc.Role
}

type CreateUserRequest struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            desc.Role
}
