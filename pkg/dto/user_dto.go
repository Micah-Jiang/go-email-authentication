package dto

import "reflect"

type UserDto struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	SmsCode  string `form:"smsCode" json:"smsCode"`
}

func (u UserDto) IsEmpty() bool {
	return reflect.DeepEqual(u, UserDto{})
}
