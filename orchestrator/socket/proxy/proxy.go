package proxy

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"time"

	"github.com/megakuul/orbstrike/orchestrator/algo"
	"github.com/megakuul/orbstrike/orchestrator/logger"
	"github.com/megakuul/orbstrike/orchestrator/proto/game"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const GAMEID_MD_KEY = "gameid"

type Server struct {
	RDB *redis.ClusterClient
	Timeout time.Duration
	GSCredentials *tls.Config
	game.UnimplementedGameServiceServer
}

func (s *Server) ProxyGameboard(clientStream game.GameService_ProxyGameboardServer) error {
	metaData, ok := metadata.FromIncomingContext(clientStream.Context())
	if !ok {
		return fmt.Errorf("Invalid proxy request: Metadata header is missing!")
	}

	gameidStrList:=metaData.Get(GAMEID_MD_KEY)
	if len(gameidStrList) != 1 {
		return fmt.Errorf("Invalid proxy request: Metadata must contain 'gameid' key at first position!")
	}

	ctx := context.Background()
	address, err := algo.FindGServer(s.RDB, &ctx, gameidStrList[0])
	if err!=nil {
		logger.WriteErrLogger(err)
		return fmt.Errorf("Internal cluster failure: Something went wrong, try again in 5 minutes!")
	} else if address=="" {
		return fmt.Errorf("Invalid proxy request: Game not found!")
	}

	var conn *grpc.ClientConn
	if s.GSCredentials!=nil {
		conn, err = grpc.Dial(address, grpc.WithTransportCredentials(
			credentials.NewTLS(s.GSCredentials),
		))
	} else {
		conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	}
	if err!=nil {
		return err
	}
	defer conn.Close()

	srv:=game.NewGameServiceClient(conn)

	srvCtx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()
	
	srvStream, err := srv.StreamGameboard(srvCtx)
	if err!=nil {
		return err
	}

	errChan:=make(chan error)
	exitChan:=make(chan struct{})

	go startWriterLoop(srvStream, clientStream, errChan, exitChan)

	go startReaderLoop(srvStream, clientStream, errChan, exitChan)

	// Wait until something is sent to errChan
	select {
	case chanErr := <-errChan:
		close(exitChan)
		return chanErr
	}
}

func startWriterLoop(
	srvStream game.GameService_StreamGameboardClient,
	clientStream game.GameService_ProxyGameboardServer,
	errChan chan<-error, exitChan <-chan struct{}) {
	for {
		res,err:=srvStream.Recv()
		if err == io.EOF {
			errChan<-nil
			return
		}
		if err!=nil {
			errChan<-err
			return
		}

		clientStream.Send(res)
		if err!=nil {
			errChan<-err
			return
		}
		select {
		case <-exitChan:
			return
		default:
		}
	}
}

func startReaderLoop(
	srvStream game.GameService_StreamGameboardClient,
	clientStream game.GameService_ProxyGameboardServer,
	errChan chan<-error, exitChan<-chan struct{}) {
	for {
		req, err := clientStream.Recv()
		if err == io.EOF {
			errChan<-nil
			return
		}
		if err!=nil {
			errChan<-err
			return
		}

		srvStream.Send(req)
		if err!=nil {
			errChan<-err
			return
		}
		
		select {
		case <-exitChan:
			return
		default:
		}
	}
}
