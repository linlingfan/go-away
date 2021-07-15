package algorithm

import (
	"fmt"
	"testing"
)

// 冒泡排序

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 4, 7, 2, 5}
	var bubbleSort = func(arr []int) {
		for i := 0; i < len(arr); i++ {
			for j := i; j < len(arr); j++ {
				if arr[j] < arr[i] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
	}
	bubbleSort(arr)
	fmt.Printf("%+v \n", arr)
}

// 快速排序

func TestFastSort(t *testing.T) {
	arr := []int{1, 4, 7, 2, 5}
	//println(maxDepth(10))
	//sort.Ints()
	partition(arr, 0, len(arr)-1)
	fmt.Printf("%+v \n", arr)
}
//func maxDepth(n int) int {
//	var depth int
//	for i := n; i > 0; i >>= 1 {
//		println(i)
//		depth++
//	}
//	return depth * 2
//}
func partition(arr []int, start, end int) {
	if start > end {
		return
	}
	i, j := start, end
	base := arr[i] // 取最左边的为基数
	for i < j {
		for i < j && arr[j] >= base {
			j--
		}
		arr[i],arr[j] = arr[j],arr[i]
		for i < j && arr[i] <= base {
			i++
		}
		arr[j],arr[i] = arr[i],arr[j]
	}
	arr[i] = base // 填入基数
	partition(arr, start, i - 1)
	partition(arr, i + 1, end)
}

// 快排递归算法 时间复杂度O(nlogn)
