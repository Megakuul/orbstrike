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

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

var ctx context.Context = context.Background()

func StartScheduler(srv *sgame.Server, config *conf.Config) {
	interval :=
		time.Duration(config.SyncIntervalMS)*time.Millisecond

	// TODO: If the application scales to 100+ instances this may not be enough uniqueness
	gserverId := time.Now().UnixNano() * int64(rand.Intn(255)) % 10000

	logger.WriteInformationLogger(
		"Initiating GSync Scheduler...",
	)
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
	
		err := srv.RDB.Set(ctx,
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
		
		err = removePlayers(srv, gserverId)
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
		srv.Mutex.Lock()
		_, exists := srv.Boards[int32(boardId)]
		if !exists {
			srv.Mutex.Unlock()
			continue
		}

		if srv.Boards[int32(boardId)].Players == nil {
            srv.Boards[int32(boardId)].Players = make(map[int32]*game.Player)
        }

		decPlayer := &game.Player{}
		err = proto.Unmarshal([]byte(playerStr), decPlayer)
		if err!=nil {
			srv.Mutex.Unlock()
			logger.WriteWarningLogger(err)
			continue
		}
		srv.Boards[int32(boardId)].Players[decPlayer.Id] = decPlayer
		srv.Mutex.Unlock()
	}
	srv.RDB.Del(ctx, "game:queue")
	
	return nil
}

func removePlayers(srv *sgame.Server, gserverId int64) error {
	exitqueue, err := srv.RDB.HGetAll(ctx, "game:exitqueue").Result()
	if err==redis.Nil {
		return nil 
	} else if err!=nil {
		return err
	}

	for playerIdStr, boardIdStr := range exitqueue {
		
		boardId, err := strconv.Atoi(boardIdStr)
		if err!=nil {
			logger.WriteWarningLogger(
				fmt.Errorf("Invalid key format in game:exitqueue"),
			)
			continue
		}

		srv.Mutex.RLock()
		_, exists := srv.Boards[int32(boardId)]
		if !exists {
			srv.Mutex.RUnlock()
			continue
		}
		
		playerId, err := strconv.Atoi(playerIdStr)
		if err!=nil {
			srv.Mutex.RUnlock()
			logger.WriteWarningLogger(
				fmt.Errorf("Invalid value format in game:exitqueue"),
			)
			continue
		}
		if _,exists = srv.Boards[int32(boardId)].Players[int32(playerId)]; !exists {
			srv.Mutex.RUnlock()
			continue
		}

		
		srv.Mutex.RUnlock()
		srv.Mutex.Lock()
		delete(srv.Boards[int32(boardId)].Players, int32(playerId))
		srv.Mutex.Unlock()
	}
	srv.RDB.Del(ctx, "game:exitqueue")
	
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
	games, err := srv.RDB.HGetAll(ctx, gamesKey).Result()
	if err!=nil {
		return err
	}

	err = srv.RDB.SAdd(ctx, "gserver:index:games", gamesKey).Err()
	if err!=nil {
		return err
	} 

	srv.Mutex.RLock()
	boardCount := len(srv.Boards)
	boardBuf := make(map[int32]*game.GameBoard)
	for strkey := range games {
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
				if err = srv.RDB.HDel(ctx,
					fmt.Sprintf("gserver:%d:games", gserverId),
					strkey,
				).Err(); err!= nil {
					logger.WriteWarningLogger(err)
				}
				continue
			} else if err!=nil {
				continue
			}

			var decBoard game.GameBoard
			err = proto.Unmarshal([]byte(encBoard), &decBoard)
			if err!=nil {
				logger.WriteWarningLogger(err)
				continue
			}
			boardBuf[int32(key)] = &decBoard
		}
	}

	if boardCount!=len(boardBuf) {
		logger.WriteInformationLogger(
			"Loaded Games have been adjusted. There are now %d loaded Games.", len(boardBuf),
		)
	}
	
	srv.Mutex.RUnlock()

	srv.Mutex.Lock()
	srv.Boards = boardBuf
	srv.Mutex.Unlock()
	
	return nil
}
