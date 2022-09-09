// Package counter is used for counting.
package counter

// Inc a counter by 1.
func (c *Counter) Inc() {
	c.value++
}

// Value returns the value of the counter.
func (c *Counter) Value() int {
	return c.value
}

// Counter is used for counting.
type Counter struct {
	value int
}

// NewCounter constructs a counter.
func NewCounter(value int) *Counter {
	return &Counter{value: value}
}
