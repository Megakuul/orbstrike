package responder

import (
	"time"
	"github.com/megakuul/orbstrike/server/socket"
)

func StartScheduler(srv *socket.Server, responders int) {
	for {
		start := time.Now()


		elapsed := time.Since(start)
        if elapsed < desiredInterval {
            time.Sleep(desiredInterval - elapsed)
        }
	}
}
