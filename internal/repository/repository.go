package repository

import (
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
}

type Usr interface {
	GetUsr(id int) (models.User, error)
	UpdateUsr(id int, input models.User) error
}

type Network interface {
	CreateNetwork(net models.Network) (int, error)
	GetNetworks() ([]models.Network, error)
	DeleteNetwork(NetId int) error
	GetNetwork(net models.Network) int
}

type Pass interface {
	ChangePassword(id int, Oldpass, Newpass string) error
	ResetPassword(pass string, email string) error
	CheckPassword(id int, password string) error
}

type Project interface {
	CreateProject(project models.Project) (int, error)
	GetProjects() ([]models.Project, error)
	GetProjectById(id int) (models.Project, error)
	UpdateProject(id int, input models.Project) error
	DeleteProject(id int) error
}

type Repository struct {
	Authorization
	Usr
	Network
	Pass
	Project
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Usr:           NewUserPostgres(db),
		Pass:          NewPassPostgres(db),
		Network:       NewNetPostgres(db),
		Project:       NewProjPostgres(db),
	}
}
