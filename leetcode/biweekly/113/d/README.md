[视频讲解](https://www.bilibili.com/video/BV1PV411N76R/) 第四题。

换根 DP 原理：[【图解】一张图秒懂换根 DP！](https://leetcode.cn/problems/sum-of-distances-in-tree/solution/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/)

本题可以先计算出以 $0$ 为根时的答案：在建图的时候，对于一条 $x\rightarrow y$ 的边，把 $(y,1)$ 加到 $x$ 的邻居，把 $(x,-1)$ 加到 $y$ 的邻居，从而可以在递归过程中统计有多少条边是需要反向的。

然后换根，假设 $y$ 是 $x$ 的儿子节点，从 $x$ 换根到 $y$，只会影响 $x$ 和 $y$ 的父子关系，其余节点不受影响，所以

$$
\textit{ans}[y] = \textit{ans}[x] + \textit{direction}
$$

如果从 $x$ 到 $y$ 不需要反向，则换根后需要反向，$\textit{direction}=1$，否则等于 $-1$，这正好就是建图时我们添加的 $1$ 和 $-1$。

```py [sol-Python3]
class Solution:
    def minEdgeReversals(self, n: int, edges: List[List[int]]) -> List[int]:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append((y, 1))
            g[y].append((x, -1))  # 从 y 到 x 需要反向

        ans = [0] * n
        def dfs(x: int, fa: int) -> None:
            for y, dir in g[x]:
                if y != fa:
                    ans[0] += dir < 0
                    dfs(y, x)
        dfs(0, -1)

        def reroot(x: int, fa: int) -> None:
            for y, dir in g[x]:
                if y != fa:
                    ans[y] = ans[x] + dir  # dir 就是从 x 换到 y 的「变化量」
                    reroot(y, x)
        reroot(0, -1)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] minEdgeReversals(int n, int[][] edges) {
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(new int[]{y, 1});
            g[y].add(new int[]{x, -1}); // 从 y 到 x 需要反向
        }

        var ans = new int[n];
        dfs(0, -1, g, ans);
        reroot(0, -1, g, ans);
        return ans;
    }

    private void dfs(int x, int fa, List<int[]>[] g, int[] ans) {
        for (var e : g[x]) {
            int y = e[0], dir = e[1];
            if (y != fa) {
                if (dir < 0) {
                    ans[0]++;
                }
                dfs(y, x, g, ans);
            }
        }
    }

    private void reroot(int x, int fa, List<int[]>[] g, int[] ans) {
        for (var e : g[x]) {
            int y = e[0], dir = e[1];
            if (y != fa) {
                ans[y] = ans[x] + dir; // dir 就是从 x 换到 y 的「变化量」
                reroot(y, x, g, ans);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<vector<pair<int, int>>> g;
    vector<int> ans;

    void dfs(int x, int fa) {
        for (auto &[y, dir] : g[x]) {
            if (y != fa) {
                ans[0] += dir < 0;
                dfs(y, x);
            }
        }
    }

    void reroot(int x, int fa) {
        for (auto &[y, dir] : g[x]) {
            if (y != fa) {
                ans[y] = ans[x] + dir; // dir 就是从 x 换到 y 的「变化量」
                reroot(y, x);
            }
        }
    }

public:
    vector<int> minEdgeReversals(int n, vector<vector<int>> &edges) {
        g.resize(n);
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            g[x].emplace_back(y, 1);
            g[y].emplace_back(x, -1); // 从 y 到 x 需要反向
        }

        ans.resize(n);
        dfs(0, -1);
        reroot(0, -1);
        return ans;
    }
};
```

```go [sol-Go]
func minEdgeReversals(n int, edges [][]int) (ans []int) {
	type pair struct{ to, dir int }
	g := make([][]pair, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], pair{y, 1})
		g[y] = append(g[y], pair{x, -1}) // 从 y 到 x 需要反向
	}

	ans = make([]int, n)
	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, e := range g[x] {
			y := e.to
			if y != fa {
				if e.dir < 0 {
					ans[0]++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)

	var reroot func(int, int)
	reroot = func(x, fa int) {
		for _, e := range g[x] {
			y := e.to
			if y != fa {
				ans[y] = ans[x] + e.dir // e.dir 就是从 x 换到 y 的「变化量」
				reroot(y, x)
			}
		}
	}
	reroot(0, -1)
	return ans
}
```

```js [sol-JavaScript]
var minEdgeReversals = function (n, edges) {
    const g = new Array(n).fill(null).map(() => []);
    for (const [x, y] of edges) {
        g[x].push([y, 1]);
        g[y].push([x, -1]); // 从 y 到 x 需要反向
    }

    const ans = new Array(n).fill(0);
    function dfs(x, fa) {
        for (const [y, dir] of g[x]) {
            if (y !== fa) {
                ans[0] += dir < 0;
                dfs(y, x);
            }
        }
    }
    dfs(0, -1);

    function reroot(x, fa) {
        for (const [y, dir] of g[x]) {
            if (y !== fa) {
                ans[y] = ans[x] + dir; // dir 就是从 x 换到 y 的「变化量」
                reroot(y, x);
            }
        }
    }
    reroot(0, -1);
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 和本题非常像的题目

- [2581. 统计可能的树根数目](https://leetcode.cn/problems/count-number-of-possible-root-nodes/)
