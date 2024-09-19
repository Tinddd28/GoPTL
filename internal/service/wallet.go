package service

import (
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/Tinddd28/GoPTL/internal/repository"
)

type WalService struct {
	repo repository.Wallet
}

func NewWalService(repo repository.Wallet) *WalService {
	return &WalService{repo: repo}
}

func (ws *WalService) CreateWalletForUser(wallet models.Wallet) (int, error) {
	return ws.repo.CreateWalletForUser(wallet)
}

func (ws *WalService) CreateWalletForProject(wallet models.Wallet) (int, error) {
	return ws.repo.CreateWalletForProject(wallet)
}

func (ws *WalService) GetWallets() ([]models.WalletForResponse, error) {
	return ws.repo.GetWallets()
}

func (ws *WalService) GetBalance(id int) (int, error) {
	return ws.repo.GetBalance(id)
}

func (ws *WalService) UpdateBalance(id, amount int) error {
	return ws.repo.UpdateBalance(id, amount)
}

func (ws *WalService) GetAddress(id int) (string, error) {
	return ws.repo.GetAddress(id)
}
