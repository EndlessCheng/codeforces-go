## 分析

对于一棵树，我们把这棵树的所有节点染成黑色或者白色，规则如下：

- 黑色节点的所有邻居都是白色。
- 白色节点的所有邻居都是黑色。

> 这个想法来自国际象棋的棋盘：所有黑色格子的四方向邻居都是白色格子，所有白色格子的四方向邻居都是黑色格子。也可以从图论的角度理解，因为树一定是二分图。

染色后，从任意节点出发，每走一步，节点的颜色都会改变。所以：

- 从某个节点走奇数步之后，一定会走到异色节点上。
- 从某个节点走偶数步之后，一定会走到同色节点上。

所以从**任意**黑色节点出发，所找到的目标节点，一定都是黑色；从**任意**白色节点出发，所找到的目标节点，一定都是白色。

不妨从节点 $0$ 开始 DFS。（你想从其他节点开始 DFS 也可以。）

## 第二棵树

对于第二棵树，我们把其中的节点分成两个集合：

- 集合 $A$：到节点 $0$ 的距离是偶数的点。其大小记作 $\textit{cnt}_2[0]$。
- 集合 $B$：到节点 $0$ 的距离是奇数的点。其大小记作 $\textit{cnt}_2[1]$。

分类讨论：

- 如果 $\textit{cnt}_2[0] > \textit{cnt}_2[1]$ ，那么第一棵树的节点 $i$ 应当连到集合 $B$ 中的任意节点，这样节点 $i$ 在第二棵树中的目标节点的个数为 $\textit{cnt}_2[0]$。
- 否则，第一棵树的节点 $i$ 应当连到集合 $A$ 中的任意节点，这样节点 $i$ 在第二棵树中的目标节点的个数为 $\textit{cnt}_2[1]$。

所以节点 $i$ 在第二棵树中，最多有

$$
\textit{max}_2 =  \max(\textit{cnt}_2[0],\textit{cnt}_2[1])
$$

个目标节点。

> 注意本题保证 $n\ge 2$ 且 $m\ge 2$。如果 $n=1$ 且 $m=1$，则不能用上式计算，需要特判这种情况。

## 第一棵树

对于第一棵树，我们把其中的节点分成两个集合：

- 集合 $A$：到节点 $0$ 的距离是偶数的点。
- 集合 $B$：到节点 $0$ 的距离是奇数的点。

分类讨论：

- 如果节点 $i$ 在集合 $A$ 中，那么它的目标节点也必然在集合 $A$ 中。
- 如果节点 $i$ 在集合 $B$ 中，那么它的目标节点也必然在集合 $B$ 中。

所以 $\textit{answer}[i]$ 等于节点 $i$ 所属集合的大小，加上 $\textit{max}_2$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1tAzoY1EUN/?t=32m17s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def count(self, edges: List[List[int]]) -> Tuple[List[List[int]], List[int]]:
        g = [[] for _ in range(len(edges) + 1)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        cnt = [0, 0]
        def dfs(x: int, fa: int, d: int) -> None:
            cnt[d] += 1
            for y in g[x]:
                if y != fa:
                    dfs(y, x, d ^ 1)
        dfs(0, -1, 0)
        return g, cnt

    def maxTargetNodes(self, edges1: List[List[int]], edges2: List[List[int]]) -> List[int]:
        _, cnt2 = self.count(edges2)
        max2 = max(cnt2)

        g, cnt1 = self.count(edges1)
        ans = [max2] * len(g)
        def dfs(x: int, fa: int, d: int) -> None:
            ans[x] += cnt1[d]
            for y in g[x]:
                if y != fa:
                    dfs(y, x, d ^ 1)
        dfs(0, -1, 0)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] maxTargetNodes(int[][] edges1, int[][] edges2) {
        List<Integer>[] g2 = buildTree(edges2);
        int[] cnt2 = new int[2];
        dfs(0, -1, 0, g2, cnt2);
        int max2 = Math.max(cnt2[0], cnt2[1]);

        List<Integer>[] g1 = buildTree(edges1);
        int[] cnt1 = new int[2];
        dfs(0, -1, 0, g1, cnt1);

        int[] ans = new int[g1.length];
        Arrays.fill(ans, max2);
        dfs1(0, -1, 0, g1, cnt1, ans);
        return ans;
    }

    private List<Integer>[] buildTree(int[][] edges) {
        List<Integer>[] g = new ArrayList[edges.length + 1];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        return g;
    }

    private void dfs(int x, int fa, int d, List<Integer>[] g, int[] cnt) {
        cnt[d]++;
        for (int y : g[x]) {
            if (y != fa) {
                dfs(y, x, d ^ 1, g, cnt);
            }
        }
    }

    private void dfs1(int x, int fa, int d, List<Integer>[] g, int[] cnt1, int[] ans) {
        ans[x] += cnt1[d];
        for (int y : g[x]) {
            if (y != fa) {
                dfs1(y, x, d ^ 1, g, cnt1, ans);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> maxTargetNodes(vector<vector<int>>& edges1, vector<vector<int>>& edges2) {
        auto count = [](vector<vector<int>>& edges) {
            vector<vector<int>> g(edges.size() + 1);
            for (auto& e : edges) {
                int x = e[0], y = e[1];
                g[x].push_back(y);
                g[y].push_back(x);
            }

            array<int, 2> cnt{};
            auto dfs = [&](this auto&& dfs, int x, int fa, int d) -> void {
                cnt[d]++;
                for (int y : g[x]) {
                    if (y != fa) {
                        dfs(y, x, d ^ 1);
                    }
                }
            };
            dfs(0, -1, 0);
            return pair(g, cnt);
        };

        auto [_, cnt2] = count(edges2);
        int max2 = max(cnt2[0], cnt2[1]);

        auto [g, cnt1] = count(edges1);
        vector<int> ans(g.size(), max2);
        auto dfs = [&](this auto&& dfs, int x, int fa, int d) -> void {
            ans[x] += cnt1[d];
            for (int y : g[x]) {
                if (y != fa) {
                    dfs(y, x, d ^ 1);
                }
            }
        };
        dfs(0, -1, 0);
        return ans;
    }
};
```

```go [sol-Go]
func count(edges [][]int) (g [][]int, cnt [2]int) {
	g = make([][]int, len(edges)+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int, int)
	dfs = func(x, fa, d int) {
		cnt[d]++
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x, d^1)
			}
		}
	}
	dfs(0, -1, 0)
	return
}

func maxTargetNodes(edges1, edges2 [][]int) []int {
	_, cnt2 := count(edges2)
	max2 := max(cnt2[0], cnt2[1])

	g, cnt1 := count(edges1)
	ans := make([]int, len(g))
	var dfs func(int, int, int)
	dfs = func(x, fa, d int) {
		ans[x] = cnt1[d] + max2
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x, d^1)
			}
		}
	}
	dfs(0, -1, 0)
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 是 $\textit{edges}_1$ 的长度，$m$ 是 $\textit{edges}_2$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 思考题

额外输入一个整数 $k$，把「距离是偶数」改成「距离是 $k$ 的倍数」，要怎么做？

请注意 $n$ 和 $m$ 小于 $k$ 的情况。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. 【本题相关】[链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
