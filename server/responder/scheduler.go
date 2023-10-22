package responder

import (
	"time"

	"github.com/megakuul/orbstrike/server/logger"
	"github.com/megakuul/orbstrike/server/proto/game"
	"github.com/megakuul/orbstrike/server/socket/sgame"
)

func StartScheduler(srv *sgame.Server) {
	pool := &Pool{
		Workers: []PoolWorker{},
		MaxQueueSize: srv.MaxChannelSize,
	}
	interval :=
		time.Duration(srv.ResponseIntervalMS)*time.Millisecond

	logger.WriteInformationLogger(
		"Initiating Response Scheduler...",
	)
	for {
		start := time.Now()
		workerCount :=
			(len(srv.SessionRequests) + srv.RequestPerWorker -1) / srv.RequestPerWorker
		
		pool.AdjustWorkers(uint(workerCount))

		srv.Mutex.RLock()
		for i, slice := range splitRequestMap(srv.SessionRequests, workerCount) {
			pool.Workers[i].AddResponse(slice, srv)
		}
		srv.Mutex.RUnlock()

		elapsed := time.Since(start)
        if elapsed < interval {
            time.Sleep(interval - elapsed)
        }
	}
}


func splitRequestMap(reqMap map[int64]*game.Move, desiredSlices int) []map[int64]*game.Move {
	if desiredSlices <= 0 {
		return []map[int64]*game.Move{}
	}
	size := (len(reqMap) + desiredSlices - 1) / desiredSlices
	
	slices := make([]map[int64]*game.Move, desiredSlices)
	for i := range slices {
		slices[i] = make(map[int64]*game.Move)
	}

	mapSliceIdx := 0
	slicesIdx := 0

	for k,v := range reqMap {
		slices[slicesIdx][k] = v
		mapSliceIdx++
		if (mapSliceIdx >= size) {
			slicesIdx++
			mapSliceIdx=0
		}
	}
	
	return slices
}
