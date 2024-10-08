package service

import (
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/Tinddd28/GoPTL/internal/repository"
)

type UserService struct {
	repo repository.Usr
}

func NewUserService(repo repository.Usr) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsr(id int) (models.UserResponse, error) {
	return s.repo.GetUsr(id)
}

func (s *UserService) UpdateUsr(id int, input models.User) error {
	if err := input.UserValidate(); err != nil {
		return err
	}
	return s.repo.UpdateUsr(id, input)
}

func (s *UserService) Verification(id int) error {
	return s.repo.Verification(id)
}
