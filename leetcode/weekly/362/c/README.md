## 方法一：枚举全排列

由于所有移走的石子个数等于所有移入的石子个数（即 $0$ 的个数），我们可以把移走的石子的坐标记录到列表 $\textit{from}$ 中（可能有重复的坐标），移入的石子的坐标记录到列表 $\textit{to}$ 中。这两个列表的长度是一样的。

枚举 $\textit{from}$ 的所有排列，与 $\textit{to}$ 匹配，即累加从 $\textit{from}[i]$ 到 $\textit{to}[i]$ 的曼哈顿距离。

所有距离之和的最小值就是答案。

如何枚举全排列请看[【基础算法精讲 16】](https://www.bilibili.com/video/BV1mY411D7f6/)。

```py [sol-Python3]
class Solution:
    def minimumMoves(self, grid: List[List[int]]) -> int:
        from_ = []
        to = []
        for i, row in enumerate(grid):
            for j, cnt in enumerate(row):
                if cnt > 1:
                    from_.extend([(i, j)] * (cnt - 1))
                elif cnt == 0:
                    to.append((i, j))

        ans = inf
        for from2 in permutations(from_):
            total = 0
            for (x1, y1), (x2, y2) in zip(from2, to):
                total += abs(x1 - x2) + abs(y1 - y2)
            ans = min(ans, total)
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumMoves(int[][] grid) {
        List<int[]> from = new ArrayList<>();
        List<int[]> to = new ArrayList<>();
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (grid[i][j] > 1) {
                    for (int k = 1; k < grid[i][j]; k++) {
                        from.add(new int[]{i, j});
                    }
                } else if (grid[i][j] == 0) {
                    to.add(new int[]{i, j});
                }
            }
        }

        int ans = Integer.MAX_VALUE;
        for (List<int[]> from2 : permutations(from)) {
            int total = 0;
            for (int i = 0; i < from2.size(); i++) {
                int[] f = from2.get(i);
                int[] t = to.get(i);
                total += Math.abs(f[0] - t[0]) + Math.abs(f[1] - t[1]);
            }
            ans = Math.min(ans, total);
        }
        return ans;
    }

    private List<List<int[]>> permutations(List<int[]> arr) {
        List<List<int[]>> result = new ArrayList<>();
        permute(arr, 0, result);
        return result;
    }

    private void permute(List<int[]> arr, int start, List<List<int[]>> result) {
        if (start == arr.size()) {
            result.add(new ArrayList<>(arr));
        }
        for (int i = start; i < arr.size(); i++) {
            swap(arr, start, i);
            permute(arr, start + 1, result);
            swap(arr, start, i);
        }
    }

    private void swap(List<int[]> arr, int i, int j) {
        int[] temp = arr.get(i);
        arr.set(i, arr.get(j));
        arr.set(j, temp);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumMoves(vector<vector<int>>& grid) {
        vector<pair<int, int>> from, to;
        for (int i = 0; i < grid.size(); i++) {
            for (int j = 0; j < grid[i].size(); j++) {
                if (grid[i][j]) {
                    for (int k = 1; k < grid[i][j]; k++) {
                        from.emplace_back(i, j);
                    }
                } else {
                    to.emplace_back(i, j);
                }
            }
        }

        int ans = INT_MAX;
        do {
            int total = 0;
            for (int i = 0; i < from.size(); i++) {
                total += abs(from[i].first - to[i].first) + abs(from[i].second - to[i].second);
            }
            ans = min(ans, total);
        } while (next_permutation(from.begin(), from.end()));
        return ans;
    }
};
```

```go [sol-Go]
type pair struct{ x, y int }

func minimumMoves(grid [][]int) int {
    var from, to []pair
    for i, row := range grid {
        for j, cnt := range row {
            if cnt > 1 {
                for k := 1; k < cnt; k++ {
                    from = append(from, pair{i, j})
                }
            } else if cnt == 0 {
                to = append(to, pair{i, j})
            }
        }
    }

    ans := math.MaxInt
    permute(from, 0, func() {
        total := 0
        for i, f := range from {
            total += abs(f.x-to[i].x) + abs(f.y-to[i].y)
        }
        ans = min(ans, total)
    })
    return ans
}

func permute(a []pair, i int, do func()) {
    if i == len(a) {
        do()
        return
    }
    permute(a, i+1, do)
    for j := i + 1; j < len(a); j++ {
        a[i], a[j] = a[j], a[i]
        permute(a, i+1, do)
        a[i], a[j] = a[j], a[i]
    }
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

```js [sol-JavaScript]
var minimumMoves = function(grid) {
    const from = [];
    const to = [];
    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid[i].length; j++) {
            if (grid[i][j] > 1) {
                for (let k = 1; k < grid[i][j]; k++) {
                    from.push([i, j]);
                }
            } else if (grid[i][j] === 0) {
                to.push([i, j]);
            }
        }
    }
    
    let ans = Infinity;
    for (let from2 of permute(from)) {
        let total = 0;
        for (let i = 0; i < from2.length; i++) {
            total += Math.abs(from2[i][0] - to[i][0]) + Math.abs(from2[i][1] - to[i][1]);
        }
        ans = Math.min(ans, total);
    }
    return ans;
};

function permute(arr) {
    const result = [];
    perm(arr, 0, result);
    return result;
}

function perm(arr, start, result) {
    if (start === arr.length) {
        result.push([...arr]);
    }
    for (let i = start; i < arr.length; i++) {
        [arr[start], arr[i]] = [arr[i], arr[start]];
        perm(arr, start + 1, result);
        [arr[start], arr[i]] = [arr[i], arr[start]];
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\cdot(mn)!)$。其中 $m=3,n=3$。
- 空间复杂度：$\mathcal{O}(mn)$。

## 方法二：最小费用最大流

更快的做法是**最小费用最大流**。即使是 $10\times 10$ 的网格也可以做。

建图规则如下：

- 从每个大于 $1$ 的格子向每个等于 $0$ 的格子连边，容量为 $1$，费用为两个格子之间的曼哈顿距离。
- 从超级源点向每个大于 $1$ 的格子连边，容量为格子的值减一（即移走的石子数），费用为 $0$。
- 从每个等于 $0$ 的格子向超级汇点连边，容量 $1$（即移入的石子数），费用为 $0$。

答案为最大流时，对应的最小费用。

```go
func minimumMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	src := m * n   // 超级源点
	dst := src + 1 // 超级汇点
	type edge struct{ to, rid, cap, cost int }
	g := make([][]edge, m*n+2)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], edge{to, len(g[to]), cap, cost})
		g[to] = append(g[to], edge{from, len(g[from]) - 1, 0, -cost})
	}
	for x, row := range grid {
		for y, v := range row {
			if v > 1 {
				addEdge(src, x*n+y, v-1, 0)
				for i, r := range grid {
					for j, w := range r {
						if w == 0 {
							addEdge(x*n+y, i*n+j, 1, abs(x-i)+abs(y-j))
						}
					}
				}
			} else if v == 0 {
				addEdge(x*n+y, dst, 1, 0)
			}
		}
	}

	// 下面是最小费用最大流模板
	const inf int = 1e9
	dist := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	spfa := func() bool {
		for i := range dist {
			dist[i] = 1e9
		}
		dist[src] = 0
		inQ := make([]bool, len(g))
		inQ[src] = true
		q := []int{src}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				if newD := dist[v] + e.cost; newD < dist[w] {
					dist[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						q = append(q, w)
						inQ[w] = true
					}
				}
			}
		}
		return dist[dst] < inf
	}
	ek := func() (maxFlow, minCost int) {
		for spfa() {
			// 沿 st-end 的最短路尽量增广
			minF := inf
			for v := dst; v != src; {
				p := fa[v]
				if c := g[p.v][p.i].cap; c < minF {
					minF = c
				}
				v = p.v
			}
			for v := dst; v != src; {
				p := fa[v]
				e := &g[p.v][p.i]
				e.cap -= minF
				g[v][e.rid].cap += minF
				v = p.v
			}
			maxFlow += minF
			minCost += dist[dst] * minF
		}
		return
	}
	_, cost := ek()
	return cost
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((mn)^4)$。其中 $m=3,n=3$。复杂度为图中节点个数、边的个数、最大流三者之积，也就是 $\mathcal{O}(mn\times (mn)^2\times mn) = \mathcal{O}((mn)^4)$。
- 空间复杂度：$\mathcal{O}((mn)^2)$。

## 相似题目

下面是一些涉及到「匹配」的题目：

- [1947. 最大兼容性评分和](https://leetcode.cn/problems/maximum-compatibility-score-sum/)
- [1349. 参加考试的最大学生数](https://leetcode.cn/problems/maximum-students-taking-exam/)
- [LCP 04. 覆盖](https://leetcode.cn/problems/broken-board-dominoes/)
- [1879. 两个数组最小的异或值之和](https://leetcode.cn/problems/minimum-xor-sum-of-two-arrays/)
- [2172. 数组的最大与和](https://leetcode.cn/problems/maximum-and-sum-of-array/)

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
