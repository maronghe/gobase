package gls_test

import (
	"sync"
	"testing"

	"github.com/changkun/gobase/gls"
)

func TestGLS(t *testing.T) {
	gls.Store("hello", "world")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		_, ok := gls.Get("hello")
		if ok {
			t.Fatalf("fail to store goroutine local data")
		}
		wg.Done()
	}()
	wg.Wait()
	v, ok := gls.Get("hello")
	if !ok {
		t.Fatalf("cannot find gls data")
	}
	if v != "world" {
		t.Fatalf("wrong gls data")
	}

	gls.Clear()
}
