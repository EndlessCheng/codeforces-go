为方便计算，我们可以先把所有节点都变成蓝色，计算对应的总代价 $\textit{base}$。然后计算把哪些节点变成红色，可以让 $\textit{base}$ 减少的尽量多，从而得到最终的最小总代价。

设子树 $x$ 的大小为 $\textit{size}[x]$。那么 $\textit{base}$ 即所有红色节点的 $\textit{size}$ 之和。

如果一个节点 $x$ 原来是红色，那么将其恢复成红色，可以让 $\textit{base}$ 减少 $\textit{size}[x]$。

如果一个节点 $x$ 是蓝色，那么将其变成红色，会让 $\textit{base}$ 增大 $\textit{size}[x]$。

于是定义节点 $x$ 的点权

$$
a[x] =
\begin{cases}
\textit{size}[x], & \textit{colors}[x] = \texttt{R}     \\
-\textit{size}[x], & \textit{colors}[x] \ne \texttt{R}     \\
\end{cases}
$$

问题变成：

- 从树中选择一个连通块，其点权和最大是多少？

枚举连通块最上面（深度最小）的节点 $x$（$x$ 一定选），我们需要知道：

- 对于 $x$ 的子树 $y$，在节点 $y$ 一定选的情况下，从子树 $y$ 中能得到的最大点权和 $f[y]$。

这是一个和原问题相似的、规模更小的子问题。

如果 $f[y] < 0$ 则不选 $y$，否则选 $y$。

于是有

$$
f[x] = a[x] + \sum_y \max(f[y], 0)
$$

答案为

$$
\textit{base} - \max(\max(f), 0)
$$

上式中的 $0$ 对应所有节点都变成蓝色的情况。

[本题视频讲解](https://www.bilibili.com/video/BV1LyW76zEvX/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minCost(self, n: int, colors: str, edges: list[list[int]]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        base = mx = 0

        def dfs(x: int, fa: int) -> tuple[int, int]:
            nonlocal base, mx
            size = 1
            f = 0
            # 遍历 x 的子树
            for y in g[x]:
                if y != fa:
                    sz, fy = dfs(y, x)
                    size += sz
                    # 如果 fy < 0，那么子树 y 都变成蓝色是最优的，这已经算入 base 中
                    f += max(fy, 0)

            # 考虑节点 x
            if colors[x] == 'R':
                base += size  # 把 x 变成蓝色的代价
                f += size  # 把 x 恢复成红色可以让 base 减少 size
            else:
                f -= size  # 把 x 变成红色会让 base 增大 size
            mx = max(mx, f)  # 最多可以让 base 减少多少

            return size, f

        dfs(0, -1)
        return base - mx
```

```java [sol-Java]
class Solution {
    private long base = 0;
    private long mx = 0;

    public long minCost(int n, String colors, int[][] edges) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        dfs(0, -1, g, colors.toCharArray());
        return base - mx;
    }

    private long[] dfs(int x, int fa, List<Integer>[] g, char[] colors) {
        int size = 1;
        long f = 0;
        // 遍历 x 的子树
        for (int y : g[x]) {
            if (y != fa) {
                long[] res = dfs(y, x, g, colors);
                size += res[0];
                // 如果 fy < 0，那么子树 y 都变成蓝色是最优的，这已经算入 base 中
                f += Math.max(res[1], 0);
            }
        }

        // 考虑节点 x
        if (colors[x] == 'R') {
            base += size; // 把 x 变成蓝色的代价
            f += size; // 把 x 恢复成红色可以让 base 减少 size
        } else {
            f -= size; // 把 x 变成红色会让 base 增大 size
        }
        mx = Math.max(mx, f); // 最多可以让 base 减少多少

        return new long[]{size, f};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(int n, string colors, vector<vector<int>>& edges) {
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        long long base = 0, mx = 0;

        auto dfs = [&](this auto&& dfs, int x, int fa) -> pair<int, long long> {
            int size = 1;
            long long f = 0;
            // 遍历 x 的子树
            for (int y : g[x]) {
                if (y != fa) {
                    auto [sz, fy] = dfs(y, x);
                    size += sz;
                    // 如果 fy < 0，那么子树 y 都变成蓝色是最优的，这已经算入 base 中
                    f += max(fy, 0LL);
                }
            }

            // 考虑节点 x
            if (colors[x] == 'R') {
                base += size;  // 把 x 变成蓝色的代价
                f += size;  // 把 x 恢复成红色可以让 base 减少 size
            } else {
                f -= size;  // 把 x 变成红色会让 base 增大 size
            }
            mx = max(mx, f);  // 最多可以让 base 减少多少

            return {size, f};
        };

        dfs(0, -1);
        return base - mx;
    }
};
```

```go [sol-Go]
func minCost(n int, colors string, edges [][]int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	base, mx := 0, 0

	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		size := 1
		f := 0
		// 遍历 x 的子树
		for _, y := range g[x] {
			if y != fa {
				sz, fy := dfs(y, x)
				size += sz
				// 如果 fy < 0，那么子树 y 都变成蓝色是最优的，这已经算入 base 中
				f += max(fy, 0)
			}
		}

		// 考虑节点 x
		if colors[x] == 'R' {
			base += size // 把 x 变成蓝色的代价
			f += size // 把 x 恢复成红色可以让 base 减少 size
		} else {
			f -= size // 把 x 变成红色会让 base 增大 size
		}
		mx = max(mx, f) // 最多可以让 base 减少多少

		return size, f
	}

	dfs(0, -1)
	return int64(base - mx)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面动态规划题单的「**十二、树形 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
