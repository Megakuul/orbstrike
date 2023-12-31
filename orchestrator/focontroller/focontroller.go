package focontroller

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/megakuul/orbstrike/orchestrator/algo"
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
		
		
		gameMaps, err := rdb.SMembers(ctx, "gserver:index:games").Result()
		if err!=nil {
			logger.WriteWarningLogger(err)
		}
		unhandledIds := algo.FindUnhandledGameMaps(rdb, &ctx, gameMaps)

		
		if len(unhandledIds)!=0 {
			algo.ReallocateGames(rdb, &ctx, unhandledIds)
		}

		awaitInterval(start)
	}
}
