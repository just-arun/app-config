package appconfig

import (
	"github.com/spf13/viper"
)

// for loading environmental variable from location
func GetConfig(path, name, fileType string, structInterface interface{}) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(fileType)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&structInterface)
	if err != nil {
		return
	}
	return
}
