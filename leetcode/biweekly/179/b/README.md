一旦我们确定了被 $\textit{pos}$ 看到的 $k$ 个人，那么：

- 这 $k$ 个人的方向是唯一确定的：在 $\textit{pos}$ 左边的人的方向为 $\texttt{L}$，在 $\textit{pos}$ 右边的人的方向为 $\texttt{R}$。
- 其余 $n-1-k$ 个人不可见，方向也是唯一确定的：在 $\textit{pos}$ 左边的人的方向为 $\texttt{R}$，在 $\textit{pos}$ 右边的人的方向为 $\texttt{L}$。

从 $n-1$ 个人中选 $k$ 个人有 $C(n-1,k)$ 种方案，每种方案，这 $n-1$ 个人的方向都是唯一确定的，而 $\textit{pos}$ 向左向右都可以，有 $2$ 种方案。所以一共有

$$
2\cdot C(n-1,k)
$$

种方案。

由于 $n$ 和 $k$ 都很大，需要**预处理阶乘及其逆元**，从而快速计算组合数。代码模板见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV1dxXSBAE6F/?t=4m31s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countVisiblePeople(self, n: int, pos: int, k: int) -> int:
        # 这样写很慢，预处理的写法见另一份代码【Python3 预处理】
        return comb(n - 1, k) * 2 % 1_000_000_007
```

```py [sol-Python3 预处理]
MOD = 1_000_000_007
MX = 100_001

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
    def countVisiblePeople(self, n: int, pos: int, k: int) -> int:
        # 把预处理的逻辑写在 class 外面，这样只会初始化一次
        return comb(n - 1, k) * 2 % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 100_001;

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

    public int countVisiblePeople(int n, int pos, int k) {
        return (int) (comb(n - 1, k) * 2 % MOD);
    }
}
```

```cpp [sol-C++]
const int MOD = 1'000'000'007;
const int MX = 100'001; // 根据题目数据范围修改

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
    int countVisiblePeople(int n, int, int k) {
        return comb(n - 1, k) * 2 % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 100_001

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
    if m < 0 || m > n {
        return 0
    }
    return fac[n] * invF[m] % mod * invF[n-m] % mod
}

func countVisiblePeople(n, _, k int) int {
    return comb(n-1, k) * 2 % mod
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面数学题单的「**§2.2 组合计数**」。

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
