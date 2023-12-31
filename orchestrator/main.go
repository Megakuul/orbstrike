package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/megakuul/orbstrike/orchestrator/conf"
	"github.com/megakuul/orbstrike/orchestrator/db"
	"github.com/megakuul/orbstrike/orchestrator/focontroller"
	"github.com/megakuul/orbstrike/orchestrator/logger"
	"github.com/megakuul/orbstrike/orchestrator/proto/auth"
	"github.com/megakuul/orbstrike/orchestrator/proto/game"
	"github.com/megakuul/orbstrike/orchestrator/socket/proxy"
	"github.com/megakuul/orbstrike/orchestrator/socket/sauth"
	"github.com/megakuul/orbstrike/orchestrator/ssl"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

var config conf.Config
var rdb *redis.ClusterClient

func main() {
	var cfgFile string
	if len(os.Args) > 1 {
		cfgFile = os.Args[1]
	} else {
		cfgFile = "orbstrike.orchestrator"
	}
	
	config, err := conf.LoadConig(cfgFile)
	if err!=nil {
		fmt.Println("Failed to load configuration!")
		fmt.Println(err)
		os.Exit(1)
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
		os.Exit(1)
	}

	rdb, err = db.StartClient(&config)
	if err!=nil {
		logger.WriteErrLogger(
			fmt.Errorf("Failed to connect to redis cluster!\n%v", err),
		)
		os.Exit(1)
	}
	defer rdb.Close()

	gstlscred, err := ssl.GetTLSClientCA(config.GSBase64SSLCA)
	if err!=nil {
		logger.WriteErrLogger(
			fmt.Errorf("Failed to get GameServer TLS certificate information!\n%v", err),
		)
		os.Exit(1)
	}
	
	proxySrv:=&proxy.Server{
		RDB: rdb,
		Timeout: time.Minute*time.Duration(config.TimeoutMin),
		GSCredentials: gstlscred,
	}

	sauthSrv:=&sauth.Server{
		RDB: rdb,
		GameLifetime: time.Duration(config.GameLifetimeMin)*time.Minute,
		DailyUserGameLimit: int64(config.DailyUserGameLimit),
		GameJoinTimeoutSec: int64(config.GameJoinTimeoutSec),
		ServerSecret: config.Secret,
	}

	tlscred, err := ssl.GetAuthCredentials(
		config.Base64SSLCertificate,
		config.Base64SSLPrivateKey,
		config.Base64SSLCA,
	)
	if err!=nil {
		logger.WriteErrLogger(
			fmt.Errorf("Failed to get Proxy TLS certificate information!\n%v", err),
		)
		os.Exit(1)
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

	go focontroller.StartScheduler(rdb, &config)
	
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
