package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func TestMainOne(t *testing.T) {
	mid := 1 / 2
	println(mid)
}

func TestBacktrack(t *testing.T) {
	nums := []int{5, 4, 6, 2}
	permute(nums)
}

func permute(nums []int) [][]int {
	var ans [][]int
	l := len(nums)
	if l == 0 {
		return ans
	}
	// 动态数组 dps（栈）
	isVisited := make([]bool, len(nums)) // 是否添加过
	stack := make([]int, 0)
	dps(nums, stack, isVisited, &ans)
	fmt.Printf("%+v \n", ans)
	return ans
}

// 回溯 + dps
func dps(nums, stack []int, isVisited []bool, ans *[][]int) {
	if len(stack) == len(nums) {
		tmp := make([]int, len(stack))
		copy(tmp, stack)
		*ans = append(*ans, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !isVisited[i] {
			isVisited[i] = true
			stack = append(stack, nums[i])
			dps(nums, stack, isVisited, ans)
			stack = stack[:len(stack)-1]
			isVisited[i] = false
		}
	}
}

// --------

func TestMerge(t *testing.T) {
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15,101}}
	intervals := [][]int{{1, 4}, {0, 4}}
	ans := merge(intervals)
	fmt.Printf("%+v \n", ans)
}

func merge(intervals [][]int) [][]int {
	l := len(intervals)
	var ans [][]int
	if l == 0 {
		return ans
	}
	if l == 1 {
		ans = append(ans, intervals[0])
		return ans
	}
	selectSort(intervals)
	fmt.Printf("%+v \n", intervals)
	temp := intervals[0]
	left, right := 0, 1
	for left <= right && right < l {
		if temp[1] >= intervals[right][0] {
			if temp[1] < intervals[right][1] {
				temp[1] = intervals[right][1]
			}
		} else {
			ans = append(ans, temp)
			temp = intervals[right]
			left = right
		}
		right++
	}
	// append last one
	ans = append(ans, temp)
	return ans
}

// 冒泡排序
func selectSort(nums [][]int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i][0] > nums[j][0] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}


// ------
func TestName(t *testing.T) {
	m:=mySqrt(6)
	println(m)
	x := math.Abs(-1.13)
	println(x)
}

func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r - l) / 2
		if mid * mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

