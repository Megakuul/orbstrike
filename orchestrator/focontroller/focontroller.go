package focontroller

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/megakuul/orbstrike/orchestrator/algo"
	"github.com/megakuul/orbstrike/orchestrator/conf"
	"github.com/megakuul/orbstrike/orchestrator/logger"
)

var ctx context.Context = context.Background()

// This will make the application only print the failover log every n iteration to not spam the logs
const FAILOVER_LOGTIMEOUT = 14

func StartScheduler(rdb *redis.ClusterClient, config *conf.Config) {
	interval :=
		time.Duration(config.FailOverIntervalMS)*time.Millisecond

	fLogCount := 0

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
		
		
		gservers, err := rdb.SMembers(ctx, "gserver:index:games").Result()
		if err!=nil {
			logger.WriteWarningLogger(err)
		}
		offlineSrvIds := algo.FindOfflineServers(rdb, &ctx, gservers)

		
		// TODO ASAP: The check below does not work, as there could be old servers that are also offline, this shouldn't block the execution.
		// gservers is not the count of servers, but the count of server "configs" left, there could be e.g. 5 gserver configs, but this is no problem when the focontroller reallocates them. I must check if there is no gserver available on another way.
		if len(gservers)-len(offlineSrvIds)<=0 {
			fLogCount++
			if fLogCount==1 {
				logger.WriteErrLogger(fmt.Errorf("No gameserver is online. Failover-Controller cannot reallocate games!"))
			} else if fLogCount > FAILOVER_LOGTIMEOUT {
				fLogCount = 0
			}
		} else {
			fLogCount = 0
			if len(offlineSrvIds)!=0 {
				algo.ReallocateGames(rdb, &ctx, offlineSrvIds)
			}
		}

		awaitInterval(start)
	}
}
