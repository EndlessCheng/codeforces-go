下午两点在B站讲这场双周赛的题目，[欢迎关注](https://space.bilibili.com/206214)~

---

排序后，用相向双指针模拟，把元素和放入哈希表中（注意不需要除以 $2$）。

```py [sol1-Python3]
class Solution:
    def distinctAverages(self, nums: List[int]) -> int:
        nums.sort()
        return len(set(nums[i] + nums[-i - 1] for i in range(len(nums) // 2)))
```

```go [sol1-Go]
func distinctAverages(nums []int) int {
	set := map[int]struct{}{}
	sort.Ints(nums)
	for i, n := 0, len(nums); i < n/2; i++ {
		set[nums[i]+nums[n-1-i]] = struct{}{}
	}
	return len(set)
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。
