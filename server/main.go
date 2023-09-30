package main

import (
	"fmt"
	"net"

	"github.com/megakuul/orbstrike/server/conf"
	"github.com/megakuul/orbstrike/server/logger"
	"github.com/megakuul/orbstrike/server/proto"
	"github.com/megakuul/orbstrike/server/socket"
	"github.com/megakuul/orbstrike/server/responder"
	"google.golang.org/grpc"
)

var config conf.Config

func main() {
	config, err := conf.LoadConig("orbstrike")
	if err!=nil {
		fmt.Println(err)
	}
	
	err = logger.InitLogger(config.LogFile, config.LogOptions, config.MaxLogSizeKB)
	if err!=nil {
		fmt.Println(err)
	}
	
	/** Example Board */
	board := &proto.GameBoard{
		Id: 187,
		Rad: 500,
		Players: map[int32]*proto.Player{
			34234: {Id: 34234, X: 10, Y: 20, Color: 1, Kills: 5, RingEnabled: true},
			32523: {Id: 32523, X: 50, Y: 80, Color: 2, Kills: 3, RingEnabled: false},
		},
	}
	/** Example Board */

	server:=&socket.Server{
		Boards: map[int32]*proto.GameBoard{
			board.Id: board,
		},
		SessionRequests: map[int64]*proto.Move{},
		SessionResponses: map[int64]error{},
	}
	
	grpcSrv := grpc.NewServer()
	proto.RegisterGameServiceServer(grpcSrv, server)

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
