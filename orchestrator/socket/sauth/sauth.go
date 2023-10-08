package sauth

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/megakuul/orbstrike/orchestrator/db"
	"github.com/megakuul/orbstrike/orchestrator/logger"
	"github.com/megakuul/orbstrike/orchestrator/proto/auth"
	"github.com/megakuul/orbstrike/orchestrator/proto/game"

	"google.golang.org/protobuf/proto"
)

var ctx context.Context = context.Background()

type Server struct {
	RDB *redis.ClusterClient
	auth.UnimplementedAuthServiceServer
}

func (s *Server) CreateGame(ctx context.Context, req *auth.CreateGameRequest) (*auth.CreateGameResponse, error) {
	gameid, err := s.RDB.Incr(ctx, "game:counter").Result()
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Game creation failed: unable to increment Game-Counter.")
	}

	encGameBoard, err := proto.Marshal(&game.GameBoard{
		Id: gameid,
		Rad: db.VBOARD_RAD,
		Players: map[int32]*game.Player{},
	})
	if err!=nil {
		logger.WriteErrLogger(err)
		return nil, fmt.Errorf("Game creation failed: unable to parse game-template.")
	}

	
	
	return nil, nil
}

func FindGameServer(rdb *redis.ClusterClient) {
	
}

func (s *Server) JoinGame(ctx context.Context, req *auth.JoinGameRequest) (*auth.JoinGameResponse, error) {
	return nil, nil
}
