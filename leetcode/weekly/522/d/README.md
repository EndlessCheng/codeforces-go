问题等价于：

- 把 $n$ 拆分成若干个奇数之和的方案数。这里的拆分考虑顺序，例如 $1+3$ 和 $3+1$ 是不同的方案。
- 由于字符串的第一个字母可以是 $\texttt{a}$，也可以是 $\texttt{b}$，方案数要乘以 $2$。

枚举最后一个连续相同段的长度为 $j=1,3,5,\ldots$，问题变成把 $n-j$ 拆分成若干个奇数之和的方案数，这是一个规模更小的子问题。

于是，定义 $f[i]$ 表示把 $i$ 拆分成若干个奇数之和的方案数。枚举最后一个连续相同段的长度为 $j=1,3,5,\ldots$，问题变成 $f[i-j]$。这些方案互斥，根据加法原理，我们有

$$
f[i] = f[i-1] + f[i-3] + f[i-5] + \cdots + f[(i+1)\bmod 2]
$$

用 $i-2$ 替换上式中的 $i$，得

$$
f[i-2] = f[i-3] + f[i-5] + \cdots + f[(i+1)\bmod 2]
$$

用 $f[i-2]$ 替换第一个式子后面的和式，得

$$
f[i] = f[i-1] + f[i-2]
$$

初始值：$f[1] = f[2] = 1$。也可以定义 $f[0] = 0$，这样更好计算。

答案：$2f[n]$。

做法同 [509. 斐波那契数](https://leetcode.cn/problems/fibonacci-number/)，见 [我的题解](https://leetcode.cn/problems/fibonacci-number/solutions/3811875/san-chong-xie-fa-shu-zu-di-tui-kong-jian-3lwj/) 中的「**三、矩阵快速幂优化**」。

代码实现时，注意取模。为什么可以在**中途取模**？原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV1FiHj6uESs/?t=22m45s)，欢迎点赞关注~

```py [sol-Python3]
MOD = 1_000_000_007

# a @ b，其中 @ 是矩阵乘法
# 更快的写法见另一份代码【NumPy】
def mul(a: list[list[int]], b: list[list[int]]) -> list[list[int]]:
    return [[sum(x * y for x, y in zip(row, col)) % MOD for col in zip(*b)]
            for row in a]

# a^n @ f
def pow_mul(a: list[list[int]], n: int, f: list[list[int]]) -> list[list[int]]:
    res = f
    while n:
        if n & 1:
            res = mul(a, res)
        a = mul(a, a)
        n >>= 1
    return res

class Solution:
    def countGoodStrings(self, n: int) -> int:
        m = [[1, 1], [1, 0]]
        f1 = [[2], [0]]  # 这里初始化成 2，就不用把答案乘以 2 了
        fn = pow_mul(m, n - 1, f1)
        return fn[0][0]
```

```py [sol-NumPy]
import numpy as np

MOD = 1_000_000_007

# a^n @ f
def pow_mul(a: np.ndarray, n: int, f: np.ndarray) -> np.ndarray:
    res = f
    while n:
        if n & 1:
            res = a @ res % MOD
        a = a @ a % MOD
        n >>= 1
    return res

class Solution:
    def countGoodStrings(self, n: int) -> int:
        m = np.array([[1, 1], [1, 0]], dtype=object)
        f1 = np.array([2, 0], dtype=object)  # 这里初始化成 2，就不用把答案乘以 2 了
        fn = pow_mul(m, n - 1, f1)
        return fn[0]
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int countGoodStrings(long n) {
        long[][] m = {
            {1, 1},
            {1, 0},
        };
        long[][] f1 = {{2}, {0}}; // 这里初始化成 2，就不用把答案乘以 2 了
        long[][] fn = powMul(m, n - 1, f1);
        return (int) fn[0][0];
    }

    // a^n * f
    private long[][] powMul(long[][] a, long n, long[][] f) {
        long[][] res = f;
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
matrix mul(const matrix& a, const matrix& b) {
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

// a^n * f
matrix pow_mul(matrix a, long long n, const matrix& f) {
    matrix res = f;
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
    int countGoodStrings(long long n) {
        matrix m = {
            {1, 1},
            {1, 0},
        };
        matrix f1 = {{2}, {0}}; // 这里初始化成 2，就不用把答案乘以 2 了
        matrix fn = pow_mul(m, n - 1, f1);
        return fn[0][0];
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

// 返回矩阵 a 和矩阵 b 相乘的结果
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

// a^n * f
func (a matrix) powMul(n int64, f matrix) matrix {
	res := f
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func countGoodStrings(n int64) (ans int) {
	m := matrix{
		{1, 1},
		{1, 0},
	}
	f1 := matrix{{2}, {0}} // 这里初始化成 2，就不用把答案乘以 2 了
	fn := m.powMul(n-1, f1)
	return fn[0][0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面动态规划题单的「**§11.6 矩阵快速幂优化 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
