package socket

import (
	"context"
	"io"
	"time"

	"github.com/megakuul/orbstrike/orchestrator/proto/game"

	"google.golang.org/grpc"
)

type Server struct {
	Timeout time.Duration
	game.UnimplementedGameServiceServer
}

func (s *Server) ProxyGameboard(clientStream game.GameService_StreamGameboardServer) error {
	initialReq, err := clientStream.Recv()
	if err == io.EOF {
		return err
	}
	if err!=nil {
		return err
	}

	_ = initialReq

	// Search the server with the gameid that is in req.Gameid
	address := "localhost:187"

	
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err!=nil {
		return err
	}
	defer conn.Close()

	srv:=game.NewGameServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()
	
	srvStream, err := srv.StreamGameboard(ctx)
	if err!=nil {
		return err
	}

	exitChan:=make(chan error)

	// TODO: Add a mechanism to exit the other loop if one fails

	go startWriterLoop(srvStream, clientStream, exitChan)

	go startReaderLoop(srvStream, clientStream, exitChan)

	// Wait until something is sent to exitChan
	select {
	case chanErr := <-exitChan:
		close(exitChan)
		return chanErr
	}
}

func startWriterLoop(
	srvStream game.GameService_StreamGameboardClient,
	clientStream game.GameService_StreamGameboardServer,
	exitChan chan<-error) {
	for {
		res,err:=srvStream.Recv()
		if err == io.EOF {
			exitChan<-nil
			return
		}
		if err!=nil {
			exitChan<-err
			return
		}

		clientStream.Send(res)
		if err!=nil {
			exitChan<-err
			return
		}
	}
}

func startReaderLoop(
	srvStream game.GameService_StreamGameboardClient,
	clientStream game.GameService_StreamGameboardServer,
	exitChan chan<-error) {
	for {
		req, err := clientStream.Recv()
		if err == io.EOF {
			exitChan<-nil
			return
		}
		if err!=nil {
			exitChan<-err
			return
		}

		srvStream.Send(req)
		if err!=nil {
			exitChan<-err
			return
		}
	}
}
