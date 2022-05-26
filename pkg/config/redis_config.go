package config

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go-email-authentication/pkg/utils"
)

func ConnectRedis() (redis.Conn, error) {
	utils.Logger.Info().Msg("start to get redis config props")
	host := utils.GetInterfaceToString(utils.Get("REDIS_HOST"))
	port := utils.GetInterfaceToString(utils.Get("REDIS_PORT"))

	c, err := redis.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return c, err
	}

	return c, nil
}

func hello(name string) string {
	return name
}
