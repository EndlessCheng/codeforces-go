## 预备知识

为方便描述，对于每一行，把这一行的 $1$ 的**列号**保存到集合中。例如这一行是 $[1,0,1,1]$，其中 $1$ 的列号集合为 $\{0,2,3\}$，等价于二进制数 $1101_{(2)}$。

两个集合有交集，等价于对应的二进制数的 AND 不为 $0$。

关于集合与位运算的知识点，请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

## 分类讨论

1. 如果答案只有 $1$ 行，根据题目要求，每一列的和至多为 $\left\lfloor 1/2 \right\rfloor = 0$，也就是这一行必须全为 $0$。
2. 如果答案有 $2$ 行，每一列的和至多为 $\left\lfloor 2/2 \right\rfloor = 1$，所以同一列至多有一个 $1$，不能有两个 $1$。从二进制角度来理解，就是这两行 AND 的结果等于 $0$。
3. 如果答案有 $3$ 行，每一列的和至多为 $\left\lfloor 3/2 \right\rfloor = 1$，这和 $2$ 行的情况是一样的。如果可以选 $3$ 行，那么必然也可以选 $2$ 行，所以无需考虑答案有 $3$ 行的情况。
4. **假定上面的情况都不存在答案**。如果答案有 $4$ 行，每一列的和至多为 $\left\lfloor 4/2 \right\rfloor = 2$，且任意两行的 AND 均不为 $0$（否则答案可以是 $2$ 行）。不妨设第一行的 $1$ 的个数最少，继续分类讨论。请大家拿出纸笔，当成一个类似**数独**的游戏来玩，考虑其他行怎么填数字（假定列数 $n=5$）：
   1. 如果第一行是 $10000$，由于任意两行的 AND 均不为 $0$，其他行的第一列必须填 $1$。但这样的话，第一列的数字和等于 $4$，不符合要求。
   2. 如果第一行是 $11000$，其他行前两列至少要有一个 $1$，那么第二行可以是 $10\texttt{\_\_\_}$，第三行可以是 $01\texttt{\_\_\_}$，但第四行无论怎么填，都会导致有一列的和超过 $2$，不符合要求。
   3. 对于第一行至少有 $3$ 个 $1$ 的情况，由于第一行的 $1$ 的个数最少，所以这 $4$ 行一共有至少 $3\cdot 4=12$ 个 $1$。但同时，由于每列至多允许有 $2$ 个 $1$，总共至多允许有 $2n=2\cdot 5 = 10$ 个 $1$，由于 $12>10$，所以不满足要求。
5. 如果答案超过 $4$ 行，类似上面的方法，可以证明答案是不存在的。

因此，**答案至多两行**。

## 算法

1. 把每一行的 $1$ 的**列号**，保存到一个二进制数中。
2. 由于至多有 $2^n\le 32$ 个不同的二进制数，而行数 $m\le 10^4$ 远大于 $32$，所以可以把二进制数去重，保存到一个哈希表 $\textit{maskToIdx}$ 中，key 为二进制数，value 为行号。
3. 如果有一行全为 $0$，返回这一行的行号。
4. 否则，写一个二重循环，枚举从 $\textit{maskToIdx}$ 中选两个 key 的所有组合，如果有两行的二进制数的 AND 的结果等于 $0$，返回这两行的行号。

## 注

本题 $n$ 至多为 $5$，而当 $n=6$ 时，存在如下合法构造：

$$
111000\\
100110\\
010101\\
001011
$$

此时就要考虑 $4$ 行的情况了。

请看 [视频讲解](https://www.bilibili.com/video/BV18u411Y7Gt/) 第四题。

## 优化前

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
				if i < j {
					return []int{i, j}
				}
				return []int{j, i}
			}
		}
	}
	return nil
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn+4^n)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(2^n)$。至多有 $2^n$ 个不同的二进制数。

## 优化

用长为 $2^n$ 的数组代替哈希表，数组元素初始化成 $-1$。

对于二重循环，由于 $x$ 和 $y$ 没有交集，可以直接枚举 $x$ 的补集的非空子集作为 $y$。

原理请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

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
			if j := maskToIdx[y]; j >= 0 {
				if i < j {
					return []int{i, j}
				}
				return []int{j, i}
			}
		}
	}
	return nil
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn+3^n)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。由于元素个数为 $k$ 的集合有 $C(n,k)$ 个，其子集有 $2^k$ 个，根据二项式定理，$\sum\limits_{k=0}^n C(n,k)2^k = (2+1)^n = 3^n$，所以二重循环的时间复杂度为 $O(3^n)$。
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
