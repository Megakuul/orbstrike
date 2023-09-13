package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pt "github.com/megakuul/orbstrike/proto"
)

type server struct {

}

func (s *server) HandleGameUpdate(ctx context.Context, mv *pt.Move) (*pt.GameBoard, error) {
	return &pt.GameBoard{
		Players: []*pt.Player{
			{X: 10, Y: 20, Color: 1, Kills: 5, RingEnabled: true},
			{X: 30, Y: 40, Color: 2, Kills: 3, RingEnabled: false},
		},
	}, nil
}

func main() {
	lis,err:= net.Listen("tcp", ":50051")
	if err!=nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pt.RegisterGameServiceServer(s, &server{})

	if err := s.Serve(lis); err!=nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
