[视频讲解](https://www.bilibili.com/video/BV1kR4y1r7Df/) 已出炉，欢迎点赞三连，在评论区分享你对这场双周赛的看法~

---

建图，枚举中心节点，选择至多 $k$ 个最大的值为正数的邻居。

```py [sol1-Python3]
class Solution:
    def maxStarSum(self, vals: List[int], edges: List[List[int]], k: int) -> int:
        g = [[] for _ in vals]
        for x, y in edges:
            if vals[y] > 0: g[x].append(vals[y])
            if vals[x] > 0: g[y].append(vals[x])
        return max(v + sum(nlargest(k, a)) for v, a in zip(vals, g))
```

```go [sol1-Go]
func maxStarSum(vals []int, edges [][]int, k int) int {
	g := make([]sort.IntSlice, len(vals))
	for _, e := range edges {
		x, y := e[0], e[1]
		if vals[y] > 0 {
			g[x] = append(g[x], vals[y])
		}
		if vals[x] > 0 {
			g[y] = append(g[y], vals[x])
		}
	}
	ans := math.MinInt32
	for i, a := range g {
		sort.Sort(sort.Reverse(a)) // 也可以用快速选择
		if len(a) > k {
			a = a[:k]
		}
		sum := vals[i]
		for _, v := range a {
			sum += v
		}
		ans = max(ans, sum)
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n+m\log n)$，其中 $n$ 为 $\textit{vals}$ 的长度，$m$ 为 $\textit{edges}$ 的长度。用快速选择算法可以做到 $O(n+m)$。
- 空间复杂度：$O(n+m)$。
