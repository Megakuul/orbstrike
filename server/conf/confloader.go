package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port int `mapstructure:"port"`
	Addr string `mapstructre:"addr"`
	HostnameSuffix string `mapstructure:"hostnamesuffix"`
	Secret string `mapstructure:"secret"`

	DBShardNodes string `mapstructure:"dbshardnodes"`
	DBUsername string `mapstructure:"dbusername"`
	DBPassword string `mapstructure:"dbpassword"`
	DBBase64SSLCertificate string `mapstructure:"db_base64_ssl_certificate"`
	DBBase64SSLPrivateKey string `mapstructure:"db_base64_ssl_privatekey"`
	DBBase64SSLCA string `mapstructure:"db_base64_ssl_ca"`
	
	LogFile string `mapstructure:"logfile"`
	LogOptions string `mapstructure:"logoptions"`
	MaxLogSizeKB int `mapstructure:"maxlogsizekb"`
	LogStdout bool `mapstructure:"logstdout"`

	MaxChannelSize int `mapstructure:"maxchannelsize"`
	RequestPerWorker int `mapstructure:"requestsperworker"`
	ResponseIntervalMS int `mapstructure:"responseintervalms"`
	SyncIntervalMS int `mapstructure:"syncintervalms"`
	DowntimeThresholdMS int `mapstructure:"downtimethresholdms"`

	Base64SSLCertificate string `mapstructure:"base64_ssl_certificate"`
	Base64SSLPrivateKey string `mapstructure:"base64_ssl_privatekey"`
	Base64SSLCA string `mapstructure:"base64_ssl_ca"`

	GameOverMessages string `mapstructure:"gameovermessages"`
}

func LoadConig(confpath string) (Config, error) {
	var config Config

	viper.SetConfigName(confpath)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("addr", "")
	viper.SetDefault("logoptions", "ERROR|WARNING")
	viper.SetDefault("logfile", "orbstrike.log")
	viper.SetDefault("maxlogsizekb", 500)
	viper.SetDefault("maxchannelsize", 15)
	viper.SetDefault("logstdout", true)
	viper.SetDefault("requestsperworker", 10)
	viper.SetDefault("responseintervalms", 15)
	viper.SetDefault("syncintervalms", 1000)
	viper.SetDefault("downtimethresholdms", 5000)

	viper.SetDefault("db_base64_ssl_certificate", "")
	viper.SetDefault("db_base64_ssl_privatekey", "")
	viper.SetDefault("db_base64_ssl_ca", "")
	
	viper.SetDefault("base64_ssl_certificate", "")
	viper.SetDefault("base64_ssl_privatekey", "")
	viper.SetDefault("base64_ssl_ca", "")
	viper.SetDefault("hostnamesuffix", "")

	viper.SetDefault("gameovermessages", "Game Over")

	err := viper.ReadInConfig()
	if err!=nil {
		if _,ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, err
		} else {
			fmt.Println("No configfile found. Values are read from environment.")
		}
	}

	if err = viper.Unmarshal(&config); err!=nil {
		return config, err
	}

	return config, nil
}
