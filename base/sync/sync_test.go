package sync

import (
	"sync"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			println(i)
			time.Sleep(time.Second * 1)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
