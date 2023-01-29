下午两点【bilibili@灵茶山艾府】直播讲题，记得关注哦~

---

### 提示 1

问题相当于把 $\textit{weights}$ 划分成 $k$ 个连续子数组，分数等于每个子数组的两端的值之和。

### 提示 2

$\textit{weights}[0]$ 和 $\textit{weights}[n-1]$ 一定在分数中。

上一个子数组的末尾和下一个子数组的开头一定**同时**在分数中。

### 提示 3

把所有 $n-1$ 个 $\textit{weights}[i]+\textit{weights}[i+1]$ 算出来，排序，那么最大的 $k-1$ 个数和最小的 $k-1$ 个数相减，即为答案。

```py [sol1-Python3]
class Solution:
    def putMarbles(self, wt: List[int], k: int) -> int:
        for i in range(len(wt) - 1):
            wt[i] += wt[i + 1]
        wt.pop()
        wt.sort()
        return sum(wt[len(wt) - k + 1:]) - sum(wt[:k - 1])
```

```go [sol1-Go]
func putMarbles(wt []int, k int) (ans int64) {
	for i, w := range wt[1:] {
		wt[i] += w
	}
	wt = wt[:len(wt)-1]
	sort.Ints(wt)
	for _, w := range wt[len(wt)-k+1:] {
		ans += int64(w)
	}
	for _, w := range wt[:k-1] {
		ans -= int64(w)
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。忽略排序的栈空间，仅用到若干额外变量。
