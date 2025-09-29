请先完成上一题 [3699. 锯齿形数组的总数 I](https://leetcode.cn/problems/number-of-zigzag-arrays-i/)，以及一些稍微简单点的矩阵快速幂优化 DP，例如 [935. 骑士拨号器](https://leetcode.cn/problems/knight-dialer/)，[我的题解](https://leetcode.cn/problems/knight-dialer/solutions/3004116/jiao-ni-yi-bu-bu-si-kao-dpcong-ji-yi-hua-x06l/)。

回顾一下，状态定义为：

- 定义 $f_0[i][j]$ 表示在第 $i$ 个数为 $j$，且第 $i-1$ 个数和第 $i$ 个数是**递增**的情况下，包含 $i$ 个数的锯齿形数组个数。
- 定义 $f_1[i][j]$ 表示在第 $i$ 个数为 $j$，且第 $i-1$ 个数和第 $i$ 个数是**递减**的情况下，包含 $i$ 个数的锯齿形数组个数。

初始值 $f_0[1][j] = f_1[1][j] = 1$。

答案为 $f_0[n]$ 和 $f_1[n]$ 的元素和。

状态转移方程为

$$
\begin{aligned}
f_0[i][j] &= \sum_{k=0}^{j-1} f_1[i-1][k]   \\
f_1[i][j] &= \sum_{k=j+1}^{r-l} f_0[i-1][k] \\
\end{aligned}
$$

为方便把状态转移方程改成矩阵乘法的形式，首先把两个数组合并成一个长为 $2k$ 的数组，其中 $k=r-l+1$。即定义 $f[i][j] = f_0[i][j]$，$f[k+i][j] = f_1[i][j]$。

相应地，系数矩阵是一个大小为 $2k\times 2k$ 的矩阵

$$
M = \begin{bmatrix}
\mathbf{0} & \mathbf{A} \\
\mathbf{B} & \mathbf{0} \\
\end{bmatrix}
$$

其中：

- $\mathbf{0}$ 是 $k\times k$ 的零矩阵。
- $\mathbf{A}$ 是 $k\times k$ 的严格下三角全 $1$ 矩阵，对应从 $f_1$ 到 $f_0$ 的转移系数。
- $\mathbf{B}$ 是 $k\times k$ 的严格上三角全 $1$ 矩阵，对应从 $f_0$ 到 $f_1$ 的转移系数。

例如 $k=3$ 时，系数矩阵为

$$
M = \begin{bmatrix}
0 & 0 & 0 & 0 & 0 & 0  \\
0 & 0 & 0 & 1 & 0 & 0 \\
0 & 0 & 0 & 1 & 1 & 0 \\
0 & 1 & 1 & 0 & 0 & 0 \\
0 & 0 & 1 & 0 & 0 & 0 \\
0 & 0 & 0 & 0 & 0 & 0 \\
\end{bmatrix}
$$

转移方程改写成矩阵乘法

$$
\begin{bmatrix}
f[i][0] \\
f[i][1] \\
\vdots \\
f[i][2k-1] \\
\end{bmatrix}
= \begin{bmatrix}
\mathbf{0} & \mathbf{A} \\
\mathbf{B} & \mathbf{0} \\
\end{bmatrix}
\begin{bmatrix}
f[i-1][0] \\
f[i-1][1] \\
\vdots \\
f[i-1][2k-1] \\
\end{bmatrix}
$$

把上式中的三个矩阵分别记作 $F[i],M,F[i-1]$，即

$$
F[i] = M\times F[i-1]
$$

那么有

$$
\begin{aligned}
F[n] &= M\times F[n-1]      \\
&= M\times M\times F[n-2]        \\
&= M\times M\times M\times  F[n-3]        \\
&\ \ \vdots  \\
&= M^{n-1}\times F[1] \\
\end{aligned}
$$

$M^{n-1}$ 可以用**快速幂**计算，原理请看[【图解】一张图秒懂快速幂](https://leetcode.cn/problems/powx-n/solution/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/)。

初始值

$$
F[1] =
\begin{bmatrix}
1 \\
1 \\
\vdots \\
1 \\
\end{bmatrix}
$$

答案为矩阵 $F[n]$ 的元素和。

[本题视频讲解](https://www.bilibili.com/video/BV156n9z7E9o/?t=33m34s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
MOD = 1_000_000_007

# a @ b，其中 @ 是矩阵乘法
# 更快的写法见另一份代码【NumPy】
def mul(a: List[List[int]], b: List[List[int]]) -> List[List[int]]:
    return [[sum(x * y for x, y in zip(row, col)) % MOD for col in zip(*b)]
            for row in a]

# a^n @ f1
def pow_mul(a: List[List[int]], n: int, f1: List[List[int]]) -> List[List[int]]:
    res = f1
    while n:
        if n & 1:
            res = mul(a, res)
        a = mul(a, a)
        n >>= 1
    return res

class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        k = r - l + 1
        m = [[0] * (k * 2) for _ in range(k * 2)]
        for i in range(k):
            for j in range(i):
                m[i][k + j] = 1
            for j in range(i + 1, k):
                m[k + i][j] = 1

        f1 = [[1] for _ in range(k * 2)]
        fn = pow_mul(m, n - 1, f1)
        return sum(row[0] for row in fn) % MOD
```

```py [sol-NumPy]
import numpy as np

MOD = 1_000_000_007

# a^n @ f1
def pow_mul(a: np.ndarray, n: int, f1: np.ndarray) -> np.ndarray:
    res = f1
    while n:
        if n & 1:
            res = a @ res % MOD
        a = a @ a % MOD
        n >>= 1
    return res

class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        k = r - l + 1
        m = np.zeros((k * 2, k * 2), dtype=object)
        for i in range(k):
            for j in range(i):
                m[i][k + j] = 1
            for j in range(i + 1, k):
                m[k + i][j] = 1

        f1 = np.ones((k * 2, 1), dtype=object)
        fn = pow_mul(m, n - 1, f1)
        return fn.sum() % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int zigZagArrays(int n, int l, int r) {
        int k = r - l + 1;
        long[][] m = new long[k * 2][k * 2];
        for (int i = 0; i < k; i++) {
            for (int j = 0; j < i; j++) {
                m[i][k + j] = 1;
            }
            for (int j = i + 1; j < k; j++) {
                m[k + i][j] = 1;
            }
        }

        long[][] f1 = new long[k * 2][1];
        for (int i = 0; i < k * 2; i++) {
            f1[i][0] = 1;
        }

        long[][] fn = powMul(m, n - 1, f1);

        long ans = 0;
        for (long[] row : fn) {
            ans += row[0];
        }
        return (int) (ans % MOD);
    }

    // a^n * f0
    private long[][] powMul(long[][] a, int n, long[][] f0) {
        long[][] res = f0;
        while (n > 0) {
            if ((n & 1) > 0) {
                res = mul(a, res);
            }
            a = mul(a, a);
            n >>= 1;
        }
        return res;
    }

    // 返回矩阵 a 和矩阵 b 相乘的结果
    private long[][] mul(long[][] a, long[][] b) {
        long[][] c = new long[a.length][b[0].length];
        for (int i = 0; i < a.length; i++) {
            for (int k = 0; k < a[i].length; k++) {
                if (a[i][k] == 0) {
                    continue;
                }
                for (int j = 0; j < b[k].length; j++) {
                    c[i][j] = (c[i][j] + a[i][k] * b[k][j]) % MOD;
                }
            }
        }
        return c;
    }
}
```

```cpp [sol-C++]
constexpr int MOD = 1'000'000'007;

using matrix = vector<vector<long long>>;

// 返回矩阵 a 和矩阵 b 相乘的结果
matrix mul(matrix& a, matrix& b) {
    int n = a.size(), m = b[0].size();
    matrix c = matrix(n, vector<long long>(m));
    for (int i = 0; i < n; i++) {
        for (int k = 0; k < a[i].size(); k++) {
            if (a[i][k] == 0) {
                continue;
            }
            for (int j = 0; j < m; j++) {
                c[i][j] = (c[i][j] + a[i][k] * b[k][j]) % MOD;
            }
        }
    }
    return c;
}

// a^n * f1
matrix pow_mul(matrix a, int n, matrix& f1) {
    matrix res = f1;
    while (n) {
        if (n & 1) {
            res = mul(a, res);
        }
        a = mul(a, a);
        n >>= 1;
    }
    return res;
}

class Solution {
public:
    int zigZagArrays(int n, int l, int r) {
        int k = r - l + 1;
        matrix m(k * 2, vector<long long>(k * 2));
        for (int i = 0; i < k; i++) {
            for (int j = 0; j < i; j++) {
                m[i][k + j] = 1;
            }
            for (int j = i + 1; j < k; j++) {
                m[k + i][j] = 1;
            }
        }

        matrix f1(k * 2, vector<long long>(1, 1));
        matrix fn = pow_mul(m, n - 1, f1);

        long long ans = 0;
        for (auto& row : fn) {
            ans += row[0];
        }
        return ans % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

type matrix [][]int

func newMatrix(n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

// 返回 a*b
func (a matrix) mul(b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod
			}
		}
	}
	return c
}

// 返回 a^n * f1
func (a matrix) powMul(n int, f1 matrix) matrix {
	res := f1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func zigZagArrays(n, l, r int) (ans int) {
	k := r - l + 1
	m := newMatrix(k*2, k*2)
	for i := range k {
		for j := range i {
			m[i][k+j] = 1
		}
		for j := i + 1; j < k; j++ {
			m[k+i][j] = 1
		}
	}

	f1 := newMatrix(k*2, 1)
	for i := range f1 {
		f1[i][0] = 1
	}

	fn := m.powMul(n-1, f1)
	for _, row := range fn {
		ans += row[0]
	}
	return ans % mod
}
```

## 利用对称性优化

和 [3699. 锯齿形数组的总数 I](https://leetcode.cn/problems/number-of-zigzag-arrays-i/) 一样，可以用对称性优化。

根据 [我的题解](https://leetcode.cn/problems/number-of-zigzag-arrays-i/solutions/3794081/qian-zhui-he-you-hua-dppythonjavacgo-by-k4ps3/) 中的「另一种写法」，我们可以只使用 $k\times k$ 大小的系数矩阵，这个矩阵在反对角线的左上全为 $1$。

```py [sol-Python3]
MOD = 1_000_000_007

# a @ b，其中 @ 是矩阵乘法
# 更快的写法见另一份代码【NumPy】
def mul(a: List[List[int]], b: List[List[int]]) -> List[List[int]]:
    return [[sum(x * y for x, y in zip(row, col)) % MOD for col in zip(*b)]
            for row in a]

# a^n @ f1
def pow_mul(a: List[List[int]], n: int, f1: List[List[int]]) -> List[List[int]]:
    res = f1
    while n:
        if n & 1:
            res = mul(a, res)
        a = mul(a, a)
        n >>= 1
    return res

class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        k = r - l + 1
        m = [[0] * k for _ in range(k)]
        for i in range(k):
            for j in range(k - 1 - i):
                m[i][j] = 1

        f1 = [[1] for _ in range(k)]
        fn = pow_mul(m, n - 1, f1)
        return sum(row[0] for row in fn) * 2 % MOD
```

```py [sol-NumPy]
import numpy as np

MOD = 1_000_000_007

# a^n @ f1
def pow_mul(a: np.ndarray, n: int, f1: np.ndarray) -> np.ndarray:
    res = f1
    while n:
        if n & 1:
            res = a @ res % MOD
        a = a @ a % MOD
        n >>= 1
    return res

class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        k = r - l + 1
        m = np.zeros((k, k), dtype=object)
        for i in range(k):
            for j in range(k - 1 - i):
                m[i][j] = 1

        f1 = np.ones((k, 1), dtype=object)
        fn = pow_mul(m, n - 1, f1)
        return fn.sum() * 2 % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int zigZagArrays(int n, int l, int r) {
        int k = r - l + 1;
        long[][] m = new long[k][k];
        for (int i = 0; i < k; i++) {
            for (int j = 0; j < k - 1 - i; j++) {
                m[i][j] = 1;
            }
        }

        long[][] f1 = new long[k][1];
        for (int i = 0; i < k; i++) {
            f1[i][0] = 1;
        }

        long[][] fn = powMul(m, n - 1, f1);

        long ans = 0;
        for (long[] row : fn) {
            ans += row[0];
        }
        return (int) (ans * 2 % MOD);
    }

    // a^n * f0
    private long[][] powMul(long[][] a, int n, long[][] f0) {
        long[][] res = f0;
        while (n > 0) {
            if ((n & 1) > 0) {
                res = mul(a, res);
            }
            a = mul(a, a);
            n >>= 1;
        }
        return res;
    }

    // 返回矩阵 a 和矩阵 b 相乘的结果
    private long[][] mul(long[][] a, long[][] b) {
        long[][] c = new long[a.length][b[0].length];
        for (int i = 0; i < a.length; i++) {
            for (int k = 0; k < a[i].length; k++) {
                if (a[i][k] == 0) {
                    continue;
                }
                for (int j = 0; j < b[k].length; j++) {
                    c[i][j] = (c[i][j] + a[i][k] * b[k][j]) % MOD;
                }
            }
        }
        return c;
    }
}
```

```cpp [sol-C++]
constexpr int MOD = 1'000'000'007;

using matrix = vector<vector<long long>>;

// 返回矩阵 a 和矩阵 b 相乘的结果
matrix mul(matrix& a, matrix& b) {
    int n = a.size(), m = b[0].size();
    matrix c = matrix(n, vector<long long>(m));
    for (int i = 0; i < n; i++) {
        for (int k = 0; k < a[i].size(); k++) {
            if (a[i][k] == 0) {
                continue;
            }
            for (int j = 0; j < m; j++) {
                c[i][j] = (c[i][j] + a[i][k] * b[k][j]) % MOD;
            }
        }
    }
    return c;
}

// a^n * f1
matrix pow_mul(matrix a, int n, matrix& f1) {
    matrix res = f1;
    while (n) {
        if (n & 1) {
            res = mul(a, res);
        }
        a = mul(a, a);
        n >>= 1;
    }
    return res;
}

class Solution {
public:
    int zigZagArrays(int n, int l, int r) {
        int k = r - l + 1;
        matrix m(k, vector<long long>(k));
        for (int i = 0; i < k; i++) {
            for (int j = 0; j < k - 1 - i; j++) {
                m[i][j] = 1;
            }
        }

        matrix f1(k, vector<long long>(1, 1));
        matrix fn = pow_mul(m, n - 1, f1);

        long long ans = 0;
        for (auto& row : fn) {
            ans += row[0];
        }
        return ans * 2 % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

type matrix [][]int

func newMatrix(n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

// 返回 a*b
func (a matrix) mul(b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod
			}
		}
	}
	return c
}

// 返回 a^n * f1
func (a matrix) powMul(n int, f1 matrix) matrix {
	res := f1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func zigZagArrays(n, l, r int) (ans int) {
	k := r - l + 1
	m := newMatrix(k, k)
	for i := range k {
		for j := range k - 1 - i {
			m[i][j] = 1
		}
	}

	f1 := newMatrix(k, 1)
	for i := range f1 {
		f1[i][0] = 1
	}

	fn := m.powMul(n-1, f1)
	for _, row := range fn {
		ans += row[0]
	}
	return ans * 2 % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((r-l)^3\log n)$。
- 空间复杂度：$\mathcal{O}((r-l)^2)$。

## 专题训练

见下面动态规划题单的「**§11.6 矩阵快速幂优化 DP**」。

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
