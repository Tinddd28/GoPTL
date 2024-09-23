package repository

import (
	"context"
	"fmt"
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WalPostgres struct {
	db *pgxpool.Pool
}

func NewWalPostgres(db *pgxpool.Pool) *WalPostgres {
	return &WalPostgres{db: db}
}

func (wp *WalPostgres) CreateWalletForUser(wallet models.Wallet) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (address, user_id, network_standard_id, balance) VALUES ($1, $2, $3, $4) returning id", walletsTable)
	err := wp.db.QueryRow(context.Background(), query, wallet.Address, wallet.UserId, wallet.NetworkId, 0).Scan(&wallet.Id)
	if err != nil {
		return 0, err
	}

	return wallet.Id, nil
}

func (wp *WalPostgres) CreateWalletForProject(wallet models.Wallet) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (adress, project_id, network_standart_id) VALUES ($1, $2, $3) returning id", walletsTable)
	err := wp.db.QueryRow(context.Background(), query, wallet.Address, wallet.ProjectId, wallet.NetworkId).Scan(&wallet.Id)
	if err != nil {
		return 0, err
	}

	return wallet.Id, nil
}

func (wp *WalPostgres) GetWallet(id int) (models.WalletForResponse, error) {
	query := fmt.Sprintf("SELECT id, address, user_id, project_id, network_standard_id, balance FROM %s WHERE id = $1", walletsTable)
	var wallet models.WalletForResponse
	err := wp.db.QueryRow(context.Background(), query, id).Scan(&wallet.Id, &wallet.Address, &wallet.UserId, &wallet.ProjectId, &wallet.NetworkId, &wallet.Balance)
	if err != nil {
		return models.WalletForResponse{}, err
	}

	return wallet, nil
}

func (wp *WalPostgres) GetWallets() ([]models.WalletForResponse, error) {
	query := fmt.Sprintf("SELECT id, address, user_id, project_id, network_standard_id, balance FROM %s", walletsTable)
	var wallets []models.WalletForResponse
	rows, err := wp.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var wallet models.WalletForResponse
		err = rows.Scan(&wallet.Id, &wallet.Address, &wallet.UserId, &wallet.ProjectId, &wallet.NetworkId, &wallet.Balance)
		if err != nil {
			return nil, err
		}

		wallets = append(wallets, wallet)
	}
	return wallets, nil
}

func (wp *WalPostgres) UpdateBalance(id, amount int) error {
	query := fmt.Sprintf("UPDATE %s SET balance = balance + $2 WHERE id = $1", walletsTable)
	_, err := wp.db.Exec(context.Background(), query, id, amount)
	if err != nil {
		return err
	}

	return nil
}

func (wp *WalPostgres) GetBalance(id int) (int, error) {
	var balance int
	query := fmt.Sprintf("SELECT balance FROM %s WHERE id = $1", walletsTable)
	err := wp.db.QueryRow(context.Background(), query, id).Scan(&balance)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (wp *WalPostgres) GetAddress(id int) (string, error) {
	var address string
	query := fmt.Sprintf("SELECT address FROM %s WHERE id = $1", walletsTable)
	err := wp.db.QueryRow(context.Background(), query, id).Scan(&address)
	if err != nil {
		return "", err
	}

	return address, nil
}
