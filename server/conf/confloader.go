package conf

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port int `mapstructure:"port"`
	LogFile string `mapstructure:"logfile"`
	LogOptions string `mapstructure:"logoptions"`
	MaxLogSizeKB int `mapstructure:"maxlogsizekb"`

	MaxChannelSize int `mapstructure:"maxchannelsize"`
	RequestPerWorker int `mapstructure:"requestsperworker"`
}

func LoadConig(confpath string) (Config, error) {
	var config Config

	viper.SetConfigName(confpath)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	
	viper.SetDefault("logoptions", "ERROR|WARNING")
	viper.SetDefault("logfile", "orbstrike.log")
	viper.SetDefault("maxlogsizekb", 500)
	viper.SetDefault("maxchannelsize", 15)
	viper.SetDefault("requestsperworker", 10)

	err := viper.ReadInConfig()
	if err!=nil {
		return config, err
	}

	if err = viper.Unmarshal(&config); err!=nil {
		return config, err
	}

	return config, nil
}
