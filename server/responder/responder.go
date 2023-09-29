package responder

import(
	"fmt"
	"math"
	
	"github.com/megakuul/orbstrike/server/socket"
	"github.com/megakuul/orbstrike/server/proto"
)

const speed = 10

func Respond(sessionRequests map[int64]*proto.Move, srv *socket.Server) {
	for sessionId, Move := range sessionRequests {
		
		srv.Mutex.Lock()
		srv.SessionResponses[sessionId] = nil
		
		curBoard, ok := srv.Boards[Move.Gameid]
		if !ok {
			srv.SessionResponses[sessionId] =
				fmt.Errorf("Game with id %d not found", Move.Gameid)
			srv.Mutex.Unlock()
			continue
		}
		curPlayer := curBoard.Players[Move.Userkey]
		if curPlayer==nil {
			srv.SessionResponses[sessionId] =
				fmt.Errorf("Player is not registered in this game\n")
			srv.Mutex.Unlock()
			continue
		}
		
		if (isOutsideMap(curPlayer.X, curPlayer.Y, curBoard.Rad)) {
			delete(curBoard.Players, curPlayer.Id)
			srv.SessionResponses[sessionId] =
				fmt.Errorf("Game Over!")
			srv.Mutex.Unlock()
			continue
		}
		
		switch (Move.Direction) {
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
		srv.Mutex.Unlock()
	}
}


func isOutsideMap(x, y, radius float64) bool {
	distanceFromCenter := math.Sqrt(x*x + y*y)
	return distanceFromCenter > radius
}
