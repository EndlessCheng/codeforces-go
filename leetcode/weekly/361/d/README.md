请看 [视频讲解](https://www.bilibili.com/video/BV1Nj411178Z/) 第四题，或者阅读[【模板讲解】树上倍增算法（以及最近公共祖先）](https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solution/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw/)

对于本题，由于 $1\le w_i \le 26$，我们可以在倍增的同时，维护从节点 $x$ 到 $x$ 的第 $2^i$ 个祖先节点这条路径上的每种边权的个数。

对于每个询问，在计算 $a$ 和 $b$ 的最近公共祖先的同时，也同样地维护从 $a$ 到 $b$ 路径上的每种边权的个数 $\textit{cnt}$。

我们可以让出现次数最多的边权保持不变，设其个数为 $\textit{maxCnt}$，那么用从 $a$ 到 $b$ 路径长度减去 $\textit{maxCnt}$，就得到了最小操作次数。

路径长度可以用深度数组 $\textit{depth}$ 算出，即

$$
(\textit{depth}[a]- \textit{depth}[\textit{lca}]) + (\textit{depth}[b]- \textit{depth}[\textit{lca}])
$$

其中 $\textit{lca}$ 是 $a$ 和 $b$ 的最近公共祖先，上式对应着一条在 $\textit{lca}$ 拐弯的路径。

> 注：另一种做法是维护从根到节点 $x$ 的路径上的每种边权的出现次数，按照计算路径长度的思路，也可以通过加加减减算出路径上的每种边权的个数。
>
> 但是，如果把问题改成维护路径上的边权最大值，这种做法就不行了，而本题解的思路仍然是可以的。

```py [sol-Python3]
class Solution:
    def minOperationsQueries(self, n: int, edges: List[List[int]], queries: List[List[int]]) -> List[int]:
        g = [[] for _ in range(n)]
        for x, y, w in edges:
            g[x].append((y, w - 1))
            g[y].append((x, w - 1))

        m = n.bit_length()
        pa = [[-1] * m for _ in range(n)]
        cnt = [[[0] * 26 for _ in range(m)] for _ in range(n)]
        depth = [0] * n
        def dfs(x: int, fa: int) -> None:
            pa[x][0] = fa
            for y, w in g[x]:
                if y != fa:
                    cnt[y][0][w] = 1
                    depth[y] = depth[x] + 1
                    dfs(y, x)
        dfs(0, -1)

        # 倍增模板
        for i in range(m - 1):
            for x in range(n):
                p = pa[x][i]
                if p != -1:
                    pp = pa[p][i]
                    pa[x][i + 1] = pp
                    for j, (c1, c2) in enumerate(zip(cnt[x][i], cnt[p][i])):
                        cnt[x][i + 1][j] = c1 + c2

        ans = []
        for x, y in queries:
            path_len = depth[x] + depth[y]  # 最后减去 depth[lca] * 2
            cw = [0] * 26
            if depth[x] > depth[y]:
                x, y = y, x

            # 使 y 和 x 在同一深度
            k = depth[y] - depth[x]
            for i in range(k.bit_length()):
                if (k >> i) & 1:  # k 二进制从低到高第 i 位是 1
                    p = pa[y][i]
                    for j, c in enumerate(cnt[y][i]):
                        cw[j] += c
                    y = p

            if y != x:
                for i in range(m - 1, -1, -1):
                    px, py = pa[x][i], pa[y][i]
                    if px != py:
                        for j, (c1, c2) in enumerate(zip(cnt[x][i], cnt[y][i])):
                            cw[j] += c1 + c2
                        x, y = px, py  # 同时上跳 2^i 步
                for j, (c1, c2) in enumerate(zip(cnt[x][0], cnt[y][0])):
                    cw[j] += c1 + c2
                x = pa[x][0]

            lca = x
            path_len -= depth[lca] * 2
            ans.append(path_len - max(cw))
        return ans
```

```java [sol-Java]
class Solution {
    public int[] minOperationsQueries(int n, int[][] edges, int[][] queries) {
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1], w = e[2] - 1;
            g[x].add(new int[]{y, w});
            g[y].add(new int[]{x, w});
        }

        int m = 32 - Integer.numberOfLeadingZeros(n); // n 的二进制长度
        var pa = new int[n][m];
        for (int i = 0; i < n; i++) {
            Arrays.fill(pa[i], -1);
        }
        var cnt = new int[n][m][26];
        var depth = new int[n];
        dfs(0, -1, g, pa, cnt, depth);

        for (int i = 0; i < m - 1; i++) {
            for (int x = 0; x < n; x++) {
                int p = pa[x][i];
                if (p != -1) {
                    int pp = pa[p][i];
                    pa[x][i + 1] = pp;
                    for (int j = 0; j < 26; j++) {
                        cnt[x][i + 1][j] = cnt[x][i][j] + cnt[p][i][j];
                    }
                }
            }
        }

        var ans = new int[queries.length];
        for (int qi = 0; qi < queries.length; qi++) {
            int x = queries[qi][0], y = queries[qi][1];
            int pathLen = depth[x] + depth[y];
            var cw = new int[26];
            if (depth[x] > depth[y]) {
                int temp = x;
                x = y;
                y = temp;
            }

            // 让 y 和 x 在同一深度
            for (int k = depth[y] - depth[x]; k > 0; k &= k - 1) {
                int i = Integer.numberOfTrailingZeros(k);
                int p = pa[y][i];
                for (int j = 0; j < 26; ++j) {
                    cw[j] += cnt[y][i][j];
                }
                y = p;
            }

            if (y != x) {
                for (int i = m - 1; i >= 0; i--) {
                    int px = pa[x][i];
                    int py = pa[y][i];
                    if (px != py) {
                        for (int j = 0; j < 26; j++) {
                            cw[j] += cnt[x][i][j] + cnt[y][i][j];
                        }
                        x = px;
                        y = py; // x 和 y 同时上跳 2^i 步
                    }
                }
                for (int j = 0; j < 26; j++) {
                    cw[j] += cnt[x][0][j] + cnt[y][0][j];
                }
                x = pa[x][0];
            }

            int lca = x;
            pathLen -= depth[lca] * 2;
            int maxCw = 0;
            for (int i = 0; i < 26; i++) {
                maxCw = Math.max(maxCw, cw[i]);
            }
            ans[qi] = pathLen - maxCw;
        }
        return ans;
    }

    private void dfs(int x, int fa, List<int[]>[] g, int[][] pa, int[][][] cnt, int[] depth) {
        pa[x][0] = fa;
        for (var e : g[x]) {
            int y = e[0], w = e[1];
            if (y != fa) {
                cnt[y][0][w] = 1;
                depth[y] = depth[x] + 1;
                dfs(y, x, g, pa, cnt, depth);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> minOperationsQueries(int n, vector<vector<int>> &edges, vector<vector<int>> &queries) {
        vector<vector<pair<int, int>>> g(n);
        for (auto &e: edges) {
            int x = e[0], y = e[1], w = e[2] - 1;
            g[x].emplace_back(y, w);
            g[y].emplace_back(x, w);
        }

        int m = 32 - __builtin_clz(n); // n 的二进制长度
        vector<vector<int>> pa(n, vector<int>(m, -1));
        vector<vector<array<int, 26>>> cnt(n, vector<array<int, 26>>(m));
        vector<int> depth(n);
        function<void(int, int)> dfs = [&](int x, int fa) {
            pa[x][0] = fa;
            for (auto [y, w]: g[x]) {
                if (y != fa) {
                    cnt[y][0][w] = 1;
                    depth[y] = depth[x] + 1;
                    dfs(y, x);
                }
            }
        };
        dfs(0, -1);

        for (int i = 0; i < m - 1; i++) {
            for (int x = 0; x < n; x++) {
                int p = pa[x][i];
                if (p != -1) {
                    int pp = pa[p][i];
                    pa[x][i + 1] = pp;
                    for (int j = 0; j < 26; ++j) {
                        cnt[x][i + 1][j] = cnt[x][i][j] + cnt[p][i][j];
                    }
                }
            }
        }

        vector<int> ans;
        for (auto &q: queries) {
            int x = q[0], y = q[1];
            int path_len = depth[x] + depth[y]; // 最后减去 depth[lca] * 2
            int cw[26]{};
            if (depth[x] > depth[y]) {
                swap(x, y);
            }

            // 让 y 和 x 在同一深度
            for (int k = depth[y] - depth[x]; k; k &= k - 1) {
                int i = __builtin_ctz(k);
                int p = pa[y][i];
                for (int j = 0; j < 26; ++j) {
                    cw[j] += cnt[y][i][j];
                }
                y = p;
            }

            if (y != x) {
                for (int i = m - 1; i >= 0; i--) {
                    int px = pa[x][i], py = pa[y][i];
                    if (px != py) {
                        for (int j = 0; j < 26; j++) {
                            cw[j] += cnt[x][i][j] + cnt[y][i][j];
                        }
                        x = px;
                        y = py; // x 和 y 同时上跳 2^i 步
                    }
                }
                for (int j = 0; j < 26; j++) {
                    cw[j] += cnt[x][0][j] + cnt[y][0][j];
                }
                x = pa[x][0];
            }

            int lca = x;
            path_len -= depth[lca] * 2;
            ans.push_back(path_len - *max_element(cw, cw + 26));
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]-1
		g[v] = append(g[v], edge{w, wt})
		g[w] = append(g[w], edge{v, wt})
	}
	
	const mx = 14 // 2^14 > 10^4
	type pair struct {
		p   int
		cnt [26]int
	}
	pa := make([][mx]pair, n)
	depth := make([]int, n)
	var build func(int, int, int)
	build = func(v, p, d int) {
		pa[v][0].p = p
		depth[v] = d
		for _, e := range g[v] {
			if w := e.to; w != p {
				pa[w][0].cnt[e.wt] = 1
				build(w, v, d+1)
			}
		}
	}
	build(0, -1, 0)

	// 倍增模板
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p.p != -1 {
				pp := pa[p.p][i]
				pa[v][i+1].p = pp.p
				for j := 0; j < 26; j++ {
					pa[v][i+1].cnt[j] = p.cnt[j] + pp.cnt[j]
				}
			} else {
				pa[v][i+1].p = -1
			}
		}
	}

	// 计算 LCA 模板（这里返回最小操作次数）
	// https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solution/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw/
	f := func(v, w int) int {
		pathLen := depth[v] + depth[w] // 最后减去 depth[lca] * 2
		cnt := [26]int{}
		if depth[v] > depth[w] {
			v, w = w, v
		}
		for i := 0; i < mx; i++ {
			if (depth[w]-depth[v])>>i&1 > 0 {
				p := pa[w][i]
				for j := 0; j < 26; j++ {
					cnt[j] += p.cnt[j]
				}
				w = p.p
			}
		}
		if w != v {
			for i := mx - 1; i >= 0; i-- {
				if pv, pw := pa[v][i], pa[w][i]; pv.p != pw.p {
					for j := 0; j < 26; j++ {
						cnt[j] += pv.cnt[j] + pw.cnt[j]
					}
					v, w = pv.p, pw.p
				}
			}
			for j := 0; j < 26; j++ {
				cnt[j] += pa[v][0].cnt[j] + pa[w][0].cnt[j]
			}
			v = pa[v][0].p
		}
		// 现在 v 是 LCA
		return pathLen - depth[v] * 2 - slices.Max(cnt[:])
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = f(q[0], q[1])
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+q)U\log n)$，其中 $q$ 为 $\textit{queries}$ 的长度，$U$ 为边权种类数。
- 空间复杂度：$\mathcal{O}(nU\log n)$。返回值的长度不计入。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
