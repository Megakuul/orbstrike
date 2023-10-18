package sauth

import (
	"context"
	"fmt"
	"time"
	"math/rand"

	"github.com/megakuul/orbstrike/orchestrator/algo"
	"github.com/megakuul/orbstrike/orchestrator/crypto"
	"github.com/megakuul/orbstrike/orchestrator/logger"
	"github.com/megakuul/orbstrike/orchestrator/proto/auth"
	"github.com/megakuul/orbstrike/orchestrator/proto/game"

	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
)

const MAX_MAP_RADIUS = 10000
const MAX_PLAYERS = 100
const MAX_SPEED = 50
const MAX_PL_RAD = 250

var ctx context.Context = context.Background()

type Server struct {
	RDB *redis.ClusterClient
	GameLifetime time.Duration
	DailyUserGameLimit int64
	GameJoinTimeoutSec int64
	ServerSecret string
	auth.UnimplementedAuthServiceServer
}

func (s *Server) CreateGame(ctx context.Context, req *auth.CreateGameRequest) (*auth.CreateGameResponse, error) {
	identifierKey := fmt.Sprintf("identifier:%s:counter", req.Identifier)
	cntr, err := s.RDB.Incr(ctx, identifierKey).Result()
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Game creation failed: unable to increment User-Counter.")
	}
	err = s.RDB.Expire(ctx, identifierKey, 24*time.Hour).Err()
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Game creation failed: unable to increment User-Counter.")
	}
	
	if cntr>s.DailyUserGameLimit {
		return nil, fmt.Errorf("Game creation failed: daily limit reached, wait 24 hours.")
	}
	
	gameid, err := s.RDB.Incr(ctx, "game:counter").Result()
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Game creation failed: unable to increment Game-Counter.")
	}

	if req.Maxplayers > MAX_PLAYERS {
		return nil,
			fmt.Errorf("Game creation failed: Game cannot have more than %d players", MAX_PLAYERS)
	}

	if req.Radius > MAX_MAP_RADIUS {
		return nil,
			fmt.Errorf("Game creation failed: Map radius must be below %d", MAX_MAP_RADIUS)
	}

	if req.Speed > MAX_SPEED {
		return nil,
			fmt.Errorf("Game creation failed: Speed must be below %d", MAX_SPEED)
	}

	if req.Playerradius > req.Playerringradius {
		return nil,
			fmt.Errorf("Game creation failed: Player ring cannot be smaller then the player")
	}

	if req.Playerradius > MAX_PL_RAD {
		return nil,
			fmt.Errorf("Game creation failed: Player radius must be below %d", MAX_PL_RAD)
	}

	defPlayer := &game.GameBoard{
		Id: int32(gameid),
		Rad: req.Radius,
		MaxPlayers: req.Maxplayers,
		Speed: req.Speed,
		PlayerRad: req.Playerradius,
		PlayerRingRad: req.Playerringradius,
		Players: map[int32]*game.Player{},
	}

	encGameBoard, err := proto.Marshal(defPlayer)
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Game creation failed: unable to parse game-template.")
	}

	err = s.RDB.SetEX(ctx,
		fmt.Sprintf("game:%d", gameid),
		encGameBoard, s.GameLifetime,
	).Err()
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Game creation failed: unable to save the game to the database.")
	}

	if err = algo.AllocateGame(s.RDB, &ctx, gameid, ""); err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Game creation failed: failed to allocate game on a gameserver.")
	}
	
	return &auth.CreateGameResponse{
		Gameid: int32(gameid),
	}, nil
}

func (s *Server) JoinGame(ctx context.Context, req *auth.JoinGameRequest) (*auth.JoinGameResponse, error) {
	playerid, err := s.RDB.Incr(ctx, "player:counter").Result()
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Player creation failed: unable to increment Player-Counter.")
	}
	encGame, err := s.RDB.Get(ctx,
		fmt.Sprintf("game:%d", req.Gameid)).Result()
	if err==redis.Nil {
		return nil, fmt.Errorf("Player creation failed: Game with ID %d does not exist.", req.Gameid)
	} else if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Player creation failed: Internal database failure, try again in 5 minutes")
	}


	var decGame game.GameBoard
	err = proto.Unmarshal([]byte(encGame), &decGame)
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Player creation failed: Failed to load Game")
	}

	// TODO: Implement algorithm for spawning players not always in the center
	encPlayer, err := proto.Marshal(&game.Player{
		Id: int32(playerid),
		Name: req.Name,
		X: 0,
		Y: 0,
		Rad: decGame.PlayerRad,
		Ringrad: decGame.PlayerRingRad,
		Color: int32(rand.Intn(8)),
		Kills: 0,
		RingEnabled: false,
		Speed: decGame.Speed,
	})
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Failed to create player template.")
	}

	userkey, err := crypto.EncryptUserKey(playerid, s.ServerSecret)
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Failed to generate userkey.")
	}

	err = s.RDB.HSet(ctx, "game:queue", encPlayer, req.Gameid).Err()
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Failed to add player to queue.")
	}

	// The part below checks if the player has been handled (taken out of the queue) and responds the client if either the player was handled or the timeout is reached.
	// TODO: If redis once implements a really simple to use key monitor feature, use this instead of this workaround implementation
	checkCount := s.GameJoinTimeoutSec
	checkTick := time.NewTicker(1 * time.Second)
	defer checkTick.Stop()
	
	for range checkTick.C {
		exists, err := s.RDB.HExists(ctx, "game:queue", string(encPlayer)).Result()
		if err==redis.Nil || !exists {
			break
		} else if err!=nil {
			logger.WriteErrLogger(err)
			return nil, fmt.Errorf("Failed to get game queue state.")
		}
		if checkCount <= 0 {
			err := s.RDB.HDel(ctx, "game:queue", string(encPlayer)).Err()
			if err!=nil {
				logger.WriteErrLogger(err)
			}
			return nil,
				fmt.Errorf("Timeout while adding player to the game! This issue can occur when the orbstrike gsync scheduler didn't handle the player in time.")
		}
		checkCount--
	}

	return &auth.JoinGameResponse{
		Userkey: userkey,
	}, nil
}

func (s *Server) ExitGame(ctx context.Context, req *auth.ExitGameRequest) (*auth.ExitGameResponse, error) {
	userid, err := crypto.DecryptUserKey(req.Userkey, s.ServerSecret)
	if userid==-1 {
		return nil, fmt.Errorf("Failed to exit game, invalid userkey provided!")
	} else if err!=nil {
		logger.WriteErrLogger(
			fmt.Errorf("Failed to delete user %d\n%v", userid, err),
		)
		return nil, nil
	}
	err = s.RDB.HSet(ctx, "game:exitqueue", userid, req.Gameid).Err()
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Failed to exit game. Try again!")
	}
	return &auth.ExitGameResponse{}, nil
}
