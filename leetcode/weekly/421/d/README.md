先求出单个字母 $\texttt{a},\texttt{b},\ldots,\texttt{z}$ 替换 $t$ 次后的长度。

## 寻找子问题

例如字母 $\texttt{a}$ 替换一次变成 $\texttt{b}$ 和 $\texttt{c}$，问题变成计算 $\texttt{b}$ 替换 $t-1$ 次后的长度，$\texttt{c}$ 替换 $t-1$ 次后的长度。二者之和即为 $\texttt{a}$ 替换 $t$ 次后的长度。

## 状态定义和状态转移方程

据此，定义 $f[i][j]$ 表示字母 $j$ 替换 $i$ 次后的长度。

上面的例子，就是 $f[i][0] = f[i-1][1] + f[i-1][2]$。

一般地，设 $c=\textit{nums}[i]$，我们有

$$
f[i][j] = \sum_{k=j+1}^{j+c} f[i-1][k\bmod 26]
$$

初始值 $f[0][j] = 1$。

答案为 $\sum\limits_{j=0}^{25} f[t][j]\cdot \textit{cnt}[j]$。其中 $\textit{cnt}[j]$ 为 $s$ 中的字母 $j$ 的出现次数。

直接计算这个 DP 的话，时间复杂度为 $\mathcal{O}(t|\Sigma|)$，这可以解决周赛第二题。对于本题，还需继续优化。

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
F[t] ={} & M\times F[t-1]      \\
={} & M\times M\times F[t-2]        \\
={} & M\times M\times M\times  F[t-3]        \\
\vdots & \\
={} & M^t\times F[0]
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
        m = pow_mul(m, t, f0)

        ans = 0
        for ch, cnt in Counter(s).items():
            ans += m[ord(ch) - ord('a')][0] * cnt
        return ans % MOD
```

```py [sol-NumPy]
import numpy as np

MOD = 1_000_000_007

# a^n @ f0
def pow(a: np.ndarray, n: int, f0: np.ndarray) -> np.ndarray:
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
        m = pow(m, t, f0)

        ans = 0
        for ch, cnt in Counter(s).items():
            ans += m[ord(ch) - ord('a')] * cnt
        return ans % MOD
```

```py [sol-NumPy 写法二]
import numpy as np

MOD = 1_000_000_007

# f0 @ a^n
def pow(f0: np.ndarray, a: np.ndarray, n: int) -> np.ndarray:
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

        m = pow(f0, m, t)
        return np.sum(m) % MOD
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
        m = powMul(m, t, f0);

        int[] cnt = new int[SIZE];
        for (char c : s.toCharArray()) {
            cnt[c - 'a']++;
        }

        long ans = 0;
        for (int i = 0; i < SIZE; i++) {
            ans += (long) m[i][0] * cnt[i];
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
        m = pow(m, t);

        int cnt[SIZE]{};
        for (char c : s) {
            cnt[c - 'a']++;
        }

        long long ans = 0;
        for (int i = 0; i < SIZE; i++) {
            // m 第 i 行的和就是 f[t][i]
            ans += reduce(m[i].begin(), m[i].end(), 0LL) * cnt[i];
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
	m = m.powMul(t, f0)

	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	for i, row := range m {
		ans += row[0] * cnt[i]
	}
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|^3\log t)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|^2)$。

更多相似题目，见下面动态规划题单中的「**§7.3 矩阵快速幂优化 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
