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
	GetNetwork(id int) (models.Network, error)
}

type Pass interface {
	ChangePassword(id int, Oldpass, Newpass string) error
	ResetPassword(pass string, email string) error
	CheckPassword(id int, password string) error
}

type Project interface {
	CreateProject(project models.Project) (int, error)
	GetProjects(offset int) ([]models.ProjectForResponse, error)
	GetProjectById(id int) (models.Project, error)
	UpdateProject(id int, input models.Project) error
	DeleteProject(id int) error
	SetUnlockToken(id, amount int) error
}

type Wallet interface {
	CreateWalletForUser(wallet models.Wallet) (int, error)
	CreateWalletForProject(wallet models.Wallet) (int, error)
	GetWallet(id int) (models.WalletForResponse, error)
	GetWallets() ([]models.WalletForResponse, error)
	UpdateBalance(id, amount int) error
	GetBalance(id int) (int, error)
	GetAddress(id int) (string, error)
}

type Repository struct {
	Authorization
	Usr
	Network
	Pass
	Project
	Wallet
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Usr:           NewUserPostgres(db),
		Pass:          NewPassPostgres(db),
		Network:       NewNetPostgres(db),
		Project:       NewProjPostgres(db),
		Wallet:        NewWalPostgres(db),
	}
}
