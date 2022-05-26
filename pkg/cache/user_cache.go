package cache

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go-email-authentication/pkg/config"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"strings"
)

func SetSendPhoneCodeCache(email, randomCode string) error {
	r, err := config.ConnectRedis()

	if err != nil {
		fmt.Println("connect redis failed")
		return err
	}

	defer func(r redis.Conn) {
		err := r.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(r)

	// 1. prevent the interface from being attacked
	smsTime, err := r.Do("Get", "sms:code:"+email)
	if err != nil {
		return err
	}

	var currentTime string = fmt.Sprintf("%s", timestamppb.Now())
	if smsTime != nil {
		timeNow, _ := strconv.ParseInt(currentTime, 10, 64)
		cacheTimeStr := strings.Split(fmt.Sprintf("%s", smsTime), "_")[1]
		cacheTime, _ := strconv.ParseInt(cacheTimeStr, 10, 64)
		if timeNow-cacheTime < 60000 {
			return fmt.Errorf("%s", "try it agin after 1 minute.")
		}
	}

	// 2. store verify code, and set the expired time for 1 minute.
	_, err = r.Do("Set", "sms:code:"+email, randomCode+"_"+currentTime, "EX", "60")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
