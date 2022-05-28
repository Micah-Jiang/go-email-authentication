package global

var (
	UserExistException  = 1000
	EmailExistException = 1001
	VerifyCodeException = 1002
	SmsCodeException    = 1003
	PasswordNotCorrect  = 1004
)

var Msg = map[int]string{
	UserExistException:  "用户名已存在",
	EmailExistException: "邮箱已被注册",
	VerifyCodeException: "验证码发送失败",
	SmsCodeException:    "验证码不正确",
	PasswordNotCorrect:  "密码验证错误",
}
