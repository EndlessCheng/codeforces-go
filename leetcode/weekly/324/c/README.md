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
        g = defaultdict(set)
        for x, y in edges:
            g[x].add(y)
            g[y].add(x)
        odd = [i for i, nb in g.items() if len(nb) % 2]
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

```java [sol1-Java]
class Solution {
    public boolean isPossible(int n, List<List<Integer>> edges) {
        var g = new Set[n + 1];
        Arrays.setAll(g, e -> new HashSet<Integer>());
        for (var e : edges) {
            int x = e.get(0), y = e.get(1);
            g[x].add(y);
            g[y].add(x);
        }
        var odd = new ArrayList<Integer>();
        for (var i = 1; i <= n; ++i)
            if (g[i].size() % 2 > 0) odd.add(i);
        var m = odd.size();
        if (m == 0) return true;
        if (m == 2) {
            int x = odd.get(0), y = odd.get(1);
            if (!g[x].contains(y)) return true;
            for (var i = 1; i <= n; ++i)
                if (i != x && i != y && !g[i].contains(x) && !g[i].contains(y))
                    return true;
            return false;
        }
        if (m == 4) {
            int a = odd.get(0), b = odd.get(1), c = odd.get(2), d = odd.get(3);
            return !g[a].contains(b) && !g[c].contains(d) ||
                    !g[a].contains(c) && !g[b].contains(d) ||
                    !g[a].contains(d) && !g[b].contains(c);
        }
        return false;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    bool isPossible(int n, vector<vector<int>> &edges) {
        unordered_set<int> g[n + 1];
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            g[x].insert(y);
            g[y].insert(x);
        }
        vector<int> odd;
        for (int i = 1; i <= n; ++i)
            if (g[i].size() % 2) odd.push_back(i);
        int m = odd.size();
        if (m == 0) return true;
        if (m == 2) {
            int x = odd[0], y = odd[1];
            if (!g[x].count(y)) return true;
            for (int i = 1; i <= n; ++i)
                if (i != x && i != y && !g[i].count(x) && !g[i].count(y))
                    return true;
            return false;
        }
        if (m == 4) {
            int a = odd[0], b = odd[1], c = odd[2], d = odd[3];
            return !g[a].count(b) && !g[c].count(d) ||
                   !g[a].count(c) && !g[b].count(d) ||
                   !g[a].count(d) && !g[b].count(c);
        }
        return false;
    }
};
```

```go [sol1-Go]
func isPossible(n int, edges [][]int) bool {
	g := map[int]map[int]bool{}
	for _, e := range edges {
		x, y := e[0], e[1]
		if g[x] == nil {
			g[x] = map[int]bool{}
		}
		g[x][y] = true
		if g[y] == nil {
			g[y] = map[int]bool{}
		}
		g[y][x] = true
	}
	odd := []int{}
	for i, nb := range g {
		if len(nb)%2 > 0 {
			odd = append(odd, i)
		}
	}
	m := len(odd)
	if m == 0 {
		return true
	}
	if m == 2 {
		x, y := odd[0], odd[1]
		if !g[x][y] {
			return true
		}
		for i := 1; i <= n; i++ {
			if i != x && i != y && !g[i][x] && !g[i][y] {
				return true
			}
		}
		return false
	}
	if m == 4 {
		a, b, c, d := odd[0], odd[1], odd[2], odd[3]
		return !g[a][b] && !g[c][d] || !g[a][c] && !g[b][d] || !g[a][d] && !g[b][c]
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$O(n+m)$，其中 $m$ 为 $\textit{edges}$ 的长度。
- 空间复杂度：$O(n+m)$。
