package proxy

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/megakuul/orbstrike/orchestrator/proto/game"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"github.com/go-redis/redis/v8"
)

const GAMEID_MD_KEY = "gameid"

const METADATA_MISSING_ERRMSG = "Invalid proxy request: Metadata header is missing"
const GAMEID_MISSING_ERRMSG = "Invalid proxy request: Metadata must contain 'gameid' key at first position"

type Server struct {
	RDB *redis.ClusterClient
	Timeout time.Duration
	game.UnimplementedGameServiceServer
}

func (s *Server) ProxyGameboard(clientStream game.GameService_ProxyGameboardServer) error {
	metaData, ok := metadata.FromIncomingContext(clientStream.Context())
	if !ok {
		return fmt.Errorf(METADATA_MISSING_ERRMSG)
	}

	gameid:=metaData.Get(GAMEID_MD_KEY)
	if len(gameid) != 1 {
		return fmt.Errorf(GAMEID_MISSING_ERRMSG)
	}
	
	s.RDB.Get(context.Background(), gameid[0]
	fmt.Println()
	
	// Search the server with the gameid that is in req.Gameid
	address := "localhost:50050"
	
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
