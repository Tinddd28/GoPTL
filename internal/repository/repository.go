package repository

import (
	"github.com/Tinddd28/GoPTL/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GetUser(email, password string) (user.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
