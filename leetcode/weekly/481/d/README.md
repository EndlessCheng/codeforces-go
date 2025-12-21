把 $\textit{group}[i]$ 叫做节点 $i$ 的点权。

**提示**：如果问题没有让我们分别计算每个相同点权组的代价，而是计算所有相同点权组的代价总和，通常可以用**贡献法**解决。

横看成岭侧成峰，考虑每条边 $x\text{-}y$ 对答案的贡献，问题变成：

- 有多少个相同点权的点对，之间的最短路径经过 $x\text{-}y$？

切断 $x\text{-}y$，树变成两个连通块 $A$ 和 $B$。我们要从 $A$ 中选一个点，$B$ 中选一个点。

设 $A$ 中点权 $i$ 有 $\textit{a}_i$ 个，$B$ 中点权 $i$ 有 $\textit{b}_i$ 个。根据**乘法原理**，有 $\textit{a}_i\cdot \textit{b}_i$ 个点权都为 $i$ 的点对，之间的最短路径经过 $x\text{-}y$。

枚举 $i=1,2,\ldots,U$（其中 $U=\max(\textit{group})\le 20$），累加得到边 $x\text{-}y$ 对答案的贡献

$$
\sum_{i=1}^{U} \textit{a}_i\cdot \textit{b}_i
$$

代码实现时，可以先统计整棵树的每个点权的出现次数列表 $\textit{total}$，然后 DFS 计算子树 $y$ 的点权的出现次数列表 $a$，那么 $b_i = \textit{total}_i - a_i$。

所有边的贡献的总和，即为最终答案。

[本题视频讲解](https://www.bilibili.com/video/BV1HsqmBwEy3/)，顺带介绍了**虚树**的思路。

```py [sol-Python3]
class Solution:
    def interactionCosts(self, n: int, edges: List[List[int]], group: List[int]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        total_cnt = Counter(group)
        ans = 0

        def dfs(x: int, fa: int) -> Dict[int, int]:
            nonlocal ans
            cnt_x = defaultdict(int)
            cnt_x[group[x]] = 1
            for y in g[x]:
                if y == fa:
                    continue
                cnt_y = dfs(y, x)
                for i, c in cnt_y.items():
                    ans += c * (total_cnt[i] - c)
                    cnt_x[i] += c
            return cnt_x

        dfs(0, -1)
        return ans
```

```java [sol-Java]
class Solution {
    private long ans = 0;

    public long interactionCosts(int n, int[][] edges, int[] group) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        int mx = 0;
        for (int x : group) {
            mx = Math.max(mx, x);
        }

        int[] total = new int[mx + 1];
        for (int x : group) {
            total[x]++;
        }

        dfs(0, -1, g, group, total, mx);
        return ans;
    }

    private int[] dfs(int x, int fa, List<Integer>[] g, int[] group, int[] total, int mx) {
        int[] cntX = new int[mx + 1];
        cntX[group[x]] = 1;
        for (int y : g[x]) {
            if (y == fa) {
                continue;
            }
            int[] cntY = dfs(y, x, g, group, total, mx);
            for (int i = 0; i <= mx; i++) {
                ans += (long) cntY[i] * (total[i] - cntY[i]);
                cntX[i] += cntY[i];
            }
        }
        return cntX;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long interactionCosts(int n, vector<vector<int>>& edges, vector<int>& group) {
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        int mx = ranges::max(group);
        vector<int> total(mx + 1);
        for (int x : group) {
            total[x]++;
        }

        long long ans = 0;

        auto dfs = [&](this auto&& dfs, int x, int fa) -> vector<int> {
            vector<int> cnt_x(mx + 1);
            cnt_x[group[x]] = 1;
            for (int y : g[x]) {
                if (y == fa) {
                    continue;
                }
                vector<int> cnt_y = dfs(y, x);
                for (int i = 0; i <= mx; i++) {
                    int c = cnt_y[i];
                    ans += 1LL * c * (total[i] - c);
                    cnt_x[i] += c;
                }
            }
            return cnt_x;
        };

        dfs(0, -1);
        return ans;
    }
};
```

```go [sol-Go]
func interactionCosts(n int, edges [][]int, group []int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	mx := slices.Max(group)
	total := make([]int, mx+1)
	for _, x := range group {
		total[x]++
	}

	var dfs func(int, int) []int
	dfs = func(x, fa int) []int {
		cntX := make([]int, mx+1)
		cntX[group[x]] = 1
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			cntY := dfs(y, x)
			for i, c := range cntY {
				ans += int64(c) * int64(total[i]-c)
				cntX[i] += c
			}
		}
		return cntX
	}
	dfs(0, -1)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{group})\le 20$。
- 空间复杂度：$\mathcal{O}(nU)$。最坏情况下，递归栈需要保存 $n$ 个长为 $\mathcal{O}(U)$ 的数组。

## 附：与 U 无关的做法——虚树

**前置知识**：[最近公共祖先](https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solution/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw/)。

对相同的 $\textit{group}[i]$ 构建 [虚树](https://oi-wiki.org/graph/virtual-tree/)，虚树上的相邻节点之间的边权，等于原树上这两点的最短距离。此时这条边的贡献为：边权乘以左右两个连通块的大小（节点个数）的乘积。

下面用单调栈建树，初次学习的同学可以用排序法，实现简单（但常数略大）。

```go [sol-Go]
func interactionCosts(n int, edges [][]int, group []int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	dfn := make([]int, n)
	ts := 0
	pa := make([][17]int, n)
	dep := make([]int, n)
	var build func(int, int)
	build = func(v, p int) {
		dfn[v] = ts
		ts++
		pa[v][0] = p
		for _, w := range g[v] {
			if w != p {
				dep[w] = dep[v] + 1
				build(w, v)
			}
		}
	}
	build(0, -1)
	mx := bits.Len(uint(n))
	for i := range mx - 1 {
		for v := range pa {
			p := pa[v][i]
			if p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	uptoDep := func(v, d int) int {
		for k := uint32(dep[v] - d); k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros32(k)]
		}
		return v
	}
	getLCA := func(v, w int) int {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		w = uptoDep(w, dep[v])
		if w == v {
			return v
		}
		for i := mx - 1; i >= 0; i-- {
			pv, pw := pa[v][i], pa[w][i]
			if pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}

	nodesMap := map[int][]int{}
	for i, x := range group {
		nodesMap[x] = append(nodesMap[x], i)
	}

	vt := make([][]int, n)   // 虚树
	isNode := make([]int, n) // 用来区分是关键节点还是 LCA
	for i := range isNode {
		isNode[i] = -1
	}
	addVtEdge := func(v, w int) {
		vt[v] = append(vt[v], w) // 往虚树上添加一条有向边
	}
	const root = 0
	st := []int{root} // 用根节点作为栈底哨兵

	for val, nodes := range nodesMap {
		// 对于相同点权的这一组关键节点 nodes，构建虚树
		slices.SortFunc(nodes, func(a, b int) int { return dfn[a] - dfn[b] })
		vt[root] = vt[root][:0] // 重置虚树
		st = st[:1]
		for _, v := range nodes {
			isNode[v] = val
			if v == root {
				continue
			}
			vt[v] = vt[v][:0]
			lca := getLCA(st[len(st)-1], v) // 路径的拐点（LCA）也加到虚树中
			// 回溯，加边
			for len(st) > 1 && dfn[lca] <= dfn[st[len(st)-2]] {
				addVtEdge(st[len(st)-2], st[len(st)-1])
				st = st[:len(st)-1]
			}
			if lca != st[len(st)-1] { // lca 不在栈中（首次遇到）
				vt[lca] = vt[lca][:0]
				addVtEdge(lca, st[len(st)-1])
				st[len(st)-1] = lca // 加到栈中
			}
			st = append(st, v)
		}
		// 最后的回溯，加边
		for i := 1; i < len(st); i++ {
			addVtEdge(st[i-1], st[i])
		}

		var dfs func(int) int
		dfs = func(v int) (size int) {
			// 如果 isNode[v] != t，那么 v 只是关键节点之间路径上的「拐点」
			if isNode[v] == val {
				size = 1
			}
			for _, w := range vt[v] {
				sz := dfs(w)
				wt := dep[w] - dep[v] // 虚树边权
				// 贡献法
				ans += int64(wt) * int64(sz) * int64(len(nodes)-sz)
				size += sz
			}
			return
		}

		rt := root
		if isNode[rt] != val && len(vt[rt]) == 1 {
			// 注意 root 只是一个哨兵，不一定在虚树上，得从真正的根节点开始
			rt = vt[rt][0]
		}
		dfs(rt)
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$。
- 空间复杂度：$\mathcal{O}(n\log n)$。

## 专题训练

见下面贪心与思维题单的「**§5.5 贡献法**」。

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
