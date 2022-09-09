// Package counter is used for counting.
package counter

import "sync"

// Inc a counter by 1.
func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// Value returns the value of the counter.
func (c *Counter) Value() int {
	return c.value
}

// Counter is used for counting.
type Counter struct {
	value int
	mu    sync.Mutex
}

// NewCounter constructs a counter.
func NewCounter(value int) *Counter {
	return &Counter{value: value}
}
