package user

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Country  string `json:"country" binding:"required"`
	Password string `json:"-" db:"hashpass"`
}
