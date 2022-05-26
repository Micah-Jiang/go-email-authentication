package utils

import (
	viper2 "github.com/spf13/viper"
	"strings"
)

func Get(fileKey string) interface{} {
	index := strings.Index(fileKey, ".")
	fileName := fileKey[0:index]
	key := fileKey[index+1:]

	viper := viper2.New()
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("mock/local/")

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	return viper.Get(key)
}
