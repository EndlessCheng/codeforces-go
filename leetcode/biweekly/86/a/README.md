下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

统计相邻数字的和，加入哈希表中，如果这些和不足 $n-1$ 个，则子数组存在。

```py [sol1-Python3]
class Solution:
    def findSubarrays(self, nums: List[int]) -> bool:
        return len(set(map(sum, pairwise(nums)))) < len(nums) - 1
```

也可以在遍历 $\textit{nums}$ 的过程中去判断。

```go [sol1-Go]
func findSubarrays(nums []int) bool {
	vis := map[int]bool{}
	for i := 1; i < len(nums); i++ {
		s := nums[i-1] + nums[i]
		if vis[s] {
			return true
		}
		vis[s] = true
	}
	return false
}
```

#### 思考题

1. 如果把子数组的长度改为一个比较大的数字 $k$ 要怎么做？

2. 如果把子数组改成子序列要怎么做？
