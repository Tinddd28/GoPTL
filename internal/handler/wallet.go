package handler

import (
	"net/http"

	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/Tinddd28/GoPTL/pkg/web3"
	"github.com/gin-gonic/gin"
)

// @Summary Create wallet for user
// @Tags wallets
// @Security ApiKeyAuth
// @Description create wallet for user
// @ModuleID createWalletForUser
// @Accept  json
// @Produce  json
// @Param wallet body models.WalletForUser true "wallet info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /wallets/create_for_user [post]
func (h *Handler) CreateWalletForUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var walletForUser models.WalletForUser

	if err := c.BindJSON(&walletForUser); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	netId := walletForUser.NetworkId
	network, err := h.services.Network.GetNetwork(netId)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, "Network not found")
		return
	}

	if !web3.CheckAddressFormat(walletForUser.Address, network) {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid address")
		return
	}
	wallet := models.Wallet{
		UserId:    userId,
		NetworkId: walletForUser.NetworkId,
		Address:   walletForUser.Address,
	}
	walletId, err := h.services.Wallet.CreateWalletForUser(wallet)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to create wallet")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": walletId,
	})
}

// @Summary Create wallet for project
// @Tags wallets
// @Security ApiKeyAuth
// @Description create wallet for project
// @ModuleID createWalletForProject
// @Accept  json
// @Produce  json
// @Param wallet body models.WalletForProject true "wallet info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /wallets/create_for_project [post]
func (h *Handler) CreateWalletForProject(c *gin.Context) {
	superUser, err := getSuperUser(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if !superUser {
		NewErrorResponse(c, http.StatusForbidden, "Not superuser")
		return
	}

	var walletForProject models.WalletForProject

	if err := c.BindJSON(&walletForProject); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	netId := walletForProject.NetworkId
	network, err := h.services.Network.GetNetwork(netId)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, "Network not found")
		return
	}

	if !web3.CheckAddressFormat(walletForProject.Address, network) {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid address")
		return
	}

	wallet := models.Wallet{
		ProjectId: walletForProject.ProjectId,
		Address:   walletForProject.Address,
		NetworkId: walletForProject.NetworkId,
	}

	walletId, err := h.services.Wallet.CreateWalletForProject(wallet)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to create wallet")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": walletId,
	})
}

// @Summary Get Wallets
// @Tags wallets
// @Decription get all wallets
// @Security ApiKeyAuth
// @ModuleID getWallets
// @Accpet json
// @Produce json
// @Success 200 {object} []models.WalletForResponse
// @Failure 500 {object} ErrorResponse
// @Router /wallets/all [get]
func (h *Handler) GetWallets(c *gin.Context) {
	wallets, err := h.services.Wallet.GetWallets()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wallets)
}

// @Summary Get Balance
// @Tags wallets
// @Security ApiKeyAuth
// @Description get balance
// @ModuleID getBalance
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} ErrorResponse
// @Router /wallets/balance [get]
func (h *Handler) GetBalance(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	balance, err := h.services.Wallet.GetBalance(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"balance": balance,
	})
}
