package goevent

import (
	"context"
	"time"
)

type Event struct {
	ch  chan struct{}
	set bool
}

func NewEvent() *Event {
	ch := make(chan struct{}, 1)
	return &Event{ch, false}
}

func (e *Event) Set() {
	if e == nil {
		return
	}
	e.set = true
	select {
	case e.ch <- struct{}{}:
	default:
	}
}

func (e *Event) Wait() {
	if e == nil {
		return
	}
	<-e.ch
	e.Set()
}

func (e *Event) WaitMax(timeout time.Duration) bool {
	if e == nil {
		return false
	}
	select {
	case <-e.ch:
		e.Set()
		return true
	case <-time.After(timeout):
		return false
	}
}

func (e *Event) WaitCtx(ctx context.Context) bool {
	if e == nil {
		return false
	}
	select {
	case <-e.ch:
		e.Set()
		return true
	case <-ctx.Done():
		return false
	}
}

func (e *Event) IsSet() bool {
	return e != nil && e.set
}
