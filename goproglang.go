//Package goproglang is my memo during read of go programming language book
package goproglang

//Counter demo encapsulation, that's object fields can access only via getter setter. see chapter 6.6
type Counter struct{ n int }

//N show current number of counter
func (c *Counter) N() int { return c.n }

//Increment increase counter by one
func (c *Counter) Increment() { c.n++ }

//Reset set counter back to zero
func (c *Counter) Reset() { c.n = 0 }
