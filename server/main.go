package main

import (
	"fmt"
	"net"

	"github.com/megakuul/orbstrike/server/conf"
	"github.com/megakuul/orbstrike/server/logger"
	"github.com/megakuul/orbstrike/server/socket"
	"github.com/megakuul/orbstrike/server/responder"
	"github.com/megakuul/orbstrike/server/proto/game"
	
	"google.golang.org/grpc"
)

var config conf.Config

func main() {
	config, err := conf.LoadConig("orbstrike.server")
	if err!=nil {
		fmt.Println(err)
	}
	
	err = logger.InitLogger(config.LogFile, config.LogOptions, config.MaxLogSizeKB)
	if err!=nil {
		fmt.Println(err)
	}
	
	/** Example Board */
	board := &game.GameBoard{
		Id: 187,
		Rad: 500,
		Players: map[int32]*game.Player{
			34234: {
				Id: 34234,
				X: 10,
				Y: 20,
				Rad: 25,
				Ringrad: 40,
				Color: 1,
				Kills: 5,
				RingEnabled: true,
				Speed: 1,
			},
			32523: {
				Id: 32523,
				X: 50,
				Y: 80,
				Rad: 25,
				Ringrad: 40,
				Color: 2,
				Kills: 3,
				RingEnabled: false,
				Speed: 1,
			},
		},
	}
	/** Example Board */

	server:=&socket.Server{
		Boards: map[int32]*game.GameBoard{
			board.Id: board,
		},
		SessionRequests: map[int64]*game.Move{},
		SessionResponses: map[int64]error{},
	}
	
	grpcSrv := grpc.NewServer()
	game.RegisterGameServiceServer(grpcSrv, server)

	go responder.StartScheduler(server, &config)
	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err!=nil {
		logger.WriteErrLogger(err)
	}
	logger.WriteInformationLogger("Listening to port %d", config.Port) 
	if err := grpcSrv.Serve(lis); err != nil {
		logger.WriteErrLogger(err)
	}
}
