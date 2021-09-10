package main

import "fmt"

var Id = 12345

func main() {
	x, y := 0, 1
	n := 10
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	println(x)
	arr := []int{1, 2, 3} // 数组是值传递，可以使用切片
	func(a []int) {
		a[0] = 2
	}(arr)

	fmt.Println(arr)
}
