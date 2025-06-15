**注**：数据范围说 `hierarchy.length == n - 1`，且 `员工 1 是所有员工的直接或间接上司`，所以输入的是一个 $n$ 点 $n-1$ 边的连通图，即树。

## 前置知识

请确认你已经掌握如下知识点：

1. **0-1 背包**，包括**空间优化**的原理。请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。
2. **状态机 DP**。请看[【基础算法精讲 21】](https://www.bilibili.com/video/BV1ho4y1W7QK/)。

## 寻找子问题

推荐看 [本题视频讲解](https://www.bilibili.com/video/BV1o1jgzJE51/?t=7m52s)，从特殊到一般，带你一步步思考。

对于节点 $x$ 来说：

- 如果不买 $x$（即不买 $\textit{present}[x]$），且预算至多为 $j$，那么问题变成：
   - 从 $x$ 的所有子树 $y$ 中能得到的最大利润之和。
   - 所有子树 $y$ 的花费总和必须 $\le j$。
   - $y$ 不能半价购买。
- 如果买 $x$，且预算至多为 $j$，那么问题变成：
   - 从 $x$ 的所有子树 $y$ 中能得到的最大利润之和，加上买 $x$ 得到的利润。
   - 所有子树 $y$ 的花费总和必须 $\le j-\textit{cost}$，其中 $\textit{cost}$ 等于 $\textit{present}[x]$ 或者 $\left\lfloor\dfrac{\textit{present}[x]}{2}\right\rfloor$。
   - $y$ 可以半价购买。

## 状态设计和状态转移

$\textit{dfs}(x)$ 返回一个长为 $(\textit{budget}+1)\times 2$ 的二维数组 $f$，其中 $f[j][k]$ 表示：

- 从子树 $x$ 中能得到的最大利润之和。
- 预算为 $j$，即花费总和 $\le j$。
- $k=0$ 表示 $x$ 不能半价购买，$k=1$ 表示 $x$ 可以半价购买。

首先计算 $x$ 的所有儿子子树 $y$ 的最大利润总和 $\textit{subF}[j][k]$。枚举 $x$ 的儿子 $y$：

- 枚举分配给当前儿子 $y$ 的预算 $j_y = 0,1,2,\ldots,j$，那么分配给前面遍历过的儿子的总预算为 $j-j_y$。
- 用前面遍历过的儿子的收益 $\textit{subF}[j-j_y][k]$ 加上当前儿子 $y$ 的收益 $\textit{dfs}(y)[j_y][k]$，更新 $\textit{subF}[j][k]$ 的最大值。注意这里用了 0-1 背包的空间优化。

然后考虑 $x$ 是否购买，计算 $f[j][k]$：

- 不买 $x$，那么分配给儿子的预算不变，仍然为 $j$，即 $f[j][k] = \textit{subF}[j][0]$，这里的 $0$ 是因为对于子树 $y$ 来说，父节点 $x$ 一定不买。
- 买 $x$，那么分配给儿子的预算要扣掉 $\textit{cost}$，即 $f[j][k] = \textit{subF}[j-\textit{cost}][1]$，这里的 $1$ 是因为对于子树 $y$ 来说，父节点 $x$ 一定买。

两种情况取最大值，得

$$
f[j][k] = \max(\textit{subF}[j][0], \textit{subF}[j-\textit{cost}][1] + \textit{future}[x] - \textit{cost})
$$

最终答案为根节点的 $f[\textit{budget}][0]$，这里的 $0$ 是因为根节点没有父节点。

## 写法一：至多

```py [sol-Python3]
# 注意！这个写法很慢，更快的写法见写法二
max = lambda a, b: b if b > a else a

class Solution:
    def maxProfit(self, n: int, present: List[int], future: List[int], hierarchy: List[List[int]], budget: int) -> int:
        g = [[] for _ in range(n)]
        for x, y in hierarchy:
            g[x - 1].append(y - 1)

        def dfs(x: int) -> List[List[int]]:
            # 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
            sub_f = [[0, 0] for _ in range(budget + 1)]
            for y in g[x]:
                fy = dfs(y)
                for j in range(budget, -1, -1):
                    # 枚举子树 y 的预算为 jy
                    # 当作一个体积为 jy，价值为 fy[jy][k] 的物品
                    for jy in range(j + 1):  
                        for k in range(2):
                            sub_f[j][k] = max(sub_f[j][k], sub_f[j - jy][k] + fy[jy][k])

            f = [[0, 0] for _ in range(budget + 1)]
            for j in range(budget + 1):
                for k in range(2):
                    cost = present[x] // (k + 1)
                    if j >= cost:
                        # 不买 x，转移来源是 sub_f[j][0]
                        # 买 x，转移来源为 sub_f[j-cost][1]，因为对于子树来说，父节点一定买
                        f[j][k] = max(sub_f[j][0], sub_f[j - cost][1] + future[x] - cost)
                    else:  # 只能不买 x
                        f[j][k] = sub_f[j][0]
            return f

        return dfs(0)[budget][0]
```

```java [sol-Java]
class Solution {
    public int maxProfit(int n, int[] present, int[] future, int[][] hierarchy, int budget) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : hierarchy) {
            g[e[0] - 1].add(e[1] - 1);
        }

        int[][] f0 = dfs(0, g, present, future, budget);
        return f0[budget][0];
    }

    private int[][] dfs(int x, List<Integer>[] g, int[] present, int[] future, int budget) {
        // 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
        int[][] subF = new int[budget + 1][2];
        for (int y : g[x]) {
            int[][] fy = dfs(y, g, present, future, budget);
            for (int j = budget; j >= 0; j--) {
                // 枚举子树 y 的预算为 jy
                // 当作一个体积为 jy，价值为 fy[jy][k] 的物品
                for (int jy = 0; jy <= j; jy++) {
                    for (int k = 0; k < 2; k++) {
                        subF[j][k] = Math.max(subF[j][k], subF[j - jy][k] + fy[jy][k]);
                    }
                }
            }
        }

        int[][] f = new int[budget + 1][2];
        for (int j = 0; j <= budget; j++) {
            for (int k = 0; k < 2; k++) {
                int cost = present[x] / (k + 1);
                if (j >= cost) {
                    // 不买 x，转移来源是 subF[j][0]
                    // 买 x，转移来源为 subF[j-cost][1]，因为对于子树来说，父节点一定买
                    f[j][k] = Math.max(subF[j][0], subF[j - cost][1] + future[x] - cost);
                } else { // 只能不买 x
                    f[j][k] = subF[j][0];
                }
            }
        }
        return f;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxProfit(int n, vector<int>& present, vector<int>& future, vector<vector<int>>& hierarchy, int budget) {
        vector<vector<int>> g(n);
        for (auto& e : hierarchy) {
            g[e[0] - 1].push_back(e[1] - 1);
        }

        auto dfs = [&](this auto&& dfs, int x) -> vector<array<int, 2>> {
            // 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
            vector<array<int, 2>> sub_f(budget + 1);
            for (int y : g[x]) {
                auto fy = dfs(y);
                for (int j = budget; j >= 0; j--) {
                    // 枚举子树 y 的预算为 jy
                    // 当作一个体积为 jy，价值为 fy[jy][k] 的物品
                    for (int jy = 0; jy <= j; jy++) {
                        for (int k = 0; k < 2; k++) {
                            sub_f[j][k] = max(sub_f[j][k], sub_f[j - jy][k] + fy[jy][k]);
                        }
                    }
                }
            }

            vector<array<int, 2>> f(budget + 1);
            for (int j = 0; j <= budget; j++) {
                for (int k = 0; k < 2; k++) {
                    int cost = present[x] / (k + 1);
                    if (j >= cost) {
                        // 不买 x，转移来源是 sub_f[j][0]
                        // 买 x，转移来源为 sub_f[j-cost][1]，因为对于子树来说，父节点一定买
                        f[j][k] = max(sub_f[j][0], sub_f[j - cost][1] + future[x] - cost);
                    } else { // 只能不买 x
                        f[j][k] = sub_f[j][0];
                    }
                }
            }
            return f;
        };

        return dfs(0)[budget][0];
    }
};
```

```go [sol-Go]
func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	g := make([][]int, n)
	for _, e := range hierarchy {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y)
	}

	var dfs func(int) [][2]int
	dfs = func(x int) [][2]int {
		// 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
		subF := make([][2]int, budget+1)
		for _, y := range g[x] {
			fy := dfs(y)
			for j := budget; j >= 0; j-- {
				// 枚举子树 y 的预算为 jy
				// 当作一个体积为 jy，价值为 resY=fy[jy][k] 的物品
				for jy, p := range fy[:j+1] {
					for k, resY := range p {
						subF[j][k] = max(subF[j][k], subF[j-jy][k]+resY)
					}
				}
			}
		}

		f := make([][2]int, budget+1)
		for j, p := range subF {
			for k := range 2 {
				cost := present[x] / (k + 1)
				if j >= cost {
					// 不买 x，转移来源是 subF[j][0]
					// 买 x，转移来源为 subF[j-cost][1]，因为对于子树来说，父节点一定买
					f[j][k] = max(p[0], subF[j-cost][1]+future[x]-cost)
				} else { // 只能不买 x
					f[j][k] = p[0]
				}
			}
		}
		return f
	}

	return dfs(0)[budget][0]
}
```

## 写法二：恰好

把状态值改成在总花费**恰好**为 $j$ 的情况下的最大利润。

同时交换 $f$ 数组的维度，改成两个长为 $\textit{budget}+1$ 的数组。

```py [sol-Python3 列表]
# 更快的写法见【Python3 字典】
fmax = lambda a, b: b if b > a else a

class Solution:
    def maxProfit(self, n: int, present: List[int], future: List[int], hierarchy: List[List[int]], budget: int) -> int:
        g = [[] for _ in range(n)]
        for x, y in hierarchy:
            g[x - 1].append(y - 1)

        def dfs(x: int) -> List[List[int, int]]:
            # 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
            sub_f = [[0] + [-inf] * budget for _ in range(2)]
            for y in g[x]:
                fy = dfs(y)
                for k, fyk in enumerate(fy):
                    nf = [0] + [-inf] * budget
                    for jy, res_y in enumerate(fyk):
                        if res_y < 0:  # 重要优化：物品价值为负数，一定不选
                            continue
                        for j in range(jy, budget + 1):
                            nf[j] = fmax(nf[j], sub_f[k][j - jy] + res_y)
                    sub_f[k] = nf

            f = [None] * 2
            for k in range(2):
                # 不买 x，转移来源为 sub_f[0]，因为对于子树来说，父节点一定不买
                f[k] = sub_f[0].copy()
                cost = present[x] // (k + 1)
                # 买 x，转移来源为 sub_f[1]，因为对于子树来说，父节点一定买
                for j in range(cost, budget + 1):
                    f[k][j] = fmax(f[k][j], sub_f[1][j - cost] + future[x] - cost)
            return f

        return max(dfs(0)[0])
```

```py [sol-Python3 字典]
fmax = lambda a, b: b if b > a else a

class Solution:
    def maxProfit(self, n: int, present: List[int], future: List[int], hierarchy: List[List[int]], budget: int) -> int:
        g = [[] for _ in range(n)]
        for x, y in hierarchy:
            g[x - 1].append(y - 1)

        def dfs(x: int) -> List[Dict[int, int]]:
            sub_f = [defaultdict(int) for _ in range(2)]
            sub_f[0][0] = sub_f[1][0] = 0
            for y in g[x]:
                fy = dfs(y)
                for k, fyk in enumerate(fy):
                    nf = defaultdict(int)
                    for j, pre_res_y in sub_f[k].items():
                        for jy, res_y in fyk.items():
                            sj = j + jy
                            if sj <= budget:
                                nf[sj] = fmax(nf[sj], pre_res_y + res_y)
                    sub_f[k] = nf

            f = [None] * 2
            for k in range(2):
                res = sub_f[0].copy()
                cost = present[x] // (k + 1)
                if cost <= budget:
                    earn = future[x] - cost
                    for j, res_y in sub_f[1].items():
                        sj = j + cost
                        if sj <= budget:
                            res[sj] = fmax(res[sj], res_y + earn)
                f[k] = res
            return f

        return max(dfs(0)[0].values())
```

```java [sol-Java]
class Solution {
    public int maxProfit(int n, int[] present, int[] future, int[][] hierarchy, int budget) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : hierarchy) {
            g[e[0] - 1].add(e[1] - 1);
        }

        int[][] f0 = dfs(0, g, present, future, budget);
        return Arrays.stream(f0[0]).max().getAsInt();
    }

    private int[][] dfs(int x, List<Integer>[] g, int[] present, int[] future, int budget) {
        // 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
        int[][] subF = new int[2][budget + 1];
        Arrays.fill(subF[0], Integer.MIN_VALUE / 2); // 表示不存在对应的花费总和
        Arrays.fill(subF[1], Integer.MIN_VALUE / 2);
        subF[0][0] = subF[1][0] = 0;
        for (int y : g[x]) {
            int[][] fy = dfs(y, g, present, future, budget);
            for (int k = 0; k < 2; k++) {
                int[] nf = new int[budget + 1];
                Arrays.fill(nf, Integer.MIN_VALUE / 2);
                nf[0] = 0;
                for (int jy = 0; jy <= budget; jy++) {
                    int resY = fy[k][jy];
                    if (resY < 0) { // 重要优化：物品价值为负数，一定不选
                        continue;
                    }
                    for (int j = jy; j <= budget; j++) {
                        nf[j] = Math.max(nf[j], subF[k][j - jy] + resY);
                    }
                }
                subF[k] = nf;
            }
        }

        int[][] f = new int[2][];
        for (int k = 0; k < 2; k++) {
            // 不买 x，转移来源为 subF[0]，因为对于子树来说，父节点一定不买
            f[k] = subF[0].clone();
            int cost = present[x] / (k + 1);
            // 买 x，转移来源为 subF[1]，因为对于子树来说，父节点一定买
            for (int j = cost; j <= budget; j++) {
                f[k][j] = Math.max(f[k][j], subF[1][j - cost] + future[x] - cost);
            }
        }
        return f;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxProfit(int n, vector<int>& present, vector<int>& future, vector<vector<int>>& hierarchy, int budget) {
        vector<vector<int>> g(n);
        for (auto& e : hierarchy) {
            g[e[0] - 1].push_back(e[1] - 1);
        }

        auto dfs = [&](this auto&& dfs, int x) -> array<vector<int>, 2> {
            // 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
            vector<int> sub_f[2]{vector<int>(budget + 1, INT_MIN / 2), vector<int>(budget + 1, INT_MIN / 2)};
            sub_f[0][0] = sub_f[1][0] = 0;
            for (int y : g[x]) {
                auto fy = dfs(y);
                for (int k = 0; k < 2; k++) {
                    vector<int> nf(budget + 1, INT_MIN / 2);
                    nf[0] = 0;
                    for (int jy = 0; jy <= budget; jy++) {
                        int res_y = fy[k][jy];
                        if (res_y < 0) { // 重要优化：物品价值为负数，一定不选
                            continue;
                        }
                        for (int j = jy; j <= budget; j++) {
                            nf[j] = max(nf[j], sub_f[k][j - jy] + res_y);
                        }
                    }
                    sub_f[k] = move(nf);
                }
            }

            array<vector<int>, 2> f;
            for (int k = 0; k < 2; k++) {
                // 不买 x，转移来源为 sub_f[0]，因为对于子树来说，父节点一定不买
                f[k] = sub_f[0];
                int cost = present[x] / (k + 1);
                // 买 x，转移来源为 sub_f[1]，因为对于子树来说，父节点一定买
                for (int j = cost; j <= budget; j++) {
                    f[k][j] = max(f[k][j], sub_f[1][j - cost] + future[x] - cost);
                }
            }
            return f;
        };

        return ranges::max(dfs(0)[0]);
    }
};
```

```go [sol-Go]
func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	g := make([][]int, n)
	for _, e := range hierarchy {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y)
	}

	var dfs func(int) [2][]int
	dfs = func(x int) [2][]int {
		// 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
		subF := [2][]int{make([]int, budget+1), make([]int, budget+1)}
		for i := 1; i <= budget; i++ {
			subF[0][i] = math.MinInt / 2 // 表示不存在对应的花费总和
			subF[1][i] = math.MinInt / 2
		}
		for _, y := range g[x] {
			fy := dfs(y)
			for k, fyk := range fy {
				nf := make([]int, budget+1)
				for i := 1; i <= budget; i++ {
					nf[i] = math.MinInt / 2
				}
				for jy, resY := range fyk {
					if resY < 0 { // 重要优化：物品价值为负数，一定不选
						continue
					}
					for j := jy; j <= budget; j++ {
						nf[j] = max(nf[j], subF[k][j-jy]+resY)
					}
				}
				subF[k] = nf
			}
		}

		f := [2][]int{}
		for k := range 2 {
			// 不买 x，转移来源为 subF[0]，因为对于子树来说，父节点一定不买
			f[k] = slices.Clone(subF[0])
			cost := present[x] / (k + 1)
			// 买 x，转移来源为 subF[1]，因为对于子树来说，父节点一定买
			for j := cost; j <= budget; j++ {
				f[k][j] = max(f[k][j], subF[1][j-cost]+future[x]-cost)
			}
		}
		return f
	}

	return slices.Max(dfs(0)[0])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\cdot \textit{budget}^2)$。有 $n-1$ 条边，每条边计算一次 $\mathcal{O}(\textit{budget}^2)$ 的转移。
- 空间复杂度：$\mathcal{O}(h\cdot \textit{budget})$，其中 $h$ 是树的高度。在随机数据下，$h=\Theta(\sqrt n)$，这个做法比在 DFS 外面创建数组更好。

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
