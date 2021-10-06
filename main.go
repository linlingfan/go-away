package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	//errgroup, c := errgroup.WithContext(context.Background())
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		id := i
		go work(ctx, &wg, id)
		//errgroup.Go(func() error {
		//	if id == 2 {
		//		time.Sleep(time.Second * 3)
		//		return errors.New("超时")
		//	}
		//	println(id)
		//	return nil
		//})

	}

	select {
	case <-ctx.Done():
		fmt.Printf("超时:%+v", ctx.Err())
	}

	wg.Wait()
	//err := errgroup.Wait()
	//if err != nil {
	//	fmt.Printf("%+v", err)
	//}
}

func work(ctx context.Context, wg *sync.WaitGroup, id int) {
	//func work(ctx context.Context, id int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("%+v", err)
		}
		wg.Done()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// 执行逻辑
			if id == 2 {
				time.Sleep(time.Second * 3)
			}
			println(id)
			return
		}
	}
}
