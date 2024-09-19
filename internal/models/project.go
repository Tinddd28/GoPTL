package models

type Project struct {
	Id           int     `json:"-" db:"id"`
	Title        string  `json:"title" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	TokenTitle   string  `json:"token_title" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
	CostPerToken float64 `json:"cost_per_token" binding:"required"`
	Image        string  `json:"image"`
}

type ProjectForm struct {
	Id           int     `json:"-" db:"id"`
	Title        string  `form:"title" binding:"required"`
	Description  string  `form:"description" binding:"required"`
	TokenTitle   string  `form:"token_title" binding:"required"`
	Amount       float64 `form:"amount" binding:"required"`
	CostPerToken float64 `form:"cost_per_token" binding:"required"`
}

type ProjectForResponse struct {
	Id            int     `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	TokenTitle    string  `json:"token_title"`
	Amount        float64 `json:"amount"`
	CostPerToken  float64 `json:"cost_per_token"`
	Image         string  `json:"image"`
	UnlockedToken int     `json:"unlocked_token"`
}

type SetUnlockToken struct {
	Id            int `json:"id" binding:"required" db:"id"`
	UnlockedToken int `json:"unlocked_token" binding:"required" db:"unlocked_token"`
}
