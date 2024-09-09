package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthPostgres struct {
	db *pgxpool.Pool
}

func NewAuthPostgres(db *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, lastname, email, country, hashpass, isactive, isverified) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", usersTable)
	row := r.db.QueryRow(context.Background(), query, user.Name, user.Lastname, user.Email, user.Country, user.Password, true, true)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var us models.User
	query := fmt.Sprintf("SELECT id, name, lastname, email, country, issuperuser FROM %s WHERE email=$1 AND hashpass=$2", usersTable)
	err := r.db.QueryRow(context.Background(), query, email, password).Scan(&us.Id, &us.Name, &us.Lastname, &us.Email, &us.Country, &us.Issuperuser)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, errors.New("models not found")
		}
		return models.User{}, err
	}
	//log_ := logger.SetupPrettyLogger()
	//log_.Info("GetUser", us)

	return us, nil
}
