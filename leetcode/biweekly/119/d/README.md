## 方法一：二进制枚举

根据 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/) 中「枚举集合」的技巧，我们可以枚举保留哪些节点，在这些节点之间连边。

连边后，根据[【模板讲解】带你发明 Floyd 算法：从记忆化搜索到递推](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/solution/dai-ni-fa-ming-floyd-suan-fa-cong-ji-yi-m8s51/) 求出任意两点之间最短路，如果保留的节点之间的最短路均不超过 $\textit{maxDistance}$，则把答案加一。

代码实现时，可以先预处理原图的邻接矩阵 $g$，这样无需每次枚举子集都要遍历 $\textit{roads}$，直接拷贝一份 $g$ 即可。

```py [sol-Python3]
class Solution:
    def numberOfSets(self, n: int, maxDistance: int, roads: List[List[int]]) -> int:
        g = [[inf] * n for _ in range(n)]
        for i in range(n):
            g[i][i] = 0
        for x, y, wt in roads:
            g[x][y] = min(g[x][y], wt)
            g[y][x] = min(g[y][x], wt)

        f = [None] * n
        def check(s: int) -> int:
            for i, row in enumerate(g):
                if s >> i & 1:
                    f[i] = row[:]

            # Floyd
            for k in range(n):
                if (s >> k & 1) == 0: continue
                for i in range(n):
                    if (s >> i & 1) == 0: continue
                    for j in range(n):
                        f[i][j] = min(f[i][j], f[i][k] + f[k][j])

            for i, di in enumerate(f):
                if (s >> i & 1) == 0: continue
                for j, dij in enumerate(di):
                    if s >> j & 1 and dij > maxDistance:
                        return 0
            return 1
        return sum(check(s) for s in range(1 << n))  # 枚举子集
```

```java [sol-Java]
class Solution {
    public int numberOfSets(int n, int maxDistance, int[][] roads) {
        int[][] g = new int[n][n];
        for (int i = 0; i < n; i++) {
            Arrays.fill(g[i], Integer.MAX_VALUE / 2); // 防止加法溢出
            g[i][i] = 0;
        }
        for (int[] e : roads) {
            int x = e[0], y = e[1], wt = e[2];
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

            // Floyd
            for (int k = 0; k < n; k++) {
                if ((s >> k & 1) == 0) continue;
                for (int i = 0; i < n; i++) {
                    if ((s >> i & 1) == 0) continue;
                    for (int j = 0; j < n; j++) {
                        f[i][j] = Math.min(f[i][j], f[i][k] + f[k][j]);
                    }
                }
            }

            for (int i = 0; i < n; i++) {
                if ((s >> i & 1) == 0) continue;
                for (int j = 0; j < n; j++) {
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
    int numberOfSets(int n, int maxDistance, vector<vector<int>> &roads) {
        vector<vector<int>> g(n, vector<int>(n, INT_MAX / 2)); // 防止加法溢出
        for (int i = 0; i < n; i++) {
            g[i][i] = 0;
        }
        for (auto &e: roads) {
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

            // Floyd
            for (int k = 0; k < n; k++) {
                if (((s >> k) & 1) == 0) continue;
                for (int i = 0; i < n; i++) {
                    if (((s >> i) & 1) == 0) continue;
                    for (int j = 0; j < n; j++) {
                        f[i][j] = min(f[i][j], f[i][k] + f[k][j]);
                    }
                }
            }

            for (int i = 0; i < n; i++) {
                if (((s >> i) & 1) == 0) continue;
                for (int j = 0; j < n; j++) {
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
			if j != i { // g[i][i] = 0
				g[i][j] = math.MaxInt / 2 // 防止加法溢出
			}
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

		// Floyd
		for k := range f {
			if s>>k&1 == 0 { continue }
			for i := range f {
				if s>>i&1 == 0 { continue }
				for j := range f {
					f[i][j] = min(f[i][j], f[i][k]+f[k][j])
				}
			}
		}

		for i, di := range f {
			if s>>i&1 == 0 { continue }
			for j, dij := range di {
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

## 方法二：进一步优化

假设现在算出了集合 $S$ 中任意点对的最短路，此时再添加一个点 $k$，我们只需要写一个 $\mathcal{O}(n^2)$ 的二重循环，枚举 $i$ 和 $j$，用 $f[i][k] + f[k][j]$ 更新 $f[i][j]$ 的最小值，就得到了集合 $S\cup \{k\}$ 中任意点对的最短路。

这意味着我们只需要 $\mathcal{O}(n^2)$ 的时间就可以算出一个集合任意点对的最短路，而不是方法一中的 $\mathcal{O}(n^3)$。



```py [sol-Python3]

```

```java [sol-Java]

```

```cpp [sol-C++]

```

```go [sol-Go]

```

## 相关题目

#### Floyd

- [2642. 设计可以求最短路径的图类](https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/) 1811
- [1334. 阈值距离内邻居最少的城市](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/) 1855
- [2101. 引爆最多的炸弹](https://leetcode.cn/problems/detonate-the-maximum-bombs/) 1880

#### 二进制枚举

- [78. 子集](https://leetcode.cn/problems/subsets/)
- [77. 组合](https://leetcode.cn/problems/combinations/)
- [1286. 字母组合迭代器](https://leetcode.cn/problems/iterator-for-combination/) 1591
- [2397. 被列覆盖的最多行数](https://leetcode.cn/problems/maximum-rows-covered-by-columns/) 1719
- [2212. 射箭比赛中的最大得分](https://leetcode.cn/problems/maximum-points-in-an-archery-competition/) 1869
- [1601. 最多可达成的换楼请求数目](https://leetcode.cn/problems/maximum-number-of-achievable-transfer-requests/) 2119
- [320. 列举单词的全部缩写](https://leetcode.cn/problems/generalized-abbreviation/)（会员题）
