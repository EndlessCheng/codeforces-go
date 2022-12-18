[视频讲解](https://www.bilibili.com/video/BV1LW4y1T7if/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

把度数为奇数的节点记到 $\textit{odd}$ 中，记 $m$ 为 $\textit{odd}$ 的长度，分类讨论：

- 如果 $m=0$，那么已经符合要求。
- 如果 $m=2$，记 $x=\textit{odd}[0],y=\textit{odd}[1]$：
    - 如果 $x$ 和 $y$ 之间没有边，那么连边之后就符合要求了。
    - 如果 $x$ 和 $y$ 之间有边，那么枚举 $[1,n]$ 的所有不为 $x$ 和 $y$ 的点 $i$，由于 $i$ 的度数一定是偶数，如果 $i$ 和 $x$ 以及 $i$ 和 $y$ 之间没有边，那么连边之后就符合要求了。
- 如果 $m=4$，记 $a=\textit{odd}[0],b=\textit{odd}[1],c=\textit{odd}[2],d=\textit{odd}[3]$：
    - 如果 $a$ 和 $b$ 以及 $c$ 和 $d$ 之间没有边，那么连边之后就符合要求了。
    - 如果 $a$ 和 $c$ 以及 $b$ 和 $d$ 之间没有边，那么连边之后就符合要求了。
    - 如果 $a$ 和 $d$ 以及 $b$ 和 $c$ 之间没有边，那么连边之后就符合要求了。
- 其余情况无法满足要求。

```py [sol1-Python3]
class Solution:
    def isPossible(self, n: int, edges: List[List[int]]) -> bool:
        g, deg = defaultdict(set), Counter()
        for x, y in edges:
            g[x].add(y)
            g[y].add(x)
            deg[x] += 1
            deg[y] += 1
        odd = [i for i, d in deg.items() if d % 2]
        m = len(odd)
        if m == 0: return True
        if m == 2:
            x, y = odd
            return x not in g[y] or any(
                i != x and i != y and x not in g[i] and y not in g[i]
                for i in range(1, n + 1))
        if m == 4:
            a, b, c, d = odd
            return b not in g[a] and d not in g[c] or \
                   c not in g[a] and d not in g[b] or \
                   d not in g[a] and c not in g[b]
        return False
```

```go [sol1-Go]
func isPossible(n int, edges [][]int) bool {
	type pair struct{ x, y int }
	has := map[pair]bool{}
	deg := make([]int, n+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		has[pair{x, y}] = true
		has[pair{y, x}] = true
		deg[x]++
		deg[y]++
	}
	odd := []int{}
	for i, d := range deg {
		if d%2 > 0 {
			odd = append(odd, i)
		}
	}
	m := len(odd)
	if m == 0 {
		return true
	}
	if m == 2 {
		x, y := odd[0], odd[1]
		if !has[pair{x, y}] {
			return true
		}
		for i := 1; i <= n; i++ {
			if i != x && i != y && !has[pair{i, x}] && !has[pair{i, y}] {
				return true
			}
		}
		return false
	}
	if m == 4 {
		a, b, c, d := odd[0], odd[1], odd[2], odd[3]
		return !has[pair{a, b}] && !has[pair{c, d}] || !has[pair{a, c}] && !has[pair{b, d}] || !has[pair{a, d}] && !has[pair{b, c}]
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$O(n+m)$，其中 $m$ 为 $\textit{edges}$ 的长度。
- 空间复杂度：$O(n+m)$。
