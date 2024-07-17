## 方法一：二进制枚举 + Floyd

根据 [从集合论到位运算](https://leetcode.cn/circle/discuss/CaOJ45/) 中「枚举集合」的技巧，枚举 $\{0,1,2,\cdots,n-1\}$ 的所有子集 $S$，作为保留的节点集合，关闭（删除）不在 $S$ 中的节点。

然后用[【模板讲解】带你发明 Floyd 算法：从记忆化搜索到递推](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/solution/dai-ni-fa-ming-floyd-suan-fa-cong-ji-yi-m8s51/) 求出任意两点之间的最短路。如果保留的节点之间的最短路均不超过 $\textit{maxDistance}$，则把答案加一。

代码实现时，可以先预处理原图的邻接矩阵 $g$，这样无需每次枚举子集都要遍历一次 $\textit{roads}$，直接拷贝一份 $g$ 中保留的节点即可。

```py [sol-Python3]
class Solution:
    def numberOfSets(self, n: int, maxDistance: int, roads: List[List[int]]) -> int:
        g = [[inf] * n for _ in range(n)]
        for x, y, wt in roads:
            g[x][y] = min(g[x][y], wt)
            g[y][x] = min(g[y][x], wt)

        f = [None] * n
        def check(s: int) -> int:
            for i, row in enumerate(g):
                if s >> i & 1:  # i 在集合 s 中
                    f[i] = row.copy()

            # Floyd 算法（只考虑在 s 中的节点）
            for k in range(n):
                if (s >> k & 1) == 0:  # k 不在集合 s 中
                    continue
                for i in range(n):
                    if (s >> i & 1) == 0 or f[i][k] == inf:
                        continue
                    for j in range(n):
                        f[i][j] = min(f[i][j], f[i][k] + f[k][j])

            # 判断保留的节点之间的最短路是否均不超过 maxDistance
            for i, di in enumerate(f):
                if (s >> i & 1) == 0:  # i 不在集合 s 中
                    continue
                for j, dij in enumerate(di[:i]):
                    if s >> j & 1 and dij > maxDistance:
                        return 0
            return 1

        # 枚举子集 s，作为保留的节点，判断这些节点否满足要求
        return sum(check(s) for s in range(1 << n))
```

```java [sol-Java]
class Solution {
    public int numberOfSets(int n, int maxDistance, int[][] roads) {
        int[][] g = new int[n][n];
        for (int[] row : g) {
            Arrays.fill(row, Integer.MAX_VALUE / 2); // 防止加法溢出
        }
        for (int[] e : roads) {
            int x = e[0];
            int y = e[1];
            int wt = e[2];
            g[x][y] = Math.min(g[x][y], wt);
            g[y][x] = Math.min(g[y][x], wt);
        }

        int ans = 0;
        int[][] f = new int[n][n];
        next:
        for (int s = 0; s < (1 << n); s++) {
            for (int i = 0; i < n; i++) {
                if ((s >> i & 1) == 1) {
                    System.arraycopy(g[i], 0, f[i], 0, n);
                }
            }

            // Floyd 算法（只考虑在 s 中的节点）
            for (int k = 0; k < n; k++) {
                if ((s >> k & 1) == 0) continue;
                for (int i = 0; i < n; i++) {
                    if ((s >> i & 1) == 0) continue;
                    for (int j = 0; j < n; j++) {
                        f[i][j] = Math.min(f[i][j], f[i][k] + f[k][j]);
                    }
                }
            }

            // 判断保留的节点之间的最短路是否均不超过 maxDistance
            for (int i = 0; i < n; i++) {
                if ((s >> i & 1) == 0) continue;
                for (int j = 0; j < i; j++) {
                    if ((s >> j & 1) == 1 && f[i][j] > maxDistance) {
                        continue next;
                    }
                }
            }
            ans++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfSets(int n, int maxDistance, vector<vector<int>>& roads) {
        vector<vector<int>> g(n, vector<int>(n, INT_MAX / 2)); // 防止加法溢出
        for (auto& e: roads) {
            int x = e[0], y = e[1], wt = e[2];
            g[x][y] = min(g[x][y], wt);
            g[y][x] = min(g[y][x], wt);
        }

        vector<vector<int>> f(n);
        auto check = [&](int s) -> bool {
            for (int i = 0; i < n; i++) {
                if ((s >> i) & 1) {
                    f[i] = g[i];
                }
            }

            // Floyd 算法（只考虑在 s 中的节点）
            for (int k = 0; k < n; k++) {
                if (((s >> k) & 1) == 0) continue;
                for (int i = 0; i < n; i++) {
                    if (((s >> i) & 1) == 0) continue;
                    for (int j = 0; j < n; j++) {
                        f[i][j] = min(f[i][j], f[i][k] + f[k][j]);
                    }
                }
            }

            // 判断保留的节点之间的最短路是否均不超过 maxDistance
            for (int i = 0; i < n; i++) {
                if (((s >> i) & 1) == 0) continue;
                for (int j = 0; j < i; j++) {
                    if ((s >> j) & 1 && f[i][j] > maxDistance) {
                        return false;
                    }
                }
            }
            return true;
        };

        int ans = 0;
        for (int s = 0; s < (1 << n); s++) { // 枚举子集
            ans += check(s);
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfSets(n, maxDistance int, roads [][]int) (ans int) {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
            g[i][j] = math.MaxInt / 2 // 防止加法溢出
		}
	}
	for _, e := range roads {
		x, y, wt := e[0], e[1], e[2]
		g[x][y] = min(g[x][y], wt)
		g[y][x] = min(g[y][x], wt)
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
next:
	for s := 0; s < 1<<n; s++ { // 枚举子集
		for i, row := range g {
			if s>>i&1 == 0 { continue }
			copy(f[i], row)
		}

		// Floyd 算法（只考虑在 s 中的节点）
		for k := range f {
			if s>>k&1 == 0 { continue }
			for i := range f {
				if s>>i&1 == 0 { continue }
				for j := range f {
					f[i][j] = min(f[i][j], f[i][k]+f[k][j])
				}
			}
		}

		// 判断保留的节点之间的最短路是否均不超过 maxDistance
		for i, di := range f {
			if s>>i&1 == 0 { continue }
			for j, dij := range di[:i] {
				if s>>j&1 > 0 && dij > maxDistance {
					continue next
				}
			}
		}
		ans++
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m + n^3\cdot 2^n)$，其中 $m$ 为 $\textit{roads}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 方法二：状压 DP + Floyd

请先阅读：[【模板讲解】带你发明 Floyd 算法：从记忆化搜索到递推](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/solution/dai-ni-fa-ming-floyd-suan-fa-cong-ji-yi-m8s51/)

在三维 Floyd DP 的基础上加以修改。

原来的定义是：$f[k][i][j]$ 表示从 $i$ 到 $j$ 的最短路长度，并且这条最短路的**中间节点**编号都 $\le k$。

为了方便在子集之间转移，把定义改成：$f[S][i][j]$ 表示从 $i$ 到 $j$ 的最短路长度，并且这条最短路的**中间节点**编号在集合 $S$ 中。

假设现在计算出了 $T=\{0,1,3\}$ 的 $f[T]$，现在要计算 $S=\{0,1,2,3\}$ 的 $f[S]$，我们可以在 $T$ 的基础上，添加节点 $k=2$，考虑节点 $k=2$ 带来的影响：

- 不影响 $i$ 到 $j$ 的最短路：$f[S][i][j] = f[T][i][j]$；
- 影响 $i$ 到 $j$ 的最短路，也就是从 $i$ 到 $k$ 再到 $j$ 更短：$[S][i][j] = f[T][i][k] + f[T][k][j]$。 

二者取最小值，得

$$
f[S][i][j] = \min(f[T][i][j], f[T][i][k] + f[T][k][j])
$$

初始值和原来的 Floyd 算法一样：$f[\varnothing][i][j] = g[i][j]$，也就是 $\textit{roads}$ 中的 $i$ 到 $j$ 这条边的长度。

> 上述方法意味着，我们只需要 $\mathcal{O}(n^2)$ 的时间，就可以「增量地」算出一个集合任意点对的最短路，而不是方法一中的 $\mathcal{O}(n^3)$。

⚠注意：该计算过程不需要保证 $i$ 和 $j$ 都是 $S$ 中的元素，$S$ 只是中间节点的编号集合，不是保留的编号集合。但是，设 $i$ 和 $j$ 都是 $S$ 中的元素，从点 $i$ 到点 $j$，在只经过 $S$ 中的节点的情况下，最短路长度就是 $f[S][i][j]$。所以在「$i$ 和 $j$ 都是 $S$ 中的元素」的前提下，$S$ 也可以作为当前保留的节点集合，$f[S][i][j]$ 可以用来与 $\textit{maxDistance}$ 比大小。

代码实现时，节点 $k$ 可以是 $S$ 中的任一元素（比如最小元素或最大元素）。

```py [sol-Python3]
class Solution:
    def numberOfSets(self, n: int, maxDistance: int, roads: List[List[int]]) -> int:
        g = [[inf] * n for _ in range(n)]
        for x, y, wt in roads:
            g[x][y] = min(g[x][y], wt)
            g[y][x] = min(g[y][x], wt)

        ans = 1  # s=0 一定满足要求
        f = [[[inf] * n for _ in range(n)] for _ in range(1 << n)]
        f[0] = g
        for s in range(1, 1 << n):
            k = s.bit_length() - 1
            t = s ^ (1 << k)
            ok = 1
            for i in range(n):
                for j in range(n):
                    f[s][i][j] = min(f[t][i][j], f[t][i][k] + f[t][k][j])  # 手动求 min 可以更快
                    if ok and j < i and s >> i & 1 and s >> j & 1 and f[s][i][j] > maxDistance:
                        ok = 0
            ans += ok
        return ans
```

```java [sol-Java]
class Solution {
    public int numberOfSets(int n, int maxDistance, int[][] roads) {
        int[][] g = new int[n][n];
        for (int[] row : g) {
            Arrays.fill(row, Integer.MAX_VALUE / 2);
        }
        for (int[] e : roads) {
            int x = e[0];
            int y = e[1];
            int wt = e[2];
            g[x][y] = Math.min(g[x][y], wt);
            g[y][x] = Math.min(g[y][x], wt);
        }

        int ans = 1; // s=0 一定满足要求
        int[][][] f = new int[1 << n][n][n];
        for (int[][] matrix : f) {
            for (int[] row : matrix) {
                Arrays.fill(row, Integer.MAX_VALUE / 2);
            }
        }
        f[0] = g;
        for (int s = 1; s < (1 << n); s++) {
            int k = Integer.numberOfTrailingZeros(s);
            int t = s ^ (1 << k);
            boolean ok = true;
            for (int i = 0; i < n; i++) {
                for (int j = 0; j < n; j++) {
                    f[s][i][j] = Math.min(f[t][i][j], f[t][i][k] + f[t][k][j]);
                    if (ok && j < i && (s >> i & 1) != 0 && (s >> j & 1) != 0 && f[s][i][j] > maxDistance) {
                        ok = false;
                    }
                }
            }
            ans += ok ? 1 : 0;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfSets(int n, int maxDistance, vector<vector<int>>& roads) {
        vector<vector<int>> g(n, vector<int>(n, INT_MAX / 2));
        for (auto& e: roads) {
            int x = e[0], y = e[1], wt = e[2];
            g[x][y] = min(g[x][y], wt);
            g[y][x] = min(g[y][x], wt);
        }

        int ans = 1; // s=0 一定满足要求
        vector<vector<vector<int>>> f(1 << n, vector<vector<int>>(n, vector<int>(n, INT_MAX / 2)));
        f[0] = g;
        for (int s = 1; s < (1 << n); s++) {
            int k = __builtin_ctz(s);
            int t = s ^ (1 << k);
            bool ok = true;
            for (int i = 0; i < n; i++) {
                for (int j = 0; j < n; j++) {
                    f[s][i][j] = min(f[t][i][j], f[t][i][k] + f[t][k][j]);
                    if (ok && j < i && (s >> i & 1) && (s >> j & 1) && f[s][i][j] > maxDistance) {
                        ok = false;
                    }
                }
            }
            ans += ok;
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfSets(n, maxDistance int, roads [][]int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
            g[i][j] = math.MaxInt / 2 // 防止加法溢出
		}
	}
	for _, e := range roads {
		x, y, wt := e[0], e[1], e[2]
		g[x][y] = min(g[x][y], wt)
		g[y][x] = min(g[y][x], wt)
	}

	ans := 1 // s=0 一定满足要求
	f := make([][][]int, 1<<n)
	for i := range f {
		f[i] = make([][]int, n)
		for j := range f[i] {
			f[i][j] = make([]int, n)
			for k := range f[i][j] {
				f[i][j][k] = math.MaxInt / 2
			}
		}
	}
	f[0] = g
	for s := uint(1); s < 1<<n; s++ {
		k := bits.TrailingZeros(s)
		t := s ^ (1 << k)
		ok := true
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[s][i][j] = min(f[t][i][j], f[t][i][k]+f[t][k][j])
				if ok && j < i && s>>i&1 != 0 && s>>j&1 != 0 && f[s][i][j] > maxDistance {
					ok = false
				}
			}
		}
		if ok {
			ans++
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m + n^2\cdot 2^n)$，其中 $m$ 为 $\textit{roads}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2\cdot 2^n)$。

## 同类题目

#### Floyd

- [2642. 设计可以求最短路径的图类](https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/) 1811
- [1334. 阈值距离内邻居最少的城市](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/) 1855
- [2976. 转换字符串的最小成本 I](https://leetcode.cn/problems/minimum-cost-to-convert-string-i/) 1882
- [2977. 转换字符串的最小成本 II](https://leetcode.cn/problems/minimum-cost-to-convert-string-ii/) 2696

#### 二进制枚举

- [78. 子集](https://leetcode.cn/problems/subsets/)
- [77. 组合](https://leetcode.cn/problems/combinations/)
- [1286. 字母组合迭代器](https://leetcode.cn/problems/iterator-for-combination/) 1591
- [2397. 被列覆盖的最多行数](https://leetcode.cn/problems/maximum-rows-covered-by-columns/) 1719
- [2212. 射箭比赛中的最大得分](https://leetcode.cn/problems/maximum-points-in-an-archery-competition/) 1869
- [2151. 基于陈述统计最多好人数](https://leetcode.cn/problems/maximum-good-people-based-on-statements/) 1980
- [1601. 最多可达成的换楼请求数目](https://leetcode.cn/problems/maximum-number-of-achievable-transfer-requests/) 2119
- [1617. 统计子树中城市之间最大距离](https://leetcode.cn/problems/count-subtrees-with-max-distance-between-cities/) 2309
- [320. 列举单词的全部缩写](https://leetcode.cn/problems/generalized-abbreviation/)（会员题）

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
