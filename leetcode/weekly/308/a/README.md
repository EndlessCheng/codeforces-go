下午 2 点在 B 站直播讲周赛（和双周赛）的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

由于题目求的是子序列，与元素在数组中的顺序无关，我们可以排序后，从小到大选择尽量多的元素，其元素和不超过询问值。

这可以用前缀和 + 二分解决。

#### 复杂度分析

- 时间复杂度：$O((n+m)\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。

```py [sol1-Python3]
class Solution:
    def answerQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        nums.sort()
        for i in range(1, len(nums)):
            nums[i] += nums[i - 1]  # 原地求前缀和
        for i, q in enumerate(queries):
            queries[i] = bisect_right(nums, q)
        return queries
```

```go [sol1-Go]
func answerQueries(nums, queries []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1] // 原地求前缀和
	}
	for i, q := range queries {
		queries[i] = sort.SearchInts(nums, q+1)
	}
	return queries
}
```

#### 思考题

把子序列改成子数组要怎么做？
