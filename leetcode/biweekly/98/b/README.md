下午两点【biIibiIi@灵茶山艾府】直播讲题，记得关注哦~

---

根据题意，修改成 $\textit{nums}$ 中的数字，可以让最小得分为 $0$。那么分数就等于最大得分。

那么从小到大排序后，我们可以：

- 修改最大的两个数为 $\textit{nums}[n-3]$，最大得分为 $\textit{nums}[n-3]-\textit{nums}[0]$；
- 修改最小的为 $\textit{nums}[1]$，最大的为 $\textit{nums}[n-2]$，最大得分为 $\textit{nums}[n-2]-\textit{nums}[1]$；
- 修改最小的两个数为 $\textit{nums}[2]$，最大得分为 $\textit{nums}[n-1]-\textit{nums}[2]$。

这样修改的理由是，修改成再更大/更小的数，不会影响最大得分了。

附：[视频讲解](https://www.bilibili.com/video/BV15D4y1G7ms/)

```py [sol1-Python3]
class Solution:
    def minimizeSum(self, a: List[int]) -> int:
        a.sort()
        return min(a[-3] - a[0], a[-2] - a[1], a[-1] - a[2])
```

```go [sol1-Go]
func minimizeSum(a []int) int {
	sort.Ints(a)
	n := len(a)
	return min(min(a[n-3]-a[0], a[n-2]-a[1]), a[n-1]-a[2])
}

func min(a, b int) int { if a > b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。手动维护或者用快速选择可以做到 $O(n)$。
- 空间复杂度：$O(1)$。忽略排序时栈的开销，仅用到若干额外变量。
