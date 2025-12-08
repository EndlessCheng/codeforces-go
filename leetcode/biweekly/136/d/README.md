**前置知识**：[【图解】一张图秒懂换根 DP！](https://leetcode.cn/problems/sum-of-distances-in-tree/solution/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/)

本题相当于对每个节点，计算以该节点为根时，树的高度（最大深度）。

其中从 $x\to y$ 的有向边的边权为 $2 - y\bmod 2$，即当 $y$ 是奇数时，边权为 $1$；当 $y$ 是偶数时，边权为 $2$。

⚠**注意**：如果 $x$ 和 $y$ 的奇偶性不同，那么从 $x\to y$ 的有向边和从 $y\to x$ 的有向边的边权是不一样的。

考虑换根 DP。

首先，通过一次 DFS，计算以 $0$ 为根节点时，树的最大深度。

在 DFS 的过程中，额外保存：

- 子树 $x$ 的**最大**深度 $\textit{maxD}$。
- 子树 $x$ 的**次大**深度 $\textit{maxD}_2$。
- 子树 $x$ 通过其儿子 $\textit{my}$ 取到的最大深度。

然后，再通过一次 DFS，计算出本题的答案。

对于节点 $x$，其答案是以下两种情况的最大值：

- 子树 $x$ 的最大深度。
- $x$ 往上走到某个节点（可以再往下拐弯）的路径长度。

对于第二种情况，可以作为 DFS 的一个参数 $\textit{fromUp}$。

如果 $x$ 的儿子 $y = \textit{my}$，那么往下传入的参数更新为

$$
\max(\textit{fromUp}, \textit{maxD}_2) + 2 - x\bmod 2
$$

如果 $x$ 的儿子 $y\ne \textit{my}$，那么往下传入的参数更新为

$$
\max(\textit{fromUp}, \textit{maxD}) + 2 - x\bmod 2
$$

**注**：我把[【图解】一张图秒懂换根 DP](https://leetcode.cn/problems/sum-of-distances-in-tree/solution/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/) 这题叫做**第一类换根 DP**，本题需要额外维护次大信息，我称其为**第二类换根 DP**。

[本题视频讲解](https://www.bilibili.com/video/BV1F4421S7XU/?t=17m28s) 第四题，欢迎点赞关注~

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def timeTaken(self, edges: List[List[int]]) -> List[int]:
        n = len(edges) + 1
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        # sub_res[x] 保存子树 x 的最大深度，次大深度，以及最大深度要往哪个儿子走
        sub_res = [None] * n
        # 计算 sub_res[x]
        def dfs(x: int, fa: int) -> None:
            max_d = max_d2 = my = 0
            for y in g[x]:
                if y == fa:
                    continue
                dfs(y, x)
                w = 2 - y % 2  # 从 x 到 y 的边权
                max_y = sub_res[y][0] + w  # 从 x 出发，往 y 方向的最大深度
                if max_y > max_d:
                    max_d2 = max_d
                    max_d = max_y
                    my = y
                elif max_y > max_d2:
                    max_d2 = max_y
            sub_res[x] = (max_d, max_d2, my)
        dfs(0, -1)

        # ans[x] 表示当 x 是树根时，整棵树的最大深度
        ans = [0] * n
        # 计算 ans[x]
        def reroot(x: int, fa: int, from_up: int) -> None:
            max_d, max_d2, my = sub_res[x]
            ans[x] = max(max_d, from_up)
            for y in g[x]:
                if y == fa:
                    continue
                # 站在 x 的角度，不往 y 走，能走多远？
                # 要么往上走（from_up），要么往除了 y 的其余子树走（mx），二者取最大值
                mx = max_d if y != my else max_d2
                w = 2 - x % 2  # 从 y 到 x 的边权
                reroot(y, x, max(from_up, mx) + w)  # 对于 y 来说，加上从 y 到 x 的边权
        reroot(0, -1, 0)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] timeTaken(int[][] edges) {
        int n = edges.length + 1;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        // subRes[x] 保存子树 x 的最大深度，次大深度，以及最大深度要往哪个儿子走
        int[][] subRes = new int[n][];
        dfs(0, -1, g, subRes);

        // ans[x] 表示当 x 是树根时，整棵树的最大深度
        int[] ans = new int[n];
        reroot(0, -1, 0, g, subRes, ans);
        return ans;
    }

    // 计算 subRes[x]
    private void dfs(int x, int fa, List<Integer>[] g, int[][] subRes) {
        int maxD = 0;
        int maxD2 = 0;
        int my = 0;
        for (int y : g[x]) {
            if (y == fa) {
                continue;
            }
            dfs(y, x, g, subRes);
            int w = 2 - y % 2; // 从 x 到 y 的边权
            int maxY = subRes[y][0] + w; // 从 x 出发，往 y 方向的最大深度
            if (maxY > maxD) {
                maxD2 = maxD;
                maxD = maxY;
                my = y;
            } else if (maxY > maxD2) {
                maxD2 = maxY;
            }
        }
        subRes[x] = new int[]{maxD, maxD2, my};
    }

    // 计算 ans[x]
    private void reroot(int x, int fa, int fromUp, List<Integer>[] g, int[][] subRes, int[] ans) {
        int maxD = subRes[x][0];
        int maxD2 = subRes[x][1];
        int my = subRes[x][2];
        ans[x] = Math.max(maxD, fromUp);
        for (int y : g[x]) {
            if (y == fa) {
                continue;
            }
            // 站在 x 的角度，不往 y 走，能走多远？
            // 要么往上走（fromUp），要么往除了 y 的其余子树走（mx），二者取最大值
            int mx = y != my ? maxD : maxD2;
            int w = 2 - x % 2; // 从 y 到 x 的边权
            reroot(y, x, Math.max(fromUp, mx) + w, g, subRes, ans);
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> timeTaken(vector<vector<int>>& edges) {
        int n = edges.size() + 1;
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        // sub_res[x] 保存子树 x 的最大深度，次大深度，以及最大深度要往哪个儿子走
        vector<tuple<int, int, int>> sub_res(n);
        // 计算 sub_res[x]
        auto dfs = [&](this auto&& dfs, int x, int fa) -> void {
            int max_d = 0, max_d2 = 0, my = 0;
            for (int y : g[x]) {
                if (y == fa) {
                    continue;
                }
                dfs(y, x);
                int w = 2 - y % 2;  // 从 x 到 y 的边权
                int max_y = get<0>(sub_res[y]) + w;  // 从 x 出发，往 y 方向的最大深度
                if (max_y > max_d) {
                    max_d2 = max_d;
                    max_d = max_y;
                    my = y;
                } else if (max_y > max_d2) {
                    max_d2 = max_y;
                }
            }
            sub_res[x] = {max_d, max_d2, my};
        };
        dfs(0, -1);

        // ans[x] 表示当 x 是树根时，整棵树的最大深度
        vector<int> ans(n);
        // 计算 ans[x]
        auto reroot = [&](this auto&& reroot, int x, int fa, int from_up) -> void {
            auto [max_d, max_d2, my] = sub_res[x];
            ans[x] = max(max_d, from_up);
            for (int y : g[x]) {
                if (y == fa) {
                    continue;
                }
                // 站在 x 的角度，不往 y 走，能走多远？
                // 要么往上走（from_up），要么往除了 y 的其余子树走（mx），二者取最大值
                int mx = y != my ? max_d : max_d2;
                int w = 2 - x % 2;  // 从 y 到 x 的边权
                reroot(y, x, max(from_up, mx) + w);  // 对于 y 来说，加上从 y 到 x 的边权
            }
        };
        reroot(0, -1, 0);
        return ans;
    }
};
```

```go [sol-Go]
func timeTaken(edges [][]int) []int {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// subRes[x] 保存子树 x 的最大深度 maxD，次大深度 maxD2，以及最大深度要往儿子 y 走
	subRes := make([]struct{ maxD, maxD2, y int }, n)
	// 计算 subRes[x]
	var dfs func(int, int)
	dfs = func(x, fa int) {
		res := &subRes[x]
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			dfs(y, x)
			w := 2 - y%2 // 从 x 到 y 的边权
			maxD := subRes[y].maxD + w // 从 x 出发，往 y 方向的最大深度
			if maxD > res.maxD {
				res.maxD2 = res.maxD
				res.maxD = maxD
				res.y = y
			} else if maxD > res.maxD2 {
				res.maxD2 = maxD
			}
		}
	}
	dfs(0, -1)

	// ans[x] 表示当 x 是树根时，整棵树的最大深度
	ans := make([]int, n)
	// 计算 ans[x]
	var reroot func(int, int, int)
	reroot = func(x, fa, fromUp int) {
		p := subRes[x]
		ans[x] = max(subRes[x].maxD, fromUp)
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			// 站在 x 的角度，不往 y 走，能走多远？
			// 要么往上走（fromUp），要么往除了 y 的其余子树走（mx），二者取最大值
			mx := p.maxD
			if y == p.y { // 对于 y 来说，上面要选次大的
				mx = p.maxD2
			}
			w := 2 - x%2 // 从 y 到 x 的边权
			reroot(y, x, max(fromUp, mx)+w) // 对于 y 来说，加上从 y 到 x 的边权
		}
	}
	reroot(0, -1, 0)
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [CF1822F. Gardening Friends](https://codeforces.com/problemset/problem/1822/F)

更多相似题目，见下面 DP 题单中的「**换根 DP**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
