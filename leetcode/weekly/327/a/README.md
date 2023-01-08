# 方法一：遍历

遍历数组，用两个变量统计。

```py [sol1-Python3]
class Solution:
    def maximumCount(self, nums: List[int]) -> int:
        less = great = 0
        for x in nums:
            if x < 0: less += 1
            elif x > 0: great += 1
        return max(less, great)
```

```go [sol1-Go]
func maximumCount(nums []int) int {
	less, great := 0, 0
	for _, x := range nums {
		if x < 0 {
			less++
		} else if x > 0 {
			great++
		}
	}
	return max(less, great)
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。

# 方法二：二分查找

二分查找 $0$，原理见 [【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

```py [sol1-Python3]
class Solution:
    def maximumCount(self, nums: List[int]) -> int:
        return max(bisect_left(nums, 0), len(nums) - bisect_right(nums, 0))
```

```go [sol1-Go]
func maximumCount(nums []int) int {
	return max(sort.SearchInts(nums, 0), len(nums)-sort.SearchInts(nums, 1))
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
