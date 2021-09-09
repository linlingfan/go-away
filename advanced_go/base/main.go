package main

import (
	"context"
	"fmt"
	"image/color"
	"sync"
	"time"
)

func main() {

	c := new(ColoredPoint)
	c.Add() // 编译期 -> c.Point.Add()

	a := [3]int{1, 2, 3}
	change(a)     // 数组是一个值 参数传递 复制整个数组
	println(a[0]) // 1

	b := &a
	prChange(b)   // 传入指针地址
	println(a[0]) // 3

	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}
	// 切面插入元素
	arr := []int{0, 2, 3}
	result := InsertIndex(1, 1, arr)
	fmt.Println(result)

	// context 用来简化对于处理单个请求的多个Goroutine之间与请求域的数据、超时和退出等操作
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	time.Sleep(time.Second)
	cancel()

	wg.Wait()
	// ---
}

func change(a [3]int) {
	a[0] = 3
}

func prChange(b *[3]int) {
	b[0] = 3
}

type Point struct{ X, Y float64 }

func (p *Point) Add() {
	println("add")
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func InsertIndex(i, v int, a []int) []int {
	a = append(a, 0)     // 切片扩展1个空间
	copy(a[i+1:], a[i:]) // a[i:]向后移动1个位置 // 无需创建临时切片
	a[i] = v             // 设置新添加的元素
	return a
}

// context

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			fmt.Println("over")
			return ctx.Err()
		}
	}
}
