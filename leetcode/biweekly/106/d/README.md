由于矩阵元素值只有 $0$ 和 $1$，对于矩阵的每一行，把这一行看成一个二进制数。

- 如果答案只有 $1$ 行，根据题目要求，每一列的和至多为 $\left\lfloor 1/2 \right\rfloor = 0$，也就是这一行必须全为 $0$。
- 如果答案有 $2$ 行，每一列的和至多为 $\left\lfloor 2/2 \right\rfloor = 1$，所以同一列至多有一个 $1$，不能有两个 $1$。从二进制角度来理解，就是这两行对应二进制数的 AND 等于 $0$。
- 如果答案有 $3$ 行，每一列的和至多为 $\left\lfloor 3/2 \right\rfloor = 1$，这和 $2$ 行的情况是一样的。如果可以选 $3$ 行，那么必然也可以选 $2$ 行，所以无需考虑答案有 $3$ 行的情况。
- 如果答案有 $4$ 行，下面细说。

⚠**注意**：以下讨论的前提是，不存在小于 $4$ 行的答案。

如果答案有 $4$ 行，则有如下性质：

**性质一**：每一列的和至多为 $\left\lfloor 4/2 \right\rfloor = 2$。

**性质二**：**任意** $2$ 行的 AND 均不为 $0$（否则答案可以是 $2$ 行）。任意是什么意思？第一行和第二行的 AND 不为 $0$、第一行和第三行的 AND 不为 $0$、第一行和第四行的 AND 不为 $0$、第二行和第三行的 AND 不为 $0$、第二行和第四行的 AND 不为 $0$、第三行和第四行的 AND 不为 $0$。一共有 $C(4,2)=6$ 个约束。
   
- 第一行和第二行的 AND 不为 $0$，根据性质二，其中一列（不妨设为第一列）的第一行和第二行都是 $1$，结合性质一，这一列的第三行和第四行都是 $0$。所以这一列从上到下是 $1100$。
- 第一行和第三行的 AND 不为 $0$，根据性质二，其中一列（不妨设为第二列）的第一行和第三行都是 $1$，结合性质一，这一列的第二行和第四行都是 $0$。所以这一列从上到下是 $1010$。
- ……
- 依此类推，有 $C(4,2)=6$ 种选择两行的方式，且**包含两个** $1$ **的列互不相同**。如果列数 $n$ 可以等于 $6$，则存在如下合法构造：

$$
111000\\
100110\\
010101\\
001011
$$

但本题 $n$ 至多为 $5$，无法满足。比如只考虑前 $5$ 列，那么在上述构造中，只需选择第三行和第四行，因为这两行的 AND 为 $0$。

继续讨论，如果答案有 $k$ 行（$k>4$）且不存在小于 $k$ 行的答案，类似地，有如下性质：

**性质一**：每一列的和至多为 $\left\lfloor k/2 \right\rfloor$。

**性质二**：任意 $k-2$ 行，必定存在一列，其元素和大于 $\left\lfloor (k-2)/2 \right\rfloor$，即大于等于 $\left\lfloor k/2 \right\rfloor$。

结合这两个性质，从 $k$ 行中任选 $k-2$ 行，一定存在一列，其元素和恰好等于 $\left\lfloor k/2 \right\rfloor$。

根据前文 $k=4$ 的讨论，我们需要至少 $C(k,k-2)$ 个互不相同的列。

但当 $k\ge 4$ 时，有

$$
C(k,k-2)=C(k,2)\ge C(4,2)=6 > n
$$

无法满足。

综上所述，在 $n\le 5$ 的数据范围下，只需考虑答案为 $1$ 行或者 $2$ 行的情况（$3$ 行的情况转换成 $2$ 行），如果不存在 $1$ 行和 $2$ 行的答案，则无解。

## 方法一

1. 遍历每一行，从左到右，算出一个长为 $n$ 的二进制数。
2. 由于至多有 $2^n\le 32$ 个不同的二进制数，而行数 $m\le 10^4$ 远大于 $32$，所以可以把二进制数去重，保存到一个哈希表 $\textit{maskToIdx}$ 中，key 为二进制数，value 为行号。
3. 如果有一行全为 $0$，返回这一行的行号。
4. 否则，写一个二重循环，枚举从 $\textit{maskToIdx}$ 中选两个数的所有组合，如果有两个数的 AND 等于 $0$，返回对应的行号。注意题目要求按**升序**返回。
5. 如果无解，返回空数组。

```py [sol-Python3]
class Solution:
    def goodSubsetofBinaryMatrix(self, grid: List[List[int]]) -> List[int]:
        mask_to_idx = {}
        for i, row in enumerate(grid):
            mask = 0
            for j, x in enumerate(row):
                mask |= x << j
            if mask == 0:
                return [i]
            mask_to_idx[mask] = i

        for x, i in mask_to_idx.items():
            for y, j in mask_to_idx.items():
                if (x & y) == 0:
                    return sorted((i, j))
        return []
```

```java [sol-Java]
class Solution {
    public List<Integer> goodSubsetofBinaryMatrix(int[][] grid) {
        Map<Integer, Integer> maskToIdx = new HashMap<>();
        for (int i = 0; i < grid.length; i++) {
            int mask = 0;
            for (int j = 0; j < grid[i].length; j++) {
                mask |= grid[i][j] << j;
            }
            if (mask == 0) {
                return List.of(i);
            }
            maskToIdx.put(mask, i);
        }

        for (Map.Entry<Integer, Integer> e1 : maskToIdx.entrySet()) {
            for (Map.Entry<Integer, Integer> e2 : maskToIdx.entrySet()) {
                if ((e1.getKey() & e2.getKey()) == 0) {
                    int i = e1.getValue();
                    int j = e2.getValue();
                    return i < j ? List.of(i, j) : List.of(j, i);
                }
            }
        }
        return List.of();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> goodSubsetofBinaryMatrix(vector<vector<int>>& grid) {
        unordered_map<int, int> mask_to_idx;
        for (int i = 0; i < grid.size(); i++) {
            int mask = 0;
            for (int j = 0; j < grid[i].size(); j++) {
                mask |= grid[i][j] << j;
            }
            if (mask == 0) {
                return {i};
            }
            mask_to_idx[mask] = i;
        }

        for (auto [x, i] : mask_to_idx) {
            for (auto [y, j] : mask_to_idx) {
                if ((x & y) == 0) {
                    return {min(i, j), max(i, j)};
                }
            }
        }
        return {};
    }
};
```

```go [sol-Go]
func goodSubsetofBinaryMatrix(grid [][]int) []int {
	maskToIdx := map[int]int{}
	for i, row := range grid {
		mask := 0
		for j, x := range row {
			mask |= x << j
		}
		if mask == 0 {
			return []int{i}
		}
		maskToIdx[mask] = i
	}

	for x, i := range maskToIdx {
		for y, j := range maskToIdx {
			if x&y == 0 {
				return []int{min(i, j), max(i, j)}
			}
		}
	}
	return nil
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn+4^n)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(2^n)$。至多有 $2^n$ 个不同的二进制数。

## 方法二

用长为 $2^n$ 的数组代替哈希表，数组元素初始化成 $-1$。

对于二重循环，由于 $x$ 和 $y$ 没有交集，可以直接枚举 $x$ 的补集的非空子集作为 $y$。

如何枚举一个集合的子集？请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def goodSubsetofBinaryMatrix(self, grid: List[List[int]]) -> List[int]:
        n = len(grid[0])
        mask_to_idx = [-1] * (1 << n)
        for i, row in enumerate(grid):
            mask = 0
            for j, x in enumerate(row):
                mask |= x << j
            if mask == 0:
                return [i]
            mask_to_idx[mask] = i

        u = (1 << n) - 1
        for x, i in enumerate(mask_to_idx):
            if i < 0:
                continue
            y = c = u ^ x
            while y:
                j = mask_to_idx[y]
                if j >= 0:
                    return sorted((i, j))
                y = (y - 1) & c
        return []
```

```java [sol-Java]
class Solution {
    public List<Integer> goodSubsetofBinaryMatrix(int[][] grid) {
        int n = grid[0].length;
        int[] maskToIdx = new int[1 << n];
        Arrays.fill(maskToIdx, -1);
        for (int i = 0; i < grid.length; i++) {
            int mask = 0;
            for (int j = 0; j < n; j++) {
                mask |= grid[i][j] << j;
            }
            if (mask == 0) {
                return List.of(i);
            }
            maskToIdx[mask] = i;
        }

        int u = (1 << n) - 1;
        for (int x = 1; x < 1 << n; x++) {
            int i = maskToIdx[x];
            if (i < 0) {
                continue;
            }
            int c = u ^ x;
            for (int y = c; y > 0; y = (y - 1) & c) {
                int j = maskToIdx[y];
                if (j >= 0) {
                    return i < j ? List.of(i, j) : List.of(j, i);
                }
            }
        }
        return List.of();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> goodSubsetofBinaryMatrix(vector<vector<int>>& grid) {
        int n = grid[0].size();
        vector<int> mask_to_idx(1 << n, -1);
        for (int i = 0; i < grid.size(); i++) {
            int mask = 0;
            for (int j = 0; j < n; j++) {
                mask |= grid[i][j] << j;
            }
            if (mask == 0) {
                return {i};
            }
            mask_to_idx[mask] = i;
        }

        int u = (1 << n) - 1;
        for (int x = 1; x < 1 << n; x++) {
            int i = mask_to_idx[x];
            if (i < 0) continue;
            int c = u ^ x;
            for (int y = c; y; y = (y - 1) & c) {
                int j = mask_to_idx[y];
                if (j >= 0) {
                    return {min(i, j), max(i, j)};
                }
            }
        }
        return {};
    }
};
```

```go [sol-Go]
func goodSubsetofBinaryMatrix(grid [][]int) []int {
	n := len(grid[0])
	maskToIdx := make([]int, 1<<n)
	for i := range maskToIdx {
		maskToIdx[i] = -1
	}
	for i, row := range grid {
		mask := 0
		for j, x := range row {
			mask |= x << j
		}
		if mask == 0 {
			return []int{i}
		}
		maskToIdx[mask] = i
	}

	u := 1<<n - 1
	for x, i := range maskToIdx {
		if i < 0 {
			continue
		}
		c := u ^ x
		for y := c; y > 0; y = (y - 1) & c {
			j := maskToIdx[y]
			if j >= 0 {
				return []int{min(i, j), max(i, j)}
			}
		}
	}
	return nil
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn+3^n)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。由于元素个数为 $k$ 的集合有 $C(n,k)$ 个，其子集有 $2^k$ 个，根据二项式定理，$\sum\limits_{k=0}^n C(n,k)2^k = (2+1)^n = 3^n$，所以二重循环的时间复杂度为 $O(3^n)$。
- 空间复杂度：$\mathcal{O}(2^n)$。

## 方法三：SOSDP（选读）

⚠**提醒**：该方法为竞赛算法，阅读前最好有一些状压 DP 的经验。

回顾方法二，相当于寻找一个**子集的子集** $Y$，满足 $\textit{maskToIdx}[Y]\ge 0$。

这可以用 **SOSDP**（Sum over Subsets Dynamic Programming）更快地计算出来。

设全集 $U=\{0,1,2,\cdots,n-1\}$。

设 $S$ 为 $U$ 的子集，$f[S]$ 定义如下：

- 如果 $S$ 不存在子集 $Y$，满足 $\textit{maskToIdx}[Y]\ge 0$，则 $f[S]=-1$。
- 如果 $S$ 存在子集 $Y$，满足 $\textit{maskToIdx}[Y]\ge 0$，则 $f[S]$ 等于任意满足要求的 $\textit{maskToIdx}[Y]$。

为方便编程，不妨取最大值，即定义

$$
f[S] = \max\limits_{Y\subseteq S} \textit{maskToIdx}[Y]
$$

先来说怎么用 $f[S]$ 计算答案。我们可以枚举 $U$ 的所有非空真子集 $S$，如果 $f[S]\ge 0$ 且 $\textit{maskToIdx}[\complement_US]\ge 0$，根据 $f$ 的定义，这意味着 $S$ 的某个子集的 $\textit{maskToIdx}$ 值和 $\complement_US$ 的 $\textit{maskToIdx}$ 值均为非负数，且这两个集合不相交，符合要求，返回答案。

然后来说怎么递推计算 $f[S]$。我们可以枚举 $S$ 中的元素 $b$，从 $S$ 中去掉 $b$，问题规模变小，这样就可以递推计算了，即

$$
f[S] = \max\limits_{b\in S} f[S\setminus \{b\}]
$$

初始值 $f[i] = \textit{maskToIdx}[i]$。

代码实现时，可以把 $\textit{maskToIdx}$ 去掉，直接在遍历 $\textit{grid}$ 的过程中初始化 $f$。

```py [sol-Python3]
class Solution:
    def goodSubsetofBinaryMatrix(self, grid: List[List[int]]) -> List[int]:
        n = len(grid[0])
        f = [-1] * (1 << n)
        for i, row in enumerate(grid):
            mask = 0
            for j, x in enumerate(row):
                mask |= x << j
            if mask == 0:
                return [i]
            f[mask] = i

        u = (1 << n) - 1
        for s in range(1, u):
            for b in range(n):
                if (s >> b & 1) == 0:
                    continue
                i = f[s] = max(f[s], f[s ^ (1 << b)])
                if i < 0:
                    continue
                j = f[u ^ s]
                if j >= 0:
                    return sorted((i, j))
        return []
```

```java [sol-Java]
class Solution {
    public List<Integer> goodSubsetofBinaryMatrix(int[][] grid) {
        int n = grid[0].length;
        int[] f = new int[1 << n];
        Arrays.fill(f, -1);
        for (int i = 0; i < grid.length; i++) {
            int mask = 0;
            for (int j = 0; j < n; j++) {
                mask |= grid[i][j] << j;
            }
            if (mask == 0) {
                return List.of(i);
            }
            f[mask] = i;
        }

        int u = (1 << n) - 1;
        for (int s = 1; s < u; s++) {
            for (int b = 0; b < n; b++) {
                if ((s >> b & 1) == 0) {
                    continue;
                }
                f[s] = Math.max(f[s], f[s ^ (1 << b)]);
                int i = f[s];
                if (i < 0) {
                    continue;
                }
                int j = f[u ^ s];
                if (j >= 0) {
                    return i < j ? List.of(i, j) : List.of(j, i);
                }
            }
        }
        return List.of();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> goodSubsetofBinaryMatrix(vector<vector<int>>& grid) {
        int n = grid[0].size();
        vector<int> f(1 << n, -1);
        for (int i = 0; i < grid.size(); i++) {
            int mask = 0;
            for (int j = 0; j < n; j++) {
                mask |= grid[i][j] << j;
            }
            if (mask == 0) {
                return {i};
            }
            f[mask] = i;
        }

        int u = (1 << n) - 1;
        for (int s = 1; s < u; s++) {
            for (int b = 0; b < n; b++) {
                if ((s >> b & 1) == 0) {
                    continue;
                }
                f[s] = max(f[s], f[s ^ (1 << b)]);
                int i = f[s];
                if (i < 0) {
                    continue;
                }
                int j = f[u ^ s];
                if (j >= 0) {
                    return {min(i, j), max(i, j)};
                }
            }
        }
        return {};
    }
};
```

```go [sol-Go]
func goodSubsetofBinaryMatrix(grid [][]int) []int {
	n := len(grid[0])
	f := make([]int, 1<<n)
	for i := range f {
		f[i] = -1
	}
	for i, row := range grid {
		mask := 0
		for j, x := range row {
			mask |= x << j
		}
		if mask == 0 {
			return []int{i}
		}
		f[mask] = i
	}

	u := 1<<n - 1
	for s := 1; s < u; s++ {
		for b := 0; b < n; b++ {
			if s>>b&1 == 0 {
				continue
			}
			f[s] = max(f[s], f[s^1<<b])
			i := f[s]
			if i < 0 {
				continue
			}
			j := f[u^s]
			if j >= 0 {
				return []int{min(i, j), max(i, j)}
			}
		}
	}
	return nil
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn+n2^n)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(2^n)$。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
