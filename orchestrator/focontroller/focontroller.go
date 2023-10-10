package focontroller

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"math/rand"

	"github.com/go-redis/redis/v8"
	"github.com/megakuul/orbstrike/orchestrator/conf"
	"github.com/megakuul/orbstrike/orchestrator/logger"
)

var ctx context.Context = context.Background()

func StartScheduler(rdb *redis.ClusterClient, config *conf.Config) {
	interval :=
		time.Duration(config.FailOverIntervalMS)*time.Millisecond

	// TODO: If the application scales to 100+ instances this may not be enough uniqueness
	orchestatorId := time.Now().UnixNano() * int64(rand.Intn(255)) % 10000

	awaitInterval := func(start time.Time) {
		elapsed := time.Since(start)
		if elapsed < interval {
			time.Sleep(interval-elapsed)
		}
	}
	
	for {
		start:=time.Now()

		err := rdb.SetNX(ctx,
			"orchestrator:focontroller", orchestatorId,
			time.Duration(config.FOControllerDowntimeThresholdMS)*time.Millisecond,
		).Err()

		orchIdStr, err := rdb.Get(ctx, "orchestrator:focontroller").Result()
		if err!=nil {
			logger.WriteWarningLogger(err)
			awaitInterval(start)
			continue
		}
		orchId, err := strconv.Atoi(orchIdStr)
		if err!=nil {
			logger.WriteWarningLogger(fmt.Errorf("FOController contains invalid orchestrator-ID."))
			awaitInterval(start)
			continue
		}
		if orchId!=int(orchestatorId) {
			awaitInterval(start)
			continue
		}
		
		
		gservers, err := rdb.SMembers(ctx, "gserver:index:games").Result()
		if err!=nil {
			logger.WriteWarningLogger(err)
		}

		offlineSrvIds := findOfflineServers(rdb, gservers)

		if len(gservers)-len(offlineSrvIds)<=0 {
			logger.WriteErrLogger(fmt.Errorf("No gameserver is online. Failover-Controller cannot reallocate games!"))
		} else {
			if len(offlineSrvIds)!=0 {
				reallocateGames(rdb, offlineSrvIds)
			}
		}

		awaitInterval(start)
	}
}

func findOfflineServers(rdb *redis.ClusterClient, gserverKeys []string) (offlineIds []int64) {
	offlineIdsBuf := []int64{}
	for _, key := range gserverKeys {
		var instanceid int64
		_, err := fmt.Sscanf(key, "gserver:%d:game", &instanceid)
		if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}

		err = rdb.Get(ctx,
			fmt.Sprintf("gserver:%d:addr", instanceid),
		).Err()
		if err==redis.Nil {
			offlineIdsBuf = append(offlineIdsBuf, instanceid)
			continue
		} else if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}
	}
	return offlineIdsBuf
}

func reallocateGames(rdb *redis.ClusterClient, offlineSrvIds []int64) error {
	for _, id := range offlineSrvIds {
		gamesKey := fmt.Sprintf("gserver:%d:games", id)

		games, err := rdb.LRange(ctx, gamesKey, 0, -1).Result()
		if err==redis.Nil {
			rdb.SRem(ctx, "gserver:index:games", gamesKey)
			continue
		} else if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}

		for _, gameKeyStr := range games {
			key, err := strconv.Atoi(gameKeyStr)
			if err!=nil {
				logger.WriteWarningLogger(
					fmt.Errorf("Invalid key format in gamekey %s", gameKeyStr),
				)
				continue
			}
			err = allocateGame(rdb, int64(key), gamesKey)
			if err!=nil {
				logger.WriteWarningLogger(err)
				continue
			}
		}
		
		err = rdb.SRem(ctx,
			"gserver:index:games",
			gamesKey,
		).Err()
		if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}
		
		err = rdb.Del(ctx,
			gamesKey,
		).Err()
		if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}
	}
	return nil
}

func allocateGame(rdb *redis.ClusterClient, gameid int64, offlinekey string) error {
	gserverKeys, err := rdb.SMembers(ctx, "gserver:index:games").Result()
	if err!=nil {
		return err
	}

	curKey := ""
	var curLoad int64 = -1
	for _, gsrvKey := range gserverKeys {
		if gsrvKey==offlinekey { continue }
		load, err := rdb.LLen(ctx, gsrvKey).Result()
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
	err = rdb.LPush(ctx, curKey, gameid).Err()
	if err!=nil {
		return err
	}
	return nil
}
