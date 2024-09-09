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
}

type Password interface {
	ChangePassword(pass models.Password) error
	ResetPassword(email string) error
}

type Repository struct {
	Authorization
	Usr
	Network
	Password
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Usr:           NewUserPostgres(db),
	}
}
