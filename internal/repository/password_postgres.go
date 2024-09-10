package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PassPostgres struct {
	db *pgxpool.Pool
}

func NewPassPostgres(db *pgxpool.Pool) *PassPostgres {
	return &PassPostgres{db: db}
}

func (pp *PassPostgres) ChangePassword(id int, Oldpass, Newpass string) error {
	if pp.CheckPassword(id, Oldpass) != nil {
		return fmt.Errorf("password is incorrect")
	}
	query := fmt.Sprintf("UPDATE %s SET hashpass=$2 WHERE id=$1", usersTable)
	res, err := pp.db.Exec(context.Background(), query, id, Newpass)
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	// Проверка количества затронутых строк
	rowsAffected := res.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected, possibly invalid user ID")
	}

	return nil
}

func (pp *PassPostgres) ResetPassword(pass string, email string) error {
	query := fmt.Sprintf("update %s set hashpass=$1 where email=$2", usersTable)
	row, err := pp.db.Exec(context.Background(), query, pass, email)
	if err != nil {
		return fmt.Errorf("failed to reset password: %w", err)
	}
	rows := row.RowsAffected()

	if rows == 0 {
		return fmt.Errorf("no rows affected, possibly invalid email")
	}

	return nil
}

func (pp *PassPostgres) CheckPassword(id int, password string) error {
	var pass string
	query := fmt.Sprintf("SELECT hashpass FROM %s WHERE id=$1", usersTable)
	_ = pp.db.QueryRow(context.Background(), query, id).Scan(&pass)
	if pass != password {
		return fmt.Errorf("password is incorrect")
	}

	return nil
}
