package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NetPostgres struct {
	db *pgxpool.Pool
}

func NewNetPostgres(db *pgxpool.Pool) *NetPostgres {
	return &NetPostgres{db: db}
}

func (np *NetPostgres) CreateNetwork(net models.Network) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, code) VALUES ($1, $2) returning id", networksTable)
	var id int
	id = np.GetNetwork(net)
	if id != 0 {
		return 0, errors.New("Network with this name of code already exists")
	}
	err := np.db.QueryRow(context.Background(), query, net.NetworkName, net.NetworkCode).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, errors.New("smth went wrong")
		}
		return 0, err
	}

	return id, nil
}

func (np *NetPostgres) GetNetworks() ([]models.Network, error) {
	query := fmt.Sprintf("SELECT id, name, code FROM %s", networksTable)
	rows, err := np.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	networks := make([]models.Network, 0)
	for rows.Next() {
		var network models.Network
		err := rows.Scan(&network.Id, &network.NetworkName, &network.NetworkCode)
		if err != nil {
			return nil, err
		}
		networks = append(networks, network)
	}
	return []models.Network{}, nil
}

func (np *NetPostgres) DeleteNetwork(NetId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", networksTable)
	_, err := np.db.Exec(context.Background(), query, NetId)
	if err != nil {
		return err
	}

	return nil
}

func (np *NetPostgres) GetNetwork(net models.Network) int {
	query := fmt.Sprintf("SELECT id, name, code FRO< %s WHERE name=$1 OR code=$2", networksTable)
	row := np.db.QueryRow(context.Background(), query, net.NetworkName, net.NetworkCode)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0
	}
	return id
}
