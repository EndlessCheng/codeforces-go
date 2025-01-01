类似 [3128. 直角三角形](https://leetcode.cn/problems/right-triangles/)，有三个元素的题目，**枚举中间**比较容易。

枚举中间的车在第 $i$ 排，那么需要知道第 $0$ 到 $i-1$ 排的车放在哪最合适，以及第 $i+1$ 到 $m-1$ 排的车放在哪最合适。

> 为什么枚举中间更容易？比较一下，如果枚举第三个车，那么需要知道第一个车和第二个车的位置关系，尤其是这两个车不能同一行。但枚举第二个车，就**自动保证**了第一个车和第三个车不会同行，我们只需关注三车不能同列。

用**前后缀分解**解决。

为了保证三个车在不同的列上，我们需要计算：

- 第 $0$ 到 $i-1$ 排的车，放在三个不同的列上的最大、次大、第三大的格子值分别是多少，以及具体放在哪列。记作 $\textit{pre}[i-1]$。
- 第 $i+1$ 到 $m-1$ 排的车，放在三个不同的列上的最大、次大、第三大的格子值分别是多少，以及具体放在哪列。记作 $\textit{suf}[i+1]$。

**枚举中间**的车在第 $i$ 排第 $j$ 列，暴力枚举 $\textit{pre}[i-1]$ 和 $\textit{suf}[i+1]$ 的最大、次大、第三大的组合，只要三个车在不同的列上，就用格子值之和更新答案的最大值。

代码实现时，可以只预处理 $\textit{suf}$，对于 $\textit{pre}$，可以在枚举中间的车的同时计算。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ZH4y1c7GA/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def maximumValueSum(self, board: List[List[int]]) -> int:
        def update(row: List[int]) -> None:
            for j, x in enumerate(row):
                for k in range(3):
                    if x > p[k][0] and all(j != j2 for _, j2 in p[:k]):
                        p[k], (x, j) = (x, j), p[k]

        m = len(board)
        suf = [None] * m
        p = [(-inf, -1)] * 3  # 最大、次大、第三大
        for i in range(m - 1, 1, -1):
            update(board[i])
            suf[i] = p[:]

        ans = -inf
        p = [(-inf, -1)] * 3  # 重置，计算 pre
        for i, row in enumerate(board[:-2]):
            update(row)
            for j2, y in enumerate(board[i + 1]):  # 第二个车
                for x, j1 in p:  # 第一个车
                    for z, j3 in suf[i + 2]:  # 第三个车
                        if j1 != j2 and j1 != j3 and j2 != j3:  # 没有同列的车
                            ans = max(ans, x + y + z)  # 注：手动 if 更快
                            break
        return ans
```

```py [sol-Python3 更快写法]
class Solution:
    def maximumValueSum(self, board: List[List[int]]) -> int:
        def update(row: List[int]) -> None:
            for j, x in enumerate(row):
                if x > p[0][0]:
                    if p[0][1] != j:  # 如果相等，仅更新最大
                        if p[1][1] != j:  # 如果相等，仅更新最大和次大
                            p[2] = p[1]
                        p[1] = p[0]
                    p[0] = (x, j)
                elif x > p[1][0] and j != p[0][1]:
                    if p[1][1] != j:  # 如果相等，仅更新次大
                        p[2] = p[1]
                    p[1] = (x, j)
                elif x > p[2][0] and j != p[0][1] and j != p[1][1]:
                    p[2] = (x, j)

        m = len(board)
        suf = [None] * m
        p = [(-inf, -1)] * 3  # 最大、次大、第三大
        for i in range(m - 1, 1, -1):
            update(board[i])
            suf[i] = p[:]

        ans = -inf
        p = [(-inf, -1)] * 3  # 重置，计算 pre
        for i, row in enumerate(board[:-2]):
            update(row)
            for j2, y in enumerate(board[i + 1]):  # 第二个车
                for x, j1 in p:  # 第一个车
                    for z, j3 in suf[i + 2]:  # 第三个车
                        if j1 != j2 and j1 != j3 and j2 != j3:  # 没有同列的车
                            ans = max(ans, x + y + z)  # 注：手动 if 更快
                            break
        return ans
```

```java [sol-Java]
class Solution {
    public long maximumValueSum(int[][] board) {
        int m = board.length;
        int n = board[0].length;
        int[][][] suf = new int[m][3][2];
        int[][] p = new int[3][2]; // 最大、次大、第三大
        for (int[] pr : p) {
            pr[0] = Integer.MIN_VALUE;
        }
        for (int i = m - 1; i > 1; i--) {
            update(board[i], p);
            for (int j = 0; j < 3; j++) {
                suf[i][j][0] = p[j][0];
                suf[i][j][1] = p[j][1];
            }
        }

        long ans = Long.MIN_VALUE;
        for (int[] pr : p) {
            pr[0] = Integer.MIN_VALUE; // 重置，计算 pre
        }
        for (int i = 1; i < m - 1; i++) {
            update(board[i - 1], p);
            for (int j = 0; j < n; j++) { // 第二个车
                for (int[] a : p) { // 第一个车
                    for (int[] b : suf[i + 1]) { // 第三个车
                        if (a[1] != j && b[1] != j && a[1] != b[1]) { // 没有同列的车
                            ans = Math.max(ans, (long) a[0] + board[i][j] + b[0]);
                            break;
                        }
                    }
                }
            }
        }
        return ans;
    }

    private void update(int[] row, int[][] p) {
        for (int j = 0; j < row.length; j++) {
            int x = row[j];
            if (x > p[0][0]) {
                if (p[0][1] != j) { // 如果相等，仅更新最大
                    if (p[1][1] != j) { // 如果相等，仅更新最大和次大
                        p[2] = p[1];
                    }
                    p[1] = p[0];
                }
                p[0] = new int[]{x, j};
            } else if (x > p[1][0] && j != p[0][1]) {
                if (p[1][1] != j) { // 如果相等，仅更新次大
                    p[2] = p[1];
                }
                p[1] = new int[]{x, j};
            } else if (x > p[2][0] && j != p[0][1] && j != p[1][1]) {
                p[2] = new int[]{x, j};
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumValueSum(vector<vector<int>>& board) {
        array<pair<int, int>, 3> p; // 最大、次大、第三大
        ranges::fill(p, pair(INT_MIN, -1));

        auto update = [&](vector<int>& row) {
            for (int j = 0; j < row.size(); j++) {
                int x = row[j];
                if (x > p[0].first) {
                    if (p[0].second != j) { // 如果相等，仅更新最大
                        if (p[1].second != j) { // 如果相等，仅更新最大和次大
                            p[2] = p[1];
                        }
                        p[1] = p[0];
                    }
                    p[0] = {x, j};
                } else if (x > p[1].first && j != p[0].second) {
                    if (p[1].second != j) { // 如果相等，仅更新次大
                        p[2] = p[1];
                    }
                    p[1] = {x, j};
                } else if (x > p[2].first && j != p[0].second && j != p[1].second) {
                    p[2] = {x, j};
                }
            }
        };

        int m = board.size(), n = board[0].size();
        vector<array<pair<int, int>, 3>> suf(m);
        for (int i = m - 1; i > 1; i--) {
            update(board[i]);
            suf[i] = p;
        }

        long long ans = LLONG_MIN;
        ranges::fill(p, pair(INT_MIN, -1)); // 重置，计算 pre
        for (int i = 1; i < m - 1; i++) {
            update(board[i - 1]);
            for (int j2 = 0; j2 < n; j2++) { // 第二个车
                for (auto& [x, j1] : p) { // 第一个车
                    for (auto& [z, j3] : suf[i + 1]) { // 第三个车
                        if (j1 != j2 && j1 != j3 && j2 != j3) { // 没有同列的车
                            ans = max(ans, (long long) x + board[i][j2] + z);
                            break;
                        }
                    }
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumValueSum(board [][]int) int64 {
	m := len(board)
	type pair struct{ x, j int }
	suf := make([][3]pair, m)
	p := [3]pair{} // 最大、次大、第三大
	for i := range p {
		p[i].x = math.MinInt
	}
	update := func(row []int) {
		for j, x := range row {
			if x > p[0].x {
				if p[0].j != j { // 如果相等，仅更新最大
					if p[1].j != j { // 如果相等，仅更新最大和次大
						p[2] = p[1]
					}
					p[1] = p[0]
				}
				p[0] = pair{x, j}
			} else if x > p[1].x && j != p[0].j {
				if p[1].j != j { // 如果相等，仅更新次大
					p[2] = p[1]
				}
				p[1] = pair{x, j}
			} else if x > p[2].x && j != p[0].j && j != p[1].j {
				p[2] = pair{x, j}
			}
		}
	}
	for i := m - 1; i > 1; i-- {
		update(board[i])
		suf[i] = p
	}

	ans := math.MinInt
	for i := range p {
		p[i].x = math.MinInt // 重置，计算 pre
	}
	for i, row := range board[:m-2] {
		update(row)
		for j, x := range board[i+1] { // 第二个车
			for _, p := range p { // 第一个车
				for _, q := range suf[i+2] { // 第三个车
					if p.j != j && q.j != j && p.j != q.j { // 没有同列的车
						ans = max(ans, p.x+x+q.x)
						break
					}
				}
			}
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{board}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(m)$。

## 附：费用流做法

建图：

- 创建一个**完全二分图**，左部为行号，右部为列号。
- 把第 $i$ 行记作节点 $i$，第 $j$ 列记作节点 $m+j$。
- 在第 $i$ 行到第 $j$ 列之间连边，容量为 $1$，费用为 $-\textit{grid}[i][j]$。因为我们求的是最小费用流，取负号转成求最大费用流，方便套模板。
- 从超级节点 $R=m+n$ 向所有行节点 $0,1,2,\cdots,m-1$ 连边，容量为 $1$，费用为 $0$。
- 从所有列节点 $m,m+1,m+2,\cdots,m+n-1$ 向超级节点 $C=m+n+1$ 连边，容量为 $1$，费用为 $0$。
- 从超级源点 $S=m+n+2$ 向 $R$ 连边，容量为 $3$，费用为 $0$。如果题目要放置 $4$ 个车，甚至 $k$ 个车，只需把这里的 $3$ 改成 $k$ 即可。 

这样建图可以保证三个车不会同行同列（否则节点 $i$ 或者节点 $m+j$ 的流量会超过 $1$，也就是超过容量）。

计算从 $S$ 到 $C$ 的最小费用流，取相反数，即为答案。

```go
func maximumValueSum(board [][]int) int64 {
	m, n := len(board), len(board[0])
	// rid 为反向边在邻接表中的下标
	type neighbor struct{ to, rid, cap, cost int }
	g := make([][]neighbor, m+n+3)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	R := m + n
	C := m + n + 1
	S := m + n + 2
	for i, row := range board {
		for j, x := range row {
			addEdge(i, m+j, 1, -x)
		}
		addEdge(R, i, 1, 0)
	}
	for j := range board[0] {
		addEdge(m+j, C, 1, 0)
	}
	addEdge(S, R, 3, 0) // 把 3 改成 k 可以支持 k 个车

	// 下面是费用流模板
	dis := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	inQ := make([]bool, len(g))
	spfa := func() bool {
		for i := range dis {
			dis[i] = math.MaxInt
		}
		dis[S] = 0
		inQ[S] = true
		q := []int{S}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				newD := dis[v] + e.cost
				if newD < dis[w] {
					dis[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						inQ[w] = true
						q = append(q, w)
					}
				}
			}
		}
		// 循环结束后所有 inQ[v] 都为 false，无需重置
		return dis[C] < math.MaxInt
	}

	minCost := 0
	for spfa() {
		minF := math.MaxInt
		for v := C; v != S; {
			p := fa[v]
			minF = min(minF, g[p.v][p.i].cap)
			v = p.v
		}
		for v := C; v != S; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
		minCost += dis[C] * minF
	}
	return int64(-minCost)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(kmn)$，其中 $k=3$，$m$ 和 $n$ 分别为 $\textit{board}$ 的行数和列数。由于这里创建的是完全二分图，算法跑 $k=3$ 次 $\mathcal{O}(mn)$ 的 SPFA 就结束了。
- 空间复杂度：$\mathcal{O}(mn)$。

## 套路：枚举中间

- [3128. 直角三角形](https://leetcode.cn/problems/right-triangles/) 1541
- [2242. 节点序列的最大得分](https://leetcode.cn/problems/maximum-score-of-a-node-sequence/) 2304
- [2867. 统计树中的合法路径数目](https://leetcode.cn/problems/count-valid-paths-in-a-tree/) 2428

更多相似题目，见下面 DP 题单中的「**专题：前后缀分解**」。

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
