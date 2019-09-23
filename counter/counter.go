package counter

import "sync"

// Counter ...
type Counter struct {
	mu    sync.Mutex
	Value int
}

// NewCounter ...
func NewCounter() *Counter {
	return &Counter{}
}

// Inc ...
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Value = c.Value + 1
}

// Val ...
func (c *Counter) Val() int {
	return c.Value
}
