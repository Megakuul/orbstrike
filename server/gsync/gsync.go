package gsync

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/megakuul/orbstrike/server/conf"
	"github.com/megakuul/orbstrike/server/logger"
	"github.com/megakuul/orbstrike/server/proto/game"
	"github.com/megakuul/orbstrike/server/socket/sgame"

	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
)

var ctx context.Context = context.Background()

func StartScheduler(srv *sgame.Server, config *conf.Config) {
	interval :=
		time.Duration(config.SyncIntervalMS)*time.Millisecond

	// TODO: If the application scales to 100+ instances this may not be enough uniqueness
	gserverId := time.Now().UnixNano() * int64(rand.Intn(255)) % 10000

	var gserverAddr string
	if config.Addr == "" {
		hostname, err := os.Hostname()
		if err!=nil {
			logger.WriteErrLogger(err)
		} else {
			gserverAddr = fmt.Sprintf("%s:%d", hostname, config.Port)
		}
	} else {
		gserverAddr = fmt.Sprintf("%s:%d", config.Addr, config.Port)
	}
	
	for {
		start:=time.Now()
	
		err := srv.RDB.SetEX(ctx,
			fmt.Sprintf("gserver:%d:addr", gserverId),
			gserverAddr,
			time.Duration(config.DowntimeThresholdMS)*time.Millisecond).Err();
		if err!=nil {
			logger.WriteErrLogger(err)
		}
		
		err = addPlayers(srv, gserverId)
		if err!=nil {
			logger.WriteErrLogger(err)
		}
		
		err = syncBoardStates(srv, gserverId)
		if err!=nil {
			logger.WriteErrLogger(err)
		}

		elapsed := time.Since(start)
        if elapsed < interval {
            time.Sleep(interval - elapsed)
        }
	}
}

func addPlayers(srv *sgame.Server, gserverId int64) error {
	queue, err := srv.RDB.HGetAll(ctx, "game:queue").Result()
	if err==redis.Nil {
		return nil 
	} else if err!=nil {
		return err
	}

	for playerStr, boardIdStr := range queue {
		boardId, err := strconv.Atoi(boardIdStr)
		if err!=nil {
			logger.WriteWarningLogger(
				fmt.Errorf("Invalid key format in game:queue"),
			)
			continue
		}
		_, exists := srv.Boards[int32(boardId)]
		if !exists {
			continue
		}

		decPlayer := &game.Player{}
		err = proto.Unmarshal([]byte(playerStr), decPlayer)
		if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}
		srv.Boards[int32(boardId)].Players[decPlayer.Id] = decPlayer
	}
	srv.RDB.Del(ctx, "game:queue")
	
	return nil
}

func syncBoardStates(srv *sgame.Server, gserverId int64) error {
	srv.Mutex.RLock()
	for k, v := range srv.Boards {
		if v==nil {
			continue
		}
		encBoard, err := proto.Marshal(v)
		if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}
		err = srv.RDB.Set(ctx,
			fmt.Sprintf("game:%d", k), encBoard,
			redis.KeepTTL).Err()
		if err!=nil {
			logger.WriteWarningLogger(err)
			continue
		}
	}
	srv.Mutex.RUnlock()

	gamesKey := fmt.Sprintf("gserver:%d:games", gserverId)
	games, err := srv.RDB.LRange(ctx,
		gamesKey, 0, -1,
	).Result()
	if err!=nil {
		return err
	}

	err = srv.RDB.SAdd(ctx, "gserver:index:games", gamesKey).Err()
	if err!=nil {
		return err
	} 
	
	srv.Mutex.RLock()
	var boardBuf map[int32]*game.GameBoard
	for _, strkey := range games {
		key, err := strconv.Atoi(strkey)
		if err!=nil {
			logger.WriteWarningLogger(
				fmt.Errorf("Invalid key format in gserver:%d:games", gserverId),
			)
			continue
		}
		val, exists := srv.Boards[int32(key)]
		if exists {
			boardBuf[int32(key)] = val
		} else {
			encBoard, err := srv.RDB.Get(ctx,
				fmt.Sprintf("game:%d", key)).Result()
			if err==redis.Nil {
				if err = srv.RDB.LRem(ctx,
					fmt.Sprintf("gserver:%d:games", gserverId),
					0, strkey,
				).Err(); err!= nil {
					logger.WriteWarningLogger(err)
				}
				continue
			} else if err!=nil {
				continue
			}

			decBoard := game.GameBoard{}
			err = proto.Unmarshal([]byte(encBoard), &decBoard)
			if err!=nil {
				logger.WriteWarningLogger(err)
				continue
			}
			boardBuf[int32(key)] = &decBoard
		}
	}
	srv.Mutex.RUnlock()

	srv.Mutex.Lock()
	srv.Boards = boardBuf
	srv.Mutex.Unlock()
	
	return nil
}
