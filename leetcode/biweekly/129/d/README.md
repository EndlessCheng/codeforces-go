## 方法一：记忆化搜索

**前置知识**：[动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含如何把记忆化搜索 1:1 翻译成递推的技巧。

先解释 $\textit{limit}$，意思是数组中至多有连续 $\textit{limit}$ 个 $0$，且至多有连续 $\textit{limit}$ 个 $1$。

看示例 3，$\textit{zero}=3,\ \textit{one}=3,\ \textit{limit}=2$。

考虑稳定数组的最后一个位置填 $0$ 还是 $1$：

- 填 $0$，问题变成剩下 $2$ 个 $0$ 和 $3$ 个 $1$ 怎么填。
- 填 $1$，问题变成剩下 $3$ 个 $0$ 和 $2$ 个 $1$ 怎么填。
- 这两个都是和原问题相似的子问题。

看上去，定义 $\textit{dfs}(i,j)$ 表示用 $i$ 个 $0$ 和 $j$ 个 $1$ 构造稳定数组的方案数？但这样定义不方便计算 $\textit{limit}$ 带来的影响。

改成定义 $\textit{dfs}(i,j,k)$ 表示用 $i$ 个 $0$ 和 $j$ 个 $1$ 构造稳定数组的方案数，其中第 $i+j$ 个位置要填 $k$，其中 $k$ 为 $0$ 或 $1$。

考虑 $\textit{dfs}(i,j,0)$ 怎么算。现在，第 $i+j$ 个位置填的是 $0$，考虑第 $i+j-1$ 个位置要填什么：

- 填 $0$，方案数就是 $\textit{dfs}(i-1,j,0)$。
- 填 $1$，方案数就是 $\textit{dfs}(i-1,j,1)$。

看上去，把这两种情况加起来，我们就得到了 $\textit{dfs}(i,j,0)$。但是，这个过程中会产生不合法的情况。

假设 $\textit{limit}=3$，$\textit{dfs}(i-1,j,0)$ 是一些以 $0$ 结尾的稳定数组（合法数组），考虑末尾 $0$ 的个数：

- 末尾有 $1$ 个 $0$：那么末尾必定是 $10$。
- 末尾有 $2$ 个 $0$：那么末尾必定是 $100$。
- 末尾有 $3$ 个 $0$：那么末尾必定是 $1000$。注意，由于末尾不能超过 $3$ 个 $0$，所以这样的稳定数组的倒数第 $4$ 个数一定是 $1$，也就是在 $\textit{dfs}(i-1,j,0)$ 中有 $\textit{dfs}(i-4,j,1)$ 个末尾是 $1000$ 的稳定数组。

若要通过 $\textit{dfs}(i-1,j,0)$ 计算 $\textit{dfs}(i,j,0)$，相当于往这 $\textit{dfs}(i-1,j,0)$ 个稳定数组的末尾再加一个 $0$：

- 末尾有 $2$ 个 $0$，即 $100$，这是合法的。
- 末尾有 $3$ 个 $0$，即 $1000$，这是合法的。
- 末尾有 $4$ 个 $0$，即 $10000$，这是不合法的，要**全部去掉**！根据上面的分析，要减去 $\textit{dfs}(i-4,j,1)$。

一般地，因为 $\textit{dfs}$ 的定义是稳定数组的方案数，只包含合法方案，所以在最后连续 $\textit{limit}$ 个位置填 $0$ 的情况下，倒数第 $\textit{limit}+1$ 个位置一定要填 $1$，这有 $\textit{dfs}(i-\textit{limit}-1,j,1)$ 种方案。对于 $\textit{dfs}(i,j,0)$ 来说，这 $\textit{dfs}(i-\textit{limit}-1,j,1)$ 个方案就是不合法方案了，要减掉，得

$$
\textit{dfs}(i,j,0) = \textit{dfs}(i-1,j,0) + \textit{dfs}(i-1,j,1) - \textit{dfs}(i-\textit{limit}-1,j,1)
$$

同理得

$$
\textit{dfs}(i,j,1) = \textit{dfs}(i,j-1,0) + \textit{dfs}(i,j-1,1) - \textit{dfs}(i,j-\textit{limit}-1,0)
$$

递归边界 1：如果 $i<0$ 或者 $j<0$，返回 $0$。也可以在递归 $\textit{dfs}(i-\textit{limit}-1,j,1)$ 前判断 $i>\textit{limit}$，在递归 $\textit{dfs}(i,j-\textit{limit}-1,0)$ 前判断 $j>\textit{limit}$。下面代码在递归前判断。

递归边界 2：如果 $i=0$，那么当 $k=1$ 且 $j\le \textit{limit}$ 的情况下返回 $1$，否则返回 $0$；如果 $j=0$，那么当 $k=0$ 且 $i\le \textit{limit}$ 的情况下返回 $1$，否则返回 $0$。

递归入口：$\textit{dfs}(\textit{zero},\textit{one},0)+\textit{dfs}(\textit{zero},\textit{one},1)$，即答案。

请看 [视频讲解](https://www.bilibili.com/video/BV16t421c7GB/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def numberOfStableArrays(self, zero: int, one: int, limit: int) -> int:
        MOD = 1_000_000_007
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int, k: int) -> int:
            if i == 0:
                return 1 if k == 1 and j <= limit else 0
            if j == 0:
                return 1 if k == 0 and i <= limit else 0
            if k == 0:
                return (dfs(i - 1, j, 0) + dfs(i - 1, j, 1) - (dfs(i - limit - 1, j, 1) if i > limit else 0)) % MOD
            else:  # else 可以去掉，这里仅仅是为了代码对齐
                return (dfs(i, j - 1, 0) + dfs(i, j - 1, 1) - (dfs(i, j - limit - 1, 0) if j > limit else 0)) % MOD
        ans = (dfs(zero, one, 0) + dfs(zero, one, 1)) % MOD
        dfs.cache_clear()  # 防止爆内存
        return ans
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int numberOfStableArrays(int zero, int one, int limit) {
        int[][][] memo = new int[zero + 1][one + 1][2];
        for (int[][] m : memo) {
            for (int[] m2 : m) {
                Arrays.fill(m2, -1); // -1 表示没有计算过
            }
        }
        return (dfs(zero, one, 0, limit, memo) + dfs(zero, one, 1, limit, memo)) % MOD;
    }

    private int dfs(int i, int j, int k, int limit, int[][][] memo) {
        if (i == 0) { // 递归边界
            return k == 1 && j <= limit ? 1 : 0;
        }
        if (j == 0) { // 递归边界
            return k == 0 && i <= limit ? 1 : 0;
        }
        if (memo[i][j][k] != -1) { // 之前计算过
            return memo[i][j][k];
        }
        if (k == 0) {
            // + MOD 保证答案非负
            memo[i][j][k] = (int) (((long) dfs(i - 1, j, 0, limit, memo) + dfs(i - 1, j, 1, limit, memo) +
                    (i > limit ? MOD - dfs(i - limit - 1, j, 1, limit, memo) : 0)) % MOD);
        } else {
            memo[i][j][k] = (int) (((long) dfs(i, j - 1, 0, limit, memo) + dfs(i, j - 1, 1, limit, memo) +
                    (j > limit ? MOD - dfs(i, j - limit - 1, 0, limit, memo) : 0)) % MOD);
        }
        return memo[i][j][k];
    }
}
```

```cpp [sol-C++]
class Solution {
    int MOD = 1'000'000'007;
    vector<vector<array<int, 2>>> memo;

    int dfs(int i, int j, int k, int limit) {
        if (i == 0) { // 递归边界
            return k == 1 && j <= limit;
        }
        if (j == 0) { // 递归边界
            return k == 0 && i <= limit;
        }
        int& res = memo[i][j][k]; // 注意这里是引用
        if (res != -1) { // 之前计算过
            return res;
        }
        if (k == 0) {
            // + MOD 保证答案非负
            res = ((long long) dfs(i - 1, j, 0, limit) + dfs(i - 1, j, 1, limit) +
                   (i > limit ? MOD - dfs(i - limit - 1, j, 1, limit) : 0)) % MOD;
        } else {
            res = ((long long) dfs(i, j - 1, 0, limit) + dfs(i, j - 1, 1, limit) +
                   (j > limit ? MOD - dfs(i, j - limit - 1, 0, limit) : 0)) % MOD;
        }
        return res;
    }

public:
    int numberOfStableArrays(int zero, int one, int limit) {
        // -1 表示没有计算过
        memo.resize(zero + 1, vector<array<int, 2>>(one + 1, {-1, -1}));
        return (dfs(zero, one, 0, limit) + dfs(zero, one, 1, limit)) % MOD;
    }
};
```

```go [sol-Go]
func numberOfStableArrays(zero, one, limit int) int {
	const mod = 1_000_000_007
	memo := make([][][2]int, zero+1)
	for i := range memo {
		memo[i] = make([][2]int, one+1)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) (res int) {
		if i == 0 { // 递归边界
			if k == 1 && j <= limit {
				return 1
			}
			return
		}
		if j == 0 { // 递归边界
			if k == 0 && i <= limit {
				return 1
			}
			return
		}
		p := &memo[i][j][k]
		if *p != -1 { // 之前计算过
			return *p
		}
		if k == 0 {
			// +mod 保证答案非负
			res = (dfs(i-1, j, 0) + dfs(i-1, j, 1)) % mod
			if i > limit {
				res = (res - dfs(i-limit-1, j, 1) + mod) % mod
			}
		} else {
			res = (dfs(i, j-1, 0) + dfs(i, j-1, 1)) % mod
			if j > limit {
				res = (res - dfs(i, j-limit-1, 0) + mod) % mod
			}
		}
		*p = res // 记忆化
		return
	}
	return (dfs(zero, one, 0) + dfs(zero, one, 1)) % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{zero}\cdot \textit{one})$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(\textit{zero}\cdot \textit{one})$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(\textit{zero}\cdot \textit{one})$。
- 空间复杂度：$\mathcal{O}(\textit{zero}\cdot \textit{one})$。有多少个状态，$\textit{memo}$ 数组的大小就是多少。

## 方法二：递推

和 $\textit{dfs}(i,j,k)$ 一样，定义 $f[i][j][k]$ 表示用 $i$ 个 $0$ 和 $j$ 个 $1$ 构造稳定数组的方案数，其中第 $i+j$ 个位置要填 $k$，其中 $k$ 为 $0$ 或 $1$。

状态转移方程：

$$
\begin{aligned}
f[i][j][0] &= f[i-1][j][0] + f[i-1][j][1] - f[i-\textit{limit}-1][j][1]     \\
f[i][j][1] &= f[i][j-1][0] + f[i][j-1][1] - f[i][j-\textit{limit}-1][0]     \\
\end{aligned}
$$

如果 $i\le \textit{limit}$ 则 $f[i-\textit{limit}-1][j][1]$ 视作 $0$，如果 $j\le \textit{limit}$ 则 $f[i][j-\textit{limit}-1][0]$ 视作 $0$。

初始值：$f[i][0][0] = f[0][j][1] = 1$，其中 $1\le i \le \min(\textit{limit}, \textit{zero}),\ 1\le j \le \min(\textit{limit}, \textit{one})$。翻译自递归边界。

答案：$f[\textit{zero}][\textit{one}][0] + f[\textit{zero}][\textit{one}][1]$。翻译自递归入口。

```py [sol-Python3]
class Solution:
    def numberOfStableArrays(self, zero: int, one: int, limit: int) -> int:
        MOD = 1_000_000_007
        f = [[[0, 0] for _ in range(one + 1)] for _ in range(zero + 1)]
        for i in range(1, min(limit, zero) + 1):
            f[i][0][0] = 1
        for j in range(1, min(limit, one) + 1):
            f[0][j][1] = 1
        for i in range(1, zero + 1):
            for j in range(1, one + 1):
                f[i][j][0] = (f[i - 1][j][0] + f[i - 1][j][1] - (f[i - limit - 1][j][1] if i > limit else 0)) % MOD
                f[i][j][1] = (f[i][j - 1][0] + f[i][j - 1][1] - (f[i][j - limit - 1][0] if j > limit else 0)) % MOD
        return sum(f[-1][-1]) % MOD
```

```java [sol-Java]
class Solution {
    public int numberOfStableArrays(int zero, int one, int limit) {
        final int MOD = 1_000_000_007;
        int[][][] f = new int[zero + 1][one + 1][2];
        for (int i = 1; i <= Math.min(limit, zero); i++) {
            f[i][0][0] = 1;
        }
        for (int j = 1; j <= Math.min(limit, one); j++) {
            f[0][j][1] = 1;
        }
        for (int i = 1; i <= zero; i++) {
            for (int j = 1; j <= one; j++) {
                // + MOD 保证答案非负
                f[i][j][0] = (int) (((long) f[i - 1][j][0] + f[i - 1][j][1] + (i > limit ? MOD - f[i - limit - 1][j][1] : 0)) % MOD);
                f[i][j][1] = (int) (((long) f[i][j - 1][0] + f[i][j - 1][1] + (j > limit ? MOD - f[i][j - limit - 1][0] : 0)) % MOD);
            }
        }
        return (f[zero][one][0] + f[zero][one][1]) % MOD;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfStableArrays(int zero, int one, int limit) {
        const int MOD = 1'000'000'007;
        vector<vector<array<int, 2>>> f(zero + 1, vector<array<int, 2>>(one + 1));
        for (int i = 1; i <= min(limit, zero); i++) {
            f[i][0][0] = 1;
        }
        for (int j = 1; j <= min(limit, one); j++) {
            f[0][j][1] = 1;
        }
        for (int i = 1; i <= zero; i++) {
            for (int j = 1; j <= one; j++) {
                // + MOD 保证答案非负
                f[i][j][0] = ((long long) f[i - 1][j][0] + f[i - 1][j][1] + (i > limit ? MOD - f[i - limit - 1][j][1] : 0)) % MOD;
                f[i][j][1] = ((long long) f[i][j - 1][0] + f[i][j - 1][1] + (j > limit ? MOD - f[i][j - limit - 1][0] : 0)) % MOD;
            }
        }
        return (f[zero][one][0] + f[zero][one][1]) % MOD;
    }
};
```

```go [sol-Go]
func numberOfStableArrays(zero, one, limit int) (ans int) {
	const mod = 1_000_000_007
	f := make([][][2]int, zero+1)
	for i := range f {
		f[i] = make([][2]int, one+1)
	}
	for i := 1; i <= min(limit, zero); i++ {
		f[i][0][0] = 1
	}
	for j := 1; j <= min(limit, one); j++ {
		f[0][j][1] = 1
	}
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			f[i][j][0] = (f[i-1][j][0] + f[i-1][j][1]) % mod
			if i > limit {
				// + mod 保证答案非负
				f[i][j][0] = (f[i][j][0] - f[i-limit-1][j][1] + mod) % mod
			}
			f[i][j][1] = (f[i][j-1][0] + f[i][j-1][1]) % mod
			if j > limit {
				f[i][j][1] = (f[i][j][1] - f[i][j-limit-1][0] + mod) % mod
			}
		}
	}
	return (f[zero][one][0] + f[zero][one][1]) % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{zero}\cdot \textit{one})$。
- 空间复杂度：$\mathcal{O}(\textit{zero}\cdot \textit{one})$。

## 方法三：容斥原理+乘法原理

回顾一个经典的组合数学问题：

- 把 $n$ 个无区别的小球，放入 $m$ 个有区别的盒子，不允许空盒，有多少种方案？

这可以用**隔板法**解决，$n$ 个小球之间有 $n-1$ 个空隙，从中选择 $m-1$ 个空隙，插入 $m-1$ 个隔板，这样就把小球分成了 $m$ 组，并且每一组都是非空的，方案数就是 $n-1$ 选 $m-1$ 的组合数 $\dbinom {n-1} {m-1}$。

- 只考虑 $0$，把 $0$ 分成 $i$ 组，方案数就是 $f_0[i] = \dbinom {\textit{zero}-1} {i-1}$；
- 只考虑 $1$，把 $1$ 分成 $i$ 组，方案数就是 $f_1[i] = \dbinom {\textit{one}-1} {i-1}$。

如何**综合考虑** $0$ 和 $1$？要如何计算方案数？

例如 $10110001$，相当于把 $0$ 分成了 $2$ 组，把 $1$ 分成了 $3$ 组。

一般地，设 $1$ 分成了 $i$ 组，那么 $0$ 会分成多少组？有哪些情况？

有如下四种情况：

- $0$ 分成 $i-1$ 组，例如 $10110001$。注意第一个数和最后一个数一定是 $1$。
- $0$ 分成 $i$ 组，且第一个数是 $0$，例如 $01010011$。注意最后一个数一定是 $1$。
- $0$ 分成 $i$ 组，且第一个数是 $1$，例如 $10100110$。注意最后一个数一定是 $0$。
- $0$ 分成 $i+1$ 组，例如 $01010110$。注意第一个数和最后一个数一定是 $0$。

注意 $0$ 和 $1$ 内部的分组方案是互相独立的，例如

$$
\begin{aligned}
&11010001\\
&10110001\\
&10100011
\end{aligned}
$$

这些例子的 $0$ 的组数不变，$1$ 的组数不变，$0$ 的分组方式也不变（都是一个 $0$ 和三个 $0$），只有 $1$ 的分组方式在变。

根据乘法原理，综合考虑 $0$ 和 $1$，把 $1$ 分成 $i$ 组总的方案数，等于上面说的四种情况（只考虑 $0$，把 $0$ 分成 $i-1,i,i+1$ 组）的方案数之和，乘以只考虑 $1$，把 $1$ 分成 $i$ 组的方案数，即

$$
(f_0[i-1] + 2\cdot f_0[i] + f_0[i+1])\cdot f_1[i]
$$

接下来，考虑 $\textit{limit}$ 带来的影响。推荐先看 [2929. 给小朋友们分糖果 II](https://leetcode.cn/problems/distribute-candies-among-children-ii/) 以及 [我的题解](https://leetcode.cn/problems/distribute-candies-among-children-ii/solution/o1-rong-chi-yuan-li-pythonjavacgo-by-end-2woj/)。

根据**容斥原理**，对于 $f_0[i] = \dbinom {\textit{zero}-1} {i-1}$，我们需要减去「至少 $1$ 组有超过 $\textit{limit}$ 个 $0$」的方案数，再加上「至少 $2$ 组有超过 $\textit{limit}$ 个 $0$」的方案数，再减去「至少 $3$ 组有超过 $\textit{limit}$ 个 $0$」的方案数，……，直到「至少 $j$ 组有超过 $\textit{limit}$ 个 $0$」的方案数，$j$ 的值见下文。

- 至少 $j$ 组有超过 $\textit{limit}$ 个 $0$，相当于先从 $i$ 组中选 $j$ 组，每组先放入 $\textit{limit}$ 个 $0$，然后再把剩下的 $\textit{zero} - j\cdot \textit{limit}$ 分成 $i$ 组（需要满足 $\textit{zero} - j\cdot \textit{limit} \ge i$），方案数为

$$
\dbinom {i} {j} \dbinom {\textit{zero} - j\cdot \textit{limit}-1} {i-1}
$$

所以

$$
f_0[i] = \dbinom {\textit{zero}-1} {i-1} + \sum_{j} (-1)^j \dbinom {i} {j} \dbinom {\textit{zero} - j\cdot \textit{limit}-1} {i-1}
$$

其中 $j\ge 1$ 且需要满足 $\textit{zero} - j\cdot \textit{limit} \ge i$，即

$$
1\le j\le \left\lfloor\dfrac{zero - i}{\textit{limit}}\right\rfloor
$$

同理有

$$
f_1[i] = \dbinom {\textit{one}-1} {i-1} + \sum_{j} (-1)^j \dbinom {i} {j} \dbinom {\textit{one} - j\cdot \textit{limit}-1} {i-1}
$$

其中

$$
1\le j\le \left\lfloor\dfrac{one - i}{\textit{limit}}\right\rfloor
$$

最终答案为

$$
\sum_{i} (f_0[i-1] + 2\cdot f_0[i] + f_0[i+1])\cdot f_1[i]
$$

其中：

1. $i\le \textit{one}$。因为 $1$ 最多分成 $\textit{one}$ 组。
2. $i-1\le \textit{zero}$。因为 $0$ 最多分成 $\textit{zero}$ 组，当 $i-1 > \textit{zero}$ 时，上式中的 $f_0[i-1] = f_0[i] = f_0[i+1] = 0$，无需累加。
3. $i\cdot \textit{limit}\ge \textit{one}$，即 $i\ge\left\lceil\dfrac{\textit{one}}{\textit{limit}}\right\rceil$。因为每组至多 $\textit{limit}$ 个 $1$，分成 $i$ 组，至多 $i\cdot \textit{limit}$ 个 $1$，这个数必须 $\ge \textit{one}$，不然剩下的 $1$ 放到哪一组都会导致组内 $1$ 的个数超过 $\textit{limit}$。

整理得

$$
\left\lceil\dfrac{\textit{one}}{\textit{limit}}\right\rceil \le i\le \min(\textit{one},\textit{zero}+1)
$$

代码实现时：

1. 上取整 $\left\lceil\dfrac{a}{b}\right\rceil$ 转换成下取整 $\left\lfloor\dfrac{a-1}{b}\right\rfloor+1$。
2. $(-1)^j$ 可以用 $1-j\bmod 2\cdot 2$ 表示，因为当 $j$ 是偶数时，该式为 $1$；当 $j$ 是奇数时，该式为 $-1$，符合 $(-1)^j$。
3. 预处理阶乘及其逆元，利用公式 $\dbinom {n} {m} = \dfrac{n!}{m!(n-m)!}$ 计算组合数。

关于取模的知识点，见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

```py [sol-Python3]
MOD = 1_000_000_007
MX = 1001

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
    def numberOfStableArrays(self, zero: int, one: int, limit: int) -> int:
        if zero > one:
            zero, one = one, zero  # 保证空间复杂度为 O(min(zero, one))
        f0 = [0] * (zero + 3)
        for i in range((zero - 1) // limit + 1, zero + 1):
            f0[i] = comb(zero - 1, i - 1)
            for j in range(1, (zero - i) // limit + 1):
                f0[i] = (f0[i] + (-1 if j % 2 else 1) * comb(i, j) * comb(zero - j * limit - 1, i - 1)) % MOD

        ans = 0
        for i in range((one - 1) // limit + 1, min(one, zero + 1) + 1):
            f1 = comb(one - 1, i - 1)
            for j in range(1, (one - i) // limit + 1):
                f1 = (f1 + (-1 if j % 2 else 1) * comb(i, j) * comb(one - j * limit - 1, i - 1)) % MOD
            ans = (ans + (f0[i - 1] + f0[i] * 2 + f0[i + 1]) * f1) % MOD
        return ans
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 1001;

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

    public int numberOfStableArrays(int zero, int one, int limit) {
        if (zero > one) {
            // swap，保证空间复杂度为 O(min(zero, one))
            int t = zero;
            zero = one;
            one = t;
        }
        long[] f0 = new long[zero + 3];
        for (int i = (zero - 1) / limit + 1; i <= zero; i++) {
            f0[i] = comb(zero - 1, i - 1);
            for (int j = 1; j <= (zero - i) / limit; j++) {
                f0[i] = (f0[i] + (1 - j % 2 * 2) * comb(i, j) * comb(zero - j * limit - 1, i - 1)) % MOD;
            }
        }

        long ans = 0;
        for (int i = (one - 1) / limit + 1; i <= Math.min(one, zero + 1); i++) {
            long f1 = comb(one - 1, i - 1);
            for (int j = 1; j <= (one - i) / limit; j++) {
                f1 = (f1 + (1 - j % 2 * 2) * comb(i, j) * comb(one - j * limit - 1, i - 1)) % MOD;
            }
            ans = (ans + (f0[i - 1] + f0[i] * 2 + f0[i + 1]) * f1) % MOD;
        }
        return (int) ((ans + MOD) % MOD); // 保证结果非负
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
const int MX = 1001;

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
    int numberOfStableArrays(int zero, int one, int limit) {
        if (zero > one) {
            swap(zero, one); // 保证空间复杂度为 O(min(zero, one))
        }
        vector<long long> f0(zero + 3);
        for (int i = (zero - 1) / limit + 1; i <= zero; i++) {
            f0[i] = comb(zero - 1, i - 1);
            for (int j = 1; j <= (zero - i) / limit; j++) {
                f0[i] = (f0[i] + (1 - j % 2 * 2) * comb(i, j) * comb(zero - j * limit - 1, i - 1)) % MOD;
            }
        }

        long long ans = 0;
        for (int i = (one - 1) / limit + 1; i <= min(one, zero + 1); i++) {
            long long f1 = comb(one - 1, i - 1);
            for (int j = 1; j <= (one - i) / limit; j++) {
                f1 = (f1 + (1 - j % 2 * 2) * comb(i, j) * comb(one - j * limit - 1, i - 1)) % MOD;
            }
            ans = (ans + (f0[i - 1] + f0[i] * 2 + f0[i + 1]) * f1) % MOD;
        }
        return (ans + MOD) % MOD; // 保证结果非负
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 1001

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

func comb(n, m int) int {
	return f[n] * invF[m] % mod * invF[n-m] % mod
}

func numberOfStableArrays(zero, one, limit int) (ans int) {
	if zero > one {
		zero, one = one, zero // 保证空间复杂度为 O(min(zero, one))
	}
	f0 := make([]int, zero+3)
	for i := (zero-1)/limit + 1; i <= zero; i++ {
		f0[i] = comb(zero-1, i-1)
		for j := 1; j <= (zero-i)/limit; j++ {
			f0[i] = (f0[i] + (1-j%2*2)*comb(i, j)*comb(zero-j*limit-1, i-1)) % mod
		}
	}

	for i := (one-1)/limit + 1; i <= min(one, zero+1); i++ {
		f1 := comb(one-1, i-1)
		for j := 1; j <= (one-i)/limit; j++ {
			f1 = (f1 + (1-j%2*2)*comb(i, j)*comb(one-j*limit-1, i-1)) % mod
		}
		ans = (ans + (f0[i-1]+f0[i]*2+f0[i+1])*f1) % mod
	}
	return (ans + mod) % mod // 保证结果非负
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

- 时间复杂度：$\mathcal{O}\left(\dfrac{\textit{zero}\cdot\textit{one}}{\textit{limit}}\right)$。忽略预处理的时间和空间。
- 空间复杂度：$\mathcal{O}(\min(\textit{zero},\textit{one}))$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
