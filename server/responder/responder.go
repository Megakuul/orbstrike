package responder

import (
	"math"
	"math/rand"

	"github.com/megakuul/orbstrike/server/crypto"
	"github.com/megakuul/orbstrike/server/proto/game"
	"github.com/megakuul/orbstrike/server/socket/sgame"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
				status.Errorf(codes.NotFound, "Game with id %d not found!", Move.Gameid)
			srv.Mutex.Unlock()
			continue
		}

		decUserkey, err := crypto.DecryptUserKey(Move.Userkey, srv.ServerSecret)
		if err!=nil {
			srv.SessionResponses[sessionId] =
				status.Errorf(codes.Unauthenticated, "Userkey is invalid!")
			srv.Mutex.Unlock()
			continue	
		}
		
		curPlayer := curBoard.Players[int32(decUserkey)]
		if curPlayer==nil {
			// When gsync scheduler doesn't handle the player in time, this will always lead to Game Over!
			srv.SessionResponses[sessionId] =
				status.Errorf(codes.Unauthenticated, srv.GameOverMessages[rand.Intn(len(srv.GameOverMessages))])
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
				status.Errorf(codes.Unauthenticated, srv.GameOverMessages[rand.Intn(len(srv.GameOverMessages))])
			srv.Mutex.Unlock()
			continue
		}

		curPlayer.RingEnabled = Move.EnableRing
		

		if (curPlayer.RingEnabled) {	
			for _,pid := range Move.HitPlayers {
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
			speed := curPlayer.Speed * float64(srv.ResponseIntervalMS) / 100;
			switch (Move.Direction) {
			case game.Move_UP:
				curPlayer.Y -= speed
			case game.Move_UP_LEFT:
				curPlayer.Y -= speed
				curPlayer.X -= speed
			case game.Move_UP_RIGHT:
				curPlayer.X += speed
				curPlayer.Y -= speed
			case game.Move_DOWN:
				curPlayer.Y += speed
			case game.Move_DOWN_LEFT:
				curPlayer.Y += speed
				curPlayer.X -= speed
			case game.Move_DOWN_RIGHT:
				curPlayer.Y += speed
				curPlayer.X += speed			
			case game.Move_LEFT:
				curPlayer.X -= speed
			case game.Move_RIGHT:
				curPlayer.X += speed
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
