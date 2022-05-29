package service

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"go-email-authentication/pkg/config"
	"go-email-authentication/pkg/dto"
	"go-email-authentication/pkg/global"
	"go-email-authentication/pkg/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

func Login(userDto dto.UserDto, c *gin.Context) {
	// 1. check username and email
	sysUser := model.QueryUserByUserNameOrEmail(userDto.Username, userDto.Email)
	if sysUser.IsEmpty() {
		c.JSON(http.StatusOK, gin.H{
			"code": global.UserExistException,
			"msg":  global.Msg[global.UserExistException],
		})
	}

	// 2. check password
	err2 := bcrypt.CompareHashAndPassword([]byte(sysUser.Password), []byte(userDto.Password))
	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": global.PasswordNotCorrect,
			"msg":  global.Msg[global.PasswordNotCorrect],
		})
	}

	// 3. return user
	c.JSON(http.StatusOK, gin.H{
		"data": sysUser,
	})
}

func Register(userDto dto.UserDto) (int, string) {
	// 0. verify sms code and process cache
	r, err := config.ConnectRedis()

	defer func(r redis.Conn) {
		err := r.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(r)

	smsInfo, err := r.Do("Get", "sms:code:"+userDto.Email)
	if err != nil {
		return http.StatusExpectationFailed, "register failed."
	}

	// delete smsCode cache
	_, err = r.Do("Del", "sms:code:"+userDto.Email)
	if err != nil {
		return http.StatusExpectationFailed, "register failed."
	}

	smsCode := strings.Split(fmt.Sprintf("%s", smsInfo), "_")[0]
	if smsCode != userDto.SmsCode {
		return global.SmsCodeException, global.Msg[global.SmsCodeException]
	}

	// 1. convert userDto to user
	user := model.SysUser{
		Email:       userDto.Email,
		Username:    userDto.Username,
		CreatedTime: time.Now(),
	}

	// 2. check email
	if model.QueryUserCountByEmail(user.Email) > 0 {
		return global.EmailExistException, global.Msg[global.EmailExistException]
	}

	// 3. check username
	if model.QueryUserCountByUserName(user.Username) > 0 {
		return global.UserExistException, global.Msg[global.UserExistException]
	}

	// 4. encode password
	//bcrypt.DefaultCost默认数值10，编码一次100ms以内，可增大数值，增加破解时间成本，例如设置为14，编码一次1s以上
	password, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusExpectationFailed, "register failed."
	}
	user.Password = string(password)

	// 5. store user
	model.InsertUser(user)
	return http.StatusOK, "registration successful"
}
