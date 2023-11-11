package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port int `mapstructure:"port"`
	TimeoutMin int `mapstructure:"timeoutmin"`
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

	GSBase64SSLCA string `mapstructure:"gameserver_base64_ssl_ca"`
	
	Base64SSLCertificate string `mapstructure:"base64_ssl_certificate"`
	Base64SSLPrivateKey string `mapstructure:"base64_ssl_privatekey"`
	Base64SSLCA string `mapstructure:"base64_ssl_ca"`

	FOControllerDowntimeThresholdMS int `mapstructure:"focontrollerdowntimethresholdms"`

	FailOverIntervalMS int `mapstructure:"failoverintervalms"`

	GameJoinTimeoutSec int `mapstructure:"gamejointimeoutsec"`
	GameLifetimeMin int `mapstructure:"gamelifetimemin"`
	DailyUserGameLimit int `mapstructure:"dailyusergamelimit"`
}

func LoadConig(confpath string) (Config, error) {
	var config Config

	viper.SetConfigName(confpath)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// AutomaticEnv() does not work properly at this time
	viper.BindEnv("port", "PORT")
    viper.BindEnv("timeoutmin", "TIMEOUTMIN")
    viper.BindEnv("secret", "SECRET")
    viper.BindEnv("dbshardnodes", "DBSHARDNODES")
    viper.BindEnv("dbusername", "DBUSERNAME")
    viper.BindEnv("dbpassword", "DBPASSWORD")
    viper.BindEnv("db_base64_ssl_certificate", "DB_BASE64_SSL_CERTIFICATE")
    viper.BindEnv("db_base64_ssl_privatekey", "DB_BASE64_SSL_PRIVATEKEY")
    viper.BindEnv("db_base64_ssl_ca", "DB_BASE64_SSL_CA")
    viper.BindEnv("logfile", "LOGFILE")
    viper.BindEnv("logoptions", "LOGOPTIONS")
    viper.BindEnv("maxlogsizekb", "MAXLOGSIZEKB")
    viper.BindEnv("logstdout", "LOGSTDOUT")
    viper.BindEnv("gameserver_base64_ssl_ca", "GAMESERVER_BASE64_SSL_CA")
    viper.BindEnv("base64_ssl_certificate", "BASE64_SSL_CERTIFICATE")
    viper.BindEnv("base64_ssl_privatekey", "BASE64_SSL_PRIVATEKEY")
    viper.BindEnv("base64_ssl_ca", "BASE64_SSL_CA")
    viper.BindEnv("focontrollerdowntimethresholdms", "FOCONTROLLERDOWNTIMETHRESHOLDMS")
    viper.BindEnv("failoverintervalms", "FAILOVERINTERVALMS")
    viper.BindEnv("gamejointimeoutsec", "GAMEJOINTIMEOUTSEC")
    viper.BindEnv("gamelifetimemin", "GAMELIFETIMEMIN")
    viper.BindEnv("dailyusergamelimit", "DAILYUSERGAMELIMIT")

	viper.SetDefault("timeoutmin", 180)
	viper.SetDefault("logoptions", "ERROR|WARNING")
	viper.SetDefault("logfile", "orbstrike.log")
	viper.SetDefault("maxlogsizekb", 500)
	viper.SetDefault("logstdout", true)
	viper.SetDefault("maxchannelsize", 15)
	viper.SetDefault("requestsperworker", 10)
	viper.SetDefault("responseintervalms", 15)
	
	viper.SetDefault("db_base64_ssl_certificate", "")
	viper.SetDefault("db_base64_ssl_privatekey", "")
	viper.SetDefault("db_base64_ssl_ca", "")

	viper.SetDefault("gameserver_base64_ssl_ca", "")
	
	viper.SetDefault("base64_ssl_certificate", "")
	viper.SetDefault("base64_ssl_privatekey", "")
	viper.SetDefault("base64_ssl_ca", "")

	viper.SetDefault("focontrollerdowntimethresholdms", 3000)

	viper.SetDefault("failoverintervalms", 1000)
	viper.SetDefault("gamelifetimemin", 180)
	viper.SetDefault("dailyusergamelimit", 999999)
	viper.SetDefault("gamejointimeoutsec", 10)

	
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
