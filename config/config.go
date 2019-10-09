package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigFile("./config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		fmt.Println("Running in debug mode")
	}
}
