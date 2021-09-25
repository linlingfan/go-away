package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

func main() {
	// 协程组 通知 实现安全退出
	//closeChan := make(chan bool)
	//var wg sync.WaitGroup
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	id := i
	//	go work(id, &wg, closeChan)
	//}
	//
	//time.Sleep(time.Second)
	//close(closeChan)
	//wg.Wait()

	// context实现退出 超时
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//var wg2 sync.WaitGroup
	//for i := 0; i < 10; i++ {
	//	wg2.Add(1)
	//	id := i
	//	go work2(ctx, id, &wg2)
	//}
	//time.Sleep(time.Second)
	//cancel()
	//wg2.Wait()

	// 协程组异常终止其他协程 errgroup 的使用
	group, c := errgroup.WithContext(context.Background())
	//var group errgroup.Group
	group.Go(func() error {
		time.Sleep(time.Second * 2)
		select {
		case <-c.Done(): // 通过
			fmt.Println("get ctx cancel and done!")
			return nil
		default:
			fmt.Println("do go...")
		}
		return nil
	})
	group.Go(func() error {
		return errors.New("123")
	})
	group.Go(func() error {
		time.Sleep(time.Second)
		return errors.New("321")
	})
	err := group.Wait()
	if err != nil {
		fmt.Printf("%+v", err)
	}

}

// 协程组安全通知退出
// 协程组实现
func work(id int, wg *sync.WaitGroup, closeChan chan bool) {
	defer wg.Done()
	for {
		select {
		case <-closeChan:
			fmt.Println("id:%d work return", id)
			return
		default:
			// 正常执行
			fmt.Println("id:%d work", id)
			time.Sleep(time.Second)
		}
	}
}

// context 上下文通知 退出
func work2(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("id:%d work2 return", id)
			return
		default:
			// 正常执行
			fmt.Println("id:%d work2", id)
			time.Sleep(time.Second)
		}
	}
}

// 手动实现一个异常返回的协程组 参照 errGroup
