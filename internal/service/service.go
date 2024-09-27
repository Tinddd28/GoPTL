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
	GetUsr(id int) (models.UserResponse, error)
	UpdateUsr(id int, input models.User) error
	Verification(id int) error
}

type Network interface {
	CreateNetwork(net models.Network) (int, error)
	GetNetworks() ([]models.Network, error)
	DeleteNetwork(NetId int) error
	GetNetwork(id int) (models.Network, error)
}

type Password interface {
	ChangePassword(id int, Oldpass, Newpass string) error
	ResetPassword(pass string, email string) error
}

type Wallet interface {
	CreateWalletForUser(wallet models.Wallet) (int, error)
	CreateWalletForProject(wallet models.Wallet) (int, error)
	GetWallets() ([]models.WalletForResponse, error)
	GetBalance(id int) (int, error)
	UpdateBalance(id, amount int) error
	GetAddress(id int) (string, error)
}

type Project interface {
	CreateProject(proj models.Project) (int, error)
	GetProjects(offset int) ([]models.ProjectForResponse, error)
	GetProjectById(id int) (models.Project, error)
	UpdateProject(id int, input models.Project) error
	DeleteProject(id int) error
	SetUnlockToken(id, amount int) error
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
		Password:      NewPassService(repos.Pass),
		Network:       NewNetService(repos.Network),
		Project:       NewProjService(repos.Project),
		Wallet:        NewWalService(repos.Wallet),
	}
}
