package models

type Network struct {
	Id          int    `json:"-" db:"id"`
	NetworkName string `json:"network_name" binding:"required"`
	NetworkCode string `json:"network_code" binding:"required"`
}
