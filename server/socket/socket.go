package socket

import (
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/megakuul/orbstrike/server/proto"
)

type Server struct {
	Boards map[int32]*proto.GameBoard
	SessionRequests map[int64]*proto.Move
	SessionResponses map[int64]error
	Mutex sync.RWMutex
	proto.UnimplementedGameServiceServer
}

func (s *Server) StreamGameboard(stream proto.GameService_StreamGameboardServer) error {
	var err error
	var req *proto.Move
	
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
			err = fmt.Errorf("No such game %d\n", req.Gameid)
			break
		}

		s.Mutex.Lock()
		s.SessionRequests[sessionId] = req
		s.Mutex.Unlock()

		s.Mutex.RLock()
		if s.SessionResponses[sessionId]!=nil {
			err = s.SessionResponses[sessionId]
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
