package sync

import "sync"

type Cond struct {
	ch     chan struct{}
	locker sync.RWMutex
}

func NewCond() *Cond { return &Cond{ch: make(chan struct{})} }
func (c *Cond) Broadcast() {
	c.locker.Lock()
	defer c.locker.Unlock()
	close(c.ch)
	c.ch = make(chan struct{})
}
func (c *Cond) Signal() {
	c.locker.RLock()
	var ch = c.ch
	c.locker.RUnlock()
	ch <- struct{}{}
}
func (c *Cond) Wait() {
	c.locker.RLock()
	var ch = c.ch
	c.locker.RUnlock()
	<-ch
}
