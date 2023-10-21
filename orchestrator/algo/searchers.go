package algo

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/megakuul/orbstrike/orchestrator/logger"
)

func FindGServer(rdb *redis.ClusterClient, ctx *context.Context, gameid string) (string, error) {
	gservers, err := rdb.SMembers(*ctx, "gserver:index:games").Result()
	if err!=nil {
		logger.WriteWarningLogger(err)
	}
	for _, key := range gservers {
		games, err:=rdb.HGetAll(*ctx, key).Result()
		if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}

		instanceid, exists := games[gameid]
		if exists {
			addr, err:=rdb.Get(*ctx,
				fmt.Sprintf("gserver:%s:addr", instanceid),
			).Result()
			if err==redis.Nil {
				return "", nil
			} else if err!=nil {
				return "", err
			}

			return addr, nil
		}
	}
	
	return "", nil
}

func FindUnhandledGameMaps(rdb *redis.ClusterClient, ctx *context.Context, gserverKeys []string) (unhandledIds []int64) {
	unhandledIdsBuf := []int64{}
	for _, key := range gserverKeys {
		var instanceid int64
		_, err := fmt.Sscanf(key, "gserver:%d:game", &instanceid)
		if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}

		err = rdb.Get(*ctx,
			fmt.Sprintf("gserver:%d:addr", instanceid),
		).Err()
		if err==redis.Nil {
			unhandledIdsBuf = append(unhandledIdsBuf, instanceid)
			continue
		} else if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}
	}
	return unhandledIdsBuf
}

