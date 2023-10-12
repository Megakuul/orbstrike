package db

import (
	"context"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/megakuul/orbstrike/orchestrator/conf"
	"github.com/megakuul/orbstrike/orchestrator/ssl"
)


func StartClient(config *conf.Config) (*redis.ClusterClient, error) {
	tlsConf, err := ssl.GetTLSClientMutual(
		config.DBBase64SSLCertificate,
		config.DBBase64SSLPrivateKey,
		config.DBBase64SSLCA,
	)
	if err!=nil {
		return nil, err
	}

	rdb:=redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: strings.Split(config.DBShardNodes, ","),
		Username: config.DBUsername,
		Password: config.DBPassword,
		TLSConfig: tlsConf,
	})

	err = rdb.Ping(context.Background()).Err()
	if err!=nil {
		return nil, err
	}

	return rdb, nil
}
