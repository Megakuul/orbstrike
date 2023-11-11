package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/megakuul/orbstrike/orchestrator/conf"
	"github.com/megakuul/orbstrike/orchestrator/ssl"
	"github.com/redis/go-redis/v9"
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
		Addrs: TrimSplit(config.DBShardNodes, ","),
		Username: config.DBUsername,
		Password: config.DBPassword,
		TLSConfig: tlsConf,
	})

	err = rdb.Ping(context.Background()).Err()
	if err!=nil {
		return nil, fmt.Errorf("%s\ncluster-nodes: %s", err, config.DBShardNodes) 
	}

	return rdb, nil
}

func TrimSplit(input string, delimiter string) ([]string) {
	slices := strings.Split(input, delimiter)
	sliceBuf := []string{}
	for _, slice := range slices {
		slice = strings.TrimSpace(slice)
		if slice!="" {
			sliceBuf = append(sliceBuf, slice)
		}
	}
	return sliceBuf
}
