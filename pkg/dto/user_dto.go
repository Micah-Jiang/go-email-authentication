package dto

import "reflect"

type UserDto struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
	SmsCode  string `form:"smsCode" json:"smsCode"`
}

func (u UserDto) IsEmpty() bool {
	return reflect.DeepEqual(u, UserDto{})
}
