由于我们只关心子序列中的最值，所以元素顺序不影响答案，可以先把数组排序。

排序后，假设 $\textit{nums}[i]$ 是子序列的最大值，这样的子序列有多少个？

我们可以从下标 $[0,i-1]$ 中选至多 $\min(k-1,i)$ 个数，作为子序列的其他元素。这样的选法总共有

$$
\sum_{j=0}^{\min(k-1,i)} \binom i j
$$

个。

由于 $k$ 很小，直接暴力计算上式。

于是 $\textit{nums}[i]$ 对答案的贡献为

$$
\textit{nums}[i] \cdot \sum_{j=0}^{\min(k-1,i)} \binom i j
$$

同理可以枚举 $\textit{nums}[i]$ 作为子序列的最小值，从右边选数字，做法同上。

**技巧**：根据对称性，$\textit{nums}[n-1-i]$ 作为最小值时，子序列个数和 $\textit{nums}[i]$ 作为最大值的个数是一样的，所以二者可以一同计算。

## 代码实现

1. 关于组合数，我们需要预处理阶乘及其逆元，然后利用公式 $C(n,m) = \dfrac{n!}{m!(n-m)!}$ 计算。
2. 关于逆元的知识点，见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)，包含费马小定理的数学证明。

具体请看 [视频讲解](https://www.bilibili.com/video/BV17RwBeqErJ/?t=7m32s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
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
    def minMaxSums(self, nums: List[int], k: int) -> int:
        nums.sort()
        ans = 0
        for i, x in enumerate(nums):
            s = sum(comb(i, j) for j in range(min(k, i + 1))) % MOD
            ans += (x + nums[-1 - i]) * s
        return ans % MOD
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

    public int minMaxSums(int[] nums, int k) {
        Arrays.sort(nums);
        int n = nums.length;
        long ans = 0;
        for (int i = 0; i < n; i++) {
            long s = 0;
            for (int j = 0; j < Math.min(k, i + 1); j++) {
                s += comb(i, j);
            }
            ans = (ans + s % MOD * (nums[i] + nums[n - 1 - i])) % MOD;
        }
        return (int) ans;
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
    int minMaxSums(vector<int>& nums, int k) {
        ranges::sort(nums);
        int n = nums.size();
        long long ans = 0;
        for (int i = 0; i < n; i++) {
            long long s = 0;
            for (int j = 0; j < min(k, i + 1); j++) {
                s += comb(i, j);
            }
            ans = (ans + s % MOD * (nums[i] + nums[n - 1 - i])) % MOD;
        }
        return ans;
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

func minMaxSums(nums []int, k int) (ans int) {
	slices.Sort(nums)
	for i, x := range nums {
		s := 0
		for j := range min(k, i+1) {
			s += comb(i, j)
		}
		ans = (ans + s%mod*(x+nums[len(nums)-1-i])) % mod
	}
	return
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(n\log n + nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 优化

对于和式

$$
s_i = \sum_{j=0}^{\min(k-1,i)} \binom i j
$$

考虑怎么通过 $s_i$ 递推得到 $s_{i+1}$：

- 如果 $i< k-1$，那么我们相当于在 $s_i$，也就是包含至多 $i$ 个数的子序列的基础上，多了一个数。对于 $s_i$ 中的每个子序列，这个数都可以选或不选，所以有
  $$
  s_{i+1} = 2\cdot s_i
  $$
- 如果 $i \ge k-1$，那么我们相当于在 $s_i$，也就是包含至多 $k-1$ 个数的子序列的基础上，多了一个数。这个数可以选也可以不选，如果选，那么前面的子序列不能包含恰好 $k-1$ 个数，要减掉，也就是减去从 $i$ 个数中选出 $k-1$ 个数的方案数，所以有
  $$
  s_{i+1} = 2\cdot s_i - \binom {i} {k-1}
  $$
  
这样可以 $\mathcal{O}(1)$ 递推和式，不需要 $\mathcal{O}(k)$ 暴力去算。

初始值 $s_0 = 1$。

代码实现时，$s_i$ 可以优化成一个变量。

```py [sol-Python3]
# 更快的写法见【预处理】
class Solution:
    def minMaxSums(self, nums: List[int], k: int) -> int:
        MOD = 1_000_000_007
        nums.sort()
        ans = 0
        s = 1
        for i, x in enumerate(nums):
            ans += (x + nums[-1 - i]) * s
            s = (s * 2 - comb(i, k - 1)) % MOD
        return ans % MOD
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
    return fac[n] * inv_f[m] * inv_f[n - m] % MOD if m <= n else 0

class Solution:
    def minMaxSums(self, nums: List[int], k: int) -> int:
        nums.sort()
        ans = 0
        s = 1
        for i, x in enumerate(nums):
            ans += (x + nums[-1 - i]) * s
            s = (s * 2 - comb(i, k - 1)) % MOD
        return ans % MOD
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

    public int minMaxSums(int[] nums, int k) {
        Arrays.sort(nums);
        int n = nums.length;
        long ans = 0;
        long s = 1;
        for (int i = 0; i < n; i++) {
            ans = (ans + s * (nums[i] + nums[n - 1 - i])) % MOD;
            s = (s * 2 - comb(i, k - 1) + MOD) % MOD;
        }
        return (int) ans;
    }

    private long comb(int n, int m) {
        return m > n ? 0 : F[n] * INV_F[m] % MOD * INV_F[n - m] % MOD;
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
    return m > n ? 0 : F[n] * INV_F[m] % MOD * INV_F[n - m] % MOD;
}

class Solution {
public:
    int minMaxSums(vector<int>& nums, int k) {
        ranges::sort(nums);
        int n = nums.size();
        long long ans = 0, s = 1;
        for (int i = 0; i < n; i++) {
            ans = (ans + s * (nums[i] + nums[n - 1 - i])) % MOD;
            s = (s * 2 - comb(i, k - 1) + MOD) % MOD;
        }
        return ans;
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
	if m > n {
		return 0
	}
	return f[n] * invF[m] % mod * invF[n-m] % mod
}

func minMaxSums(nums []int, k int) (ans int) {
	slices.Sort(nums)
	s := 1
	for i, x := range nums {
		ans = (ans + s*(x+nums[len(nums)-1-i])) % mod
		s = (s*2 - comb(i, k-1) + mod) % mod
	}
	return
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

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
