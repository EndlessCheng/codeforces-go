## 题意

从 $[0,n-1]$ 中选 $m$ 个下标 $I_0,I_1,\ldots,I_{m-1}$，要求

$$
\sum_{j=0}^{m-1} 2^{I_{j}}
$$

的二进制中恰好有 $k$ 个 $1$。

定义

$$
f(I) = \prod_{j=0}^{m-1} \textit{nums}[I_j]
$$

计算所有 $f(I)$ 的总和。

## 式子变形

首先把式子简化，然后再确定思考方向。

假设选了 $j$ 个下标 $i$，对答案的贡献是多少？

分两部分：

1. 对于所有包含 $j$ 个下标 $i$ 的 $I$，贡献为 $\textit{nums}[i]^j$，这个值乘以其他下标的贡献。
2. 总共有 $m!$ 个不同的下标排列，其中有重复的，根据**可重集排列数**，要除以 $j!$。

例如 $n=2$，$m=7$，选了 $3$ 个下标 $0$ 和 $4$ 个下标 $1$，那么：

1. 第一部分的贡献为 $\textit{nums}[0]^3\cdot \textit{nums}[1]^4$。
2. 第二部分的贡献为 $\dfrac{7!}{3!4!}$。

这种选法（$3$ 个下标 $0$ 和 $4$ 个下标 $1$）如果合法（二进制恰好有 $k$ 个 $1$），那么对答案的贡献为

$$
\dfrac{7!}{3!4!}\cdot \textit{nums}[0]^3\cdot \textit{nums}[1]^4
$$

把 $7!$ 提出来，变形得

$$
7!\cdot \left(\dfrac{\textit{nums}[0]^3}{3!} \right)\cdot \left(\dfrac{\textit{nums}[1]^4}{4!} \right)
$$

这样每一项就可以独立计算了。

一般地，设我们分别选了 $c_0,c_1,c_2,\ldots,c_{n-1}$ 个下标 $0,1,2,\ldots,n-1$，那么

$$
f(I) = \prod_{j=0}^{m-1} \textit{nums}[I_j] = \prod_{i=0}^{n-1} \textit{nums}[i]^{c_i}
$$

枚举 $c_i$，把 $m!$ 提出来，答案为

$$
\sum_{I}f(I) =  \sum_{I}\prod_{i=0}^{n-1} \textit{nums}[i]^{c_i} = m! \sum_{c_0+\cdots+c_{n-1} = m} \prod_{i=0}^{n-1} \dfrac{\textit{nums}[i]^{c_i}}{c_i!}
$$

现在还剩下一个难点，如何判断 $\displaystyle\sum\limits_{j=0}^{m-1} 2^{I_{j}} = \displaystyle\sum\limits_{i=0}^{n-1} c_i 2^i$ 的二进制中恰好有 $k$ 个 $1$？

## 关键思路

从小到大枚举下标 $i=0,1,2,\ldots,n-1$，在计算 $S=\displaystyle\sum\limits_{i=0}^{n-1} c_i 2^i$ 的过程中，比如现在枚举的下标 $i=6$，那么后续加到 $S$ 中的数一定 $\ge 2^6$，一定不会影响 $S$ 的小于 $i$ 的比特位，我们可以**提前统计这些比特位中的 $1$**！

换句话说，在递归过程中只需保存 $S$ 右移 $i$ 位的结果，而不是原始的 $S$，从而大幅减少状态个数！

## 思路

我们需要知道如下信息：

- 当前枚举的下标 $i$。
- 还剩下 $\textit{leftM}$ 个下标需要选。
- 当前累加的 $S$，其右移 $i$ 位的结果是 $x$。即 $\left\lfloor\dfrac{S}{2^i}\right\rfloor = x$。
- 去掉右移掉的 $1$ 后，$S$ 还需包含恰好 $\textit{leftK}$ 个 $1$。

定义 $\textit{dfs}(i,\textit{leftM},x,\textit{leftK})$ 表示在上述情况下，剩余元素的贡献。

枚举选 $j=0,1,2,\ldots,\textit{leftM}$ 个下标 $i$，接下来要解决的问题是：

- 当前枚举的下标是 $i+1$。
- 还剩下 $\textit{leftM}-j$ 个下标需要选。
- 当前累加的 $S' = S + j\cdot 2^i$，其右移 $i+1$ 位的结果是 $\left\lfloor\dfrac{S'}{2^{i+1}}\right\rfloor =  \left\lfloor\dfrac{x+j}{2}\right\rfloor$。
- 去掉右移掉的 $1$ 后，$S$ 还需包含恰好 $\textit{leftK}-\textit{bit}$ 个 $1$，其中 $\textit{bit}=(x+j)\bmod 2$。

如果 $\textit{bit}\le \textit{leftK}$，那么可以递归到 $r=\textit{dfs}(i+1, \textit{leftM}-j,\left\lfloor\tfrac{x+j}{2}\right\rfloor,\textit{leftK}-\textit{bit})$。把 $r$ 乘以 $\dfrac{\textit{nums}[i]^j}{j!}$，累加到 $\textit{dfs}(i,\textit{leftM},x,\textit{leftK})$ 中，得

$$
\textit{dfs}(i,\textit{leftM},x,\textit{leftK}) = \sum_{j=0}^{\textit{leftM}} \textit{dfs}(i+1, \textit{leftM}-j,\left\lfloor\tfrac{x+j}{2}\right\rfloor,\textit{leftK}-\textit{bit})\cdot \dfrac{\textit{nums}[i]^j}{j!}
$$

其中枚举的 $j$ 还需要满足 $\textit{bit}\le \textit{leftK}$。

递归边界：$i=n$ 时，如果 $\textit{leftM}=0$ 且 $x$ 的二进制中恰好有 $\textit{leftK}$ 个 $1$，那么找到了一个合法序列，返回 $1$，否则返回 $0$。此外，如果 $\textit{leftM}=0$ 或者 $\textit{leftK}=0$，那么后续无法选任何数，可以直接返回，处理方式同 $i=n$。

递归入口：$\textit{dfs}(0,m,0,k)\cdot m!$，即答案。

代码实现时，预处理 $\textit{powV}[i][j] = \textit{nums}[i]^j$，从而加速计算过程。

代码实现时，预处理阶乘及其逆元，从而加速计算过程，原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

## 写法一：记忆化搜索

原理见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

[本题视频讲解](https://www.bilibili.com/video/BV1avVwz5EbY/?t=32m38s)，欢迎点赞关注~

**可行性剪枝**：设 $x$ 的二进制中有 $c_1$ 个 $1$，无论后面怎么选下标，在 $x$ 的基础上，至多增加 $\textit{leftM}$ 个 $1$。如果此时 $c_1+\textit{leftM} < \textit{leftK}$，那么这种状态一定无法得到合法序列，直接返回 $0$。

```py [sol-Python3]
MOD = 1_000_000_007
MX = 31

fac = [0] * MX  # fac[i] = i!
fac[0] = 1
for i in range(1, MX):
    fac[i] = fac[i - 1] * i % MOD

inv_f = [0] * MX  # inv_f[i] = i!^-1
inv_f[-1] = pow(fac[-1], -1, MOD)
for i in range(MX - 1, 0, -1):
    inv_f[i - 1] = inv_f[i] * i % MOD

class Solution:
    def magicalSum(self, m: int, k: int, nums: List[int]) -> int:
        n = len(nums)
        pow_v = [[1] * (m + 1) for _ in range(n)]
        for i, v in enumerate(nums):
            for j in range(1, m + 1):
                pow_v[i][j] = pow_v[i][j - 1] * v % MOD

        @cache
        def dfs(i: int, left_m: int, x: int, left_k: int) -> int:
            c1 = x.bit_count()
            if c1 + left_m < left_k:  # 可行性剪枝
                return 0
            if i == n or left_m == 0 or left_k == 0:  # 无法继续选数字
                return 1 if left_m == 0 and c1 == left_k else 0
            res = 0
            for j in range(left_m + 1):  # 枚举 I 中有 j 个下标 i
                # 这 j 个下标 i 对 S 的贡献是 j * pow(2, i)
                # 由于 x = S >> i，转化成对 x 的贡献是 j
                bit = (x + j) & 1  # 取最低位，提前从 left_k 中减去，其余进位到 x 中
                r = dfs(i + 1, left_m - j, (x + j) >> 1, left_k - bit)
                res += r * pow_v[i][j] * inv_f[j]
            return res % MOD

        return dfs(0, m, 0, k) * fac[m] % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 31;

    private static final long[] F = new long[MX]; // F[i] = i!
    private static final long[] INV_F = new long[MX]; // INV_F[i] = i!^-1

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

    public int magicalSum(int m, int k, int[] nums) {
        int n = nums.length;
        int[][] powV = new int[n][m + 1];
        for (int i = 0; i < n; i++) {
            powV[i][0] = 1;
            for (int j = 1; j <= m; j++) {
                powV[i][j] = (int) ((long) powV[i][j - 1] * nums[i] % MOD);
            }
        }

        int[][][][] memo = new int[n][m + 1][m / 2 + 1][k + 1];
        for (int[][][] a : memo) {
            for (int[][] b : a) {
                for (int[] c : b) {
                    Arrays.fill(c, -1);
                }
            }
        }
        return (int) (dfs(0, m, 0, k, powV, memo) * F[m] % MOD);
    }

    private long dfs(int i, int leftM, int x, int leftK, int[][] powV, int[][][][] memo) {
        int c1 = Integer.bitCount(x);
        if (c1 + leftM < leftK) { // 可行性剪枝
            return 0;
        }
        if (i == powV.length || leftM == 0 || leftK == 0) { // 无法继续选数字
            return leftM == 0 && c1 == leftK ? 1 : 0;
        }
        if (memo[i][leftM][x][leftK] != -1) {
            return memo[i][leftM][x][leftK];
        }
        long res = 0;
        for (int j = 0; j <= leftM; j++) { // 枚举 I 中有 j 个下标 i
            // 这 j 个下标 i 对 S 的贡献是 j * pow(2, i)
            // 由于 x = S >> i，转化成对 x 的贡献是 j
            int bit = (x + j) & 1; // 取最低位，提前从 leftK 中减去，其余进位到 x 中
            long r = dfs(i + 1, leftM - j, (x + j) >> 1, leftK - bit, powV, memo);
            res = (res + r * powV[i][j] % MOD * INV_F[j]) % MOD;
        }
        return memo[i][leftM][x][leftK] = (int) res;
    }
}
```

```cpp [sol-C++]
const int MOD = 1'000'000'007;
const int MX = 31;

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

class Solution {
public:
    int magicalSum(int m, int k, vector<int>& nums) {
        int n = nums.size();
        vector pow_v(n, vector<int>(m + 1));
        for (int i = 0; i < n; i++) {
            pow_v[i][0] = 1;
            for (int j = 1; j <= m; j++) {
                pow_v[i][j] = 1LL * pow_v[i][j - 1] * nums[i] % MOD;
            }
        }

        vector memo(n, vector(m + 1, vector(m / 2 + 1, vector<int>(k + 1, -1))));
        auto dfs = [&](this auto&& dfs, int i, int left_m, int x, int left_k) -> int {
            int c1 = popcount((uint32_t) x);
            if (c1 + left_m < left_k) { // 可行性剪枝
                return 0;
            }
            if (i == n || left_m == 0 || left_k == 0) { // 无法继续选数字
                return left_m == 0 && c1 == left_k;
            }
            int& res = memo[i][left_m][x][left_k]; // 注意这里是引用
            if (res != -1) {
                return res;
            }
            res = 0;
            for (int j = 0; j <= left_m; j++) { // 枚举 I 中有 j 个下标 i
                // 这 j 个下标 i 对 S 的贡献是 j * pow(2, i)
                // 由于 x = S >> i，转化成对 x 的贡献是 j
                int bit = (x + j) & 1; // 取最低位，提前从 left_k 中减去，其余进位到 x 中
                int r = dfs(i + 1, left_m - j, (x + j) >> 1, left_k - bit);
                res = (res + 1LL * r * pow_v[i][j] % MOD * INV_F[j]) % MOD;
            }
            return res;
        };
        return 1LL * dfs(0, m, 0, k) * F[m] % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 31

var fac [mx]int  // fac[i] = i!
var invF [mx]int // invF[i] = i!^-1

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

func magicalSum(m, k int, nums []int) int {
	n := len(nums)
	powV := make([][]int, n)
	for i, v := range nums {
		powV[i] = make([]int, m+1)
		powV[i][0] = 1
		for j := 1; j <= m; j++ {
			powV[i][j] = powV[i][j-1] * v % mod
		}
	}

	memo := make([][][][]int, n)
	for i := range memo {
		memo[i] = make([][][]int, m+1)
		for j := range memo[i] {
			memo[i][j] = make([][]int, m/2+1)
			for p := range memo[i][j] {
				memo[i][j][p] = make([]int, k+1)
				for q := range memo[i][j][p] {
					memo[i][j][p][q] = -1
				}
			}
		}
	}
	var dfs func(int, int, int, int) int
	dfs = func(i, leftM, x, leftK int) (res int) {
		c1 := bits.OnesCount(uint(x))
		if c1+leftM < leftK { // 可行性剪枝
			return
		}
		if i == n || leftM == 0 || leftK == 0 { // 无法继续选数字
			if leftM == 0 && c1 == leftK {
				return 1
			}
			return
		}
		p := &memo[i][leftM][x][leftK]
		if *p != -1 {
			return *p
		}
		for j := range leftM + 1 { // 枚举 I 中有 j 个下标 i
			// 这 j 个下标 i 对 S 的贡献是 j * pow(2, i)
			// 由于 x = S >> i，转化成对 x 的贡献是 j
			bit := (x + j) & 1 // 取最低位，提前从 leftK 中减去，其余进位到 x 中
			r := dfs(i+1, leftM-j, (x+j)>>1, leftK-bit)
			res = (res + r*powV[i][j]%mod*invF[j]) % mod
		}
		*p = res
		return
	}
	return dfs(0, m, 0, k) * fac[m] % mod
}
```

## 写法二：递推

把记忆化搜索 1:1 翻译成递推。

DP 数组的初始值怎么写？就是记忆化搜索的递归边界。

```py [sol-Python3]
MOD = 1_000_000_007
MX = 31

fac = [0] * MX  # fac[i] = i!
fac[0] = 1
for i in range(1, MX):
    fac[i] = fac[i - 1] * i % MOD

inv_f = [0] * MX  # inv_f[i] = i!^-1
inv_f[-1] = pow(fac[-1], -1, MOD)
for i in range(MX - 1, 0, -1):
    inv_f[i - 1] = inv_f[i] * i % MOD

class Solution:
    def magicalSum(self, m: int, k: int, nums: List[int]) -> int:
        n = len(nums)
        pow_v = [[1] * (m + 1) for _ in range(n)]
        for i, v in enumerate(nums):
            for j in range(1, m + 1):
                pow_v[i][j] = pow_v[i][j - 1] * v % MOD

        f = [[[[0] * (k + 1) for _ in range(m // 2 + 1)] for _ in range(m + 1)] for _ in range(n + 1)]
        for x in range(m // 2 + 1):
            c1 = x.bit_count()
            if c1 <= k:
                f[n][0][x][c1] = 1

        for i in range(n - 1, -1, -1):
            for left_m in range(m + 1):
                for x in range(m // 2 + 1):
                    for left_k in range(k + 1):
                        res = 0
                        for j in range(min(left_m, m - x) + 1):
                            bit = (x + j) & 1
                            if bit <= left_k:
                                r = f[i + 1][left_m - j][(x + j) >> 1][left_k - bit]
                                res += r * pow_v[i][j] * inv_f[j]
                        f[i][left_m][x][left_k] = res % MOD
        return f[0][m][0][k] * fac[m] % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 31;

    private static final long[] F = new long[MX]; // F[i] = i!
    private static final long[] INV_F = new long[MX]; // INV_F[i] = i!^-1

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

    public int magicalSum(int m, int k, int[] nums) {
        int n = nums.length;
        int[][] powV = new int[n][m + 1];
        for (int i = 0; i < n; i++) {
            powV[i][0] = 1;
            for (int j = 1; j <= m; j++) {
                powV[i][j] = (int) ((long) powV[i][j - 1] * nums[i] % MOD);
            }
        }

        int[][][][] f = new int[n + 1][m + 1][m / 2 + 1][k + 1];
        for (int x = 0; x <= m / 2; x++) {
            int c1 = Integer.bitCount(x);
            if (c1 <= k) {
                f[n][0][x][c1] = 1;
            }
        }

        for (int i = n - 1; i >= 0; i--) {
            for (int leftM = 0; leftM <= m; leftM++) {
                for (int x = 0; x <= m / 2; x++) {
                    for (int leftK = 0; leftK <= k; leftK++) {
                        long res = 0;
                        for (int j = 0; j <= Math.min(leftM, m - x); j++) {
                            int bit = (x + j) & 1;
                            if (bit <= leftK) {
                                long r = f[i + 1][leftM - j][(x + j) >> 1][leftK - bit];
                                res = (res + r * powV[i][j] % MOD * INV_F[j]) % MOD;
                            }
                        }
                        f[i][leftM][x][leftK] = (int) res;
                    }
                }
            }
        }
        return (int) ((long) f[0][m][0][k] * F[m] % MOD);
    }
}
```

```cpp [sol-C++]
const int MOD = 1'000'000'007;
const int MX = 31;

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

class Solution {
public:
    int magicalSum(int m, int k, vector<int>& nums) {
        int n = nums.size();
        vector pow_v(n, vector<int>(m + 1));
        for (int i = 0; i < n; i++) {
            pow_v[i][0] = 1;
            for (int j = 1; j <= m; j++) {
                pow_v[i][j] = 1LL * pow_v[i][j - 1] * nums[i] % MOD;
            }
        }

        vector f(n + 1, vector(m + 1, vector(m / 2 + 1, vector<int>(k + 1))));
        for (uint32_t x = 0; x <= m / 2; x++) {
            int c1 = popcount(x);
            if (c1 <= k) {
                f[n][0][x][c1] = 1;
            }
        }

        for (int i = n - 1; i >= 0; i--) {
            for (int left_m = 0; left_m <= m; left_m++) {
                for (int x = 0; x <= m / 2; x++) {
                    for (int left_k = 0; left_k <= k; left_k++) {
                        long long res = 0;
                        for (int j = 0; j <= min(left_m, m - x); j++) {
                            int bit = (x + j) & 1;
                            if (bit <= left_k) {
                                int r = f[i + 1][left_m - j][(x + j) >> 1][left_k - bit];
                                res = (res + 1LL * r * pow_v[i][j] % MOD * INV_F[j]) % MOD;
                            }
                        }
                        f[i][left_m][x][left_k] = res;
                    }
                }
            }
        }
        return 1LL * f[0][m][0][k] * F[m] % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 31

var fac [mx]int  // fac[i] = i!
var invF [mx]int // invF[i] = i!^-1

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

func magicalSum(m, k int, nums []int) int {
	n := len(nums)
	powV := make([][]int, n)
	for i, v := range nums {
		powV[i] = make([]int, m+1)
		powV[i][0] = 1
		for j := 1; j <= m; j++ {
			powV[i][j] = powV[i][j-1] * v % mod
		}
	}

	f := make([][][][]int, n+1)
	for i := range f {
		f[i] = make([][][]int, m+1)
		for j := range f[i] {
			f[i][j] = make([][]int, m/2+1)
			for x := range f[i][j] {
				f[i][j][x] = make([]int, k+1)
			}
		}
	}
	for x := range m/2 + 1 {
		c1 := bits.OnesCount(uint(x))
		if c1 <= k {
			f[n][0][x][c1] = 1
		}
	}

	for i := n - 1; i >= 0; i-- {
		for leftM := range m + 1 {
			for x := range m/2 + 1 {
				for leftK := range k + 1 {
					res := 0
					for j := range min(leftM, m-x) + 1 {
						bit := (x + j) & 1
						if bit <= leftK {
							r := f[i+1][leftM-j][(x+j)>>1][leftK-bit]
							res = (res + r*powV[i][j]%mod*invF[j]) % mod
						}
					}
					f[i][leftM][x][leftK] = res
				}
			}
		}
	}
	return f[0][m][0][k] * fac[m] % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm^3k)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(nm^2k)$。**注**：使用滚动数组可以优化至 $\mathcal{O}(nm + m^2k)$。其中 $\mathcal{O}(nm)$ 是 $\textit{powV}$ 的空间。

## 相似题目

[3343. 统计平衡排列的数目](https://leetcode.cn/problems/count-number-of-balanced-permutations/)

更多相似题目，见下面动态规划题单的「**§7.6 多维 DP**」和数学题单的「**§2.2 组合计数**」。

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
