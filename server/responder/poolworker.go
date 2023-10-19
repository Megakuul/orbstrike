package responder

import (
	"fmt"

	"github.com/megakuul/orbstrike/server/logger"
	"github.com/megakuul/orbstrike/server/mutex"
	"github.com/megakuul/orbstrike/server/proto/game"
	"github.com/megakuul/orbstrike/server/socket/sgame"
)

type Pool struct {
	Workers []PoolWorker
	MaxQueueSize int
}

func (p *Pool)AdjustWorkers(desiredCount uint) {
	if len(p.Workers) == int(desiredCount) {
		return
	}

	adjustCount := int(desiredCount) - len(p.Workers)
	if adjustCount > 0 {
		for i:=0;i<adjustCount;i++ {
			worker:=InitPoolWorker(p.MaxQueueSize)
			p.Workers = append(p.Workers, *worker)

			go worker.StartPoolWorker()
		}
	} else if adjustCount < 0 {
		for i:=0;i< -adjustCount;i++ {
			lastIndex := len(p.Workers)-i-1
			p.Workers[lastIndex].StopPoolWorker()
		}
		p.Workers = p.Workers[:len(p.Workers)+adjustCount]
	}
	logger.WriteInformationLogger(
		"Response Worker Pool has been adjusted. There are now %d workers running.",
		len(p.Workers),
	)
}

const LOAD_FACTOR = 0.8

type RespondType func()

type PoolWorker struct {
	Queue chan RespondType
	Mutex *mutex.RWMutex
}

func InitPoolWorker(maxQueueSize int) *PoolWorker {
	return &PoolWorker{
		Queue: make(chan RespondType, maxQueueSize),
		Mutex: &mutex.RWMutex{},
	}
}

func (p *PoolWorker)AddResponse(sessionRequests map[int64]*game.Move, srv *sgame.Server) {
	p.Mutex.RLock()
	if float64(len(p.Queue)) / float64(cap(p.Queue)) > LOAD_FACTOR {
		go logger.WriteWarningLogger(
			fmt.Errorf("Channel load: %d/%d exceeds the loadfactor",
				len(p.Queue), cap(p.Queue)),
		)
	}
	p.Mutex.RUnlock()
	
	p.Queue<-func() {
		Respond(sessionRequests, srv)
	}
}

func (p *PoolWorker)StopPoolWorker() {
	close(p.Queue)
}

func (p *PoolWorker)StartPoolWorker() {
	
	for {
		select {
		case task, ok := <-p.Queue:
			if ok {
				task()
			} else {
				break
			}
		}
	}
}
