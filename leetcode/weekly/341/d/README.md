## 本题视频讲解

见[【周赛 341】](https://www.bilibili.com/video/BV1ng4y1T7QA/)。

## 方法一：暴力 DFS 每条路径

### 提示 1

如果知道每个点总共被经过多少次，就可以仿照 [337. 打家劫舍 III](https://leetcode.cn/problems/house-robber-iii/) 计算答案了（下面会细说）。

注意到数据范围比较小，可以对每个 $\textit{trips}[i]$ 都跑一遍 DFS，把从 $\textit{start}$ 到 $\textit{end}$ 的路径上的点 $x$ 的经过次数 $\textit{cnt}[x]$ 都加一。这一技巧在之前的双周赛中出现过，见 [2467. 树上最大得分和路径](https://leetcode.cn/problems/most-profitable-path-in-a-tree/)。

### 提示 2

既然知道了每个点会被经过多少次，把 $\textit{price}[i]$ 更新成 $\textit{price}[i]\cdot \textit{cnt}[i]$，问题就变成计算减半后的 $\textit{price}[i]$ 之和的最小值。

随便选一个节点出发 DFS，比如节点 $0$。对于节点 $x$ 及其儿子 $y$，分类讨论：

- 如果 $\textit{price}[x]$ 不变，那么 $\textit{price}[y]$ 可以减半，也可以不变，取这两种情况的最小值；
- 如果 $\textit{price}[x]$ 减半，那么 $\textit{price}[y]$ 只能不变。

因此子树 $x$ 需要返回两个值：

- $\textit{price}[x]$ 不变时的子树 $x$ 的最小价值总和；
- $\textit{price}[x]$ 减半时的子树 $x$ 的最小价值总和。

答案就是根节点不变/减半的最小值。

```py [sol1-Python3]
class Solution:
    def minimumTotalPrice(self, n: int, edges: List[List[int]], price: List[int], trips: List[List[int]]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)  # 建树

        cnt = [0] * n
        for start, end in trips:
            def dfs(x: int, fa: int) -> bool:
                if x == end:  # 到达终点（注意树只有唯一的一条简单路径）
                    cnt[x] += 1  # 统计从 start 到 end 的路径上的点经过了多少次
                    return True  # 找到终点
                for y in g[x]:
                    if y != fa and dfs(y, x):
                        cnt[x] += 1  # 统计从 start 到 end 的路径上的点经过了多少次
                        return True  # 找到终点
                return False  # 未找到终点
            dfs(start, -1)

        # 类似 337. 打家劫舍 III https://leetcode.cn/problems/house-robber-iii/
        def dfs(x: int, fa: int) -> (int, int):
            not_halve = price[x] * cnt[x]  # x 不变
            halve = not_halve // 2  # x 减半
            for y in g[x]:
                if y != fa:
                    nh, h = dfs(y, x)  # 计算 y 不变/减半的最小价值总和
                    not_halve += min(nh, h)  # x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
                    halve += nh  # x 减半，那么 y 只能不变
            return not_halve, halve
        return min(dfs(0, -1))
```

```java [sol1-Java]
class Solution {
    private List<Integer>[] g;
    private int[] price, cnt;
    private int end;

    public int minimumTotalPrice(int n, int[][] edges, int[] price, int[][] trips) {
        g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x); // 建树
        }
        this.price = price;

        cnt = new int[n];
        for (var t : trips) {
            end = t[1];
            path(t[0], -1);
        }

        var p = dfs(0, -1);
        return Math.min(p[0], p[1]);
    }

    private boolean path(int x, int fa) {
        if (x == end) { // 到达终点（注意树只有唯一的一条简单路径）
            ++cnt[x]; // 统计从 start 到 end 的路径上的点经过了多少次
            return true; // 找到终点
        }
        for (var y : g[x])
            if (y != fa && path(y, x)) {
                ++cnt[x]; // 统计从 start 到 end 的路径上的点经过了多少次
                return true; // 找到终点
            }
        return false; // 未找到终点
    }

    // 类似 337. 打家劫舍 III https://leetcode.cn/problems/house-robber-iii/
    private int[] dfs(int x, int fa) {
        int notHalve = price[x] * cnt[x]; // x 不变
        int halve = notHalve / 2; // x 减半
        for (var y : g[x])
            if (y != fa) {
                var p = dfs(y, x); // 计算 y 不变/减半的最小价值总和
                notHalve += Math.min(p[0], p[1]); // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
                halve += p[0]; // x 减半，那么 y 只能不变
            }
        return new int[]{notHalve, halve};
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int minimumTotalPrice(int n, vector<vector<int>> &edges, vector<int> &price, vector<vector<int>> &trips) {
        vector<vector<int>> g(n);
        for (auto &e: edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x); // 建树
        }

        int cnt[n]; memset(cnt, 0, sizeof(cnt));
        for (auto &t: trips) {
            int end = t[1];
            function<bool(int, int)> dfs = [&](int x, int fa) -> bool {
                if (x == end) { // 到达终点（注意树只有唯一的一条简单路径）
                    ++cnt[x]; // 统计从 start 到 end 的路径上的点经过了多少次
                    return true; // 找到终点
                }
                for (int y: g[x])
                    if (y != fa && dfs(y, x)) {
                        ++cnt[x]; // 统计从 start 到 end 的路径上的点经过了多少次
                        return true; // 找到终点
                    }
                return false; // 未找到终点
            };
            dfs(t[0], -1);
        }

        // 类似 337. 打家劫舍 III https://leetcode.cn/problems/house-robber-iii/
        function<pair<int, int>(int, int)> dfs = [&](int x, int fa) -> pair<int, int> {
            int not_halve = price[x] * cnt[x]; // x 不变
            int halve = not_halve / 2; // x 减半
            for (int y: g[x])
                if (y != fa) {
                    auto [nh, h] = dfs(y, x); // 计算 y 不变/减半的最小价值总和
                    not_halve += min(nh, h); // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
                    halve += nh; // x 减半，那么 y 只能不变
                }
            return {not_halve, halve};
        };
        auto [nh, h] = dfs(0, -1);
        return min(nh, h);
    }
};
```

```go [sol1-Go]
func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建树
	}

	cnt := make([]int, n)
	for _, t := range trips {
		end := t[1]
		var dfs func(int, int) bool
		dfs = func(x, fa int) bool {
			if x == end { // 到达终点（注意树只有唯一的一条简单路径）
				cnt[x]++ // 统计从 start 到 end 的路径上的点经过了多少次
				return true // 找到终点
			}
			for _, y := range g[x] {
				if y != fa && dfs(y, x) {
					cnt[x]++ // 统计从 start 到 end 的路径上的点经过了多少次
					return true
				}
			}
			return false // 未找到终点
		}
		dfs(t[0], -1)
	}

	// 类似 337. 打家劫舍 III https://leetcode.cn/problems/house-robber-iii/
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
	nh, h := dfs(0, -1)
	return min(nh, h)
}

func min(a, b int) int { if a > b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(nm)$，其中 $m$ 为 $\textit{trips}$ 的长度。
- 空间复杂度：$O(n)$。

## 方法二：Tarjan 离线 LCA + 树上差分

核心思路：利用**树上差分**打标记，再通过一次 DFS 算出 $\textit{cnt}$ 值。

从 $x=\textit{start}$ 到 $y=\textit{end}$ 的路径可以视作从 $x$ 向上到某个点「拐弯」，再向下到达 $y$。（拐弯的点也可能就是 $x$ 或 $y$）

这个拐弯的点就是 $x$ 和 $y$ 的 $\textit{lca}$（最近公共祖先）。

把路径视作 $x-\textit{lca}'-\textit{lca}-y$，其中 $\textit{lca}'$ 是 $\textit{lca}$ 的儿子。由于更新的是点，拆分成 $x-\textit{lca}'$ 和 $y-\textit{lca}$。那么自底向上更新差分 $\textit{diff}$ 值：

- 对于 $x-\textit{lca}'$，更新 $\textit{diff}[x]$ 加一，$\textit{diff}[\textit{lca}]$ 减一；
- 对于 $y-\textit{lca}$，更新 $\textit{diff}[y]$ 加一，$\textit{diff}[\textit{father}[\textit{lca}]]$ 减一，这里 $\textit{father}[i]$ 表示 $i$ 的父节点。

最近公共祖先，用 **Tarjan 离线算法**计算，解释见代码注释。

然后 DFS，在递归的「归」的过程中累加 $\textit{diff}$，计算出 $\textit{cnt}$ 值。

```py [sol2-Python3]
class Solution:
    def minimumTotalPrice(self, n: int, edges: List[List[int]], price: List[int], trips: List[List[int]]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)  # 建树

        qs = [[] for _ in range(n)]
        for s, e in trips:
            qs[s].append(e)  # 路径端点分组
            if s != e:
                qs[e].append(s)

        # 并查集模板
        pa = list(range(n))
        def find(x: int) -> int:
            if x != pa[x]:
                pa[x] = find(pa[x])
            return pa[x]

        diff = [0] * n
        father = [0] * n
        color = [0] * n
        def tarjan(x: int, fa: int) -> None:
            father[x] = fa
            color[x] = 1  # 递归中
            for y in g[x]:
                if color[y] == 0:  # 未递归
                    tarjan(y, x)
                    pa[y] = x  # 相当于把 y 的子树节点全部 merge 到 x
            for y in qs[x]:
                # color[y] == 2 意味着 y 所在子树已经遍历完
                # 也就意味着 y 已经 merge 到它和 x 的 lca 上了
                if y == x or color[y] == 2:  # 从 y 向上到达 lca 然后拐弯向下到达 x
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

```java [sol2-Java]
class Solution {
    private List<Integer>[] g, qs;
    private int[] diff, father, color, price;

    public int minimumTotalPrice(int n, int[][] edges, int[] price, int[][] trips) {
        g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x); // 建树
        }

        qs = new ArrayList[n];
        Arrays.setAll(qs, e -> new ArrayList<>());
        for (var t : trips) {
            int x = t[0], y = t[1];
            qs[x].add(y); // 路径端点分组
            if (x != y) qs[y].add(x);
        }

        pa = new int[n];
        for (int i = 1; i < n; ++i)
            pa[i] = i;

        diff = new int[n];
        father = new int[n];
        color = new int[n];
        tarjan(0, -1);

        this.price = price;
        var p = dfs(0, -1);
        return Math.min(p[0], p[1]);
    }

    // 并查集模板
    private int[] pa;

    private int find(int x) {
        if (pa[x] != x)
            pa[x] = find(pa[x]);
        return pa[x];
    }

    private void tarjan(int x, int fa) {
        father[x] = fa;
        color[x] = 1; // 递归中
        for (int y : g[x])
            if (color[y] == 0) { // 未递归
                tarjan(y, x);
                pa[y] = x; // 相当于把 y 的子树节点全部 merge 到 x
            }
        for (int y : qs[x])
            // color[y] == 2 意味着 y 所在子树已经遍历完
            // 也就意味着 y 已经 merge 到它和 x 的 lca 上了
            if (y == x || color[y] == 2) { // 从 y 向上到达 lca 然后拐弯向下到达 x
                ++diff[x];
                ++diff[y];
                int lca = find(y);
                --diff[lca];
                int f = father[lca];
                if (f >= 0) {
                    --diff[f];
                }
            }
        color[x] = 2; // 递归结束
    }

    private int[] dfs(int x, int fa) {
        int notHalve = 0, halve = 0, cnt = diff[x];
        for (int y : g[x])
            if (y != fa) {
                var p = dfs(y, x); // 计算 y 不变/减半的最小价值总和
                notHalve += Math.min(p[0], p[1]); // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
                halve += p[0]; // x 减半，那么 y 只能不变
                cnt += p[2]; // 自底向上累加差分值
            }
        notHalve += price[x] * cnt; // x 不变
        halve += price[x] * cnt / 2; // x 减半
        return new int[]{notHalve, halve, cnt};
    }
}
```

```cpp [sol2-C++]
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
            if (x != y) qs[y].push_back(x);
        }

        // 并查集模板
        int pa[n];
        iota(pa, pa + n, 0);
        function<int(int)> find = [&](int x) -> int { return pa[x] == x ? x : pa[x] = find(pa[x]); };

        int diff[n], father[n], color[n];
        memset(diff, 0, sizeof(diff));
        memset(color, 0, sizeof(color));
        function<void(int, int)> tarjan = [&](int x, int fa) {
            father[x] = fa;
            color[x] = 1; // 递归中
            for (int y: g[x])
                if (color[y] == 0) { // 未递归
                    tarjan(y, x);
                    pa[y] = x; // 相当于把 y 的子树节点全部 merge 到 x
                }
            for (int y: qs[x])
                // color[y] == 2 意味着 y 所在子树已经遍历完
                // 也就意味着 y 已经 merge 到它和 x 的 lca 上了
                if (y == x || color[y] == 2) { // 从 y 向上到达 lca 然后拐弯向下到达 x
                    ++diff[x];
                    ++diff[y];
                    int lca = find(y);
                    --diff[lca];
                    int f = father[lca];
                    if (f >= 0) {
                        --diff[f];
                    }
                }
            color[x] = 2; // 递归结束
        };
        tarjan(0, -1);

        function<tuple<int, int, int>(int, int)> dfs = [&](int x, int fa) -> tuple<int, int, int> {
            int not_halve = 0, halve = 0, cnt = diff[x];
            for (int y: g[x])
                if (y != fa) {
                    auto [nh, h, c] = dfs(y, x); // 计算 y 不变/减半的最小价值总和
                    not_halve += min(nh, h); // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
                    halve += nh; // x 减半，那么 y 只能不变
                    cnt += c; // 自底向上累加差分值
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

```go [sol2-Go]
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
	pa := make([]int, n)
	for i := range pa {
		pa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if pa[x] != x {
			pa[x] = find(pa[x])
		}
		return pa[x]
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
				pa[y] = x // 相当于把 y 的子树节点全部 merge 到 x
			}
		}
		for _, y := range qs[x] {
			// color[y] == 2 意味着 y 所在子树已经遍历完
			// 也就意味着 y 已经 merge 到它和 x 的 lca 上了
			if y == x || color[y] == 2 { // 从 y 向上到达 lca 然后拐弯向下到达 x
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

func min(a, b int) int { if a > b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n+m\alpha)$，其中 $m$ 为 $\textit{trips}$ 的长度，$\alpha$ 为并查集的常数，可视作 $O(1)$。
- 空间复杂度：$O(n+m)$。
