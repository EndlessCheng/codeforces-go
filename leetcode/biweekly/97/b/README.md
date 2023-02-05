下午两点【bilibili@灵茶山艾府】直播讲题，记得关注哦~

---

贪心，优先选小的。

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

### 复杂度分析

- 时间复杂度：$O(m+n)$，其中 $m$ 为 $\textit{banned}$ 的长度。
- 空间复杂度：$O(m)$。
