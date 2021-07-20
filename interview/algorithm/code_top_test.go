package algorithm

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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
	m := mySqrt(6)
	println(m)
	x := math.Abs(-1.13)
	println(x)
}

func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

// --------
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return ans
	}
	ans = make([][]int, 0, 0)
	var pathArr []int
	dfs(root, targetSum, pathArr)
	return ans
}

func dfs(node *TreeNode, targetSum int, pathArr []int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		pathArr = append(pathArr, node.Val)
		sum := 0
		for _, value := range pathArr {
			sum = sum + value
		}
		if sum == targetSum {
			var temp []int
			copy(temp, pathArr)
			ans = append(ans, pathArr)
		}
		return
	}
	pathArr = append(pathArr, node.Val)
	dfs(node.Left, targetSum, pathArr)
	dfs(node.Right, targetSum, pathArr)
}

func TestTreeNodeDFS(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	for i := 0; i < 4; i++ {

	}
	pathSum(root, 22)
}

func TestCopyArr(t *testing.T) {
	arr := []int{1, 2, 4}
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	fmt.Printf("%+v \n", arrCopy)
	fmt.Printf("%+v \n", arr)
	fmt.Printf("%+v \n", arr[:len(arr)-1])
}

// -------
// 滑动窗口
// 窗口大小
// 校验t是否全部在窗口

func TestMinWindow(t *testing.T) {
	n := math.MaxInt64
	println(n)
	// "ADOBECODEBANC"
	//"ABC"
	// ""brfwhqcbnvcwmesnzjrkfvsbnmwruvzxpewyxyzkwbiigawflprrmkuayqtunsscyqbvqahmepvkstrdkkuwwoyugrpuljwkzzhpapeggbsaujegvacsrdmxwbtiksulrbdtshzbirbechkykoqbreyqguasmugdxjzssytasweugquqemryrrozlqtzmzdjkedccyewvitdantfcmdtdturlpqxrvgbzmgekkchvvkvjfcnrfgxacqlrgiiermfdulpeoxbjhzunbhejpdjttxnchxgajnfdzvknsnyzmtoocvxuyinshaupycrjpxpouzdmmcnzxcpbxovvwjpgjlvjnotfnivhjzvvdjwcnndkbzhykzetkbfaryfficorbkogtniiamscnbcxaisnlxpcxujfwyyaipgdcrleffisimpluhyhfnmejvmhtkjdzsidjtioaqijdxuzeuhlhtssqebmolpqdatovdewumfzjbypvhdecvytbzfpkwhpwdfsgjqavxbgycyvjbzmyxhzymqlkachdibjrtmpqnxqpvlzymnyluzjiswsszbvheeaxsppgpfvkswlgljrhjdpbzktalqhqwippzfxomsbfrmnevfmgkdhlncmbuvtrfiifpfvczwjqiyxgfjmpowdearfvmymsiuzazuearprbnqjtbwkpu"
	//"ntugap"
	str := "brfwhqcbnvcwmesnzjrkfvsbnmwruvzxpewyxyzkwbiigawflprrmkuayqtunsscyqbvqahmepvkstrdkkuwwoyugrpuljwkzzhpapeggbsaujegvacsrdmxwbtiksulrbdtshzbirbechkykoqbreyqguasmugdxjzssytasweugquqemryrrozlqtzmzdjkedccyewvitdantfcmdtdturlpqxrvgbzmgekkchvvkvjfcnrfgxacqlrgiiermfdulpeoxbjhzunbhejpdjttxnchxgajnfdzvknsnyzmtoocvxuyinshaupycrjpxpouzdmmcnzxcpbxovvwjpgjlvjnotfnivhjzvvdjwcnndkbzhykzetkbfaryfficorbkogtniiamscnbcxaisnlxpcxujfwyyaipgdcrleffisimpluhyhfnmejvmhtkjdzsidjtioaqijdxuzeuhlhtssqebmolpqdatovdewumfzjbypvhdecvytbzfpkwhpwdfsgjqavxbgycyvjbzmyxhzymqlkachdibjrtmpqnxqpvlzymnyluzjiswsszbvheeaxsppgpfvkswlgljrhjdpbzktalqhqwippzfxomsbfrmnevfmgkdhlncmbuvtrfiifpfvczwjqiyxgfjmpowdearfvmymsiuzazuearprbnqjtbwkpu"
	t1 := "ntugap"
	s := minWindow(str, t1)
	println(s)
}

func minWindow(s string, t string) string {
	windowLen := len(t)
	sLen := len(s)
	if sLen == 0 || sLen < windowLen {
		return ""
	}
	tMap := make(map[string]int, 0)
	for _, v := range t {
		_, ok := tMap[string(v)]
		if ok {
			tMap[string(v)] = tMap[string(v)] + 1
		} else {
			tMap[string(v)] = 1
		}
	}
	isHad := false
	windowRight := windowLen
	for windowRight <= sLen && windowLen <= sLen {
		sl := s[windowRight-windowLen : windowRight]
		sMap := make(map[string]int, 0)
		for _, v := range sl {
			_, ok := sMap[string(v)]
			if ok {
				sMap[string(v)] = sMap[string(v)] + 1
			} else {
				sMap[string(v)] = 1
			}
		}
		for key, v := range tMap {
			if strings.Contains(sl, key) && sMap[key] >= v {
				isHad = true
			} else {
				isHad = false
				break
			}
		}
		if isHad {
			return sl
		}

		if windowRight == sLen {
			windowLen++
			windowRight = windowLen
		} else {
			windowRight++
		}
	}
	return ""
}

// ----- version

func TestCompareVersion(t *testing.T) {
	num, err := strconv.Atoi("001")
	if err != nil {
		fmt.Printf("%+v \n", err)
	}
	println(num)
}
