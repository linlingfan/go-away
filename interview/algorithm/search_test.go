package algorithm

import (
	"fmt"
	"testing"
)

func TestBFS(t *testing.T) {
	arr := []int{1, 2, 3}
	arr = arr[1:]
	println(arr[0])
	println(len(arr))
}

// 二位数组的基本操作
func TestSecondArr(t *testing.T) {
	var arr [][]int
	for i := 0; i < 2; i++ {
		// 每一行数组
		arr = append(arr, []int{1, 2})
	}
	println(len(arr))
	println(arr[0][0])

	for i, value := range arr {
		fmt.Printf("%d,:%+v \n", i, value)
	}
}
