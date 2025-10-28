## 方法一：矩阵快速幂优化 DP

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

### 优化前

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
                m[i, k + j] = 1
            for j in range(i + 1, k):
                m[k + i, j] = 1

        f1 = np.ones((k * 2,), dtype=object)
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

### 利用对称性优化

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
                m[i, j] = 1

        f1 = np.ones((k,), dtype=object)
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

## 方法二：Berlekamp-Massey 算法 + Kitamasa 算法

见我的知乎科普文章：

[Berlekamp-Massey 算法：如何预测数列的下一项？](https://zhuanlan.zhihu.com/p/1966417899825665440)

[Kitamasa 算法：更快地计算线性递推的第 n 项](https://zhuanlan.zhihu.com/p/1964051212304364939)

矩阵快速幂优化 DP 的题，一般可以结合这两个算法优化。通用三步：

1. 首先，用上一题 [3699. 锯齿形数组的总数 I](https://leetcode.cn/problems/number-of-zigzag-arrays-i/) 的算法，计算 $n=2,\ldots,2k+1$ 时的答案。其中 $k=r-l+1$ 是系数矩阵的行数和列数。
2. 然后，用 Berlekamp-Massey 算法找规律，得到线性递推式。
3. 最后，根据线性递推式，用 Kitamasa 算法算出答案。

```py [sol-Python3]
MOD = 1_000_000_007

class Solution:
    # 给定数列的前 m 项 a，返回符合 a 的最短常系数齐次线性递推式的系数 coef（模 MOD 意义下）
    # 设 coef 长为 k，当 n >= k 时，有递推式 f(n) = coef[0] * f(n-1) + coef[1] * f(n-2) + ... + coef[k-1] * f(n-k)  （注意 coef 的顺序）
    # 初始值 f(n) = a[n]  (0 <= n < k)
    # 时间复杂度 O(m^2)，其中 m 是 a 的长度
    def berlekampMassey(self, a: List[int]) -> List[int]:
        pre_c = []
        pre_i, pre_d = -1, 0
        coef = []

        for i, v in enumerate(a):
            # d = a[i] - 递推式算出来的值
            d = (v - sum(c * a[i - 1 - j] for j, c in enumerate(coef))) % MOD
            if d == 0:  # 递推式正确
                continue

            # 首次算错，初始化 coef 为 i+1 个 0
            if pre_i < 0:
                coef = [0] * (i + 1)
                pre_i, pre_d = i, d
                continue

            bias = i - pre_i
            old_len = len(coef)
            new_len = bias + len(pre_c)
            if new_len > old_len:  # 递推式变长了
                tmp = coef[:]
                coef += [0] * (new_len - old_len)

            # 历史错误为 pre_d = a[pre_i] - sum_j pre_c[j]*a[pre_i-1-j]
            # 现在 a[i] = sum_j coef[j]*a[i-1-j] + d
            # 联立得 a[i] = sum_j coef[j]*a[i-1-j] + d/pre_d * (a[pre_i] - sum_j pre_c[j]*a[pre_i-1-j])
            # 其中 a[pre_i] 的系数 d/pre_d 位于当前（i）的 bias-1 = i-pre_i-1 处
            delta = d * pow(pre_d, -1, MOD) % MOD
            coef[bias - 1] = (coef[bias - 1] + delta) % MOD
            for j, c in enumerate(pre_c):
                coef[bias + j] = (coef[bias + j] - delta * c) % MOD

            if new_len > old_len:
                pre_c = tmp
                pre_i, pre_d = i, d

        # 计算完后，可能 coef 的末尾有 0，这些 0 不能去掉
        # 比如数列 (1,2,4,2,4,2,4,...) 的系数为 [0,1,0]，表示 f(n) = 0*f(n-1) + 1*f(n-2) + 0*f(n-3) = f(n-2)   (n >= 3)
        # 如果把末尾的 0 去掉，变成 [0,1]，就表示 f(n) = 0*f(n-1) + f(n-2) = f(n-2)   (n >= 2)
        # 看上去一样，但按照这个式子算出来的数列是错误的 (1,2,1,2,1,2,...)

        return coef

    # 给定常系数齐次线性递推式 f(n) = coef[k-1] * f(n-1) + coef[k-2] * f(n-2) + ... + coef[0] * f(n-k)
    # 以及初始值 f(i) = a[i] (0 <= i < k)
    # 返回 f(n) % mod，其中参数 n 从 0 开始
    # 注意 coef 的顺序
    # 时间复杂度 O(k^2 log n)，其中 k 是 coef 的长度
    def kitamasa(self, coef: List[int], a: List[int], n: int) -> int:
        if n < len(a):
            return a[n] % MOD

        k = len(coef)
        # 特判 k = 0, 1 的情况
        if k == 0:
            return 0
        if k == 1:
            return a[0] * pow(coef[0], n, MOD) % MOD

        # 已知 f(n) 的各项系数为 a，f(m) 的各项系数为 b
        # 计算并返回 f(n+m) 的各项系数 c
        def compose(a: List[int], b: List[int]) -> List[int]:
            c = [0] * k
            for v in a:
                for j, w in enumerate(b):
                    c[j] = (c[j] + v * w) % MOD
                # 原地计算下一组系数，比如已知 f(4) 的各项系数，现在要计算 f(5) 的各项系数
                # 倒序遍历，避免提前覆盖旧值
                bk1 = b[-1]
                for i in range(k - 1, 0, -1):
                    b[i] = (b[i - 1] + bk1 * coef[i]) % MOD
                b[0] = bk1 * coef[0] % MOD
            return c

        # 计算 res_c，以表出 f(n) = res_c[k-1] * a[k-1] + res_c[k-2] * a[k-2] + ... + res_c[0] * a[0]
        res_c = [0] * k
        c = [0] * k
        res_c[0] = c[1] = 1
        while n > 0:
            if n % 2:
                res_c = compose(c, res_c)
            # 由于会修改 compose 的第二个参数，这里把 c 复制一份再传入
            c = compose(c, c[:])
            n //= 2

        return sum(c * v for c, v in zip(res_c, a)) % MOD

    # 见上一题 3699. 锯齿形数组的总数 I
    def zigZagArraysInit(self, l: int, r: int) -> List[int]:
        k = r - l + 1
        f = [1] * k

        a = [0] * (k * 2)  # 注意 a 不包含 n=0 和 n=1 的项，所以下面 kitamasa 传入的是 n-2
        for i in range(k * 2):
            pre = 0
            for j, v in enumerate(f):
                f[j] = pre % MOD
                pre += v
            a[i] = sum(f) * 2 % MOD
            f.reverse()
        return a

    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        a = self.zigZagArraysInit(l, r)
        coef = self.berlekampMassey(a)
        coef.reverse()  # 注意 kitamasa 入参的顺序
        return self.kitamasa(coef, a, n - 2)
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int zigZagArrays(int n, int l, int r) {
        int[] a = zigZagArraysInit(l, r);

        List<Integer> coef = berlekampMassey(a);
        int k = coef.size();
        int[] c = new int[k];
        for (int i = 0; i < k; i++) {
            c[i] = coef.get(k - 1 - i); // 注意 kitamasa 入参的顺序
        }

        return kitamasa(c, a, n - 2);
    }

    // 见上一题 3699. 锯齿形数组的总数 I
    private int[] zigZagArraysInit(int l, int r) {
        int k = r - l + 1;
        int[] f = new int[k];
        Arrays.fill(f, 1);

        int[] a = new int[k * 2]; // 注意 a 不包含 n=0 和 n=1 的项，所以 kitamasa 传入的是 n-2
        for (int i = 0; i < a.length; i++) {
            long pre = 0;
            long s = 0;
            for (int j = 0; j < k; j++) {
                int v = f[j];
                f[j] = (int) (pre % MOD);
                pre += v;
                s += f[j];
            }
            a[i] = (int) (s * 2 % MOD);
            reverse(f);
        }
        return a;
    }

    // 给定数列的前 m 项 a，返回符合 a 的最短常系数齐次线性递推式的系数 coef（模 MOD 意义下）
    // 设 coef 长为 k，当 n >= k 时，有递推式：
    // f(n) = coef[0]*f(n-1) + coef[1]*f(n-2) + ... + coef[k-1]*f(n-k)
    // 初始值 f(n) = a[n] (0 <= n < k)
    // 时间复杂度 O(m^2)，其中 m 是 a 的长度
    private List<Integer> berlekampMassey(int[] a) {
        List<Integer> coef = new ArrayList<>();
        List<Integer> preC = new ArrayList<>();
        int preI = -1;
        int preD = 0;

        for (int i = 0; i < a.length; i++) {
            // d = a[i] - 递推式算出来的值
            long d = a[i];
            for (int j = 0; j < coef.size(); j++) {
                d = (d - (long) coef.get(j) * a[i - 1 - j]) % MOD;
            }
            if (d == 0) { // 递推式正确
                continue;
            }

            // 首次算错，初始化 coef 为 i+1 个 0
            if (preI < 0) {
                coef = new ArrayList<>(Collections.nCopies(i + 1, 0));
                preI = i;
                preD = (int) d;
                continue;
            }

            int bias = i - preI;
            int oldLen = coef.size();
            int newLen = bias + preC.size();
            List<Integer> tmp = null;
            if (newLen > oldLen) { // 递推式变长了
                tmp = new ArrayList<>(coef);
                coef.addAll(Collections.nCopies(newLen - oldLen, 0));
            }

            // 历史错误为 preD = a[preI] - sum_j preC[j]*a[preI-1-j]
            // 现在 a[i] = sum_j coef[j]*a[i-1-j] + d
            // 联立得 a[i] = sum_j coef[j]*a[i-1-j] + d/preD * (a[preI] - sum_j preC[j]*a[preI-1-j])
            // 其中 a[preI] 的系数 d/preD 位于当前（i）的 bias-1 = i-preI-1 处
            long delta = d * pow(preD, MOD - 2) % MOD;
            coef.set(bias - 1, (int) ((coef.get(bias - 1) + delta) % MOD));
            for (int j = 0; j < preC.size(); j++) {
                coef.set(bias + j, (int) ((coef.get(bias + j) - delta * preC.get(j)) % MOD));
            }

            if (newLen > oldLen) {
                preC = tmp;
                preI = i;
                preD = (int) d;
            }
        }

        // 计算完后，可能 coef 的末尾有 0，这些 0 不能去掉
        // 比如数列 (1,2,4,2,4,2,4,...) 的系数为 [0,1,0]，表示 f(n) = 0*f(n-1) + 1*f(n-2) + 0*f(n-3) = f(n-2)   (n >= 3)
        // 如果把末尾的 0 去掉，变成 [0,1]，就表示 f(n) = 0*f(n-1) + f(n-2) = f(n-2)   (n >= 2)
        // 看上去一样，但按照这个式子算出来的数列是错误的 (1,2,1,2,1,2,...)

        return coef;
    }

    // 给定常系数齐次线性递推式 f(n) = coef[k-1] * f(n-1) + ... + coef[0] * f(n-k)
    // 以及初始值 f(i) = a[i] (0 <= i < k)
    // 返回 f(n) % MOD，其中参数 n 从 0 开始
    // 注意 coef 的顺序
    // 时间复杂度 O(k^2 log n)，其中 k 是 coef 的长度
    private int kitamasa(int[] coef, int[] a, long n) {
        if (n < a.length) {
            return a[(int) n] % MOD;
        }

        int k = coef.length;
        // 特判 k = 0, 1 的情况
        if (k == 0) {
            return 0;
        }
        if (k == 1) {
            return (int) ((long) a[0] * pow(coef[0], n) % MOD);
        }

        // 计算 resC，以表出 f(n) = resC[k-1] * a[k-1] + ... + resC[0] * a[0]
        int[] resC = new int[k];
        int[] c = new int[k];
        resC[0] = c[1] = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                resC = compose(coef, c, resC);
            }
            // 由于会修改 compose 的第三个参数，这里把 c 复制一份再传入
            c = compose(coef, c, c.clone());
        }

        long ans = 0;
        for (int i = 0; i < k; i++) {
            ans = (ans + (long) resC[i] * a[i]) % MOD;
        }

        return (int) ((ans + MOD) % MOD); // 保证返回值非负
    }

    // 已知 f(n) 的各项系数为 a，f(m) 的各项系数为 b
    // 计算并返回 f(n+m) 的各项系数 c
    private int[] compose(int[] coef, int[] a, int[] b) {
        int k = a.length;
        int[] c = new int[k];
        for (int v : a) {
            for (int j = 0; j < k; j++) {
                c[j] = (int) ((c[j] + (long) v * b[j]) % MOD);
            }
            // 原地计算下一组系数，比如已知 f(4) 的各项系数，现在要计算 f(5) 的各项系数
            // 倒序遍历，避免提前覆盖旧值
            long bk1 = b[k - 1];
            for (int i = k - 1; i > 0; i--) {
                b[i] = (int) ((b[i - 1] + bk1 * coef[i]) % MOD);
            }
            b[0] = (int) (bk1 * coef[0] % MOD);
        }
        return c;
    }

    private int pow(long x, long n) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return (int) res;
    }

    private void reverse(int[] a) {
        for (int i = 0, j = a.length - 1; i < j; i++, j--) {
            int tmp = a[i];
            a[i] = a[j];
            a[j] = tmp;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int MOD = 1'000'000'007;

    int pow(long long x, int n) {
        long long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }

    // 给定数列的前 m 项 a，返回符合 a 的最短常系数齐次线性递推式的系数 coef（模 MOD 意义下）
    // 设 coef 长为 k，当 n >= k 时，有递推式 f(n) = coef[0] * f(n-1) + coef[1] * f(n-2) + ... + coef[k-1] * f(n-k)  （注意 coef 的顺序）
    // 初始值 f(n) = a[n]  (0 <= n < k)
    // 时间复杂度 O(m^2)，其中 m 是 a 的长度
    vector<int> berlekampMassey(const vector<int>& a) {
        vector<int> coef, pre_c, tmp;
        int pre_i = -1, pre_d = 0;

        for (int i = 0; i < a.size(); i++) {
            // d = a[i] - 递推式算出来的值
            long long d = a[i];
            for (int j = 0; j < coef.size(); j++) {
                d = (d - 1LL * coef[j] * a[i - 1 - j]) % MOD;
            }
            if (d == 0) { // 递推式正确
                continue;
            }

            // 首次算错，初始化 coef 为 i+1 个 0
            if (pre_i < 0) {
                coef.resize(i + 1);
                pre_i = i;
                pre_d = d;
                continue;
            }

            int bias = i - pre_i;
            int old_len = coef.size();
            int new_len = bias + pre_c.size();
            if (new_len > old_len) { // 递推式变长了
                tmp = coef;
                coef.resize(new_len);
            }

            // 历史错误为 pre_d = a[pre_i] - sum_j pre_c[j]*a[pre_i-1-j]
            // 现在 a[i] = sum_j coef[j]*a[i-1-j] + d
            // 联立得 a[i] = sum_j coef[j]*a[i-1-j] + d/pre_d * (a[pre_i] - sum_j pre_c[j]*a[pre_i-1-j])
            // 其中 a[pre_i] 的系数 d/pre_d 位于当前（i）的 bias-1 = i-pre_i-1 处
            long long delta = d * pow(pre_d, MOD - 2) % MOD;
            coef[bias - 1] = (coef[bias - 1] + delta) % MOD;
            for (int j = 0; j < pre_c.size(); j++) {
                coef[bias + j] = (coef[bias + j] - delta * pre_c[j]) % MOD;
            }

            if (new_len > old_len) {
                pre_c = move(tmp);
                pre_i = i;
                pre_d = d;
            }
        }

        // 计算完后，可能 coef 的末尾有 0，这些 0 不能去掉
        // 比如数列 (1,2,4,2,4,2,4,...) 的系数为 [0,1,0]，表示 f(n) = 0*f(n-1) + 1*f(n-2) + 0*f(n-3) = f(n-2)   (n >= 3)
        // 如果把末尾的 0 去掉，变成 [0,1]，就表示 f(n) = 0*f(n-1) + f(n-2) = f(n-2)   (n >= 2)
        // 看上去一样，但按照这个式子算出来的数列是错误的 (1,2,1,2,1,2,...)

        return coef;
    }

    // 给定常系数齐次线性递推式 f(n) = coef[k-1] * f(n-1) + coef[k-2] * f(n-2) + ... + coef[0] * f(n-k)
    // 以及初始值 f(i) = a[i] (0 <= i < k)
    // 返回 f(n) % mod，其中参数 n 从 0 开始
    // 注意 coef 的顺序
    // 时间复杂度 O(k^2 log n)，其中 k 是 coef 的长度
    int kitamasa(const vector<int>& coef, const vector<int>& a, long long n) {
        if (n < a.size()) {
            return a[n] % MOD;
        }

        int k = coef.size();
        // 特判 k = 0, 1 的情况
        if (k == 0) {
            return 0;
        }
        if (k == 1) {
            return 1LL * a[0] * pow(coef[0], n) % MOD;
        }

        // 已知 f(n) 的各项系数为 A，f(m) 的各项系数为 B
        // 计算并返回 f(n+m) 的各项系数 C
        auto compose = [&](const vector<int>& A, vector<int> B) -> vector<int> {
            vector<int> C(k);
            for (int v : A) {
                for (int j = 0; j < k; j++) {
                    C[j] = (C[j] + 1LL * v * B[j]) % MOD;
                }
                // 原地计算下一组系数，比如已知 f(4) 的各项系数，现在要计算 f(5) 的各项系数
                // 倒序遍历，避免提前覆盖旧值
                int bk1 = B.back();
                for (int i = k - 1; i > 0; i--) {
                    B[i] = (B[i - 1] + 1LL * bk1 * coef[i]) % MOD;
                }
                B[0] = 1LL * bk1 * coef[0] % MOD;
            }
            return C;
        };

        // 计算 res_c，以表出 f(n) = res_c[k-1] * a[k-1] + res_c[k-2] * a[k-2] + ... + res_c[0] * a[0]
        vector<int> res_c(k), c(k);
        res_c[0] = c[1] = 1;
        for (; n > 0; n /= 2) {
            if (n % 2) {
                res_c = compose(c, move(res_c));
            }
            c = compose(c, c);
        }

        long long ans = 0;
        for (int i = 0; i < k; i++) {
            ans = (ans + 1LL * res_c[i] * a[i]) % MOD;
        }

        return (ans + MOD) % MOD; // 保证返回值非负
    }

    // 见上一题 3699. 锯齿形数组的总数 I
    vector<int> zigZagArraysInit(int l, int r) {
        int k = r - l + 1;
        vector<int> f(k, 1);
        vector<int> a(k * 2); // 注意 a 不包含 n=0 和 n=1 的项，所以下面 kitamasa 传入的是 n-2
        for (int i = 0; i < a.size(); i++) {
            long long pre = 0, s = 0;
            for (int j = 0; j < f.size(); j++) {
                int v = f[j];
                f[j] = pre % MOD;
                pre += v;
                s += f[j];
            }
            a[i] = s * 2 % MOD;
            ranges::reverse(f);
        }
        return a;
    }

public:
    int zigZagArrays(int n, int l, int r) {
        vector<int> a = zigZagArraysInit(l, r);
        vector<int> coef = berlekampMassey(a);
        ranges::reverse(coef); // 注意 kitamasa 入参的顺序
        return kitamasa(coef, a, n - 2);
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

// 给定数列的前 m 项 a，返回符合 a 的最短常系数齐次线性递推式的系数 coef（模 mod 意义下）
// 设 coef 长为 k，当 n >= k 时，有递推式 f(n) = coef[0] * f(n-1) + coef[1] * f(n-2) + ... + coef[k-1] * f(n-k)  （注意 coef 的顺序）
// 初始值 f(n) = a[n]  (0 <= n < k)
// 时间复杂度 O(m^2)，其中 m 是 a 的长度
func berlekampMassey(a []int) (coef []int) {
	var preC []int
	preI, preD := -1, 0

	for i, v := range a {
		// d = a[i] - 递推式算出来的值
		d := v
		for j, c := range coef {
			d = (d - c*a[i-1-j]) % mod
		}
		if d == 0 { // 递推式正确
			continue
		}

		// 首次算错，初始化 coef 为 i+1 个 0
		if preI < 0 {
			coef = make([]int, i+1)
			preI, preD = i, d
			continue
		}

		bias := i - preI
		oldLen := len(coef)
		newLen := bias + len(preC)
		var tmp []int
		if newLen > oldLen { // 递推式变长了
			tmp = slices.Clone(coef)
			coef = slices.Grow(coef, newLen-oldLen)[:newLen] // coef.resize(newLen)
		}

		// 历史错误为 preD = a[preI] - sum_j preC[j]*a[preI-1-j]
		// 现在 a[i] = sum_j coef[j]*a[i-1-j] + d
		// 联立得 a[i] = sum_j coef[j]*a[i-1-j] + d/preD * (a[preI] - sum_j preC[j]*a[preI-1-j])
		// 其中 a[preI] 的系数 d/preD 位于当前（i）的 bias-1 = i-preI-1 处
		delta := d * pow(preD, mod-2) % mod // pow(preD, mod-2) 为 preD 的逆元
		coef[bias-1] = (coef[bias-1] + delta) % mod
		for j, c := range preC {
			coef[bias+j] = (coef[bias+j] - delta*c) % mod
		}

		if newLen > oldLen {
			preC = tmp
			preI, preD = i, d
		}
	}

	// 计算完后，可能 coef 的末尾有 0，这些 0 不能去掉
	// 比如数列 (1,2,4,2,4,2,4,...) 的系数为 [0,1,0]，表示 f(n) = 0*f(n-1) + 1*f(n-2) + 0*f(n-3) = f(n-2)   (n >= 3)
	// 如果把末尾的 0 去掉，变成 [0,1]，就表示 f(n) = 0*f(n-1) + f(n-2) = f(n-2)   (n >= 2)
	// 看上去一样，但按照这个式子算出来的数列是错误的 (1,2,1,2,1,2,...)

	return
}

// 给定常系数齐次线性递推式 f(n) = coef[k-1] * f(n-1) + coef[k-2] * f(n-2) + ... + coef[0] * f(n-k)
// 以及初始值 f(i) = a[i] (0 <= i < k)
// 返回 f(n) % mod，其中参数 n 从 0 开始
// 注意 coef 的顺序
// 时间复杂度 O(k^2 log n)，其中 k 是 coef 的长度
func kitamasa(coef, a []int, n int) (ans int) {
	defer func() { ans = (ans + mod) % mod }() // 保证结果非负
	if n < len(a) {
		return a[n] % mod
	}

	k := len(coef)
	// 特判 k = 0, 1 的情况
	if k == 0 {
		return 0
	}
	if k == 1 {
		return a[0] * pow(coef[0], n) % mod
	}

	// 已知 f(n) 的各项系数为 a，f(m) 的各项系数为 b
	// 计算并返回 f(n+m) 的各项系数 c
	compose := func(a, b []int) []int {
		c := make([]int, k)
		for _, v := range a {
			for j, w := range b {
				c[j] = (c[j] + v*w) % mod
			}
			// 原地计算下一组系数，比如已知 f(4) 的各项系数，现在要计算 f(5) 的各项系数
			// 倒序遍历，避免提前覆盖旧值
			bk1 := b[k-1]
			for i := k - 1; i > 0; i-- {
				b[i] = (b[i-1] + bk1*coef[i]) % mod
			}
			b[0] = bk1 * coef[0] % mod
		}
		return c
	}

	// 计算 resC，以表出 f(n) = resC[k-1] * a[k-1] + resC[k-2] * a[k-2] + ... + resC[0] * a[0]
	resC := make([]int, k)
	resC[0] = 1
	c := make([]int, k)
	c[1] = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			resC = compose(c, resC)
		}
		// 由于会修改 compose 的第二个参数，这里把 c 复制一份再传入
		c = compose(c, slices.Clone(c))
	}

	for i, c := range resC {
		ans = (ans + c*a[i]) % mod
	}

	return
}

// 见上一题 3699. 锯齿形数组的总数 I
func zigZagArraysInit(l, r int) []int {
	k := r - l + 1
	f := make([]int, k)
	for i := range f {
		f[i] = 1
	}

	a := make([]int, k*2) // 注意 a 不包含 n=0 和 n=1 的项，所以下面 kitamasa 传入的是 n-2
	for i := range a {
		pre := 0
		s := 0
		for j, v := range f {
			f[j] = pre % mod
			pre += v
			s += f[j]
		}
		a[i] = s * 2 % mod
		slices.Reverse(f)
	}
	return a
}

func zigZagArrays(n, l, r int) int {
	a := zigZagArraysInit(l, r)
	coef := berlekampMassey(a)
	slices.Reverse(coef) // 注意 kitamasa 入参的顺序
	return kitamasa(coef, a, n-2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((r-l)^2\log n)$。
- 空间复杂度：$\mathcal{O}(r-l)$。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
