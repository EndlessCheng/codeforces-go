用一个哈希表统计出现过的数字，一边遍历，一边看 $-\textit{nums}[i]$ 是否在哈希表中。

```py [sol1-Python3]
class Solution:
    def findMaxK(self, nums: List[int]) -> int:
        ans = -1
        s = set()
        for x in nums:
            if -x in s: ans = max(ans, abs(x))
            s.add(x)
        return ans
```

```go [sol1-Go]
func findMaxK(nums []int) int {
	ans := -1
	has := map[int]bool{}
	for _, x := range nums {
		if abs(x) > ans && has[-x] {
			ans = abs(x)
		}
		has[x] = true
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。
