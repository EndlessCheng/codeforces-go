[视频讲解](https://www.bilibili.com/video/BV1d54y1M7Qg/) 第二题。

对于一个负数来说，它后面的前缀和都会把这个负数加进去。

由于要统计的是正数前缀和，那么把负数尽量放在后面，能统计到尽量多的正数前缀和。

同时，绝对值小的负数应该排在负数的前面，尽量在前缀和减为负数前还能多统计一些正数。

```py [sol1-Python3]
class Solution:
    def maxScore(self, nums: List[int]) -> int:
        nums.sort(reverse=True)
        return sum(s > 0 for s in accumulate(nums))
```

```go [sol1-Go]
func maxScore(nums []int) (ans int) {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	sum := 0
	for _, x := range nums {
		sum += x
		if sum <= 0 {
			break
		}
		ans++
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。忽略排序时的栈开销，仅用到若干额外变量。
