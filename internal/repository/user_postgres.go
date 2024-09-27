package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPostgres struct {
	db *pgxpool.Pool
}

func NewUserPostgres(db *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetUsr(id int) (models.UserResponse, error) {
	var us models.UserResponse

	query := fmt.Sprintf("SELECT id, name, lastname, email, country, isverified FROM %s WHERE id=$1", usersTable)
	err := r.db.QueryRow(context.Background(), query, id).Scan(&us.Id, &us.Name, &us.Lastname, &us.Email, &us.Country, &us.Isverified)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.UserResponse{}, errors.New("user not found")
		}
		return models.UserResponse{}, err
	}
	return us, nil
}

func (r *UserPostgres) UpdateUsr(id int, input models.User) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, lastname=$2, email=$3, country=$4 WHERE id=$5", usersTable)
	_, err := r.db.Exec(context.Background(), query, input.Name, input.Lastname, input.Email, input.Country, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("user not found")
		}
	}
	return nil
}

func (r *UserPostgres) Verification(id int) error {
	var email string
	query := fmt.Sprintf("UPDATE %s SET isverified=$1 WHERE id=$2 returning email", usersTable)
	err := r.db.QueryRow(context.Background(), query, true, id).Scan(&email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("user not found")
		}
	}

	return nil
}
