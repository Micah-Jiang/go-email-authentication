package cache

import (
	"fmt"
	"go-email-authentication/pkg/config"
	"go-email-authentication/pkg/utils"
)

func SetSendPhoneCodeCache(email, randomCode string) {
	r, err := config.ConnectRedis()

	if err != nil {
		utils.Logger.Info().Msg("connect redis failed")
	}

	defer r.Close()

	_, err = r.Do("Set", "verify code"+email, randomCode)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = r.Do("Expire", "verify code"+email, 60)
	if err != nil {
		fmt.Println(err)
		return
	}
}
