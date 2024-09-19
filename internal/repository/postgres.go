package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

//TRC20
//TON
//ERC20
//BEP20
//SOL

const (
	usersTable        = "users"
	walletsTable      = "wallets"
	transactionsTable = "transactions"
	projectsTable     = "projects"
	networksTable     = "network_standards"
	entriesTable      = "entries"
	limitCount        = 10
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %v", err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("unable to ping the database: %v", err)
	}

	return db, nil
}
