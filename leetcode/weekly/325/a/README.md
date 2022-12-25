欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)，下午两点在B站讲这场周赛的题目。

---

直接遍历每个 $\textit{words}[i]$，如果它等于 $\textit{target}$，那么用 

$$
\min(|i-\textit{startIndex}|, n-|i-\textit{startIndex}|)
$$

更新答案的最小值。

```py [sol1-Python3]
class Solution:
    def closetTarget(self, words: List[str], target: str, startIndex: int) -> int:
        ans = n = len(words)
        for i, w in enumerate(words):
            if w == target:
                ans = min(ans, abs(i - startIndex), n - abs(i - startIndex))
        return ans if ans < n else -1
```

```go [sol1-Go]
func closetTarget(words []string, target string, startIndex int) int {
	n := len(words)
	ans := n
	for i, s := range words {
		if s == target {
			ans = min(ans, min(abs(i-startIndex), n-abs(i-startIndex)))
		}
	}
	if ans == n {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$O(nL)$，其中 $n$ 为 $\textit{words}$ 的长度，$L$ 为 $\textit{target}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
