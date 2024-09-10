package repository

import (
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NetPostgres struct {
	db *pgxpool.Pool
}

func NewNetPostgres(db *pgxpool.Pool) *NetPostgres {
	return &NetPostgres{db: db}
}

func (np *NetPostgres) CreateNetwork(net models.Network) (int, error) {
	return 0, nil
}

func (np *NetPostgres) GetNetworks() ([]models.Network, error) {
	return []models.Network{}, nil
}

func (np *NetPostgres) DeleteNetwork(NetId int) error {
	return nil
}
