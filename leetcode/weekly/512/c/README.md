## 总体思路

1. 计算长为 $k$ 和为 $n$ 的正整数序列个数。
2. 减去其中乘积为奇数的序列个数。

## 正整数序列个数

想象有 $n$ 个小球排成一行，我们要把这 $n$ 个小球划分成 $k$ 个非空段。

根据「隔板法」，$n$ 个小球有 $n-1$ 个空隙，选择 $k-1$ 个空隙插入隔板，即可分成 $k$ 个非空段，方案数为组合数

$$
\binom {n-1} {k-1}
$$

## 乘积为奇数的序列个数

序列只要有一个偶数，乘积就是偶数。

所以只有全为奇数的序列，乘积才不是偶数。减去这种序列的方案数，即为答案。

**问题**：把 $n$ 拆分成 $k$ 个正奇数的方案数。

如何把这个问题变成我们熟悉的问题？

设 $n=x_1+x_2+\cdots+x_k$，其中 $x_i$ 是正奇数。

正奇数 $x_i$ 可以表示成 $2y_i-1$，其中 $y_i$ 是正整数。

> **注**：如果表示成 $2y_i+1$，那么 $y_i$ 就是非负整数了，无法直接转化成「正整数序列个数」问题。

那么

$$
\begin{aligned}
n &= x_1+x_2+\cdots+x_k         \\
  &= (2y_1-1) + (2y_2-1) + \cdots + (2y_k-1)           \\
  &= 2(y_1+y_2+\cdots+y_k) - k           \\
\end{aligned}
$$

移项得

$$
y_1+y_2+\cdots+y_k = \dfrac{n+k}{2}
$$

**问题变成**：把 $\dfrac{n+k}{2}$ 拆分成 $k$ 个正整数的方案数。

如果 $n+k$ 是奇数，那么 $\dfrac{n+k}{2}$ 是小数，无法由 $k$ 个正整数相加得到，所以方案数是 $0$。

如果 $n+k$ 是偶数，那么和前文的做法一样，由隔板法可得，方案数为组合数

$$
\binom {\frac{n+k}{2} - 1} {k-1}
$$

综上所述，答案为

$$
\begin{cases}
\dbinom {n-1} {k-1}, & (n+k)\bmod 2 = 1     \\
\dbinom {n-1} {k-1} - \dbinom {\frac{n+k}{2} - 1} {k-1}, & (n+k)\bmod 2 = 0     \\
\end{cases}
$$

代码实现时，可以**预处理阶乘及其逆元**，从而 $\mathcal{O}(1)$ 计算组合数。代码模板见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV1Ps3j6nE3D/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countValidSequences(self, n: int, k: int) -> int:
        # 更快的做法见【Python3 预处理】
        ans = comb(n - 1, k - 1)
        if (n + k) % 2 == 0:
            ans -= comb((n + k) // 2 - 1, k - 1)
        return ans % 1_000_000_007
```

```py [sol-Python3 预处理]
MOD = 1_000_000_007
MX = 500_000

fac = [0] * MX  # fac[i] = i!
fac[0] = 1
for i in range(1, MX):
    fac[i] = fac[i - 1] * i % MOD

inv_f = [0] * MX  # inv_f[i] = i!^-1
inv_f[-1] = pow(fac[-1], -1, MOD)
for i in range(MX - 1, 0, -1):
    inv_f[i - 1] = inv_f[i] * i % MOD

# 从 n 个数中选 m 个数的方案数
def comb(n: int, m: int) -> int:
    return fac[n] * inv_f[m] * inv_f[n - m] % MOD


class Solution:
    def countValidSequences(self, n: int, k: int) -> int:
        ans = comb(n - 1, k - 1)
        if (n + k) % 2 == 0:
            ans = (ans - comb((n + k) // 2 - 1, k - 1)) % MOD
        return ans
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 500_000;
    private static final long[] F = new long[MX]; // F[i] = i!
    private static final long[] INV_F = new long[MX]; // INV_F[i] = i!^-1 = pow(i!, MOD-2)
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        F[0] = 1;
        for (int i = 1; i < MX; i++) {
            F[i] = F[i - 1] * i % MOD;
        }

        INV_F[MX - 1] = pow(F[MX - 1], MOD - 2);
        for (int i = MX - 1; i > 0; i--) {
            INV_F[i - 1] = INV_F[i] * i % MOD;
        }
    }

    private long pow(long x, int n) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }

    // 从 n 个数中选 m 个数的方案数
    private long comb(int n, int m) {
        return F[n] * INV_F[m] % MOD * INV_F[n - m] % MOD;
    }

    public int countValidSequences(int n, int k) {
        long ans = comb(n - 1, k - 1);
        if ((n + k) % 2 == 0) {
            ans = (ans - comb((n + k) / 2 - 1, k - 1) + MOD) % MOD; // +MOD 保证答案非负
        }
        return (int) ans;
    }
}
```

```cpp [sol-C++]
const int MOD = 1'000'000'007;
const int MX = 500'000;

long long F[MX]; // F[i] = i!
long long INV_F[MX]; // INV_F[i] = i!^-1 = qpow(i!, MOD-2)

long long qpow(long long x, int n) {
    long long res = 1;
    for (; n; n /= 2) {
        if (n % 2) {
            res = res * x % MOD;
        }
        x = x * x % MOD;
    }
    return res;
}

auto init = [] {
    F[0] = 1;
    for (int i = 1; i < MX; i++) {
        F[i] = F[i - 1] * i % MOD;
    }

    INV_F[MX - 1] = qpow(F[MX - 1], MOD - 2);
    for (int i = MX - 1; i; i--) {
        INV_F[i - 1] = INV_F[i] * i % MOD;
    }
    return 0;
}();

// 从 n 个数中选 m 个数的方案数
long long comb(int n, int m) {
    return F[n] * INV_F[m] % MOD * INV_F[n - m] % MOD;
}

class Solution {
public:
    int countValidSequences(int n, int k) {
        long long ans = comb(n - 1, k - 1);
        if ((n + k) % 2 == 0) {
            ans = (ans - comb((n + k) / 2 - 1, k - 1) + MOD) % MOD; // +MOD 保证答案非负
        }
        return ans;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 500_000

var fac [mx]int  // fac[i] = i!
var invF [mx]int // invF[i] = i!^-1 = pow(i!, mod-2)

func init() {
	fac[0] = 1
	for i := 1; i < mx; i++ {
		fac[i] = fac[i-1] * i % mod
	}

	invF[mx-1] = pow(fac[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

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

// 从 n 个数中选 m 个数的方案数
func comb(n, m int) int {
	return fac[n] * invF[m] % mod * invF[n-m] % mod
}

func countValidSequences(n, k int) int {
	ans := comb(n-1, k-1)
	if (n+k)%2 == 0 {
		ans = (ans - comb((n+k)/2-1, k-1) + mod) % mod // +mod 保证答案非负
	}
	return ans
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 附：其他推导方法

计算乘积为奇数的序列个数。

仍然想象有 $n$ 个小球排成一行。

1. 先拿出 $k$ 个小球。
2. 把剩余的 $n-k$ 个小球划分成 $k$ 段（可以为空），且每一段的小球个数都是偶数。
3. 之前拿出了 $k$ 个小球，每段再放入 $1$ 个小球，即可让每一段的小球个数都是正奇数。

其中第二步的方案数是多少？

如果 $n-k$ 是奇数，那么无法划分，方案数是 $0$。

否则，把小球两两一组，两个小球视作一个大球，问题变成：把 $\dfrac{n-k}{2}$ 个大球划分成 $k$ 段（可以为空）的方案数。

1. 先放入 $k$ 个大球。
2. 把这 $\dfrac{n-k}{2} + k$ 个大球分成 $k$ 个**非空**段。
3. 从每一段再拿出 $1$ 个大球，就把 $\dfrac{n-k}{2}$ 个大球划分成 $k$ 段（可以为空）了。

其中第二步用隔板法解决，方案数为组合数

$$
\binom {\frac{n-k}{2} + k - 1} {k-1} = \binom {\frac{n+k}{2} - 1} {k-1}
$$

## 专题训练

见下面数学题单的「**§2.2 组合计数**」。

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
