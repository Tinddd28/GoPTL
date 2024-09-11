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
	return ns.repo.CreateNetwork(net)
}

func (ns *NetService) GetNetworks() ([]models.Network, error) {
	networks, err := ns.repo.GetNetworks()
	if err != nil {
		return nil, err
	}
	return networks, nil
}

func (ns *NetService) DeleteNetwork(NetId int) error {
	err := ns.repo.DeleteNetwork(NetId)
	if err != nil {
		return err
	}
	return nil
}
