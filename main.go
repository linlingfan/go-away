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

	arr2 := arr[:2]
	arr2 = append(arr2, 4)    // 容量cap 没超过；修改底层数组 3->4; 如果超过底层数组容量；分配一个更大的底层数组；
	arr2 = append(arr2, 4, 0) // 如果超过底层数组容量；分配一个更大的底层数组 2,2,4,0；
	fmt.Println(arr2)
	//arr = append(arr,0)
	fmt.Println(arr)
}
