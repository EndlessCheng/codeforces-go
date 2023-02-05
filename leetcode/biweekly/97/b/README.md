贪心，优先选小的。

为了快速判断一个数是否在数组 $\textit{banned}$ 中，可以将其转换成哈希表。

附：[视频讲解](https://www.bilibili.com/video/BV1rM4y1X7z9/)

```py [sol1-Python3]
class Solution:
    def maxCount(self, banned: List[int], n: int, maxSum: int) -> int:
        ans, s = 0, set(banned)
        for i in range(1, n + 1):
            if i > maxSum: break
            if i not in s:
                maxSum -= i
                ans += 1
        return ans
```

```go [sol1-Go]
func maxCount(banned []int, n, maxSum int) (ans int) {
	has := map[int]bool{}
	for _, v := range banned {
		has[v] = true
	}
	for i := 1; i <= n && i <= maxSum; i++ {
		if !has[i] {
			maxSum -= i
			ans++
		}
	}
	return
}
```

### 思考题

如果 $\textit{n}$ 和 $\textit{maxSum}$ 都很大呢（比如达到 $10^{18}$）？

欢迎在评论区发表你的做法。

### 复杂度分析

- 时间复杂度：$O(m+n)$，其中 $m$ 为 $\textit{banned}$ 的长度。
- 空间复杂度：$O(m)$。
