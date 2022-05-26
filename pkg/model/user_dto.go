package model

type UserDto struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	SmsCode  string `form:"smsCode" json:"smsCode" binding:"required"`
}
