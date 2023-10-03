package main

import (
	"fmt"
	"net"
	
	"github.com/megakuul/orbstrike/orchestrator/logger"
	"github.com/megakuul/orbstrike/orchestrator/conf"
	"github.com/megakuul/orbstrike/orchestrator/ssl"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const GAMEID_KEY = "gameid"


func Interceptor(srv interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {

	fmt.Println("Haai")
	ctx:=stream.Context()
	md,_ := metadata.FromIncomingContext(ctx)
	gameid, ok := md[GAMEID_KEY]
	if ok {
		fmt.Println(gameid)
	}

	return nil
}

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

	var srv *grpc.Server
	
	if tlscred!=nil {
		// Encrypted Socket
		srv = grpc.NewServer(grpc.Creds(tlscred), grpc.StreamInterceptor(Interceptor))
	} else {
		// Unencrypted Socket
		srv = grpc.NewServer(grpc.StreamInterceptor(Interceptor))
		logger.WriteWarningLogger(fmt.Errorf(
			"gRPC socket does not run with TLS and is not encrypted!",
		))
	}
	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err!=nil {
		logger.WriteErrLogger(err)
	}
	
	logger.WriteInformationLogger("Listening to port %d", config.Port)
	if err := srv.Serve(lis); err!=nil {
		logger.WriteErrLogger(err)
	}
}
