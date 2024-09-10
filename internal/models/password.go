package models

type Password struct {
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	OldPassword     string `json:"old_password" binding:"required"`
}

type PassReset struct {
	Email string `json:"email" binding:"required"`
}
