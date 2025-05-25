先把上一题 [3558. 给边赋权值的方案数 I](https://leetcode.cn/problems/number-of-ways-to-assign-edge-weights-i/description/) 做了。

本题需要快速计算两点之间的距离（边数），套 LCA 模板即可，[原理讲解](https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solution/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw/)。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class LcaBinaryLifting:
    def __init__(self, edges: List[List[int]]):
        n = len(edges) + 1
        m = n.bit_length()
        g = [[] for _ in range(n)]
        for x, y in edges:
            x -= 1  # 节点编号从 0 开始
            y -= 1
            g[x].append(y)
            g[y].append(x)

        depth = [0] * n
        pa = [[-1] * m for _ in range(n)]

        def dfs(x: int, fa: int) -> None:
            pa[x][0] = fa
            for y in g[x]:
                if y != fa:
                    depth[y] = depth[x] + 1
                    dfs(y, x)

        dfs(0, -1)

        for i in range(m - 1):
            for x in range(n):
                if (p := pa[x][i]) != -1:
                    pa[x][i + 1] = pa[p][i]
        self.depth = depth
        self.pa = pa

    def get_kth_ancestor(self, node: int, k: int) -> int:
        for i in range(k.bit_length()):
            if k >> i & 1:  # k 二进制从低到高第 i 位是 1
                node = self.pa[node][i]
        return node

    # 返回 x 和 y 的最近公共祖先（节点编号从 0 开始）
    def get_lca(self, x: int, y: int) -> int:
        if self.depth[x] > self.depth[y]:
            x, y = y, x
        # 使 y 和 x 在同一深度
        y = self.get_kth_ancestor(y, self.depth[y] - self.depth[x])
        if y == x:
            return x
        for i in range(len(self.pa[x]) - 1, -1, -1):
            px, py = self.pa[x][i], self.pa[y][i]
            if px != py:
                x, y = px, py  # 同时往上跳 2**i 步
        return self.pa[x][0]

    # 返回 x 到 y 的距离（最短路长度）
    def get_dis(self, x: int, y: int) -> int:
        return self.depth[x] + self.depth[y] - self.depth[self.get_lca(x, y)] * 2

# 预处理 2 的幂
MOD = 1_000_000_007
POW2 = [0] * 10 ** 5
POW2[0] = 1
for i in range(1, len(POW2)):
    POW2[i] = POW2[i - 1] * 2 % MOD

class Solution:
    def assignEdgeWeights(self, edges: List[List[int]], queries: List[List[int]]) -> List[int]:
        g = LcaBinaryLifting(edges)
        return [POW2[g.get_dis(x - 1, y - 1) - 1] if x != y else 0 for x, y in queries]
```

```java [sol-Java]
class LcaBinaryLifting {
    private final int[] depth;
    private final int[][] pa;

    public LcaBinaryLifting(int[][] edges) {
        int n = edges.length + 1;
        int m = 32 - Integer.numberOfLeadingZeros(n); // n 的二进制长度
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0] - 1;
            int y = e[1] - 1;
            g[x].add(y);
            g[y].add(x);
        }

        depth = new int[n];
        pa = new int[n][m];
        dfs(g, 0, -1);

        for (int i = 0; i < m - 1; i++) {
            for (int x = 0; x < n; x++) {
                int p = pa[x][i];
                pa[x][i + 1] = p < 0 ? -1 : pa[p][i];
            }
        }
    }

    private void dfs(List<Integer>[] g, int x, int fa) {
        pa[x][0] = fa;
        for (int y : g[x]) {
            if (y != fa) {
                depth[y] = depth[x] + 1;
                dfs(g, y, x);
            }
        }
    }

    public int getKthAncestor(int node, int k) {
        for (; k > 0; k &= k - 1) {
            node = pa[node][Integer.numberOfTrailingZeros(k)];
        }
        return node;
    }

    // 返回 x 和 y 的最近公共祖先（节点编号从 0 开始）
    public int getLCA(int x, int y) {
        if (depth[x] > depth[y]) {
            int tmp = y;
            y = x;
            x = tmp;
        }
        y = getKthAncestor(y, depth[y] - depth[x]); // 使 y 和 x 在同一深度
        if (y == x) {
            return x;
        }
        for (int i = pa[x].length - 1; i >= 0; i--) {
            int px = pa[x][i], py = pa[y][i];
            if (px != py) {
                x = px;
                y = py; // 同时往上跳 2^i 步
            }
        }
        return pa[x][0];
    }

    // 返回 x 到 y 的距离（最短路长度）
    public int getDis(int x, int y) {
        return depth[x] + depth[y] - depth[getLCA(x, y)] * 2;
    }
}

class Solution {
    private static final int MOD = 1_000_000_007;

    public int[] assignEdgeWeights(int[][] edges, int[][] queries) {
        LcaBinaryLifting g = new LcaBinaryLifting(edges);

        // 预处理 2 的幂
        int[] pow2 = new int[edges.length + 1];
        pow2[0] = 1;
        for (int i = 1; i < pow2.length; i++) {
            pow2[i] = pow2[i - 1] * 2 % MOD;
        }

        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int x = queries[i][0] - 1;
            int y = queries[i][1] - 1;
            if (x != y) {
                ans[i] = pow2[g.getDis(x, y) - 1];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class LcaBinaryLifting {
    vector<int> depth;
    vector<vector<int>> pa;

public:
    LcaBinaryLifting(vector<vector<int>>& edges) {
        int n = edges.size() + 1;
        int m = bit_width((unsigned) n); // n 的二进制长度
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0] - 1, y = e[1] - 1;
            g[x].push_back(y);
            g[y].push_back(x);
        }

        depth.resize(n);
        pa.resize(n, vector<int>(m, -1));
        auto dfs = [&](this auto&& dfs, int x, int fa) -> void {
            pa[x][0] = fa;
            for (int y : g[x]) {
                if (y != fa) {
                    depth[y] = depth[x] + 1;
                    dfs(y, x);
                }
            }
        };
        dfs(0, -1);

        for (int i = 0; i < m - 1; i++) {
            for (int x = 0; x < n; x++) {
                if (int p = pa[x][i]; p != -1) {
                    pa[x][i + 1] = pa[p][i];
                }
            }
        }
    }

    int get_kth_ancestor(int node, int k) {
        for (; k; k &= k - 1) {
            node = pa[node][countr_zero((unsigned) k)];
        }
        return node;
    }

    // 返回 x 和 y 的最近公共祖先（节点编号从 0 开始）
    int get_lca(int x, int y) {
        if (depth[x] > depth[y]) {
            swap(x, y);
        }
        y = get_kth_ancestor(y, depth[y] - depth[x]); // 使 y 和 x 在同一深度
        if (y == x) {
            return x;
        }
        for (int i = pa[x].size() - 1; i >= 0; i--) {
            int px = pa[x][i], py = pa[y][i];
            if (px != py) {
                x = px;
                y = py; // 同时往上跳 2^i 步
            }
        }
        return pa[x][0];
    }

    // 返回 x 到 y 的距离（最短路长度）
    int get_dis(int x, int y) {
        return depth[x] + depth[y] - depth[get_lca(x, y)] * 2;
    }
};

const int MOD = 1'000'000'007;
const int MX = 100'000;

int pow2[MX];

auto init = [] {
    // 预处理 2 的幂
    pow2[0] = 1;
    for (int i = 1; i < MX; i++) {
        pow2[i] = pow2[i - 1] * 2 % MOD;
    }
    return 0;
}();

class Solution {
public:
    vector<int> assignEdgeWeights(vector<vector<int>>& edges, vector<vector<int>>& queries) {
        LcaBinaryLifting g(edges);
        vector<int> ans(queries.size());
        for (int i = 0; i < queries.size(); i++) {
            int x = queries[i][0] - 1;
            int y = queries[i][1] - 1;
            if (x != y) {
                ans[i] = pow2[g.get_dis(x, y) - 1];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

var pow2 = [1e5]int{1}

func init() {
	// 预处理 2 的幂
	for i := 1; i < len(pow2); i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}
}

func assignEdgeWeights(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	const mx = 17
	pa := make([][mx]int, n)
	dep := make([]int, n)
	var dfs func(int, int)
	dfs = func(x, p int) {
		pa[x][0] = p
		for _, y := range g[x] {
			if y != p {
				dep[y] = dep[x] + 1
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)
	for i := range mx - 1 {
		for x := range pa {
			if p := pa[x][i]; p != -1 {
				pa[x][i+1] = pa[p][i]
			} else {
				pa[x][i+1] = -1
			}
		}
	}
	uptoDep := func(x, d int) int {
		for k := uint(dep[x] - d); k > 0; k &= k - 1 {
			x = pa[x][bits.TrailingZeros(k)]
		}
		return x
	}
	getLCA := func(x, y int) int {
		if dep[x] > dep[y] {
			x, y = y, x
		}
		y = uptoDep(y, dep[x])
		if y == x {
			return x
		}
		for i := mx - 1; i >= 0; i-- {
			if pv, pw := pa[x][i], pa[y][i]; pv != pw {
				x, y = pv, pw
			}
		}
		return pa[x][0]
	}
	getDis := func(x, y int) int { return dep[x] + dep[y] - dep[getLCA(x, y)]*2 }

	ans := make([]int, len(queries))
	for i, q := range queries {
		if q[0] != q[1] {
			ans[i] = pow2[getDis(q[0]-1, q[1]-1)-1]
		}
	}
	return ans
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}((n+q)\log n)$，其中 $n$ 是 $\textit{edges}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n\log n)$。返回值不计入。

更多相似题目，见下面树题单的「**§3.8 倍增算法、最近公共祖先（LCA）**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. 【本题相关】[链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
