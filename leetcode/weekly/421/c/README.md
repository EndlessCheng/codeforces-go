## 方法一：多维 DP

### 1) 寻找子问题

规定空子序列的 GCD 等于 $0$。

考虑从右往左选数字，讨论 $\textit{nums}[n-1]$ **选或不选**：

- 不选，不把 $\textit{nums}[n-1]$ 加到任何子序列中，那么需要解决的问题为：在 $\textit{nums}[0]$ 到 $\textit{nums}[n-2]$ 中选数字，且当前两个子序列的 GCD 分别为 $0,0$ 时，最终可以得到的合法子序列对的个数。
- 选，把 $\textit{nums}[n-1]$ 加到第一个子序列中，那么需要解决的问题为：在 $\textit{nums}[0]$ 到 $\textit{nums}[n-2]$ 中选数字，且当前两个子序列的 GCD 分别为 $\textit{nums}[n-1],0$ 时，最终可以得到的合法子序列对的个数。
- 选，把 $\textit{nums}[n-1]$ 加到第二个子序列中，那么需要解决的问题为：在 $\textit{nums}[0]$ 到 $\textit{nums}[n-2]$ 中选数字，且当前两个子序列的 GCD 分别为 $0,\textit{nums}[n-1]$ 时，最终可以得到的合法子序列对的个数。

由于选或不选都会把原问题变成一个**和原问题相似的、规模更小的子问题**，所以可以用**递归**解决。

> 注 1：从右往左思考，主要是为了方便把递归翻译成递推。从左往右思考也是可以的。
>
> 注 2：动态规划有「**选或不选**」和「**枚举选哪个**」两种基本思考方式。子序列相邻无关一般是「选或不选」，子序列相邻相关（例如 LIS 问题）一般是「枚举选哪个」。本题用到的是「选或不选」。

### 2) 状态定义与状态转移方程

根据上面的讨论，我们需要在递归过程中跟踪以下信息：

- $i$：当前考虑 $\textit{nums}[i]$ 选或不选。
- $j$：第一个子序列的 GCD 值。
- $k$：第二个子序列的 GCD 值。

因此，定义状态为 $\textit{dfs}(i,j,k)$，表示在 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 中选数字，且当前两个子序列的 GCD 分别为 $j,k$ 时，最终可以得到的合法子序列对的个数。

接下来，思考如何从一个状态转移到另一个状态。

讨论 $x=\textit{nums}[i]$ **选或不选**：

- 不选，不把 $x$ 加到任何子序列中，那么需要解决的问题为：在 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 中选数字，且当前两个子序列的 GCD 分别为 $j,k$ 时，最终可以得到的合法子序列对的个数，即 $\textit{dfs}(i-1,j,k)$。
- 选，把 $x$ 加到第一个子序列中，那么需要解决的问题为：在 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 中选数字，且当前两个子序列的 GCD 分别为 $\text{GCD}(j,x),k$ 时，最终可以得到的合法子序列对的个数，即 $\textit{dfs}(i-1,\text{GCD}(j,x),k)$。
- 选，把 $x$ 加到第二个子序列中，那么需要解决的问题为：在 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 中选数字，且当前两个子序列的 GCD 分别为 $j,\text{GCD}(k,x)$ 时，最终可以得到的合法子序列对的个数，即 $\textit{dfs}(i-1,j,\text{GCD}(k,x))$。

三种选法互斥，根据加法原理，三者相加得

$$
\textit{dfs}(i,j,k) = \textit{dfs}(i-1,j,k) + \textit{dfs}(i-1,\text{GCD}(j,x),k) + \textit{dfs}(i-1,j,\text{GCD}(k,x))
$$

**递归边界**：$\textit{dfs}(-1,j,j)=1,\ \textit{dfs}(-1,j,k)=0\ (j\ne k)$。

**递归入口**：$\textit{dfs}(n-1,0,0)-1$，也就是答案。减一是去掉两个子序列都为空的情况。

### 3) 递归搜索 + 保存递归返回值 = 记忆化搜索

视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,j,k)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1hn1MYhEtC/?t=6m35s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def subsequencePairCount(self, nums: List[int]) -> int:
        MOD = 1_000_000_007
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int, k: int) -> int:
            if i < 0:
                return 1 if j == k else 0
            return (dfs(i - 1, j, k) + dfs(i - 1, gcd(j, nums[i]), k) + dfs(i - 1, j, gcd(k, nums[i]))) % MOD
        return (dfs(len(nums) - 1, 0, 0) - 1) % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int subsequencePairCount(int[] nums) {
        int n = nums.length;
        int m = 0;
        for (int x : nums) {
            m = Math.max(m, x);
        }
        int[][][] memo = new int[n][m + 1][m + 1];
        for (int[][] mat : memo) {
            for (int[] row : mat) {
                Arrays.fill(row, -1); // -1 表示没有计算过
            }
        }
        return (dfs(n - 1, 0, 0, nums, memo) - 1 + MOD) % MOD; // +MOD 防止减一后变成负数
    }

    int dfs(int i, int j, int k, int[] nums, int[][][] memo) {
        if (i < 0) {
            return j == k ? 1 : 0;
        }
        if (memo[i][j][k] < 0) {
            long res = (long) dfs(i - 1, j, k, nums, memo) +
                       dfs(i - 1, gcd(j, nums[i]), k, nums, memo) +
                       dfs(i - 1, j, gcd(k, nums[i]), nums, memo);
            memo[i][j][k] = (int) (res % MOD);
        }
        return memo[i][j][k];
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int subsequencePairCount(vector<int>& nums) {
        const int MOD = 1'000'000'007;
        int n = nums.size();
        int m = ranges::max(nums);
        vector<vector<vector<int>>> memo(n, vector<vector<int>>(m + 1, vector<int>(m + 1, -1))); // -1 表示没有计算过
        auto dfs = [&](auto&& dfs, int i, int j, int k) -> int {
            if (i < 0) {
                return j == k;
            }
            int& res = memo[i][j][k]; // 注意这里是引用
            if (res < 0) {
                res = ((long long) dfs(dfs, i - 1, j, k) +
                       dfs(dfs, i - 1, gcd(j, nums[i]), k) +
                       dfs(dfs, i - 1, j, gcd(k, nums[i]))) % MOD;
            }
            return res;
        };
        return (dfs(dfs, n - 1, 0, 0) - 1 + MOD) % MOD; // +MOD 防止减一后变成负数
    }
};
```

```go [sol-Go]
func subsequencePairCount2(nums []int) int {
	const mod = 1_000_000_007
	n := len(nums)
	m := slices.Max(nums)
	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, m+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, m+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1 // -1 表示没有计算过
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i < 0 {
			if j == k {
				return 1
			}
			return 0
		}
		p := &memo[i][j][k]
		if *p < 0 {
			*p = (dfs(i-1, j, k) + dfs(i-1, gcd(j, nums[i]), k) + dfs(i-1, j, gcd(k, nums[i]))) % mod
		}
		return *p
	}
	// 减去两个子序列都是空的情况
	return (dfs(n-1, 0, 0) - 1 + mod) % mod // +mod 防止减一后变成负数
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU^2\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(nU^2)$，单个状态的计算时间为 $\mathcal{O}(\log U)$，所以总的时间复杂度为 $\mathcal{O}(nU^2\log U)$。
- 空间复杂度：$\mathcal{O}(nU^2)$。保存多少状态，就需要多少空间。

### 4) 1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i+1][j][k]$ 的定义和 $\textit{dfs}(i,j,k)$ 的定义是一样的，都表示在 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 中选数字，且当前两个子序列的 GCD 分别为 $j,k$ 时，最终可以得到的合法子序列对的个数。这里 $+1$ 是为了把 $\textit{dfs}(-1,j,k)$ 这个状态也翻译过来，这样我们可以把 $f[0][j][k]$ 作为初始值。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i+1][j][k] = f[i][j][k] + f[i][\text{GCD}(j,x)][k] + f[i][j][\text{GCD}(k,x)]
$$

初始值 $f[0][j][j]=1\ (j\ge 0)$，其余为 $0$，翻译自递归边界 $\textit{dfs}(-1,j,j)=1$。注意这里初始化 $f[0][0][0]=0$，这样最后返回答案的时候，就不需要再减一了。

答案为 $f[n][0][0]$，翻译自递归入口 $\textit{dfs}(n-1,0,0)$。

#### 答疑

**问**：为什么同样的逻辑，本题记忆化搜索比递推快一些？

**答**：记忆化搜索不一定会计算所有状态，而递推会把所有状态都算一遍。

```py [sol-Python3]
class Solution:
    def subsequencePairCount(self, nums: List[int]) -> int:
        MOD = 1_000_000_007
        n = len(nums)
        m = max(nums)
        f = [[[0] * (m + 1) for _ in range(m + 1)] for _ in range(n + 1)]
        for j in range(1, m + 1):
            f[0][j][j] = 1
        for i, x in enumerate(nums):
            for j in range(m + 1):
                for k in range(m + 1):
                    f[i + 1][j][k] = (f[i][j][k] + f[i][gcd(j, x)][k] + f[i][j][gcd(k, x)]) % MOD
        return f[n][0][0]
```

```java [sol-Java]
class Solution {
    public int subsequencePairCount(int[] nums) {
        final int MOD = 1_000_000_007;
        int n = nums.length;
        int m = 0;
        for (int x : nums) {
            m = Math.max(m, x);
        }
        int[][][] f = new int[n + 1][m + 1][m + 1];
        for (int j = 1; j <= m; j++) {
            f[0][j][j] = 1;
        }
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            for (int j = 0; j <= m; j++) {
                for (int k = 0; k <= m; k++) {
                    f[i + 1][j][k] = (int) (((long) f[i][j][k] + f[i][gcd(j, x)][k] + f[i][j][gcd(k, x)]) % MOD);
                }
            }
        }
        return f[n][0][0];
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int subsequencePairCount(vector<int>& nums) {
        const int MOD = 1'000'000'007;
        int n = nums.size();
        int m = ranges::max(nums);
        vector<vector<vector<int>>> f(n + 1, vector<vector<int>>(m + 1, vector<int>(m + 1)));
        for (int j = 1; j <= m; j++) {
            f[0][j][j] = 1;
        }
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            for (int j = 0; j <= m; j++) {
                for (int k = 0; k <= m; k++) {
                    f[i + 1][j][k] = ((long long) f[i][j][k] + f[i][gcd(j, x)][k] + f[i][j][gcd(k, x)]) % MOD;
                }
            }
        }
        return f[n][0][0];
    }
};
```

```go [sol-Go]
func subsequencePairCount(nums []int) int {
	const mod = 1_000_000_007
	n := len(nums)
	m := slices.Max(nums)
	f := make([][][]int, n+1)
	for i := range f {
		f[i] = make([][]int, m+1)
		for j := range f[i] {
			f[i][j] = make([]int, m+1)
		}
	}
	for j := 1; j <= m; j++ {
		f[0][j][j] = 1
	}
	for i, x := range nums {
		for j := 0; j <= m; j++ {
			for k := 0; k <= m; k++ {
				f[i+1][j][k] = (f[i][j][k] + f[i][gcd(j, x)][k] + f[i][j][gcd(k, x)]) % mod
			}
		}
	}
	return f[n][0][0]
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU^2\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(nU^2)$。

注：可以预处理 $200$ 内所有数对的 GCD，加快计算效率。另外可以用滚动数组优化空间。

## 方法二：倍数容斥

### 总体思路

1. 定义 $f[g_1][g_2]$ 表示第一个子序列的 GCD 是 $g_1$ 的倍数，第二个子序列的 GCD 是 $g_2$ 的倍数的子序列对的个数。
2. 利用**倍数容斥**，求出第一个子序列的 GCD 恰好等于 $i$，且第二个子序列的 GCD 也恰好等于 $i$ 的子序列对的个数。

### 计算 $f$

定义 $c[i]$ 表示 $\textit{nums}$ 中的 $i$ 的个数。

定义 $\textit{cnt}[i]$ 表示 $\textit{nums}$ 中的 $i$ 的倍数的个数，即

$$
\textit{cnt}[i] = c[i]+c[2i]+c[3i] +\cdots = \sum_{j=1}^{\left\lfloor m/i\right\rfloor} c[j\cdot i]
$$

其中 $m=\max(\textit{nums})$。

计算 $f[g_1][g_2]$：

- 设 $l=\text{LCM}(g_1,g_2),\ c=\textit{cnt}[l],\ c_1 = \textit{cnt}[g_1],\ c_2 = \textit{cnt}[g_2]$。
- 对于既是 $g_1$ 倍数又是 $g_2$ 倍数的元素（这有 $c$ 个），和方法一的讨论一样，可以不选，也可以选择并放入第一个子序列，或者放入第二个子序列，方案数为 $3^c$。
- 对于是 $g_1$ 倍数但不是 $l$ 倍数的元素（这有 $c_1-c$ 个），可以不选，也可以选择并放入第一个子序列，方案数为 $2^{c_1-c}$。
- 对于是 $g_2$ 倍数但不是 $l$ 倍数的元素（这有 $c_2-c$ 个），可以不选，也可以选择并放入第二个子序列，方案数为 $2^{c_2-c}$。
- 三者互相独立，一共有 $3^c\cdot 2^{c_1+c_2-2c}$ 个方案。

但其中有不合法的方案：

- 减去第一个子序列为空的方案数，也就是所有元素都放入了第二个子序列，方案数为 $2^{c_2}$。
- 减去第二个子序列为空的方案数，也就是所有元素都放入了第一个子序列，方案数为 $2^{c_1}$。
- 注意两个子序列都为空的情况，重复减去了，所以要再加回来，方案数为 $1$，因为所有元素都不选的方案只有一个。

综上所述：

$$
f[g_1][g_2] = 3^c\cdot 2^{c_1+c_2-2c} - 2^{c_1} - 2^{c_2} + 1
$$

### 倍数容斥

计算第一个子序列的 GCD 恰好等于 $i$，且第二个子序列的 GCD 也恰好等于 $i$ 的子序列对的个数。

这可以用**倍数容斥**。

为方便大家理解，先来计算第一个子序列的 GCD 恰好等于 $i$，第二个子序列的 GCD 随意（$1$ 的倍数）的子序列对的个数。

- 从 $f[i][1]$ 开始，也就是第一个子序列的 GCD 是 $i$ 的倍数的方案数。
- 从中减去第一个子序列的 GCD 恰好等于 $2i,3i,4i,\cdots$ 的方案数。
- 减去 $f[2i][1]$。
- 减去 $f[3i][1]$。
- 忽略 $f[4i][1]$，因为 $f[4i][1]$ 已经在 $f[2i][1]$ 中了。
- 减去 $f[5i][1]$。
- 加上 $f[6i][1]$，因为 $6i$ 既是 $2i$ 的倍数，又是 $3i$ 的倍数，多减了一次。
- 以此类推，每个 $f[j\cdot i][1]$ 的前面都有一个系数 $-1$、$0$ 或者 $1$，这正是**莫比乌斯函数** $\mu(j)$。详细介绍请查阅初等数论书籍。

所以第一个子序列的 GCD 恰好等于 $i$，第二个子序列的 GCD 随意（$1$ 的倍数）的子序列对的个数为

$$
\sum_{j=1}^{\left\lfloor m/i\right\rfloor} \mu(j)f[j\cdot i][1]
$$

同理，第一个子序列的 GCD 恰好等于 $i$，且第二个子序列的 GCD 也恰好等于 $i$ 的子序列对的个数为

$$
\sum_{j=1}^{\left\lfloor m/i\right\rfloor} \sum_{k=1}^{\left\lfloor m/i\right\rfloor} \mu(j)\mu(k)f[j\cdot i][k\cdot i]
$$

最终答案为

$$
\sum_{i=1}^{m}\sum_{j=1}^{\left\lfloor m/i\right\rfloor} \sum_{k=1}^{\left\lfloor m/i\right\rfloor} \mu(j)\mu(k)f[j\cdot i][k\cdot i]
$$

```py [sol-Python3]
MOD = 1_000_000_007
MX = 201

lcms = [[lcm(i, j) for j in range(MX)] for i in range(MX)]

pow2 = [1] * MX
pow3 = [1] * MX
for i in range(1, MX):
    pow2[i] = pow2[i - 1] * 2 % MOD
    pow3[i] = pow3[i - 1] * 3 % MOD

mu = [0] * MX
mu[1] = 1
for i in range(1, MX):
    for j in range(i * 2, MX, i):
        mu[j] -= mu[i]

class Solution:
    def subsequencePairCount(self, nums: List[int]) -> int:
        m = max(nums)
        # cnt[i] 表示 nums 中的 i 的倍数的个数
        cnt = [0] * (m + 1)
        for x in nums:
            cnt[x] += 1
        for i in range(1, m + 1):
            for j in range(i * 2, m + 1, i):
                cnt[i] += cnt[j]  # 统计 i 的倍数的个数

        f = [[0] * (m + 1) for _ in range(m + 1)]
        for g1 in range(1, m + 1):
            for g2 in range(1, m + 1):
                l = lcms[g1][g2]
                c = cnt[l] if l <= m else 0
                c1, c2 = cnt[g1], cnt[g2]
                f[g1][g2] = (pow3[c] * pow2[c1 + c2 - c * 2] - pow2[c1] - pow2[c2] + 1) % MOD

        # 倍数容斥
        return sum(mu[j] * mu[k] * f[j * i][k * i]
                   for i in range(1, m + 1)
                   for j in range(1, m // i + 1)
                   for k in range(1, m // i + 1)) % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 201;

    private static final int[][] lcms = new int[MX][MX];
    private static final int[] pow2 = new int[MX];
    private static final int[] pow3 = new int[MX];
    private static final int[] mu = new int[MX];

    static {
        for (int i = 1; i < MX; i++) {
            for (int j = 1; j < MX; j++) {
                lcms[i][j] = lcm(i, j);
            }
        }

        pow2[0] = pow3[0] = 1;
        for (int i = 1; i < MX; i++) {
            pow2[i] = pow2[i - 1] * 2 % MOD;
            pow3[i] = (int) ((long) pow3[i - 1] * 3 % MOD);
        }

        mu[1] = 1;
        for (int i = 1; i < MX; i++) {
            for (int j = i * 2; j < MX; j += i) {
                mu[j] -= mu[i];
            }
        }
    }

    public int subsequencePairCount(int[] nums) {
        int m = 0;
        for (int x : nums) {
            m = Math.max(m, x);
        }

        // cnt[i] 表示 nums 中的 i 的倍数的个数
        int[] cnt = new int[m + 1];
        for (int x : nums) {
            cnt[x]++;
        }
        for (int i = 1; i <= m; i++) {
            for (int j = i * 2; j <= m; j += i) {
                cnt[i] += cnt[j]; // 统计 i 的倍数的个数
            }
        }

        int[][] f = new int[m + 1][m + 1];
        for (int g1 = 1; g1 <= m; g1++) {
            for (int g2 = 1; g2 <= m; g2++) {
                int l = lcms[g1][g2];
                int c = l <= m ? cnt[l] : 0;
                int c1 = cnt[g1];
                int c2 = cnt[g2];
                f[g1][g2] = (int) (((long) pow3[c] * pow2[c1 + c2 - c * 2] - pow2[c1] - pow2[c2] + 1) % MOD);
            }
        }

        // 倍数容斥
        long ans = 0;
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= m / i; j++) {
                for (int k = 1; k <= m / i; k++) {
                    ans += mu[j] * mu[k] * f[j * i][k * i];
                }
            }
        }
        return (int) ((ans % MOD + MOD) % MOD); // 保证 ans 非负
    }

    private static int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }

    private static int lcm(int a, int b) {
        return a / gcd(a, b) * b;
    }
}
```

```cpp [sol-C++]
const int MOD = 1'000'000'007;
const int MX = 201;

int lcms[MX][MX], pow2[MX], pow3[MX], mu[MX];

auto init = [] {
    for (int i = 1; i < MX; i++) {
        for (int j = 1; j < MX; j++) {
            lcms[i][j] = lcm(i, j);
        }
    }

    pow2[0] = pow3[0] = 1;
    for (int i = 1; i < MX; i++) {
        pow2[i] = pow2[i - 1] * 2 % MOD;
        pow3[i] = (long long) pow3[i - 1] * 3 % MOD;
    }

    mu[1] = 1;
    for (int i = 1; i < MX; i++) {
        for (int j = i * 2; j < MX; j += i) {
            mu[j] -= mu[i];
        }
    }
    return 0;
}();

class Solution {
public:
    int subsequencePairCount(vector<int>& nums) {
        int m = ranges::max(nums);
        // cnt[i] 表示 nums 中的 i 的倍数的个数
        vector<int> cnt(m + 1);
        for (int x : nums) {
            cnt[x]++;
        }
        for (int i = 1; i <= m; i++) {
            for (int j = i * 2; j <= m; j += i) {
                cnt[i] += cnt[j]; // 统计 i 的倍数的个数
            }
        }

        vector<vector<int>> f(m + 1, vector<int>(m + 1));
        for (int g1 = 1; g1 <= m; g1++) {
            for (int g2 = 1; g2 <= m; g2++) {
                int l = lcms[g1][g2];
                int c = l <= m ? cnt[l] : 0;
                int c1 = cnt[g1], c2 = cnt[g2];
                f[g1][g2] = ((long long) pow3[c] * pow2[c1 + c2 - c * 2] - pow2[c1] - pow2[c2] + 1) % MOD;
            }
        }

        // 倍数容斥
        long long ans = 0;
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= m / i; j++) {
                for (int k = 1; k <= m / i; k++) {
                    ans += mu[j] * mu[k] * f[j * i][k * i];
                }
            }
        }
        return (ans % MOD + MOD) % MOD; // 保证 ans 非负
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 201

var lcms [mx][mx]int
var pow2, pow3, mu [mx]int

func init() {
	for i := 1; i < mx; i++ {
		for j := 1; j < mx; j++ {
			lcms[i][j] = lcm(i, j)
		}
	}

	pow2[0], pow3[0] = 1, 1
	for i := 1; i < mx; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
		pow3[i] = pow3[i-1] * 3 % mod
	}

	mu[1] = 1
	for i := 1; i < mx; i++ {
		for j := i * 2; j < mx; j += i {
			mu[j] -= mu[i]
		}
	}
}

func subsequencePairCount(nums []int) int {
	m := slices.Max(nums)
	// cnt[i] 表示 nums 中的 i 的倍数的个数
	cnt := make([]int, m+1)
	for _, x := range nums {
		cnt[x]++
	}
	for i := 1; i <= m; i++ {
		for j := i * 2; j <= m; j += i {
			cnt[i] += cnt[j] // 统计 i 的倍数的个数
		}
	}

	f := make([][]int, m+1)
	for g1 := 1; g1 <= m; g1++ {
		f[g1] = make([]int, m+1)
		for g2 := 1; g2 <= m; g2++ {
			l := lcms[g1][g2]
			c := 0
			if l <= m {
				c = cnt[l]
			}
			c1, c2 := cnt[g1], cnt[g2]
			f[g1][g2] = (pow3[c]*pow2[c1+c2-c*2] - pow2[c1] - pow2[c2] + 1) % mod
		}
	}

	// 倍数容斥
	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= m/i; j++ {
			for k := 1; k <= m/i; k++ {
				ans += mu[j] * mu[k] * f[j*i][k*i]
			}
		}
	}
	return (ans%mod + mod) % mod // 保证 ans 非负
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
```

#### 复杂度分析

预处理的时间忽略不计。

- 时间复杂度：$\mathcal{O}(n+U^2)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。最后的三重循环，循环次数 $\sum\limits_{i} \left(\dfrac{U}{i}\right)^2 = U^2\sum\limits_{i} \dfrac{1}{i^2} < U^2\cdot \dfrac{\pi}{6}$，即 $\mathcal{O}(U^2)$。
- 空间复杂度：$\mathcal{O}(U^2)$。

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
