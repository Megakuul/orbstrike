package conf

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port int `mapstructure:"port"`
	TimeoutMin int `mapstructure:"timeoutmin"`

	LogFile string `mapstructure:"logfile"`
	LogOptions string `mapstructure:"logoptions"`
	MaxLogSizeKB int `mapstructure:"maxlogsizekb"`
	
	Base64SSLCertificate string `mapstructure:"base64_ssl_certificate`
	Base64SSLPrivateKey string `mapstructure:"base64_ssl_privatekey`
	Base64SSLCA string `mapstructure:"base64_ssl_ca`
}

func LoadConig(confpath string) (Config, error) {
	var config Config

	viper.SetConfigName(confpath)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("timeoutmin", 180)
	viper.SetDefault("logoptions", "ERROR|WARNING")
	viper.SetDefault("logfile", "orbstrike.log")
	viper.SetDefault("maxlogsizekb", 500)
	viper.SetDefault("maxchannelsize", 15)
	viper.SetDefault("requestsperworker", 10)
	viper.SetDefault("responseintervalms", 15)

	viper.SetDefault("base64_ssl_certificate", "")
	viper.SetDefault("base64_ssl_privatekey", "")
	viper.SetDefault("base64_ssl_ca", "")

	err := viper.ReadInConfig()
	if err!=nil {
		return config, err
	}

	if err = viper.Unmarshal(&config); err!=nil {
		return config, err
	}

	return config, nil
}
