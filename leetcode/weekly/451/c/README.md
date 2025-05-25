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

- 枚举分配给儿子 $y$ 的预算 $j_y = 0,1,2,\ldots,j$。
- 用前面遍历过的儿子的收益 $\textit{subF}[j-j_y][k]$ 加上儿子 $y$ 的收益 $\textit{dfs}(y)[j_y][k]$，更新 $\textit{subF}[j][k]$ 的最大值。注意这里用到了 0-1 背包的空间优化。

然后考虑 $x$ 是否购买，计算 $f[j][k]$：

- 不买 $x$，那么分配给儿子的预算不变，仍然为 $j$，即 $f[j][k] = \textit{subF}[j][0]$，这里的 $0$ 是因为对于子树 $y$ 来说，父节点 $x$ 一定不买。
- 买 $x$，那么分配给儿子的预算要扣掉 $\textit{cost}$，即 $f[j][k] = \textit{subF}[j-\textit{cost}][1]$，这里的 $1$ 是因为对于子树 $y$ 来说，父节点 $x$ 一定买。

两种情况取最大值，即为 $f[j][k]$。

最终答案为根节点的 $f[\textit{budget}][0]$，这里的 $0$ 是因为根节点没有父节点。

```py [sol-Python3 列表]
# 注意！这个写法很慢，推荐先理解列表的写法，然后看【Python3 字典】的写法
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
                    for jy in range(j + 1):  # 枚举子树 y 的预算为 jy
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

```py [sol-Python3 字典]
fmax = lambda a, b: b if b > a else a

class Solution:
    def maxProfit(self, n: int, present: List[int], future: List[int], hierarchy: List[List[int]], budget: int) -> int:
        g = [[] for _ in range(n)]
        for x, y in hierarchy:
            g[x - 1].append(y - 1)

        def dfs(x: int) -> List[Dict[int, int]]:
            # 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
            sub_f = [defaultdict(int), defaultdict(int)]
            sub_f[0][0] = sub_f[1][0] = 0
            for y in g[x]:
                fy = dfs(y)
                for k, fyk in enumerate(fy):
                    nf = defaultdict(int)
                    for j, v in sub_f[k].items():
                        for jy, vy in fyk.items():  # 枚举子树 y 的预算为 jy
                            s = j + jy
                            if s <= budget:
                                nf[s] = fmax(nf[s], v + vy)
                    sub_f[k] = nf

            f = [None] * 2
            for k in range(2):
                res = sub_f[0].copy()  # 不买 x
                cost = present[x] // (k + 1)
                if cost <= budget:
                    earn = future[x] - cost
                    # 买 x，转移来源为 sub_f[1]，因为对于子树来说，父节点一定买
                    for j, v in sub_f[1].items():
                        j += cost
                        if j <= budget:
                            res[j] = fmax(res[j], v + earn)
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
            int x = e[0] - 1, y = e[1] - 1;
            g[x].add(y);
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
                for (int jy = 0; jy <= j; jy++) { // 枚举子树 y 的预算为 jy
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
            int x = e[0] - 1, y = e[1] - 1;
            g[x].push_back(y);
        }

        auto dfs = [&](this auto&& dfs, int x) -> vector<array<int, 2>> {
            // 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
            vector<array<int, 2>> sub_f(budget + 1);
            for (int y : g[x]) {
                auto fy = dfs(y);
                for (int j = budget; j >= 0; j--) {
                    for (int jy = 0; jy <= j; jy++) { // 枚举子树 y 的预算为 jy
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
				for jy, p := range fy[:j+1] { // 枚举子树 y 的预算为 jy
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

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\cdot \textit{budget}^2)$。有 $n-1$ 条边，每条边计算一次 $\mathcal{O}(\textit{budget}^2)$ 的转移。
- 空间复杂度：$\mathcal{O}(n\cdot \textit{budget})$。

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
