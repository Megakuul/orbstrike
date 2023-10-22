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

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"

	"net/http"
)

var config conf.Config
var rdb *redis.ClusterClient

func main() {
	var cfgFile string
	if len(os.Args) > 1 {
		cfgFile = os.Args[1]
	} else {
		cfgFile = "orbstrike.server"
	}
	
	config, err := conf.LoadConig(cfgFile)
	if err!=nil {
		fmt.Println("Failed to load configuration!")
		fmt.Println(err)
	}
	
	err = logger.InitLogger(
		config.LogFile,
		config.LogOptions,
		config.MaxLogSizeKB,
		config.LogStdout,
	)
	if err!=nil {
		fmt.Println("Failed to initialize logger!")
		fmt.Println(err)
	}

	logger.WriteInformationLogger(
		"Connecting to database cluster...",
	)
	rdb, err = db.StartClient(&config)
	if err!=nil {
		logger.WriteErrLogger(
			fmt.Errorf("Failed to connect to redis cluster!\n%v", err),
		)
		os.Exit(1)
	}
	defer rdb.Close()
	
	tlscred, err := ssl.GetAuthCredentials(
		config.Base64SSLCertificate,
		config.Base64SSLPrivateKey,
		config.Base64SSLCA,
	)
	if err!=nil {
		logger.WriteErrLogger(
			fmt.Errorf("Failed to get TLS certificate information!\n%v", err),
		)
		os.Exit(1)
	}

	server:=&sgame.Server{
		RDB: rdb,
		Boards: map[int32]*game.GameBoard{},
		SessionRequests: map[int64]*game.Move{},
		SessionResponses: map[int64]error{},
		ServerSecret: config.Secret,

		MaxChannelSize: config.MaxChannelSize,
		ResponseIntervalMS: config.ResponseIntervalMS,
		RequestPerWorker: config.RequestPerWorker,
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
	
	go responder.StartScheduler(server)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err!=nil {
		logger.WriteErrLogger(
			fmt.Errorf("Failed to launch tcp-socket on %d\n%v", config.Port, err),
		)
	}
	logger.WriteInformationLogger("Listening to port %d", config.Port) 
	if err := grpcSrv.Serve(lis); err != nil {
		logger.WriteErrLogger(
			fmt.Errorf("Failed to launch gRPC endpoint!\n%v", err),
		)
	}
}
