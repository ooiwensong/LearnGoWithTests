package sync

import "sync"

type Counter struct {
	mu    sync.Mutex // mutual exclusion lock; only allows one goroutine to mutate a value at one time
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
