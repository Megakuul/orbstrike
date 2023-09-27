package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"math"
	
	"github.com/megakuul/orbstrike/server/proto"
	"google.golang.org/grpc"
)

type server struct {
	boards map[int32]*proto.GameBoard
	proto.UnimplementedGameServiceServer
}

func (s *server) StreamGameboard(stream proto.GameService_StreamGameboardServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err!=nil{
			fmt.Println(err)
			return err
		}

		const speed = 10;
		
		curBoard, ok := s.boards[req.Gameid]
		if !ok {
			return fmt.Errorf("No such game %d\n", req.Gameid)
		}
		
		curPlayer := curBoard.Players[req.Userkey]
		if curPlayer==nil {
			return fmt.Errorf("Player is not registered in this game\n")
		}

		if (isOutsideMap(curPlayer.X, curPlayer.Y, curBoard.Rad)) {
			delete(curBoard.Players, curPlayer.Id)
			return fmt.Errorf("Game Over!\n")
		}
		
		switch (req.Direction) {
		case proto.Move_UP:
			curPlayer.Y -= speed
		case proto.Move_UP_LEFT:
			curPlayer.Y -= speed
			curPlayer.X -= speed
		case proto.Move_UP_RIGHT:
			curPlayer.X += speed
			curPlayer.Y -= speed
		case proto.Move_DOWN:
			curPlayer.Y += speed
		case proto.Move_DOWN_LEFT:
			curPlayer.Y += speed
			curPlayer.X -= speed
		case proto.Move_DOWN_RIGHT:
			curPlayer.Y += speed
			curPlayer.X += speed			
		case proto.Move_LEFT:
			curPlayer.X -= speed
		case proto.Move_RIGHT:
			curPlayer.X += speed
		}


		if err = stream.Send(curBoard); err!= nil {
			return err
		}
	}
}

func isOutsideMap(x, y, radius float64) bool {
	distanceFromCenter := math.Sqrt(x*x + y*y)
	return distanceFromCenter > radius
}

func main() {
	board := &proto.GameBoard{
		Id: 187,
		Rad: 500,
		Players: map[int32]*proto.Player{
			34234: {Id: 34234, X: 10, Y: 20, Color: 1, Kills: 5, RingEnabled: true},
			32523: {Id: 32523, X: 50, Y: 80, Color: 2, Kills: 3, RingEnabled: false},
		},
	}

	grpcSrv := grpc.NewServer()
	proto.RegisterGameServiceServer(grpcSrv, &server{
		boards: map[int32]*proto.GameBoard{
			board.Id: board,
		},
	})
	
	lis, err := net.Listen("tcp", ":8080")
	if err!=nil {
		log.Fatalf("[ORBSTRIKE-SERVER PANIC]: %v", err)
	}
	fmt.Println("[ORBSTRIKE-SERVER INFO]: Listening to port 8080")
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("[ORBSTRIKE-SERVER PANIC]: %v", err)
	}
}
