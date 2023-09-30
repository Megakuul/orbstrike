package responder

import (
	"time"

	"github.com/megakuul/orbstrike/server/proto"
	"github.com/megakuul/orbstrike/server/socket"
	"github.com/megakuul/orbstrike/server/conf"
)

func StartScheduler(srv *socket.Server, config *conf.Config) {
	pool := &Pool{
		Workers: []PoolWorker{},
		MaxQueueSize: config.MaxChannelSize,
	}
	interval :=
		time.Duration(config.ResponseIntervalMS)*time.Millisecond
	
	for {
		start := time.Now()
		workerCount :=
			(len(srv.SessionRequests) + config.RequestPerWorker -1) / config.RequestPerWorker
		
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


func splitRequestMap(reqMap map[int64]*proto.Move, desiredSlices int) []map[int64]*proto.Move {
	if desiredSlices <= 0 {
		return []map[int64]*proto.Move{}
	}
	size := (len(reqMap) + desiredSlices - 1) / desiredSlices
	
	slices := make([]map[int64]*proto.Move, desiredSlices)
	for i := range slices {
		slices[i] = make(map[int64]*proto.Move)
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
