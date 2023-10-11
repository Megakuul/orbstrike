package sauth

import (
	"context"
	"fmt"
	"time"

	"github.com/megakuul/orbstrike/orchestrator/algo"
	"github.com/megakuul/orbstrike/orchestrator/crypto"
	"github.com/megakuul/orbstrike/orchestrator/logger"
	"github.com/megakuul/orbstrike/orchestrator/proto/auth"
	"github.com/megakuul/orbstrike/orchestrator/proto/game"

	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
)

var ctx context.Context = context.Background()

type Server struct {
	RDB *redis.ClusterClient
	GameLifetime time.Duration
	DailyUserGameLimit int64
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

	defPlayer := &game.GameBoard{
		Id: int32(gameid),
		Rad: 600,
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

	// TODO: For a better user experience, here I could check if the game exists

	// TODO: Dont hardcode the players specifications
	encPlayer, err := proto.Marshal(&game.Player{
		Id: int32(playerid),
		Name: req.Name,
		X: 0,
		Y: 0,
		Rad: 50,
		Ringrad: 70,
		Color: 1,
		Kills: 0,
		RingEnabled: true,
		Speed: 2,
	})
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Failed to create player template.")
	}

	err = s.RDB.HSet(ctx, "game:queue", encPlayer, req.Gameid).Err()
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Failed to add player to queue.")
	}

	userkey, err := crypto.EncryptUserKey(playerid, s.ServerSecret)
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Failed to generate userkey.")
	}
	
	return &auth.JoinGameResponse{
		Userkey: userkey,
	}, nil
}
