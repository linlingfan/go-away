package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(eg *sync.WaitGroup, cannel chan bool, param int) {
	defer eg.Done()
	for {
		select {
		default:
			fmt.Println("hello:%d", param)
			time.Sleep(time.Second )
			// 正常工作
		case <-cannel:
			// 退出
			println("退出：%d",param)
			return
		}
	}
}

func main() {
	cancel := make(chan bool)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		param := i
		wg.Add(1)
		go worker(&wg, cancel, param)
	}
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}
