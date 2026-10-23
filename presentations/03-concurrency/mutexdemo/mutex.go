package mutexdemo

import "sync"

var (
	n  int
	mu sync.Mutex
)

func increment() {
	mu.Lock()
	n++ // критическая секция: только один за раз
	mu.Unlock()
}
