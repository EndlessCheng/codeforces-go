**前置题目**：[104. 二叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)。

设从 $1$ 到 $x$ 的路径中有 $k$ 条边。由于边权只能是 $1$ 或 $2$，必须有奇数个 $1$，才能使边权之和是奇数。

**定理**：从 $k$ 个不同元素中，选奇数个数，有 $2^{k-1}$ 种选法。

请看 [文字证明](https://zhuanlan.zhihu.com/p/1909852852114948837) 或者 [视频讲解](https://www.bilibili.com/video/BV1cqjgzdEPP/?t=9m12s)，欢迎点赞关注~

怎么计算 $2^{k-1}\bmod M$？可以用循环，也可以用 [快速幂](https://leetcode.cn/problems/powx-n/solution/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/)。

```py [sol-Python3]
class Solution:
    def assignEdgeWeights(self, edges: List[List[int]]) -> int:
        MOD = 1_000_000_007
        n = len(edges) + 1
        g = [[] for _ in range(n + 1)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        def dfs(x: int, fa: int) -> int:
            d = 0
            for y in g[x]:
                if y != fa:  # 不递归到父节点
                    d = max(d, dfs(y, x) + 1)
            return d

        k = dfs(1, 0)
        return pow(2, k - 1, MOD)
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int assignEdgeWeights(int[][] edges) {
        int n = edges.length + 1;
        List<Integer>[] g = new ArrayList[n + 1];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        int k = dfs(1, 0, g);
        return (int) pow(2, k - 1);
    }

    private int dfs(int x, int fa, List<Integer>[] g) {
        int d = 0;
        for (int y : g[x]) {
            if (y != fa) { // 不递归到父节点
                d = Math.max(d, dfs(y, x, g) + 1);
            }
        }
        return d;
    }

    private long pow(long x, int n) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1'000'000'007;

    long long qpow(long long x, int n) {
        long long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }

public:
    int assignEdgeWeights(vector<vector<int>>& edges) {
        int n = edges.size() + 1;
        vector<vector<int>> g(n + 1);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        auto dfs = [&](this auto&& dfs, int x, int fa) -> int {
            int d = 0;
            for (int y : g[x]) {
                if (y != fa) { // 不递归到父节点
                    d = max(d, dfs(y, x) + 1);
                }
            }
            return d;
        };

        int k = dfs(1, 0);
        return qpow(2, k - 1);
    }
};
```

```go [sol-Go]
func assignEdgeWeights(edges [][]int) int {
	n := len(edges) + 1
	g := make([][]int, n+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) int
	dfs = func(x, fa int) (d int) {
		for _, y := range g[x] {
			if y != fa { // 不递归到父节点
				d = max(d, dfs(y, x)+1)
			}
		}
		return
	}

	k := dfs(1, 0)
	return pow(2, k-1)
}

func pow(x, n int) int {
	const mod = 1_000_000_007
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
