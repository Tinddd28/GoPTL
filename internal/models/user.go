package models

import "errors"

type User struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" binding:"required"`
	Lastname    string `json:"lastname" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Country     string `json:"country" binding:"required"`
	Password    string `json:"-" db:"hashpass"`
	Issuperuser bool   `json:"-" db:"issuperuser"`
	Isactive    bool   `json:"-" db:"isactive"`
	Isverified  bool   `json:"-" db:"isverified"`
	CreatedAt   string `json:"-" db:"created_at"`
	UpdatedAt   string `json:"-" db:"updated_at"`
}

type UserResponse struct {
	Id         int    `json:"-"`
	Name       string `json:"name"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Country    string `json:"country"`
	Isverified bool   `json:"isverified"`
}

func (us User) UserValidate() error {
	if us.Name == "" && us.Lastname == "" && us.Email == "" && us.Country == "" {
		return errors.New("updated structure is empty")
	}
	return nil
}
