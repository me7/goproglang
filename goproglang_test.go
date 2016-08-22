package goproglang

import "testing"

func TestCounter(t *testing.T) {
	c := Counter{}
	c.n = 13
	c.Increment()
	if c.N() != 1 {
		t.Errorf("Increment fail: expect 2, get %d", c.N())
	}
}
