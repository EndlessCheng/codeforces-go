**前置知识**：[【图解】一张图秒懂换根 DP！](https://leetcode.cn/problems/sum-of-distances-in-tree/solution/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/)

本题相当于对每个节点，计算以该节点为根时，树的最大深度。

其中从 $x\rightarrow y$ 的有向边的边权为 $2 - y\bmod 2$，即当 $y$ 是奇数时，边权为 $1$；当 $y$ 是偶数时，边权为 $2$。

⚠**注意**：如果 $x$ 和 $y$ 的奇偶性不同，那么从 $x\rightarrow y$ 的有向边和从 $y\rightarrow x$ 的有向边的边权是不一样的。

考虑换根 DP。

首先，通过一次 DFS，计算以 $0$ 为根节点时，树的最大深度。

在 DFS 的过程中，额外保存：

- 子树 $x$ 的**最大**深度 $\textit{maxD}$。
- 子树 $x$ 的**次大**深度 $\textit{maxD}_2$。
- 子树 $x$ 通过其儿子 $\textit{my}$ 取到的最大深度。

然后，再通过一次 DFS，计算出本题的答案。

对于节点 $x$，其答案是以下两种情况的最大值：

- 子树 $x$ 的最大深度。
- $x$ 往上走到某个节点再往下拐弯的路径长度。

对于第二种情况，可以作为 DFS 的一个参数 $\textit{fromUp}$。

如果 $x$ 的儿子 $y = \textit{my}$，那么往下传入的参数更新为

$$
\max(\textit{fromUp}, \textit{maxD}_2) + 2 - x\bmod 2
$$

如果 $x$ 的儿子 $y\ne \textit{my}$，那么往下传入的参数更新为

$$
\max(\textit{fromUp}, \textit{maxD}) + 2 - x\bmod 2
$$

注：我把[【图解】一张图秒懂换根 DP](https://leetcode.cn/problems/sum-of-distances-in-tree/solution/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/) 这题叫做**第一类换根 DP**，本题需要额外维护次大信息，我称其为**第二类换根 DP**。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1F4421S7XU/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def timeTaken(self, edges: List[List[int]]) -> List[int]:
        g = [[] for _ in range(len(edges) + 1)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        # nodes[x] 保存子树 x 的最大深度 max_d，次大深度 max_d2，以及最大深度要往儿子 my 走
        nodes = [None] * len(g)
        def dfs(x: int, fa: int) -> int:
            max_d = max_d2 = my = 0
            for y in g[x]:
                if y == fa:
                    continue
                depth = dfs(y, x) + 2 - y % 2  # 从 x 出发，往 my 方向的最大深度
                if depth > max_d:
                    max_d2 = max_d
                    max_d = depth
                    my = y
                elif depth > max_d2:
                    max_d2 = depth
            nodes[x] = (max_d, max_d2, my)
            return max_d
        dfs(0, -1)

        ans = [0] * len(g)
        def reroot(x: int, fa: int, from_up: int) -> None:
            max_d, max_d2, my = nodes[x]
            ans[x] = max(from_up, max_d)
            w = 2 - x % 2  # 从 y 到 x 的边权
            for y in g[x]:
                if y != fa:
                    reroot(y, x, max(from_up, max_d2 if y == my else max_d) + w)
        reroot(0, -1, 0)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] timeTaken(int[][] edges) {
        List<Integer>[] g = new ArrayList[edges.length + 1];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        // nodes[x] 保存子树 x 的最大深度 maxD，次大深度 maxD2，以及最大深度要往儿子 my 走
        int[][] nodes = new int[g.length][3];
        dfs(0, -1, g, nodes);

        int[] ans = new int[g.length];
        reroot(0, -1, 0, g, nodes, ans);
        return ans;
    }

    private int dfs(int x, int fa, List<Integer>[] g, int[][] nodes) {
        int maxD = 0;
        int maxD2 = 0;
        int my = 0;
        for (int y : g[x]) {
            if (y == fa) {
                continue;
            }
            int depth = dfs(y, x, g, nodes) + 2 - y % 2; // 从 x 出发，往 my 方向的最大深度
            if (depth > maxD) {
                maxD2 = maxD;
                maxD = depth;
                my = y;
            } else if (depth > maxD2) {
                maxD2 = depth;
            }
        }
        nodes[x][0] = maxD;
        nodes[x][1] = maxD2;
        nodes[x][2] = my;
        return maxD;
    }

    private void reroot(int x, int fa, int fromUp, List<Integer>[] g, int[][] nodes, int[] ans) {
        int maxD = nodes[x][0];
        int maxD2 = nodes[x][1];
        int my = nodes[x][2];
        ans[x] = Math.max(fromUp, maxD);
        for (int y : g[x]) {
            if (y != fa) {
                reroot(y, x, Math.max(fromUp, (y == my ? maxD2 : maxD)) + 2 - x % 2, g, nodes, ans);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> timeTaken(vector<vector<int>>& edges) {
        vector<vector<int>> g(edges.size() + 1);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        // nodes[x] 保存子树 x 的最大深度 max_d，次大深度 max_d2，以及最大深度要往儿子 my 走
        vector<tuple<int, int, int>> nodes(g.size());
        auto dfs = [&](auto&& dfs, int x, int fa) -> int {
            int max_d = 0, max_d2 = 0, my = 0;
            for (int y : g[x]) {
                if (y == fa) {
                    continue;
                }
                int depth = dfs(dfs, y, x) + 2 - y % 2; // 从 x 出发，往 my 方向的最大深度
                if (depth > max_d) {
                    max_d2 = max_d;
                    max_d = depth;
                    my = y;
                } else if (depth > max_d2) {
                    max_d2 = depth;
                }
            }
            nodes[x] = {max_d, max_d2, my};
            return max_d;
        };
        dfs(dfs, 0, -1);

        vector<int> ans(g.size());
        auto reroot = [&](auto&& reroot, int x, int fa, int from_up) -> void {
            auto& [max_d, max_d2, my] = nodes[x];
            ans[x] = max(from_up, max_d);
            for (int y : g[x]) {
                if (y != fa) {
                    reroot(reroot, y, x, max(from_up, (y == my ? max_d2 : max_d)) + 2 - x % 2);
                }
            }
        };
        reroot(reroot, 0, -1, 0);
        return ans;
    }
};
```

```go [sol-Go]
func timeTaken(edges [][]int) []int {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// nodes[x] 保存子树 x 的最大深度 maxD，次大深度 maxD2，以及最大深度要往儿子 y 走
	nodes := make([]struct{ maxD, maxD2, y int }, len(g))
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		p := &nodes[x]
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			maxD := dfs(y, x) + 2 - y%2 // 从 x 出发，往 y 方向的最大深度
			if maxD > p.maxD {
				p.maxD2 = p.maxD
				p.maxD = maxD
				p.y = y
			} else if maxD > p.maxD2 {
				p.maxD2 = maxD
			}
		}
		return p.maxD
	}
	dfs(0, -1)

	ans := make([]int, len(g))
	var reroot func(int, int, int)
	reroot = func(x, fa, fromUp int) {
		p := nodes[x]
		ans[x] = max(fromUp, p.maxD)
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			w := 2 - x%2 // 从 y 到 x 的边权
			if y == p.y { // 对于 y 来说，上面要选次大的
				reroot(y, x, max(fromUp, p.maxD2)+w)
			} else { // 对于 y 来说，上面要选最大的
				reroot(y, x, max(fromUp, p.maxD)+w)
			}
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

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
