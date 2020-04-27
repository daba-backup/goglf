package window

import (
	"sync"
)

var m sync.Mutex

func Lock() {
	m.Lock()
}
func Unlock() {
	m.Unlock()
}
