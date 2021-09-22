# 算法小记

### 动态规划！
- [最长上升子序列 300](https://leetcode-cn.com/problems/longest-increasing-subsequence/)
- [1143. 最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)
- [53. 最大子序和](https://leetcode-cn.com/problems/maximum-subarray/)  
- [买股票最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)

### 贪心算法

### BFS、DFS遍历

1. BFS遍历 [二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)
2. DFS遍历 [求根节点到叶节点数字之和](https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/)


### 递归算法

### 双指针(快慢指针、左右指针)
- 双指针技巧再分为两类，一类是「快慢指针」，一类是「左右指针」。前者解决主要解决链表中的问题，比如典型的判定链表中是否包含环；后者主要解决数组（或者字符串）中的问题，比如二分查找。

[参考](https://github.com/labuladong/fucking-algorithm/blob/master/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/%E5%8F%8C%E6%8C%87%E9%92%88%E6%8A%80%E5%B7%A7.md)
- [是否为环形链表141](https://leetcode-cn.com/problems/linked-list-cycle/)

- 快慢指针
    1. 判定链表中是否含有环
    2. 已知链表中含有环，返回这个环的起始位置
    3. 寻找链表的中点（链表二分法）
    4. 寻找链表的倒数第 k 个元素（fast 先k步）

- 左右指针（左右指针在数组中实际是指两个索引值，一般初始化为 left = 0, right = len(arr) - 1 。）
    1. 二分查找
    2. 两数之和
    3. 反转数组
    4. 滑动窗口法（解决 大类子字符串匹配）[详情](https://labuladong.gitee.io/algo/2/22/54/)

### 回溯算法 （递归）
[参考](https://github.com/labuladong/fucking-algorithm/blob/master/%E9%AB%98%E9%A2%91%E9%9D%A2%E8%AF%95%E7%B3%BB%E5%88%97/%E5%AD%90%E9%9B%86%E6%8E%92%E5%88%97%E7%BB%84%E5%90%88.md)
 - [全排列 46](https://leetcode-cn.com/problems/permutations/)

```
// 模板
result = []
func backtrack(路径，选择列表) {
	if 满足结束条件 {
		result.add(路径)
	}
	return

	for 选择 in 选择列表 {
		做选择
		backtrack(路径，选择列表)
		撤销选择
	}
}

```
- [子集](https://leetcode-cn.com/problems/subsets/)
- [组合](https://leetcode-cn.com/problems/combinations/) [组合总和](https://leetcode-cn.com/problems/combination-sum/)
- [全排列](https://leetcode-cn.com/problems/permutations/)
### 滑动窗口
- [详情](https://labuladong.gitee.io/algo/2/22/54/)

### 树的遍历 中序遍历 前序遍历 后续遍历

- 

### 洗牌算法

### 递归算法
[参考](https://github.com/labuladong/fucking-algorithm/blob/master/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/%E9%80%92%E5%BD%92%E8%AF%A6%E8%A7%A3.md)


### 其他
- [31. 下一个排列](https://leetcode-cn.com/problems/next-permutation/)
- [146. LRU 缓存机制](https://leetcode-cn.com/problems/lru-cache/)
- [42. 接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)