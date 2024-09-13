package service

import (
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/Tinddd28/GoPTL/internal/repository"
)

type ProjService struct {
	repo repository.Project
}

func NewProjService(repo repository.Project) *ProjService {
	return &ProjService{repo: repo}
}

func (ps *ProjService) CreateProject(proj models.Project) (int, error) {
	return ps.repo.CreateProject(proj)
}

func (ps *ProjService) GetProjects(offset int) ([]models.Project, error) {
	projects, err := ps.repo.GetProjects(offset)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (ps *ProjService) GetProjectById(id int) (models.Project, error) {
	project, err := ps.repo.GetProjectById(id)
	if err != nil {
		return models.Project{}, err
	}
	return project, nil
}

func (ps *ProjService) UpdateProject(id int, input models.Project) error {
	return ps.repo.UpdateProject(id, input)
}

func (ps *ProjService) DeleteProject(id int) error {
	return ps.repo.DeleteProject(id)
}
