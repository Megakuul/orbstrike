package main

import (
	"fmt"
	"net"
	"os"

	"github.com/megakuul/orbstrike/server/conf"
	"github.com/megakuul/orbstrike/server/db"
	"github.com/megakuul/orbstrike/server/gsync"
	"github.com/megakuul/orbstrike/server/logger"
	"github.com/megakuul/orbstrike/server/proto/game"
	"github.com/megakuul/orbstrike/server/responder"
	"github.com/megakuul/orbstrike/server/socket/sgame"
	"github.com/megakuul/orbstrike/server/ssl"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

var config conf.Config
var rdb *redis.ClusterClient

func main() {
	config, err := conf.LoadConig("orbstrike.server")
	if err!=nil {
		fmt.Println(err)
	}
	
	err = logger.InitLogger(config.LogFile, config.LogOptions, config.MaxLogSizeKB)
	if err!=nil {
		fmt.Println(err)
	}

	rdb, err = db.StartClient(&config)
	if err!=nil {
		logger.WriteErrLogger(err)
		os.Exit(1)
	}
	defer rdb.Close()
	
	tlscred, err := ssl.GetAuthCredentials(
		config.Base64SSLCertificate,
		config.Base64SSLPrivateKey,
		config.Base64SSLCA,
	)
	if err!=nil {
		logger.WriteErrLogger(err)
		os.Exit(1)
	}

	server:=&sgame.Server{
		RDB: rdb,
		Boards: map[int32]*game.GameBoard{},
		SessionRequests: map[int64]*game.Move{},
		SessionResponses: map[int64]error{},
		ServerSecret: config.Secret,
	}
	
	var grpcSrv *grpc.Server
	if tlscred!=nil {
		// Encrypted Socket
		grpcSrv = grpc.NewServer(grpc.Creds(tlscred))
	} else {
		// Unencrypted Socket
		grpcSrv = grpc.NewServer()
		logger.WriteWarningLogger(fmt.Errorf(
			"gRPC socket does not run with TLS and is not encrypted!",
		))
	}

	game.RegisterGameServiceServer(grpcSrv, server)

	go gsync.StartScheduler(server, &config)
	
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
