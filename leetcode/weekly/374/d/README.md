首先考虑两个相邻感冒小朋友之间的序列有多少个。

设这两个相邻感冒小朋友之间有 $k$ 个没有感冒的小朋友。

这 $k$ 个小朋友中，如果是被左边的人传染的，记作 L；如果是被右边的人传染的，记作 R。

按照传染感冒的顺序，我们可以得到一个 LR 序列。问题相当于求 LR 序列有多少个，这相当于从 $k$ 个位置中选一个子集，全部填 L，其余位置填 R。

但是考虑到最后感冒的那个小朋友既可以记作 L，又可以记作 R，所以有 

$$
2^{k-1}
$$

种方案。

特殊情况：

- 对于最左边感冒小朋友传染他左边的人，序列只有一种。
- 对于最右边感冒小朋友传染他右边的人，序列只有一种。

然后考虑不同的感冒序列之间如何「合并」。

## 思路一：组合数

假设有三个感冒序列，长度分别为 $k_1,k_2,k_3$，长度之和为 $s$。

- 先从 $s$ 个位置中选 $k_1$ 个位置放第一个感冒序列，这有 $C(s,k_1)$ 种放法。
- 然后从 $s-k_1$ 个位置中选 $k_2$ 个位置放第二个感冒序列，这有 $C(s-k_1,k_2)$ 种放法。
- 然后从 $s-k_1-k_2$ 个位置中选 $k_3$ 个位置放第三个感冒序列，这有 $C(s-k_1-k_2,k_3)$ 种放法。

根据乘法原理，把所有放法相乘，再乘上每种感冒序列的方案，即为答案。

代码实现时，所有的 $2^{k-1}$ 要全部乘起来，可以只记录指数，最后再计算 $2$ 的幂。

[本题视频讲解](https://www.bilibili.com/video/BV1og4y1Z7SZ/)

```py [sol-Python3]
MOD = 1_000_000_007
MX = 100_000

# 组合数模板
fac = [0] * MX
fac[0] = 1
for i in range(1, MX):
    fac[i] = fac[i - 1] * i % MOD

inv_fac = [0] * MX
inv_fac[MX - 1] = pow(fac[MX - 1], -1, MOD)
for i in range(MX - 1, 0, -1):
    inv_fac[i - 1] = inv_fac[i] * i % MOD

def comb(n: int, k: int) -> int:
    return fac[n] * inv_fac[k] % MOD * inv_fac[n - k] % MOD

class Solution:
    def numberOfSequence(self, n: int, a: List[int]) -> int:
        m = len(a)
        total = n - m
        ans = comb(total, a[0]) * comb(total - a[0], n - a[-1] - 1) % MOD
        total -= a[0] + n - a[-1] - 1
        e = 0
        for p, q in pairwise(a):
            k = q - p - 1
            if k:
                e += k - 1
                ans = ans * comb(total, k) % MOD
                total -= k
        return ans * pow(2, e, MOD) % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 100_000;

    // 组合数模板
    private static final long[] FAC = new long[MX];
    private static final long[] INV_FAC = new long[MX];

    static {
        FAC[0] = 1;
        for (int i = 1; i < MX; i++) {
            FAC[i] = FAC[i - 1] * i % MOD;
        }
        INV_FAC[MX - 1] = pow(FAC[MX - 1], MOD - 2);
        for (int i = MX - 1; i > 0; i--) {
            INV_FAC[i - 1] = INV_FAC[i] * i % MOD;
        }
    }

    private static long comb(int n, int k) {
        return FAC[n] * INV_FAC[k] % MOD * INV_FAC[n - k] % MOD;
    }

    public int numberOfSequence(int n, int[] a) {
        int m = a.length;
        int total = n - m;
        long ans = comb(total, a[0]) * comb(total - a[0], n - a[m - 1] - 1) % MOD;
        total -= a[0] + n - a[m - 1] - 1;
        int e = 0;
        for (int i = 1; i < m; i++) {
            int k = a[i] - a[i - 1] - 1;
            if (k > 0) {
                e += k - 1;
                ans = ans * comb(total, k) % MOD;
                total -= k;
            }
        }
        return (int) (ans * pow(2, e) % MOD);
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

long long q_pow(long long x, int n) {
    long long res = 1;
    for (; n > 0; n /= 2) {
        if (n % 2) {
            res = res * x % MOD;
        }
        x = x * x % MOD;
    }
    return res;
}

// 组合数模板
long long fac[MX], inv_fac[MX];

auto init = [] {
    fac[0] = 1;
    for (int i = 1; i < MX; i++) {
        fac[i] = fac[i - 1] * i % MOD;
    }
    inv_fac[MX - 1] = q_pow(fac[MX - 1], MOD - 2);
    for (int i = MX - 1; i > 0; i--) {
        inv_fac[i - 1] = inv_fac[i] * i % MOD;
    }
    return 0;
}();

long long comb(int n, int k) {
    return fac[n] * inv_fac[k] % MOD * inv_fac[n - k] % MOD;
}

class Solution {
public:
    int numberOfSequence(int n, vector<int>& a) {
        int m = a.size();
        int total = n - m;
        long long ans = comb(total, a[0]) * comb(total - a[0], n - a.back() - 1) % MOD;
        total -= a[0] + n - a.back() - 1;
        int e = 0;
        for (int i = 1; i < m; i++) {
            int k = a[i] - a[i - 1] - 1;
            if (k) {
                e += k - 1;
                ans = ans * comb(total, k) % MOD;
                total -= k;
            }
        }
        return ans * q_pow(2, e) % MOD;
    }
};
```

```go [sol-Go]
// 组合数模板
const mod = 1_000_000_007
const mx = 100_000

var fac, invFac [mx]int

func init() {
	fac[0] = 1
	for i := 1; i < mx; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	invFac[mx-1] = pow(fac[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invFac[i-1] = invFac[i] * i % mod
	}
}

func comb(n, k int) int {
	return fac[n] * invFac[k] % mod * invFac[n-k] % mod
}

func numberOfSequence(n int, a []int) int {
	m := len(a)
	total := n - m
	ans := comb(total, a[0]) * comb(total-a[0], n-a[m-1]-1) % mod
	total -= a[0] + n - a[m-1] - 1
	e := 0
	for i := 1; i < m; i++ {
		k := a[i] - a[i-1] - 1
		if k > 0 {
			e += k - 1
			ans = ans * comb(total, k) % mod
			total -= k
		}
	}
	return ans * pow(2, e) % mod
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
```

## 思路二：可重集排列数

序列的总长度为 $n-m$，排列数为 

$$
(n-m)!
$$

其中每一段序列（两个感冒小朋友之间）的小朋友，我们不关心其顺序，那么设第 $i$ 段序列的长度为 $k_i$，方案数要除以 $k_i!$。

所以答案为 $2^e$ 乘以

$$
\dfrac{(n-m)!}{\prod\limits_i k_i!}
$$

```py [sol-Python3]
MOD = 1_000_000_007
MX = 100_000

fac = [0] * MX
fac[0] = 1
for i in range(1, MX):
    fac[i] = fac[i - 1] * i % MOD

inv_fac = [0] * MX
inv_fac[MX - 1] = pow(fac[MX - 1], -1, MOD)
for i in range(MX - 1, 0, -1):
    inv_fac[i - 1] = inv_fac[i] * i % MOD

class Solution:
    def numberOfSequence(self, n: int, a: List[int]) -> int:
        ans = fac[n - len(a)] * inv_fac[a[0]] * inv_fac[n - 1 - a[-1]] % MOD
        e = 0
        for p, q in pairwise(a):
            k = q - p - 1
            if k > 1:
                e += k - 1
                ans = ans * inv_fac[k] % MOD
        return ans * pow(2, e, MOD) % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 100_000;

    private static final long[] FAC = new long[MX];
    private static final long[] INV_FAC = new long[MX];

    static {
        FAC[0] = 1;
        for (int i = 1; i < MX; i++) {
            FAC[i] = FAC[i - 1] * i % MOD;
        }
        INV_FAC[MX - 1] = pow(FAC[MX - 1], MOD - 2);
        for (int i = MX - 1; i > 0; i--) {
            INV_FAC[i - 1] = INV_FAC[i] * i % MOD;
        }
    }

    public int numberOfSequence(int n, int[] a) {
        int m = a.length;
        long ans = FAC[n - m] * INV_FAC[a[0]] % MOD * INV_FAC[n - 1 - a[m - 1]] % MOD;
        int e = 0;
        for (int i = 1; i < m; i++) {
            int k = a[i] - a[i - 1] - 1;
            if (k > 1) {
                e += k - 1;
                ans = ans * INV_FAC[k] % MOD;
            }
        }
        return (int) (ans * pow(2, e) % MOD);
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

long long q_pow(long long x, int n) {
    long long res = 1;
    for (; n > 0; n /= 2) {
        if (n % 2) {
            res = res * x % MOD;
        }
        x = x * x % MOD;
    }
    return res;
}

long long fac[MX], inv_fac[MX];

auto init = [] {
    fac[0] = 1;
    for (int i = 1; i < MX; i++) {
        fac[i] = fac[i - 1] * i % MOD;
    }
    inv_fac[MX - 1] = q_pow(fac[MX - 1], MOD - 2);
    for (int i = MX - 1; i > 0; i--) {
        inv_fac[i - 1] = inv_fac[i] * i % MOD;
    }
    return 0;
}();

class Solution {
public:
    int numberOfSequence(int n, vector<int>& a) {
        int m = a.size();
        long ans = fac[n - m] * inv_fac[a[0]] % MOD * inv_fac[n - 1 - a[m - 1]] % MOD;
        int e = 0;
        for (int i = 0; i < m - 1; i++) {
            int k = a[i + 1] - a[i] - 1;
            if (k > 1) {
                e += k - 1;
                ans = ans * inv_fac[k] % MOD;
            }
        }
        return ans * q_pow(2, e) % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 100_000

var fac, invFac [mx]int

func init() {
	fac[0] = 1
	for i := 1; i < mx; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	invFac[mx-1] = pow(fac[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invFac[i-1] = invFac[i] * i % mod
	}
}

func numberOfSequence(n int, a []int) int {
	m := len(a)
	ans := fac[n-m] * invFac[a[0]] % mod * invFac[n-1-a[m-1]] % mod
	e := 0
	for i := 1; i < m; i++ {
		k := a[i] - a[i-1] - 1
		if k > 0 {
			e += k - 1
			ans = ans * invFac[k] % mod
		}
	}
	return ans * pow(2, e) % mod
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
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m)$，其中 $m$ 为 $\textit{sick}$ 的长度。预处理的时间忽略不计。计算 $2^e$ 的时间忽略不计（预处理的话可以做到 $\mathcal{O}(1)$）。
- 空间复杂度：$\mathcal{O}(1)$。预处理的空间忽略不计。

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
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
