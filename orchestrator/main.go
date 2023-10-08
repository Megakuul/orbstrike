package main

import (
	"fmt"
	"net"
	"time"
	"os"

	"github.com/megakuul/orbstrike/orchestrator/conf"
	"github.com/megakuul/orbstrike/orchestrator/logger"
	"github.com/megakuul/orbstrike/orchestrator/proto/auth"
	"github.com/megakuul/orbstrike/orchestrator/proto/game"
	"github.com/megakuul/orbstrike/orchestrator/socket/proxy"
	"github.com/megakuul/orbstrike/orchestrator/socket/sauth"
	"github.com/megakuul/orbstrike/orchestrator/ssl"
	"github.com/megakuul/orbstrike/orchestrator/db"

	"google.golang.org/grpc"
	"github.com/go-redis/redis/v8"
)

var config conf.Config
var rdb *redis.ClusterClient

func main() {
	config, err := conf.LoadConig("orbstrike.orchestrator")
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	err = logger.InitLogger(config.LogFile, config.LogOptions, config.MaxLogSizeKB)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
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

	proxySrv:=&proxy.Server{
		RDB: rdb,
		Timeout: time.Minute*time.Duration(config.TimeoutMin),
	}

	sauthSrv:=&sauth.Server{
		RDB: rdb,
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

	game.RegisterGameServiceServer(grpcSrv, proxySrv)

	auth.RegisterAuthServiceServer(grpcSrv, sauthSrv)
	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err!=nil {
		logger.WriteErrLogger(err)
		os.Exit(1)
	}
	
	logger.WriteInformationLogger("Listening to port %d", config.Port)
	if err := grpcSrv.Serve(lis); err!=nil {
		logger.WriteErrLogger(err)
		os.Exit(1)
	}
}
