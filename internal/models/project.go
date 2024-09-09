package models

type Project struct {
	Id           int     `json:"-" db:"id"`
	Title        string  `json:"title" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	TokenTitle   string  `json:"token_title" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
	CostPerToken float64 `json:"cost_per_token" binding:"required"`
	Image        string  `json:"image" binding:"required"`
}
