package algo

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/megakuul/orbstrike/orchestrator/logger"
)


func AllocateGame(rdb *redis.ClusterClient, ctx *context.Context, gameid int64, offlinekey string) error {
	gserverKeys, err := rdb.SMembers(*ctx, "gserver:index:games").Result()
	if err!=nil {
		return err
	}

	curKey := ""
	var curLoad int64 = -1
	for _, gsrvKey := range gserverKeys {
		if gsrvKey==offlinekey { continue }
		load, err := rdb.HLen(*ctx, gsrvKey).Result()
		if err!=nil {
			continue
		}
		if curLoad<0 || load < curLoad {
			curLoad=load
			curKey=gsrvKey
			if curLoad==0 { break }
		} else { continue }
	}
	if curKey=="" {
		return fmt.Errorf("Failed to allocate game [%d]. No gameserver is available...", gameid)
	}

	var instanceid int64
	_, err = fmt.Sscanf(curKey, "gserver:%d:games", &instanceid)
	if err!=nil {
		return fmt.Errorf("Failure while allocating game: Invalid gserver:x:games key")
	}
	
	err = rdb.HSet(*ctx, curKey, gameid, instanceid).Err()
	if err!=nil {
		return err
	}
	return nil
}


func ReallocateGames(rdb *redis.ClusterClient, ctx *context.Context, offlineSrvIds []int64) error {
	for _, id := range offlineSrvIds {
		gamesKey := fmt.Sprintf("gserver:%d:games", id)

		games, err := rdb.HGetAll(*ctx, gamesKey).Result()
		if err==redis.Nil {
			rdb.SRem(*ctx, "gserver:index:games", gamesKey)
			continue
		} else if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}

		for gameKeyStr := range games {
			key, err := strconv.Atoi(gameKeyStr)
			if err!=nil {
				logger.WriteWarningLogger(
					fmt.Errorf("Invalid key format in gamekey %s", gameKeyStr),
				)
				continue
			}
			err = AllocateGame(rdb, ctx, int64(key), gamesKey)
			if err!=nil {
				logger.WriteWarningLogger(err)
				continue
			}
		}
		
		err = rdb.SRem(*ctx,
			"gserver:index:games",
			gamesKey,
		).Err()
		if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}
		
		err = rdb.Del(*ctx,
			gamesKey,
		).Err()
		if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}
	}
	return nil
}
