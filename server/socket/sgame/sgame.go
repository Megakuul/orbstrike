package sgame

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/megakuul/orbstrike/server/proto/game"
	"github.com/megakuul/orbstrike/server/mutex"
)

type Server struct {
	RDB *redis.ClusterClient
	Boards map[int32]*game.GameBoard
	SessionRequests map[int64]*game.Move
	SessionResponses map[int64]error
	Mutex mutex.RWMutex
	ServerSecret string

	MaxChannelSize int
	ResponseIntervalMS int
	RequestPerWorker int
	game.UnimplementedGameServiceServer
}

func (s *Server) StreamGameboard(stream game.GameService_StreamGameboardServer) error {
	var err error
	var req *game.Move
	
	// TODO: If the application scales to 100+ instances this may not be enough uniqueness
	sessionId := time.Now().UnixNano() * int64(rand.Intn(255))
	
	for {
		req, err = stream.Recv()
		if err == io.EOF {
			err = nil
			break
		}
		if err!=nil{
			break
		}

		curBoard, ok := s.Boards[req.Gameid]
		if !ok {
			err = fmt.Errorf("No such game %d", req.Gameid)
			break
		}

		s.Mutex.Lock()
		s.SessionRequests[sessionId] = req
		s.Mutex.Unlock()

		s.Mutex.RLock()
		if s.SessionResponses[sessionId]!=nil {
			err = s.SessionResponses[sessionId]
			s.Mutex.RUnlock()
			break
		}
		s.Mutex.RUnlock()

		if err = stream.Send(curBoard); err!= nil {
			break
		}
		
	}

	s.Mutex.Lock()
	delete(s.SessionRequests, sessionId)
	delete(s.SessionResponses, sessionId)
	s.Mutex.Unlock()
	
	return err
}
