package utils

import (
	"os"

	"github.com/spf13/viper"
)

// InitConfig xxx
func InitConfig() {
	workName, _ := os.Getwd()
	viper.AddConfigPath(workName + "/config")
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
