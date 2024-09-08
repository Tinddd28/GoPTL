package service

import (
	"github.com/Tinddd28/GoPTL/internal/repository"
	"github.com/Tinddd28/GoPTL/internal/user"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
