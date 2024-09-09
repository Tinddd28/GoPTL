package service

import (
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/Tinddd28/GoPTL/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, bool, error)
}

type Usr interface {
	GetUsr(id int) (models.User, error)
	UpdateUsr(id int, input models.User) error
}

type Network interface {
	CreateNetwork(net models.Network) (int, error)
	GetNetworks() ([]models.Network, error)
	DeleteNetwork(NetId int) error
}

type Password interface {
	ChangePassword(pass models.Password) error
	ResetPassword(email string) error
}

type Wallet interface {
	CreateWalletForUser(wallet models.WalletForUser) (int, error)
	CreateWalletForProject(wallet models.WalletForProject) (int, error)
	GetWallets() ([]models.Wallet, error)
	GetBalance(id int) (int, error)
}

type Project interface {
	CreateProject(project models.Project) (int, error)
	GetProjects() ([]models.Project, error)
	GetProjectById(id int) (models.Project, error)
	UpdateProject(id int, input models.Project) error
	DeleteProject(id int) error
}

type Service struct {
	Authorization
	Usr
	Network
	Password
	Wallet
	Project
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Usr:           NewUserService(repos.Usr),
	}
}
