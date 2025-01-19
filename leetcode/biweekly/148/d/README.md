本质上，答案是一堆 $|x_i-x_j|+|y_i-y_j|$ 之和。拆分成 $|x_i-x_j|$ 之和、$|y_i-y_j|$ 之和。

看看示例 2 是怎么算的。答案包含 $|1-0|$，$|3-1|$ 等式子。想一想，$|1-0|$ 在答案中出现了多少次？

![3426.png](https://pic.leetcode.cn/1737249404-HvZLla-3426.png)

出现了 $2$ 次（图一和图二）。为什么？因为当我们在前两个位置放置棋子后，剩余的两个位置还可以放剩余的一个棋子，方案数为组合数 $\binom {4-2} {3-2} = \binom 2 1 = 2$。

继续，有这样几类式子：

- $|1-0|,|2-1|,|3-2|$，结果为 $1$。
- $|2-0|,|3-1|$，结果为 $2$。
- $|3-0|$，结果为 $3$。

每个式子在答案中的出现次数都是 $2$，所以示例 2 的答案为

$$
(1\cdot 3 + 2\cdot 2 + 3\cdot 1)\cdot \binom 2 1 = 20
$$

一般地，如果 $m=1$，那么绝对差为 $d$ 的式子 $|x_i-x_j|$ 有 $n-d$ 种，每个式子的出现次数为 $\binom {mn-2} {k-2}$，表示在其余 $mn-2$ 个位置中选择 $k-2$ 个位置放置剩余的 $k-2$ 个棋子。

所有 $|x_i-x_j|$ 之和为

$$
\binom {mn-2} {k-2} \sum_{d=1}^{n-1} d\cdot(n-d)
$$

推广到 $m$ 为任意数的情况。由于两个棋子处于同一列的情况下 $|x_i-x_j|=0$，所以只考虑两个棋子不同列的情况，那么每个棋子都可以在 $m$ 行中任选一行放置，所以上式要额外乘以 $m^2$，即

$$
\binom {mn-2} {k-2} m^2 \sum_{d=1}^{n-1} d\cdot(n-d)
$$

这就是所有 $|x_i-x_j|$ 之和。

同理，所有 $|y_i-y_j|$ 之和为

$$
\binom {mn-2} {k-2} n^2 \sum_{d=1}^{m-1} d\cdot(m-d)
$$

进一步地，

$$
\begin{aligned}
    & \sum_{d=1}^{n-1} d\cdot(n-d)      \\
={} & \sum_{d=1}^{n-1} (nd-d^2)        \\
={} & n\sum_{d=1}^{n-1} d- \sum_{d=1}^{n-1} d^2        \\
={} & n\cdot \dfrac{n(n-1)}{2}- \dfrac{n(n-1)(2n-1)}{6}        \\
={} & \dfrac{(n+1)n(n-1)}{6}        \\
={} & \binom {n+1} 3        \\
\end{aligned}
$$

所以最终答案为

$$
\binom {mn-2} {k-2} \left(m^2 \binom {n+1} 3 + n^2\binom {m+1} 3\right)
$$

## 代码实现

1. 关于组合数，我们需要预处理阶乘及其逆元，然后利用公式 $C(n,m) = \dfrac{n!}{m!(n-m)!}$ 计算。
2. 关于逆元的知识点，见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)，包含费马小定理的数学证明。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1xBwBeEEie/?t=18m01s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def distanceSum(self, m: int, n: int, k: int) -> int:
        MOD = 1_000_000_007
        return comb(m * n - 2, k - 2) % MOD * (m * m * comb(n + 1, 3) + n * n * comb(m + 1, 3)) % MOD
```

```py [sol-Python3 预处理]
MOD = 1_000_000_007
MX = 100_000

fac = [0] * MX  # f[i] = i!
fac[0] = 1
for i in range(1, MX):
    fac[i] = fac[i - 1] * i % MOD

inv_f = [0] * MX  # inv_f[i] = i!^-1
inv_f[-1] = pow(fac[-1], -1, MOD)
for i in range(MX - 1, 0, -1):
    inv_f[i - 1] = inv_f[i] * i % MOD

def comb(n: int, m: int) -> int:
    return fac[n] * inv_f[m] * inv_f[n - m] % MOD

class Solution:
    def distanceSum(self, m: int, n: int, k: int) -> int:
        return comb(m * n - 2, k - 2) * (m * n * (m * (n * n - 1) + n * (m * m - 1))) // 6 % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 100_000;

    private static final long[] F = new long[MX]; // f[i] = i!
    private static final long[] INV_F = new long[MX]; // inv_f[i] = i!^-1

    static {
        F[0] = 1;
        for (int i = 1; i < MX; i++) {
            F[i] = F[i - 1] * i % MOD;
        }

        INV_F[MX - 1] = pow(F[MX - 1], MOD - 2);
        for (int i = MX - 1; i > 0; i--) {
            INV_F[i - 1] = INV_F[i] * i % MOD;
        }
    }

    public int distanceSum(int m, int n, int k) {
        return (int) ((m * n * (m * ((long) n * n - 1) + n * ((long) m * m - 1))) / 6 % MOD * comb(m * n - 2, k - 2) % MOD);
    }

    private long comb(int n, int m) {
        return F[n] * INV_F[m] % MOD * INV_F[n - m] % MOD;
    }

    private static long pow(long x, int n) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }
}
```

```cpp [sol-C++]
const int MOD = 1'000'000'007;
const int MX = 100'000;

long long F[MX]; // F[i] = i!
long long INV_F[MX]; // INV_F[i] = i!^-1

long long pow(long long x, int n) {
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

    INV_F[MX - 1] = pow(F[MX - 1], MOD - 2);
    for (int i = MX - 1; i; i--) {
        INV_F[i - 1] = INV_F[i] * i % MOD;
    }
    return 0;
}();

long long comb(int n, int m) {
    return F[n] * INV_F[m] % MOD * INV_F[n - m] % MOD;
}

class Solution {
public:
    int distanceSum(int m, int n, int k) {
        return (m * n * (m * (1LL * n * n - 1) + n * (1LL * m * m - 1))) / 6 % MOD * comb(m * n - 2, k - 2) % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 100_000

var f [mx]int    // f[i] = i!
var invF [mx]int // invF[i] = i!^-1

func init() {
	f[0] = 1
	for i := 1; i < mx; i++ {
		f[i] = f[i-1] * i % mod
	}

	invF[mx-1] = pow(f[mx-1], mod-2)
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

func comb(n, m int) int {
	return f[n] * invF[m] % mod * invF[n-m] % mod
}

func distanceSum(m, n, k int) int {
	return (m * n * (m*(n*n-1) + n*(m*m-1))) / 6 % mod * comb(m*n-2, k-2) % mod
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. 【本题相关】[数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
