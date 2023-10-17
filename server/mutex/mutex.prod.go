// +build !debug

package mutex

import (
	"sync"
)

type Mutex = sync.Mutex
type RWMutex = sync.RWMutex
