package utils

import "github.com/spf13/viper"

type Server struct {
	Port string `mapstruct:"port"`
	Host string `mapstruct:"host"`
}
type Config struct {
	Server `mapstruct:"server"`
}

func LoadConfig() (c *Config, err error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	err = v.ReadInConfig()
	if err != nil {
		return
	}
	c = &Config{}
	v.Unmarshal(c)
	return
}
