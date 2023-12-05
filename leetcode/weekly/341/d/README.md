## 方法一：暴力 DFS 每条路径

对每个 $\textit{trips}[i]$ 都 DFS 一次这棵树，在 DFS 的过程中，把从 $\textit{start}$ 到 $\textit{end}$ 的路径上的每个点 $x$ 的经过次数 $\textit{cnt}[x]$ 都加一。

既然知道了每个点会被经过多少次，把 $\textit{price}[i]$ 更新成 $\textit{price}[i]\cdot \textit{cnt}[i]$，问题就转换成计算减半后的 $\textit{price}[i]$ 之和的最小值。注意 $\textit{cnt}[i]=0$ 时 $\textit{price}[i]$ 会被更新成 $0$，我们无需考虑没有经过的节点。

对于转换后的问题，做法和 [337. 打家劫舍 III](https://leetcode.cn/problems/house-robber-iii/) 是类似的，请看视频讲解：[【基础算法精讲 24】](https://www.bilibili.com/video/BV1vu4y1f7dn/)

我们随便选一个节点出发 DFS（比如节点 $0$）。在 DFS 的过程中，对于节点 $x$ 及其儿子 $y$，分类讨论：

- 如果 $\textit{price}[x]$ 不变，那么 $\textit{price}[y]$ 可以减半，也可以不变，取这两种情况的最小值；
- 如果 $\textit{price}[x]$ 减半，那么 $\textit{price}[y]$ 只能不变。

因此子树 $x$ 需要返回两个值：

- $\textit{price}[x]$ 不变时的子树 $x$ 的最小价值总和；
- $\textit{price}[x]$ 减半时的子树 $x$ 的最小价值总和。

答案就是根节点不变/减半的最小值。

#### 答疑

**问**：代码实现时，如何找到从 $\textit{start}$ 到 $\textit{end}$ 的路径？

**答**：以 $\textit{start}$ 为树根 DFS，找到 $\textit{end}$ 时，$\textit{end}$ 及其祖先节点就恰好组成了从 $\textit{start}$ 到 $\textit{end}$ 的路径。据此可以在递归的「归」当中去更新 $\textit{cnt}$。

```py [sol-Python3]
class Solution:
    def minimumTotalPrice(self, n: int, edges: List[List[int]], price: List[int], trips: List[List[int]]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        cnt = [0] * n
        for start, end in trips:
            def dfs(x: int, fa: int) -> bool:
                if x == end:
                    cnt[x] += 1
                    return True  # 找到 end
                for y in g[x]:
                    if y != fa and dfs(y, x):
                        cnt[x] += 1  # x 是 end 的祖先节点，也就在路径上
                        return True
                return False  # 未找到 end
            dfs(start, -1)

        # 类似 337. 打家劫舍 III
        def dfs(x: int, fa: int) -> (int, int):
            not_halve = price[x] * cnt[x]  # x 不变
            halve = not_halve // 2  # x 减半
            for y in g[x]:
                if y != fa:
                    nh, h = dfs(y, x)  # 计算 y 不变/减半的最小价值总和
                    not_halve += min(nh, h)  # x 不变，那么 y 可以不变或者减半，取这两种情况的最小值
                    halve += nh  # x 减半，那么 y 只能不变
            return not_halve, halve
        return min(dfs(0, -1))
```

```java [sol-Java]
class Solution {
    private List<Integer>[] g;
    private int[] price, cnt;
    private int end;

    public int minimumTotalPrice(int n, int[][] edges, int[] price, int[][] trips) {
        g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        cnt = new int[n];
        for (int[] t : trips) {
            end = t[1];
            dfs(t[0], -1);
        }

        this.price = price;
        int[] res = dp(0, -1);
        return Math.min(res[0], res[1]);
    }

    private boolean dfs(int x, int fa) {
        if (x == end) {
            cnt[x]++;
            return true; // 找到 end
        }
        for (int y : g[x]) {
            if (y != fa && dfs(y, x)) {
                cnt[x]++; // x 是 end 的祖先节点，也就在路径上
                return true;
            }
        }
        return false; // 未找到 end
    }

    // 类似 337. 打家劫舍 III
    private int[] dp(int x, int fa) {
        int notHalve = price[x] * cnt[x]; // x 不变
        int halve = notHalve / 2; // x 减半
        for (int y : g[x]) {
            if (y != fa) {
                int[] res = dp(y, x); // 计算 y 不变/减半的最小价值总和
                notHalve += Math.min(res[0], res[1]); // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
                halve += res[0]; // x 减半，那么 y 只能不变
            }
        }
        return new int[]{notHalve, halve};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumTotalPrice(int n, vector<vector<int>> &edges, vector<int> &price, vector<vector<int>> &trips) {
        vector<vector<int>> g(n);
        for (auto &e: edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x); // 建树
        }

        vector<int> cnt(n);
        for (auto &t: trips) {
            int end = t[1];
            function<bool(int, int)> dfs = [&](int x, int fa) -> bool {
                if (x == end) {
                    cnt[x]++;
                    return true; // 找到 end
                }
                for (int y: g[x]) {
                    if (y != fa && dfs(y, x)) {
                        cnt[x]++; // x 是 end 的祖先节点，也就在路径上
                        return true;
                    }
                }
                return false; // 未找到 end
            };
            dfs(t[0], -1);
        }

        // 类似 337. 打家劫舍 III
        function<pair<int, int>(int, int)> dfs = [&](int x, int fa) -> pair<int, int> {
            int not_halve = price[x] * cnt[x]; // x 不变
            int halve = not_halve / 2; // x 减半
            for (int y: g[x]) {
                if (y != fa) {
                    auto [nh, h] = dfs(y, x); // 计算 y 不变/减半的最小价值总和
                    not_halve += min(nh, h); // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
                    halve += nh; // x 减半，那么 y 只能不变
                }
            }
            return {not_halve, halve};
        };
        auto [nh, h] = dfs(0, -1);
        return min(nh, h);
    }
};
```

```go [sol-Go]
func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	cnt := make([]int, n)
	for _, t := range trips {
		end := t[1]
		var dfs func(int, int) bool
		dfs = func(x, fa int) bool {
			if x == end {
				cnt[x]++
				return true // 找到 end
			}
			for _, y := range g[x] {
				if y != fa && dfs(y, x) {
					cnt[x]++ // x 是 end 的祖先节点，也就在路径上
					return true
				}
			}
			return false // 未找到 end
		}
		dfs(t[0], -1)
	}

	// 类似 337. 打家劫舍 III
	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		notHalve := price[x] * cnt[x] // x 不变
		halve := notHalve / 2 // x 减半
		for _, y := range g[x] {
			if y != fa {
				nh, h := dfs(y, x) // 计算 y 不变/减半的最小价值总和
				notHalve += min(nh, h) // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
				halve += nh // x 减半，那么 y 只能不变
			}
		}
		return notHalve, halve
	}
	return min(dfs(0, -1))
}
```

```js [sol-JavaScript]
var minimumTotalPrice = function (n, edges, price, trips) {
    const g = Array(n).fill(null).map(() => []);
    for (const [x, y] of edges) {
        g[x].push(y);
        g[y].push(x);
    }

    const cnt = Array(n).fill(0);
    for (const [start, end] of trips) {
        function dfs(x, fa) {
            if (x === end) {
                cnt[x]++;
                return true; // 找到 end
            }
            for (const y of g[x]) {
                if (y !== fa && dfs(y, x)) {
                    cnt[x]++; // x 是 end 的祖先节点，也就在路径上
                    return true;
                }
            }
            return false; // 未找到 end
        }
        dfs(start, -1);
    }

    // 类似 337. 打家劫舍 III
    function dp(x, fa) {
        let not_halve = price[x] * cnt[x]; // x 不变
        let halve = Math.floor(not_halve / 2); // x 减半
        for (const y of g[x]) {
            if (y !== fa) {
                const [nh, h] = dp(y, x); // 计算 y 不变/减半的最小价值总和
                not_halve += Math.min(nh, h); // x 不变，那么 y 可以不变或者减半，取这两种情况的最小值
                halve += nh; // x 减半，那么 y 只能不变
            }
        }
        return [not_halve, halve];
    }
    return Math.min(...dp(0, -1));
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_total_price(n: i32, edges: Vec<Vec<i32>>, price: Vec<i32>, trips: Vec<Vec<i32>>) -> i32 {
        let n = n as usize;
        let mut g = vec![vec![]; n];
        for e in &edges {
            let x = e[0] as usize;
            let y = e[1] as usize;
            g[x].push(y);
            g[y].push(x);
        }

        fn dfs(x: usize, fa: usize, cnt: &mut Vec<i32>, g: &Vec<Vec<usize>>, end: usize) -> bool {
            if x == end {
                cnt[x] += 1;
                return true; // 找到 end
            }
            for &y in &g[x] {
                if y != fa && dfs(y, x, cnt, g, end) {
                    cnt[x] += 1; // x 是 end 的祖先节点，也就在路径上
                    return true;
                }
            }
            false // 未找到 end
        }
        let mut cnt = vec![0; n];
        for t in &trips {
            dfs(t[0] as usize, n, &mut cnt, &g, t[1] as usize);
        }

        // 类似 337. 打家劫舍 III
        fn dp(x: usize, fa: usize, price: &Vec<i32>, cnt: &Vec<i32>, g: &Vec<Vec<usize>>) -> (i32, i32) {
            let mut not_halve = price[x] * cnt[x]; // x 不变
            let mut halve = not_halve / 2; // x 减半
            for &y in &g[x] {
                if y != fa {
                    let (nh, h) = dp(y, x, price, cnt, g); // 计算 y 不变/减半的最小价值总和
                    not_halve += nh.min(h); // x 不变，那么 y 可以不变或者减半，取这两种情况的最小值
                    halve += nh; // x 减半，那么 y 只能不变
                }
            }
            (not_halve, halve)
        }
        let (nh, h) = dp(0, 0, &price, &cnt, &g);
        nh.min(h)
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $m$ 为 $\textit{trips}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二（选读）：Tarjan 离线 LCA + 树上差分

**前置知识**：并查集、Tarjan 离线求 LCA、树上差分。

数组上的区间加一操作，我们可以用 [差分数组](https://leetcode.cn/circle/discuss/FfMCgb/) 解决（请至少完成一道差分数组题目再往下读）。这一思想同样可以用到树上，把树上的一条路径上的节点值加一，也可以用差分数组解决。

从 $x=\textit{start}$ 到 $y=\textit{end}$ 的路径可以视作从 $x$ 向上到某个点「拐弯」，再向下到达 $y$。

这个拐弯的点是 $x$ 和 $y$ 的 $\textit{lca}$（最近公共祖先）。注意拐弯的点也可能就是 $x$ 或 $y$。

设路径为 $x-z-\textit{lca}-y$，其中 $z$ 是 $\textit{lca}$ 往 $x$ 方向的儿子。由于更新的是点，拆分成 $x-z$ 和 $y-\textit{lca}$ 这两段路径。

把路径上的点的 $\textit{cnt}$ 加一，转换成对差分数组 $\textit{diff}$ 的两个数的更新。规定把下面的点加一，把上面的点减一：

- 对于 $x-z$，把 $\textit{diff}[x]$ 加一，$\textit{diff}[\textit{lca}]$ 减一。注意，如果 $x$ 就是 $\textit{lca}$，那么 $z$ 是不存在的，而差分操作刚好对 $\textit{diff}[x]$ 加一再减一，没有变化。所以我们无需特判 $x$ 就是 $\textit{lca}$ 的情况。
- 对于 $y-\textit{lca}$，把 $\textit{diff}[y]$ 加一，$\textit{diff}[\textit{father}[\textit{lca}]]$ 减一，其中 $\textit{father}[\textit{lca}]$ 表示 $\textit{lca}$ 的父节点。

最近公共祖先 $\textit{lca}$ 可以用 Tarjan 离线算法计算，见代码注释。

更新完 $\textit{diff}$ 后，DFS 这棵树，在递归的「归」的过程中自底向上累加 $\textit{diff}$，计算出 $\textit{cnt}$ 值。这个过程可以和计算答案的过程合在一起。

注：求最近公共祖先不止一种方式，也可以用树上倍增实现，请看 [【模板讲解】树上倍增](https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solution/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw/)

```py [sol-Python3]
class Solution:
    def minimumTotalPrice(self, n: int, edges: List[List[int]], price: List[int], trips: List[List[int]]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        qs = [[] for _ in range(n)]
        for s, e in trips:
            qs[s].append(e)  # 路径端点分组
            if s != e:
                qs[e].append(s)

        # 并查集模板
        root = list(range(n))
        def find(x: int) -> int:
            if x != root[x]:
                root[x] = find(root[x])
            return root[x]

        diff = [0] * n
        father = [0] * n
        color = [0] * n
        def tarjan(x: int, fa: int) -> None:
            father[x] = fa
            color[x] = 1  # 递归中
            for y in g[x]:
                if color[y] == 0:  # 未递归
                    tarjan(y, x)
                    root[y] = x  # 相当于把 y 的子树节点全部 merge 到 x
            for y in qs[x]:
                # color[y] == 2 意味着 y 所在子树已经遍历完
                # 也就意味着 y 已经 merge 到它和 x 的 lca 上了
                # 此时 find(y) 就是 x 和 y 的 lca
                if y == x or color[y] == 2:
                    diff[x] += 1
                    diff[y] += 1
                    lca = find(y)
                    diff[lca] -= 1
                    if father[lca] >= 0:
                        diff[father[lca]] -= 1
            color[x] = 2  # 递归结束
        tarjan(0, -1)

        def dfs(x: int, fa: int) -> (int, int, int):
            not_halve, halve, cnt = 0, 0, diff[x]
            for y in g[x]:
                if y != fa:
                    nh, h, c = dfs(y, x)  # 计算 y 不变/减半的最小价值总和
                    not_halve += min(nh, h)  # x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
                    halve += nh  # x 减半，那么 y 只能不变
                    cnt += c  # 自底向上累加差分值
            not_halve += price[x] * cnt  # x 不变
            halve += price[x] * cnt // 2  # x 减半
            return not_halve, halve, cnt
        return min(dfs(0, -1)[:2])
```

```java [sol-Java]
class Solution {
    private List<Integer>[] g, qs;
    private int[] diff, father, color, price;

    public int minimumTotalPrice(int n, int[][] edges, int[] price, int[][] trips) {
        g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        qs = new ArrayList[n];
        Arrays.setAll(qs, e -> new ArrayList<>());
        for (int[] t : trips) {
            int x = t[0], y = t[1];
            qs[x].add(y); // 路径端点分组
            if (x != y) {
                qs[y].add(x);
            }
        }

        root = new int[n];
        for (int i = 1; i < n; i++) {
            root[i] = i;
        }

        diff = new int[n];
        father = new int[n];
        color = new int[n];
        tarjan(0, -1);

        this.price = price;
        int[] res = dfs(0, -1);
        return Math.min(res[0], res[1]);
    }

    // 并查集模板
    private int[] root;

    private int find(int x) {
        if (root[x] != x) {
            root[x] = find(root[x]);
        }
        return root[x];
    }

    private void tarjan(int x, int fa) {
        father[x] = fa;
        color[x] = 1; // 递归中
        for (int y : g[x]) {
            if (color[y] == 0) { // 未递归
                tarjan(y, x);
                root[y] = x; // 相当于把 y 的子树节点全部 merge 到 x
            }
        }
        for (int y : qs[x]) {
            // color[y] == 2 意味着 y 所在子树已经遍历完
            // 也就意味着 y 已经 merge 到它和 x 的 lca 上了
            // 此时 find(y) 就是 x 和 y 的 lca
            if (y == x || color[y] == 2) {
                diff[x]++;
                diff[y]++;
                int lca = find(y);
                diff[lca]--;
                int f = father[lca];
                if (f >= 0) {
                    diff[f]--;
                }
            }
        }
        color[x] = 2; // 递归结束
    }

    private int[] dfs(int x, int fa) {
        int notHalve = 0, halve = 0, cnt = diff[x];
        for (int y : g[x]) {
            if (y != fa) {
                int[] res = dfs(y, x); // 计算 y 不变/减半的最小价值总和
                notHalve += Math.min(res[0], res[1]); // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
                halve += res[0]; // x 减半，那么 y 只能不变
                cnt += res[2]; // 自底向上累加差分值
            }
        }
        notHalve += price[x] * cnt; // x 不变
        halve += price[x] * cnt / 2; // x 减半
        return new int[]{notHalve, halve, cnt};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumTotalPrice(int n, vector<vector<int>> &edges, vector<int> &price, vector<vector<int>> &trips) {
        vector<vector<int>> g(n);
        for (auto &e: edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x); // 建树
        }

        vector<vector<int>> qs(n);
        for (auto &t: trips) {
            int x = t[0], y = t[1];
            qs[x].push_back(y); // 路径端点分组
            if (x != y) {
                qs[y].push_back(x);
            }
        }

        // 并查集模板
        vector<int> root(n);
        iota(root.begin(), root.end(), 0);
        function<int(int)> find = [&](int x) -> int { return root[x] == x ? x : root[x] = find(root[x]); };

        vector<int> diff(n), father(n), color(n);
        function<void(int, int)> tarjan = [&](int x, int fa) {
            father[x] = fa;
            color[x] = 1; // 递归中
            for (int y: g[x]) {
                if (color[y] == 0) { // 未递归
                    tarjan(y, x);
                    root[y] = x; // 相当于把 y 的子树节点全部 merge 到 x
                }
            }
            for (int y: qs[x]) {
                // color[y] == 2 意味着 y 所在子树已经遍历完
                // 也就意味着 y 已经 merge 到它和 x 的 lca 上了
                // 此时 find(y) 就是 x 和 y 的 lca
                if (y == x || color[y] == 2) {
                    diff[x]++;
                    diff[y]++;
                    int lca = find(y);
                    diff[lca]--;
                    int f = father[lca];
                    if (f >= 0) {
                        diff[f]--;
                    }
                }
            }
            color[x] = 2; // 递归结束
        };
        tarjan(0, -1);

        function<tuple<int, int, int>(int, int)> dfs = [&](int x, int fa) -> tuple<int, int, int> {
            int not_halve = 0, halve = 0, cnt = diff[x];
            for (int y: g[x]) {
                if (y != fa) {
                    auto [nh, h, c] = dfs(y, x); // 计算 y 不变/减半的最小价值总和
                    not_halve += min(nh, h); // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
                    halve += nh; // x 减半，那么 y 只能不变
                    cnt += c; // 自底向上累加差分值
                }
            }
            not_halve += price[x] * cnt; // x 不变
            halve += price[x] * cnt / 2; // x 减半
            return {not_halve, halve, cnt};
        };
        auto [nh, h, _] = dfs(0, -1);
        return min(nh, h);
    }
};
```

```go [sol-Go]
func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建树
	}

	qs := make([][]int, n)
	for _, t := range trips {
		x, y := t[0], t[1]
		qs[x] = append(qs[x], y) // 路径端点分组
		if x != y {
			qs[y] = append(qs[y], x)
		}
	}

	// 并查集模板
	root := make([]int, n)
	for i := range root {
		root[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if root[x] != x {
			root[x] = find(root[x])
		}
		return root[x]
	}

	diff := make([]int, n)
	father := make([]int, n)
	color := make([]int8, n)
	var tarjan func(int, int)
	tarjan = func(x, fa int) {
	father[x] = fa
		color[x] = 1 // 递归中
		for _, y := range g[x] {
			if color[y] == 0 { // 未递归
				tarjan(y, x)
				root[y] = x // 相当于把 y 的子树节点全部 merge 到 x
			}
		}
		for _, y := range qs[x] {
			// color[y] == 2 意味着 y 所在子树已经遍历完
			// 也就意味着 y 已经 merge 到它和 x 的 lca 上了
			// 此时 find(y) 就是 x 和 y 的 lca
			if y == x || color[y] == 2 {
				diff[x]++
				diff[y]++
				lca := find(y)
				diff[lca]--
				if f := father[lca]; f >= 0 {
					diff[f]--
				}
			}
		}
		color[x] = 2 // 递归结束
	}
	tarjan(0, -1)

	var dfs func(int, int) (int, int, int)
	dfs = func(x, fa int) (notHalve, halve, cnt int) {
		cnt = diff[x]
		for _, y := range g[x] {
			if y != fa {
				nh, h, c := dfs(y, x)  // 计算 y 不变/减半的最小价值总和
				notHalve += min(nh, h) // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
				halve += nh            // x 减半，那么 y 只能不变
				cnt += c               // 自底向上累加差分值
			}
		}
		notHalve += price[x] * cnt  // x 不变
		halve += price[x] * cnt / 2 // x 减半
		return
	}
	nh, h, _ := dfs(0, -1)
	return min(nh, h)
}
```

```js [sol-JavaScript]
var minimumTotalPrice = function (n, edges, price, trips) {
    const g = Array(n).fill(null).map(() => []);
    for (const [x, y] of edges) {
        g[x].push(y);
        g[y].push(x);
    }

    const qs = Array(n).fill(null).map(() => []);
    for (const [s, e] of trips) {
        qs[s].push(e); // 路径端点分组
        if (s !== e) {
            qs[e].push(s);
        }
    }

    // 并查集模板
    const root = [...Array(n).keys()];
    function find(x) {
        if (x !== root[x]) {
            root[x] = find(root[x]);
        }
        return root[x];
    }

    const diff = Array(n).fill(0);
    const father = Array(n).fill(0);
    const color = Array(n).fill(0);
    function tarjan(x, fa) {
        father[x] = fa;
        color[x] = 1; // 递归中
        for (const y of g[x]) {
            if (color[y] === 0) { // 未递归
                tarjan(y, x);
                root[y] = x; // 相当于把 y 的子树节点全部 merge 到 x
            }
        }
        for (const y of qs[x]) {
            // color[y] == 2 意味着 y 所在子树已经遍历完
            // 也就意味着 y 已经 merge 到它和 x 的 lca 上了
            // 此时 find(y) 就是 x 和 y 的 lca
            if (y === x || color[y] === 2) {
                diff[x] += 1;
                diff[y] += 1;
                const lca = find(y);
                diff[lca] -= 1;
                if (father[lca] >= 0) {
                    diff[father[lca]] -= 1;
                }
            }
        }
        color[x] = 2; // 递归结束
    }
    tarjan(0, -1);
    
    function dfs(x, fa) {
        let not_halve = 0;
        let halve = 0;
        let cnt = diff[x];
        for (const y of g[x]) {
            if (y !== fa) {
                const [nh, h, c] = dfs(y, x); // 计算 y 不变/减半的最小价值总和
                not_halve += Math.min(nh, h); // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
                halve += nh; // x 减半，那么 y 只能不变
                cnt += c; // 自底向上累加差分值
            }
        }
        not_halve += price[x] * cnt; // x 不变
        halve += Math.floor(price[x] * cnt / 2); // x 减半
        return [not_halve, halve, cnt];
    }
    const [nh, h, _] = dfs(0, -1);
    return Math.min(nh, h);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_total_price(n: i32, edges: Vec<Vec<i32>>, price: Vec<i32>, trips: Vec<Vec<i32>>) -> i32 {
        let n = n as usize;
        let mut g = vec![vec![]; n];
        for e in &edges {
            let x = e[0] as usize;
            let y = e[1] as usize;
            g[x].push(y);
            g[y].push(x);
        }

        let mut qs = vec![vec![]; n];
        for t in &trips {
            let s = t[0] as usize;
            let e = t[1] as usize;
            qs[s].push(e); // 路径端点分组
            if s != e {
                qs[e].push(s);
            }
        }

        // 并查集模板
        let mut root: Vec<usize> = (0..n).collect();
        fn find(x: usize, root: &mut Vec<usize>) -> usize {
            if x != root[x] {
                root[x] = find(root[x], root);
            }
            root[x]
        }

        let mut diff = vec![0; n];
        let mut father = vec![0; n];
        let mut color = vec![0; n];
        fn tarjan(x: usize, fa: usize, diff: &mut Vec<i32>, father: &mut Vec<usize>, color: &mut Vec<i32>, root: &mut Vec<usize>, g: &Vec<Vec<usize>>, qs: &Vec<Vec<usize>>) {
            father[x] = fa;
            color[x] = 1; // 递归中
            for &y in &g[x] {
                if color[y] == 0 { // 未递归
                    tarjan(y, x, diff, father, color, root, g, qs);
                    root[y] = x; // 相当于把 y 的子树节点全部 merge 到 x
                }
            }
            for &y in &qs[x] {
                // color[y] == 2 意味着 y 所在子树已经遍历完
                // 也就意味着 y 已经 merge 到它和 x 的 lca 上了
                // 此时 find(y) 就是 x 和 y 的 lca
                if y == x || color[y] == 2 {
                    diff[x] += 1;
                    diff[y] += 1;
                    let lca = find(y, root);
                    diff[lca] -= 1;
                    if father[lca] != g.len() {
                        diff[father[lca]] -= 1;
                    }
                }
            }
            color[x] = 2; // 递归结束
        }
        tarjan(0, n, &mut diff, &mut father, &mut color, &mut root, &g, &qs);

        fn dfs(x: usize, fa: usize, price: &Vec<i32>, diff: &Vec<i32>, g: &Vec<Vec<usize>>) -> (i32, i32, i32) {
            let mut not_halve = 0;
            let mut halve = 0;
            let mut cnt = diff[x];
            for &y in &g[x] {
                if y != fa {
                    let (nh, h, c) = dfs(y, x, price, diff, g); // 计算 y 不变/减半的最小价值总和
                    not_halve += nh.min(h); // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
                    halve += nh; // x 减半，那么 y 只能不变
                    cnt += c; // 自底向上累加差分值
                }
            }
            not_halve += price[x] * cnt; // x 不变
            halve += price[x] * cnt / 2; // x 减半
            (not_halve, halve, cnt)
        }
        let (nh, h, _) = dfs(0, 0, &price, &diff, &g);
        nh.min(h)
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m\alpha)$，其中 $m$ 为 $\textit{trips}$ 的长度，$\alpha$ 可视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(n+m)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

更多精彩题解，请看 [往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
