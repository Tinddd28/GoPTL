package service

import (
	"github.com/Tinddd28/GoPTL/internal/repository"
	"github.com/Tinddd28/GoPTL/pkg/hash"
)

type PassService struct {
	repo repository.Pass
}

func NewPassService(repo repository.Pass) *PassService {
	return &PassService{repo: repo}
}

func (s *PassService) ChangePassword(id int, Oldpass, Newpass string) error {
	oldpass := hash.GeneratePassHash(Oldpass)
	newpass := hash.GeneratePassHash(Newpass)

	if err := s.repo.CheckPassword(id, oldpass); err != nil {
		return err
	}

	if err := s.repo.ChangePassword(id, oldpass, newpass); err != nil {
		return err
	}
	return nil
}

func (s *PassService) ResetPassword(pass string, email string) error {
	hashpass := hash.GeneratePassHash(pass)
	if err := s.repo.ResetPassword(hashpass, email); err != nil {
		return err
	}
	return nil
}
