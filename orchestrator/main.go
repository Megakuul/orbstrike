package main

import (
	"fmt"
	"net"
	"time"

	"github.com/megakuul/orbstrike/orchestrator/conf"
	"github.com/megakuul/orbstrike/orchestrator/logger"
	"github.com/megakuul/orbstrike/orchestrator/proto/game"
	"github.com/megakuul/orbstrike/orchestrator/socket"
	"github.com/megakuul/orbstrike/orchestrator/ssl"

	"google.golang.org/grpc"
)

var config conf.Config


func main() {
	config, err := conf.LoadConig("orbstrike.orchestrator")
	if err!=nil {
		fmt.Println(err)
	}
	
	err = logger.InitLogger(config.LogFile, config.LogOptions, config.MaxLogSizeKB)
	if err!=nil {
		fmt.Println(err)
	}

	tlscred, err := ssl.GetAuthCredentials(
		config.Base64SSLCertificate,
		config.Base64SSLPrivateKey,
		config.Base64SSLCA,
	)
	if err!=nil {
		logger.WriteErrLogger(err)
	}

	server:=&socket.Server{
		Timeout: time.Minute*time.Duration(config.TimeoutMin),
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
	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err!=nil {
		logger.WriteErrLogger(err)
	}
	
	logger.WriteInformationLogger("Listening to port %d", config.Port)
	if err := grpcSrv.Serve(lis); err!=nil {
		logger.WriteErrLogger(err)
	}
}
