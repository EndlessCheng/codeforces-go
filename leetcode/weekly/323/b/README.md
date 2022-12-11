[视频讲解](https://www.bilibili.com/video/BV1QK41167cr/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

由于数组元素至少为 $2$，平方 $k$ 次后，元素至少为 $2^{2^{k-1}}$。

因此只要暴力枚举初始项，不断平方即可，至多循环 $\log\log U$ 次，这里 $U=max(\textit{nums})$。

检查元素是否在数组中，可以用哈希表预处理。

```py [sol1-Python3]
class Solution:
    def longestSquareStreak(self, nums: List[int]) -> int:
        ans, s = 0, set(nums)
        for x in s:
            cnt = 0
            while x in s:
                cnt += 1
                x *= x
            ans = max(ans, cnt)
        return ans if ans > 1 else -1
```

```go [sol1-Go]
func longestSquareStreak(nums []int) (ans int) {
	set := map[int]bool{}
	for _, x := range nums {
		set[x] = true
	}
	for x := range set {
		cnt := 1
		for x *= x; set[x]; x *= x {
			cnt++
		}
		ans = max(ans, cnt)
	}
	if ans == 1 {
		return -1
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n\log\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
- 空间复杂度：$O(n)$。
