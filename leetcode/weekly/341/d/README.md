下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

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
func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) (ans int) {
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

### 更快的做法

本题更快的做法是利用**树上差分**打标记，再通过一次 DFS 算出 $\textit{cnt}$。

用**离线 LCA** 实现的话可以做到完美的 $O(n+m)$ 线性时间复杂度。

直播结束后我会添加这个做法的代码。
