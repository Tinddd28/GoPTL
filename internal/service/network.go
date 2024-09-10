package service

import (
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/Tinddd28/GoPTL/internal/repository"
)

type NetService struct {
	repo repository.Network
}

func NewNetService(repo repository.Network) *NetService {
	return &NetService{repo: repo}
}

func (ns *NetService) CreateNetwork(net models.Network) (int, error) {
	return 0, nil
}

func (ns *NetService) GetNetworks() ([]models.Network, error) {
	return []models.Network{}, nil
}

func (ns *NetService) DeleteNetwork(NetId int) error {
	return nil
}
