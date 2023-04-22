package config

import (
	"github.com/spf13/viper"
)

var m, n *viper.Viper

func init() {
	m = viper.New()
	m.SetConfigType("yaml")
	m.SetConfigName("mysql")
	m.AddConfigPath("config/environments/")

}

func GetMysqlConfig() *viper.Viper {
	if err := m.ReadInConfig(); err != nil {
		return nil
	}
	return m
}
