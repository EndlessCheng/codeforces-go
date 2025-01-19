没学过换根 DP 的同学，请先看[【图解】一张图秒懂换根 DP！](https://leetcode.cn/problems/sum-of-distances-in-tree/solution/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/)

本题如果只求以 $0$ 为根时的猜对次数 $\textit{cnt}_0$，把 $\textit{guesses}$ 转成哈希表，DFS 一次这棵树就可以算出来。

如果枚举 $0$ 到 $n-1$ 的每个点作为树根，就需要 DFS $n$ 次，需要 $\mathcal{O}(n^2)$ 的时间，怎么优化呢？

注意到，如果节点 $x$ 和节点 $y$ 相邻，那么从「以 $x$ 为根的树」变成「以 $y$ 为根的树」，就只有 $x$ 和 $y$ 的父子关系改变了，其余相邻节点的父子关系没有变化。所以只有 $[x,y]$ 和 $[y,x]$ 这两个猜测的正确性变了，其余猜测的正确性不变。

因此，在计算出 $\textit{cnt}_0$ 后，我们可以再次从 $0$ 出发，DFS 这棵树。从节点 $x$ 递归到节点 $y$ 时：

- 如果有猜测 $[x,y]$，那么猜对次数减一。
- 如果有猜测 $[y,x]$，那么猜对次数加一。

DFS 的同时，统计猜对次数 $\ge k$ 的节点个数，即为答案。

```py [sol-Python3]
class Solution:
    def rootCount(self, edges: List[List[int]], guesses: List[List[int]], k: int) -> int:
        g = [[] for _ in range(len(edges) + 1)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)  # 建图

        s = {(x, y) for x, y in guesses}  # guesses 转成哈希表 s

        ans = cnt0 = 0
        def dfs(x: int, fa: int) -> None:
            nonlocal cnt0
            for y in g[x]:
                if y != fa:
                    cnt0 += (x, y) in s  # 以 0 为根时，猜对了
                    dfs(y, x)
        dfs(0, -1)

        def reroot(x: int, fa: int, cnt: int) -> None:
            nonlocal ans
            ans += cnt >= k  # 此时 cnt 就是以 x 为根时的猜对次数
            for y in g[x]:
                if y != fa:
                    reroot(y, x, cnt - ((x, y) in s) + ((y, x) in s))
        reroot(0, -1, cnt0)
        return ans
```

```java [sol-Java]
class Solution {
    private List<Integer>[] g;
    private Set<Long> s = new HashSet<>();
    private int k, ans, cnt0;

    public int rootCount(int[][] edges, int[][] guesses, int k) {
        this.k = k;
        g = new ArrayList[edges.length + 1];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x); // 建图
        }

        for (int[] e : guesses) { // guesses 转成哈希表
            s.add((long) e[0] << 32 | e[1]); // 两个 4 字节 int 压缩成一个 8 字节 long
        }

        dfs(0, -1);
        reroot(0, -1, cnt0);
        return ans;
    }

    private void dfs(int x, int fa) {
        for (int y : g[x]) {
            if (y != fa) {
                if (s.contains((long) x << 32 | y)) { // 以 0 为根时，猜对了
                    cnt0++;
                }
                dfs(y, x);
            }
        }
    }

    private void reroot(int x, int fa, int cnt) {
        if (cnt >= k) { // 此时 cnt 就是以 x 为根时的猜对次数
            ans++;
        }
        for (int y : g[x]) {
            if (y != fa) {
                int c = cnt;
                if (s.contains((long) x << 32 | y)) c--; // 原来是对的，现在错了
                if (s.contains((long) y << 32 | x)) c++; // 原来是错的，现在对了
                reroot(y, x, c);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int rootCount(vector<vector<int>>& edges, vector<vector<int>>& guesses, int k) {
        using LL = long long;

        vector<vector<int>> g(edges.size() + 1);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x); // 建图
        }

        unordered_set<LL> s;
        for (auto& e : guesses) { // guesses 转成哈希表
            s.insert((LL) e[0] << 32 | e[1]); // 两个 4 字节数压缩成一个 8 字节数
        }

        int ans = 0, cnt0 = 0;
        auto dfs = [&](this auto&& dfs, int x, int fa) -> void {
            for (int y : g[x]) {
                if (y != fa) {
                    cnt0 += s.count((LL) x << 32 | y); // 以 0 为根时，猜对了
                    dfs(y, x);
                }
            }
        };
        dfs(0, -1);

        auto reroot = [&](this auto&& reroot, int x, int fa, int cnt) -> void {
            ans += cnt >= k; // 此时 cnt 就是以 x 为根时的猜对次数
            for (int y : g[x]) {
                if (y != fa) {
                    reroot(y, x, cnt
                        - s.count((LL) x << 32 | y)   // 原来是对的，现在错了
                        + s.count((LL) y << 32 | x)); // 原来是错的，现在对了
                }
            }
        };
        reroot(0, -1, cnt0);
        return ans;
    }
};
```

```go [sol-Go]
func rootCount(edges [][]int, guesses [][]int, k int) (ans int) {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v) // 建图
	}

	type pair struct{ x, y int }
	s := make(map[pair]int, len(guesses))
	for _, p := range guesses { // guesses 转成哈希表
		s[pair{p[0], p[1]}] = 1
	}

	cnt0 := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if s[pair{x, y}] == 1 { // 以 0 为根时，猜对了
					cnt0++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)

	var reroot func(int, int, int)
	reroot = func(x, fa, cnt int) {
		if cnt >= k { // 此时 cnt 就是以 x 为根时的猜对次数
			ans++
		}
		for _, y := range g[x] {
			if y != fa {
				reroot(y, x, cnt-s[pair{x, y}]+s[pair{y, x}])
			}
		}
	}
	reroot(0, -1, cnt0)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 为 $\textit{edges}$ 的长度加一，$m$ 为 $\textit{guesses}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 思考题

如果把「$u$ 是 $v$ 的父节点」改成「$u$ 是 $v$ 的**祖先节点**」，要怎么做呢？（解答见 [视频](https://www.bilibili.com/video/BV1dY4y1C77x/)）

如果改成「$\textit{guesses}[i]$ 猜对会得到 $\textit{score}[i]$ 分，计算的是以每个点为根时的得分之和」，要怎么做呢？（本题相当于 $\textit{score}[i]$ 均为 $1$）

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
