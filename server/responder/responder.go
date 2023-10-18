package responder

import (
	"fmt"
	"math"

	"github.com/megakuul/orbstrike/server/crypto"
	"github.com/megakuul/orbstrike/server/proto/game"
	"github.com/megakuul/orbstrike/server/socket/sgame"
)

func Respond(sessionRequests map[int64]*game.Move, srv *sgame.Server) {
	for sessionId, Move := range sessionRequests {
		srv.Mutex.Lock()
		// Those sessions wait to be collected and handled by the socket and can be skipped
		if srv.SessionResponses[sessionId]!=nil {
			srv.Mutex.Unlock()
			continue
		}
		
		srv.SessionResponses[sessionId] = nil
		
		curBoard, ok := srv.Boards[Move.Gameid]
		if !ok {
			srv.SessionResponses[sessionId] =
				fmt.Errorf("Game with id %d not found", Move.Gameid)
			srv.Mutex.Unlock()
			continue
		}

		decUserkey, err := crypto.DecryptUserKey(Move.Userkey, srv.ServerSecret)
		if err!=nil {
			srv.SessionResponses[sessionId] =
				fmt.Errorf("UNAUTHORIZED! Userkey is invalid.")
			srv.Mutex.Unlock()
			continue	
		}
		
		curPlayer := curBoard.Players[int32(decUserkey)]
		if curPlayer==nil {
			srv.SessionResponses[sessionId] =
				fmt.Errorf("Player is not registered in this game! This issue can occur when the orbstrike gsync scheduler didn't handle the player in time.")
			srv.Mutex.Unlock()
			continue
		}
		curPlayerPos := PlayerPosition{
			X: curPlayer.X,
			Y: curPlayer.Y,
			RAD: curPlayer.Rad,
		}
		
		if curPlayerPos.isOutsideMap(curBoard.Rad) {
			delete(curBoard.Players, curPlayer.Id)
			srv.SessionResponses[sessionId] =
				fmt.Errorf("Game Over!")
			srv.Mutex.Unlock()
			continue
		}

		curPlayer.RingEnabled = sessionRequests[sessionId].EnableRing
		

		if (curPlayer.RingEnabled) {	
			for _,pid := range sessionRequests[sessionId].HitPlayers {
				curTar, ok := curBoard.Players[pid]
				if !ok {
					continue
				}
				
				if (!curTar.RingEnabled) {
					if curPlayerPos.isPlayerTouched(&PlayerPosition{
						X: curTar.X,
						Y: curTar.Y,
						RAD: curTar.Rad,
					}) {
						delete(curBoard.Players, pid)
						curPlayer.Kills++
					}
				}
			}
		} else {
			switch (Move.Direction) {
			case game.Move_UP:
				curPlayer.Y -= curPlayer.Speed
			case game.Move_UP_LEFT:
				curPlayer.Y -= curPlayer.Speed
				curPlayer.X -= curPlayer.Speed
			case game.Move_UP_RIGHT:
				curPlayer.X += curPlayer.Speed
				curPlayer.Y -= curPlayer.Speed
			case game.Move_DOWN:
				curPlayer.Y += curPlayer.Speed
			case game.Move_DOWN_LEFT:
				curPlayer.Y += curPlayer.Speed
				curPlayer.X -= curPlayer.Speed
			case game.Move_DOWN_RIGHT:
				curPlayer.Y += curPlayer.Speed
				curPlayer.X += curPlayer.Speed			
			case game.Move_LEFT:
				curPlayer.X -= curPlayer.Speed
			case game.Move_RIGHT:
				curPlayer.X += curPlayer.Speed
			}
		}
		srv.Mutex.Unlock()
	}
}

type PlayerPosition struct {
	X float64
	Y float64
	RAD float64
}

func (p *PlayerPosition)isPlayerTouched(target *PlayerPosition) bool {
	dx := p.X - target.X
	dy := p.Y - target.Y

	// We square the radius because its more efficient then squarerooting the distance
	distance := dx*dx + dy*dy
	radiusSum := p.RAD + target.RAD
	return distance <= radiusSum*radiusSum
}

func (p *PlayerPosition)isOutsideMap(mapradius float64) bool {
	distanceFromCenter := math.Sqrt(p.X*p.X + p.Y*p.Y)
	return distanceFromCenter > mapradius
}
