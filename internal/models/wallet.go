package models

type WalletForUser struct {
	Id        int    `json:"-" db:"id"`
	Address   string `json:"address" binding:"required" db:"address"`
	UserId    int    `json:"-" db:"user_id"`
	NetworkId int    `json:"network_id" binding:"required" db:"network_id"`
}

type WalletForProject struct {
	Id        int    `json:"-" db:"id"`
	Address   string `json:"address" binding:"required" db:"address"`
	ProjectId int    `json:"project_id" binding:"required" db:"project_id"`
	NetworkId int    `json:"network_id" binding:"required" db:"network_id"`
}

type Wallet struct {
	Id        int    `json:"-" db:"id"`
	Address   string `json:"address" binding:"required" db:"address"`
	UserId    int    `json:"user_id" db:"user_id"`
	ProjectId int    `json:"project_id" db:"project_id"`
	NetworkId int    `json:"network_id" binding:"required" db:"network_id"`
}

type WalletForResponse struct {
	Id        int     `json:"id" db:"id"`
	Address   string  `json:"address" binding:"required" db:"address"`
	UserId    int     `json:"user_id" db:"user_id"`
	ProjectId int     `json:"project_id" db:"project_id"`
	NetworkId int     `json:"network_id" binding:"required" db:"network_id"`
	Balance   float64 `json:"balance" db:"balance"`
}
