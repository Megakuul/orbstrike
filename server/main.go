package main

import (
	"io"
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/megakuul/orbstrike/server/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGameServiceServer
}

func (s *server) StreamGameboard(stream proto.GameService_StreamGameboardServer) error {
	for {
		direction, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		_ = direction
		if err!=nil{
			return err
		}

		board := &proto.GameBoard{
			Players: []*proto.Player{
				{X: 10, Y: 20, Color: 1, Kills: 5, RingEnabled: true},
				{X: 30, Y: 40, Color: 2, Kills: 3, RingEnabled: false},
			},
		}

		if err = stream.Send(board); err!= nil {
			return err
		}
	}
}

func main() {
	grpcSrv := grpc.NewServer()
	proto.RegisterGameServiceServer(grpcSrv, &server{})

	wrpSrv := grpcweb.WrapServer(grpcSrv,
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
	)

	httpSrv := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			if wrpSrv.IsGrpcWebRequest(req) || wrpSrv.IsAcceptableGrpcCorsRequest(req) {
				wrpSrv.ServeHTTP(res, req)
			}
		}),
	}

	log.Println("Listening on ", httpSrv.Addr)
	log.Fatal(httpSrv.ListenAndServe())
}
