先求出单个字母 $\texttt{a},\texttt{b},\ldots,\texttt{z}$ 替换 $t$ 次后的长度。

## 寻找子问题

例如字母 $\texttt{a}$ 替换一次变成 $\texttt{b}$ 和 $\texttt{c}$，问题变成计算 $\texttt{b}$ 替换 $t-1$ 次后的长度，$\texttt{c}$ 替换 $t-1$ 次后的长度。二者之和即为 $\texttt{a}$ 替换 $t$ 次后的长度。

## 状态定义和状态转移方程

据此，定义 $f[i][j]$ 表示字母 $j$ 替换 $i$ 次后的长度。

上面的例子，就是 $f[i][0] = f[i-1][1] + f[i-1][2]$。

一般地，设 $c=\textit{nums}[j]$，我们有

$$
f[i][j] = \sum_{k=j+1}^{j+c} f[i-1][k\bmod 26]
$$

初始值 $f[0][j] = 1$。

答案为 $\sum\limits_{j=0}^{25} f[t][j]\cdot \textit{cnt}[j]$。其中 $\textit{cnt}[j]$ 为 $s$ 中的字母 $j$ 的出现次数。

这可以解决 [3335. 字符串转换后的长度 I](https://leetcode.cn/problems/total-characters-in-string-after-transformations-i/)，但对于本题，还需继续优化。

**注**：本题直接计算这个 DP 是 $\mathcal{O}(t|\Sigma|^2)$ 的。用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 优化，可以做到 $\mathcal{O}(t|\Sigma|)$。

## 矩阵快速幂优化

以示例 1 为例（也相当于周赛第二题），其中 $\textit{nums}=[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2]$。

我们有

$$
\begin{aligned}
f[i][0] &= f[i-1][1]     \\
f[i][1] &= f[i-1][2]     \\
f[i][2] &= f[i-1][3]     \\
\vdots \\
f[i][23] &= f[i-1][24]     \\
f[i][24] &= f[i-1][25]     \\
f[i][25] &= f[i-1][0] + f[i-1][1]     \\
\end{aligned}
$$

用矩阵乘法表示，即

$$
\begin{bmatrix}
f[i][0] \\
f[i][1] \\
f[i][2] \\
\vdots \\
f[i][23] \\
f[i][24] \\
f[i][25] \\
\end{bmatrix}
= \begin{bmatrix}
0 & 1 & 0 & 0 & \cdots & 0 & 0 \\
0 & 0 & 1 & 0 & \cdots & 0 & 0 \\
0 & 0 & 0 & 1 & \cdots & 0 & 0 \\
\vdots & \vdots & \vdots & \vdots & \ddots & \vdots & \vdots \\
0 & 0 & 0 & 0 & \cdots & 1 & 0 \\
0 & 0 & 0 & 0 & \cdots & 0 & 1 \\
1 & 1 & 0 & 0 & \cdots & 0 & 0 \\
\end{bmatrix}
\begin{bmatrix}
f[i-1][0] \\
f[i-1][1] \\
f[i-1][2] \\
\vdots \\
f[i-1][23] \\
f[i-1][24] \\
f[i-1][25] \\
\end{bmatrix}
$$

把上式中的三个矩阵分别记作 $F[i],M,F[i-1]$，即

$$
F[i] = M\times F[i-1]
$$

那么有

$$
\begin{aligned}
F[t] &= M\times F[t-1]      \\
&= M\times M\times F[t-2]        \\
&= M\times M\times M\times  F[t-3]        \\
&\ \ \vdots  \\
&= M^t\times F[0]  \\
\end{aligned}
$$

其中 $F[0]$ 是一个长为 $26$ 的列向量，值全为 $1$（对应着 $f$ 数组的初始值 $f[0][j] = 1$）。

$M^t$ 可以用**快速幂**计算，原理请看[【图解】一张图秒懂快速幂](https://leetcode.cn/problems/powx-n/solution/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/)。

根据矩阵乘法的运算法则，$f[t][j]$ 等于矩阵 $M^t$ 的第 $j$ 行与列向量 $F[0]$ 计算点积。由于 $F[0]$ 全为 $1$，所以 $f[t][j]$ 也等于 $M^t$ 第 $j$ 行的元素和。

一般地，枚举 $i=0,1,\ldots, 25$ 以及 $j=i+1,i+2,\ldots,i+\textit{nums}[i]$，初始化 $M[i][j\bmod 26]=1$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1hn1MYhEtC/?t=21m27s)，欢迎点赞关注~

```py [sol-Python3]
MOD = 1_000_000_007

# a @ b，其中 @ 是矩阵乘法
def mul(a: List[List[int]], b: List[List[int]]) -> List[List[int]]:
    return [[sum(x * y for x, y in zip(row, col)) % MOD for col in zip(*b)]
            for row in a]

# a^n @ f0
def pow_mul(a: List[List[int]], n: int, f0: List[List[int]]) -> List[List[int]]:
    res = f0
    while n:
        if n & 1:
            res = mul(a, res)
        a = mul(a, a)
        n >>= 1
    return res

class Solution:
    def lengthAfterTransformations(self, s: str, t: int, nums: List[int]) -> int:
        SIZE = 26
        f0 = [[1] for _ in range(SIZE)]
        m = [[0] * SIZE for _ in range(SIZE)]
        for i, c in enumerate(nums):
            for j in range(i + 1, i + c + 1):
                m[i][j % SIZE] = 1
        mt = pow_mul(m, t, f0)

        ans = 0
        for ch, cnt in Counter(s).items():
            ans += mt[ord(ch) - ord('a')][0] * cnt
        return ans % MOD
```

```py [sol-NumPy]
import numpy as np

MOD = 1_000_000_007

# a^n @ f0
def pow_mul(a: np.ndarray, n: int, f0: np.ndarray) -> np.ndarray:
    res = f0
    while n:
        if n & 1:
            res = a @ res % MOD
        a = a @ a % MOD
        n >>= 1
    return res

class Solution:
    def lengthAfterTransformations(self, s: str, t: int, nums: List[int]) -> int:
        SIZE = 26
        f0 = np.ones((SIZE,), dtype=object)
        m = np.zeros((SIZE, SIZE), dtype=object)
        for i, c in enumerate(nums):
            for j in range(i + 1, i + c + 1):
                m[i, j % SIZE] = 1
        mt = pow_mul(m, t, f0)

        ans = 0
        for ch, cnt in Counter(s).items():
            ans += mt[ord(ch) - ord('a')] * cnt
        return ans % MOD
```

```py [sol-NumPy 写法二]
import numpy as np

MOD = 1_000_000_007

# f0 @ a^n
def pow_mul(f0: np.ndarray, a: np.ndarray, n: int) -> np.ndarray:
    res = f0
    while n:
        if n & 1:
            res = res @ a % MOD
        a = a @ a % MOD
        n >>= 1
    return res

class Solution:
    def lengthAfterTransformations(self, s: str, t: int, nums: List[int]) -> int:
        cnt = Counter(s)
        f0 = np.array([cnt[c] for c in ascii_lowercase], dtype=object)

        SIZE = 26
        m = np.zeros((SIZE, SIZE), dtype=object)
        for i, c in enumerate(nums):
            for j in range(i + 1, i + c + 1):
                m[i, j % SIZE] = 1

        mt = pow_mul(f0, m, t)
        return np.sum(mt) % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int lengthAfterTransformations(String s, int t, List<Integer> nums) {
        final int SIZE = 26;

        int[][] f0 = new int[SIZE][1];
        for (int i = 0; i < SIZE; i++) {
            f0[i][0] = 1;
        }

        int[][] m = new int[SIZE][SIZE];
        for (int i = 0; i < SIZE; i++) {
            int c = nums.get(i);
            for (int j = i + 1; j <= i + c; j++) {
                m[i][j % SIZE] = 1;
            }
        }

        int[][] mt = powMul(m, t, f0);

        int[] cnt = new int[SIZE];
        for (char c : s.toCharArray()) {
            cnt[c - 'a']++;
        }

        long ans = 0;
        for (int i = 0; i < SIZE; i++) {
            ans += (long) mt[i][0] * cnt[i];
        }
        return (int) (ans % MOD);
    }

    // a^n * f0
    private int[][] powMul(int[][] a, int n, int[][] f0) {
        int[][] res = f0;
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
    private int[][] mul(int[][] a, int[][] b) {
        int[][] c = new int[a.length][b[0].length];
        for (int i = 0; i < a.length; i++) {
            for (int k = 0; k < a[i].length; k++) {
                if (a[i][k] == 0) {
                    continue;
                }
                for (int j = 0; j < b[k].length; j++) {
                    c[i][j] = (int) ((c[i][j] + (long) a[i][k] * b[k][j]) % MOD);
                }
            }
        }
        return c;
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int MOD = 1'000'000'007;
    static constexpr int SIZE = 26;

    using Matrix = array<array<int, SIZE>, SIZE>;

    // 返回矩阵 a 和矩阵 b 相乘的结果
    Matrix mul(Matrix& a, Matrix& b) {
        Matrix c{};
        for (int i = 0; i < SIZE; i++) {
            for (int k = 0; k < SIZE; k++) {
                if (a[i][k] == 0) {
                    continue;
                }
                for (int j = 0; j < SIZE; j++) {
                    c[i][j] = (c[i][j] + (long long) a[i][k] * b[k][j]) % MOD;
                }
            }
        }
        return c;
    }

    // 返回 n 个矩阵 a 相乘的结果
    Matrix pow(Matrix a, int n) {
        Matrix res = {};
        for (int i = 0; i < SIZE; i++) {
            res[i][i] = 1; // 单位矩阵
        }
        while (n) {
            if (n & 1) {
                res = mul(res, a);
            }
            a = mul(a, a);
            n >>= 1;
        }
        return res;
    }

public:
    int lengthAfterTransformations(string s, int t, vector<int>& nums) {
        Matrix m{};
        for (int i = 0; i < SIZE; i++) {
            for (int j = i + 1; j <= i + nums[i]; j++) {
                m[i][j % SIZE] = 1;
            }
        }
        Matrix mt = pow(m, t);

        int cnt[SIZE]{};
        for (char c : s) {
            cnt[c - 'a']++;
        }

        long long ans = 0;
        for (int i = 0; i < SIZE; i++) {
            // m 第 i 行的和就是 f[t][i]
            ans += reduce(mt[i].begin(), mt[i].end(), 0LL) * cnt[i];
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

// a^n * f0
func (a matrix) powMul(n int, f0 matrix) matrix {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func lengthAfterTransformations(s string, t int, nums []int) (ans int) {
	const size = 26
	f0 := newMatrix(size, 1)
	for i := range f0 {
		f0[i][0] = 1
	}
	m := newMatrix(size, size)
	for i, c := range nums {
		for j := i + 1; j <= i+c; j++ {
			m[i][j%size] = 1
		}
	}
	mt := m.powMul(t, f0)

	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	for i, row := range mt {
		ans += row[0] * cnt[i]
	}
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|^3\log t)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|^2)$。

## Berlekamp-Massey 算法 + Kitamasa 算法

见我的知乎科普文章：

[Berlekamp-Massey 算法：如何预测数列的下一项？](https://zhuanlan.zhihu.com/p/1966417899825665440)

[Kitamasa 算法：更快地计算线性递推的第 n 项](https://zhuanlan.zhihu.com/p/1964051212304364939)

矩阵快速幂优化 DP 的题，一般可以结合这两个算法优化。通用三步：

1. 首先，用本文开头的 DP，结合前缀和优化，计算 $t=1,2,\ldots,52$ 时的答案。其中 $52$ 等于系数矩阵的阶 $26$ 乘以 $2$。
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
    # 返回 f(n) % MOD，其中参数 n 从 0 开始
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

    # 计算 t = 1,2,...,52 的答案
    def lengthAfterTransformationsInit(self, nums: List[int], cnt: List[int]) -> List[int]:
        k = 26
        f = [1] * k
        a = [0] * (k * 2)
        for i in range(k * 2):
            # 计算 f + f 的前缀和
            pre = list(accumulate(f + f, initial=0))
            s = 0
            for j, (num, c) in enumerate(zip(nums, cnt)):
                f[j] = (pre[j + num + 1] - pre[j + 1]) % MOD
                s += f[j] * c
            a[i] = s % MOD
        return a

    def lengthAfterTransformations(self, s: str, t: int, nums: List[int]) -> int:
        cnt = Counter(s)
        cnt = [cnt[c] for c in ascii_lowercase]
        a = self.lengthAfterTransformationsInit(nums, cnt)
        coef = self.berlekampMassey(a)
        coef.reverse()  # 注意 kitamasa 入参的顺序
        return self.kitamasa(coef, a, t - 1)
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int SIZE = 26;

    public int lengthAfterTransformations(String s, int t, List<Integer> Nums) {
        int[] nums = new int[SIZE];
        for (int i = 0; i < SIZE; i++) {
            nums[i] = Nums.get(i);
        }

        int[] cnt = new int[SIZE];
        for (char c : s.toCharArray()) {
            cnt[c - 'a']++;
        }

        int[] a = lengthAfterTransformationsInit(nums, cnt);
        List<Integer> coef = berlekampMassey(a);
        int k = coef.size();
        int[] c = new int[k];
        for (int i = 0; i < k; i++) {
            c[i] = coef.get(k - 1 - i); // 注意 kitamasa 入参的顺序
        }

        return kitamasa(c, a, t - 1);
    }

    // 计算 t = 1,2,...,52 的答案
    private int[] lengthAfterTransformationsInit(int[] nums, int[] cnt) {
        long[] f = new long[SIZE];
        Arrays.fill(f, 1);
        long[] sum = new long[SIZE * 2 + 1];

        int[] a = new int[SIZE * 2];
        for (int i = 0; i < SIZE * 2; i++) {
            // 计算 f + f 的前缀和
            for (int j = 0; j < SIZE * 2; j++) {
                sum[j + 1] = sum[j] + f[j % SIZE];
            }
            long s = 0;
            for (int j = 0; j < SIZE; j++) {
                f[j] = (sum[j + nums[j] + 1] - sum[j + 1]) % MOD;
                s += f[j] * cnt[j];
            }
            a[i] = (int) (s % MOD);
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
    // 返回 f(n) % MOD，其中参数 n 从 0 开始
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

    // 计算 t = 1,2,...,52 的答案
    vector<int> lengthAfterTransformationsInit(vector<int>& nums, array<int, 26>& cnt) {
        constexpr int k = 26;
        vector<long long> f(k, 1);
        vector<long long> sum(k * 2 + 1);
        vector<int> a(k * 2);
        for (int i = 0; i < k * 2; i++) {
            // 计算 f + f 的前缀和
            for (int j = 0; j < k * 2; j++) {
                sum[j + 1] = sum[j] + f[j % k];
            }
            long long s = 0;
            for (int j = 0; j < k; j++) {
                f[j] = (sum[j + nums[j] + 1] - sum[j + 1]) % MOD;
                s += f[j] * cnt[j];
            }
            a[i] = s % MOD;
        }
        return a;
    }

public:
    int lengthAfterTransformations(string s, int t, vector<int>& nums) {
        array<int, 26> cnt{};
        for (char c : s) {
            cnt[c - 'a']++;
        }
        vector<int> a = lengthAfterTransformationsInit(nums, cnt);
        vector<int> coef = berlekampMassey(a);
        ranges::reverse(coef); // 注意 kitamasa 入参的顺序
        return kitamasa(coef, a, t - 1);
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

	// 手动找规律用
	for i, c := range coef {
		if c < -mod/2 {
			c += mod
		} else if c > mod/2 {
			c -= mod
		}
		coef[i] = c
	}

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

// 计算 t = 1,2,...,52 的答案
func lengthAfterTransformationsInit(nums, cnt []int) []int {
	const k = 26
	f := [k]int{}
	for i := range f {
		f[i] = 1
	}
	sum := [k*2 + 1]int{}

	a := make([]int, k*2)
	for i := range a {
		// 计算 f + f 的前缀和
		for j := range k * 2 {
			sum[j+1] = sum[j] + f[j%k]
		}
		s := 0
		for j, c := range nums {
			f[j] = (sum[j+c+1] - sum[j+1]) % mod
			s += f[j] * cnt[j]
		}
		a[i] = s % mod
	}
	return a
}

func lengthAfterTransformations(s string, t int, nums []int) int {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	a := lengthAfterTransformationsInit(nums, cnt[:])
	coef := berlekampMassey(a)
	slices.Reverse(coef) // 注意 kitamasa 入参的顺序
	return kitamasa(coef, a, t-1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|^2\log t)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
