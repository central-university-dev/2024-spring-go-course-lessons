package ex3race_test

import (
	"sync"
	"testing"
)

func IncCounter(cnt *int, wg *sync.WaitGroup) {
	defer wg.Done()

	for range 10000 {
		*cnt++
	}
}

func TestIncCounter(t *testing.T) {
	cnt := 0
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go IncCounter(&cnt, wg)
	go IncCounter(&cnt, wg)
	wg.Wait()

	if cnt != 20000 {
		t.Errorf("expected 2000, but got %d", cnt)
	}
}
