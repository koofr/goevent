package goevent

import (
	"testing"
	"time"
)

func TestNilObject(t *testing.T) {
	var event *Event = nil
	event.Set()
	if event.IsSet() {
		t.Error("nil event is set")
	}
	event.Wait()
}

func TestTryWait(t *testing.T) {
	e := NewEvent()
	if e.IsSet() == true {
		t.Error("unset event returns true on 'IsSet'")
	}
	e.Set()
	go func() {
		if e.IsSet() == false {
			t.Error("set event returns false on 'IsSet' on forked goroutine")
		}
	}()
	if e.IsSet() == false {
		t.Error("set event returns false on 'IsSet'")
	}
}

func TestEventWaitMaxTimeout(t *testing.T) {
	e := NewEvent()
	ok := e.WaitMax(10 * time.Millisecond)
	if ok {
		t.Error("unset event fired")
	}
}

func TestEventWaitMax(t *testing.T) {
	e := NewEvent()
	go func() {
		time.Sleep(1 * time.Millisecond)
		e.Set()
	}()

	ok := e.WaitMax(10 * time.Millisecond)
	if !ok {
		t.Error("Event didn't fire")
	}
}
