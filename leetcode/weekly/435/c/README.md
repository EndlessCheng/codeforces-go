## 一、寻找子问题

原问题：

- 把 $\textit{nums}$ 中的某些数变成 $\textit{target}$ 中各个元素的倍数。

考虑 $x=\textit{nums}[n-1]$：

- 不修改，那么问题变成把 $\textit{nums}[0]$ 到 $\textit{nums}[n-1]$ 中的某些数变成 $\textit{target}$ 中各个元素的倍数。
- 修改，考虑从 $\textit{target}$ 中选出一个非空子集 $\textit{sub}$，把 $x$ 变成 $\textit{sub}$ 中的所有数的倍数，也就是 $\textit{sub}$ 的 LCM（最小公倍数）的倍数。问题变成把 $\textit{nums}[0]$ 到 $\textit{nums}[n-1]$ 中的某些数变成剩余集合 $\textit{target}\setminus \textit{sub}$ 中各个元素的倍数。其中 $\setminus$ 符号表示集合的差。

## 二、状态设计与状态转移方程

根据上面的讨论，我们需要在递归过程中跟踪以下信息：

- $i$：可以修改的是 $\textit{nums}[0]$ 到 $\textit{nums}[n-1]$ 中的数。
- $j$：$\textit{target}$ 剩余元素的下标组成的集合。

因此，定义状态为 $\textit{dfs}(i,j)$，表示把 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 中的某些数变成 $j$ 中相应元素的倍数，需要的最少操作次数。

考虑 $x=\textit{nums}[i]$ 改或不改：

- 不改：问题变成把 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 中的某些数变成 $j$ 中相应元素的倍数，需要的最少操作次数，即 $\textit{dfs}(i-1,j)$。
- 改：枚举 $j$ 的非空子集 $\textit{sub}$，把 $x$ 变成 $\textit{sub}$ 中的所有数的倍数，也就是 $\textit{sub}$ 的 LCM（最小公倍数）的倍数。接下来要解决的问题是，把 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 中的某些数变成 $j \setminus \textit{sub}$ 中相应元素的倍数，需要的最少操作次数，即 $\textit{dfs}(i-1,j\setminus \textit{sub})$。

这两种情况取最小值，就得到了 $\textit{dfs}(i,j)$，即

$$
\textit{dfs}(i,j) = \min\left(\textit{dfs}(i-1,j),\ \min\limits_{\textit{sub}\subseteq j} \textit{dfs}(i-1,j\setminus \textit{sub}) + \text{op}(\textit{nums}[i], \text{LCM}(\textit{sub})) \right)
$$

其中 $\text{op}(x,a)$ 为把 $x$ 通过 $+1$ 操作变成 $a$ 的倍数的最少操作次数，这等于

$$
(a - x\bmod a) \bmod a
$$

上式兼顾了 $x$ 是 $a$ 的倍数和 $x$ 不是 $a$ 的倍数两种情况。

**递归边界**：

- $\textit{dfs}(i,\varnothing)=0$。
- $\textit{dfs}(-1,j)=\infty$，其中 $j\ne \varnothing$。用 $\infty$ 表示不合法的状态，从而保证 $\min$ 不会取到不合法的状态。

**递归入口**：$\textit{dfs}(n-1,U)$，这是原问题，也是答案。其中全集 $U=\{0,1,2,\ldots,m-1\}$，其中 $m$ 是 $\textit{target}$ 的长度。

**细节**：

1. $4$ 个两两互质的数的 LCM 至多约为 $(10^4)^4 = 10^{16}$，这超过了 32 位整数最大值，要用 64 位整数。
2. 用二进制表示集合，原理见 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。
3. 预处理 $\textit{target}$ 所有子集的 LCM 以加快计算效率。这可以再写一个状压 DP：对于子集 $S$ 的 LCM，可以取 $S$ 最高位对应数字，与 $S$ 去掉最高位后的子集的 LCM，二者计算 LCM。这比直接暴力枚举所有子集，再遍历子集计算 LCM 的时间复杂度更优，具体见后面的时间复杂度分析。

[本题视频讲解](https://www.bilibili.com/video/BV1D5F6eRECp/?t=14m29s)，欢迎点赞关注~

## 三、记忆化搜索

```py [sol-Python3]
class Solution:
    def minimumIncrements(self, nums: List[int], target: List[int]) -> int:
        # 预处理 target 的所有子集的 LCM
        m = len(target)
        lcms = [1] * (1 << m)
        for i, t in enumerate(target):
            bit = 1 << i
            for mask in range(bit):
                lcms[bit | mask] = lcm(t, lcms[mask])

        @cache
        def dfs(i: int, j: int) -> int:
            if j == 0:
                return 0
            if i < 0:  # 不能有剩余元素
                return inf
            # 不修改 nums[i]
            res = dfs(i - 1, j)
            # 枚举 j 的所有非空子集 sub，把 nums[i] 改成 lcms[sub] 的倍数
            sub = j
            while sub:
                l = lcms[sub]
                res = min(res, dfs(i - 1, j ^ sub) + (l - nums[i] % l) % l)
                sub = (sub - 1) & j
            return res
        return dfs(len(nums) - 1, (1 << m) - 1)
```

```java [sol-Java]
class Solution {
    public int minimumIncrements(int[] nums, int[] target) {
        int n = nums.length;
        int m = target.length;

        // 预处理 target 的所有子集的 LCM
        long[] lcms = new long[1 << m];
        lcms[0] = 1;
        for (int i = 0; i < m; i++) {
            int bit = 1 << i;
            for (int mask = 0; mask < bit; mask++) {
                lcms[bit | mask] = lcm(target[i], lcms[mask]);
            }
        }

        long[][] memo = new long[n][1 << m];
        for (long[] row : memo) {
            Arrays.fill(row, -1);
        }
        return (int) dfs(n - 1, (1 << m) - 1, nums, lcms, memo);
    }

    private long dfs(int i, int j, int[] nums, long[] lcms, long[][] memo) {
        if (j == 0) {
            return 0;
        }
        if (i < 0) { // 不能有剩余元素
            return Long.MAX_VALUE / 2; // 防止溢出
        }
        if (memo[i][j] != -1) {
            return memo[i][j];
        }
        long res = dfs(i - 1, j, nums, lcms, memo); // 不修改 nums[i]
        for (int sub = j; sub > 0; sub = (sub - 1) & j) { // 枚举 j 的所有非空子集
            long l = lcms[sub];
            res = Math.min(res, dfs(i - 1, j ^ sub, nums, lcms, memo) + (l - nums[i] % l) % l);
        }
        return memo[i][j] = res;
    }

    private long lcm(long a, long b) {
        return a / gcd(a, b) * b;
    }

    private long gcd(long a, long b) {
        return b == 0 ? a : gcd(b, a % b);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumIncrements(vector<int>& nums, vector<int>& target) {
        int n = nums.size();
        int m = target.size();

        // 预处理 target 的所有子集的 LCM
        vector<long long> lcms(1 << m);
        lcms[0] = 1;
        for (int i = 0; i < m; i++) {
            int bit = 1 << i;
            for (int mask = 0; mask < bit; mask++) {
                lcms[bit | mask] = lcm(target[i], lcms[mask]);
            }
        }

        vector memo(n, vector<long long>(1 << m, -1));
        auto dfs = [&](this auto&& dfs, int i, int j) -> long long {
            if (j == 0) {
                return 0;
            }
            if (i < 0) { // 不能有剩余元素
                return LLONG_MAX / 2; // 防止溢出
            }
            long long& res = memo[i][j]; // 注意这里是引用
            if (res != -1) {
                return res;
            }
            res = dfs(i - 1, j); // 不修改 nums[i]
            for (int sub = j; sub; sub = (sub - 1) & j) { // 枚举 j 的所有非空子集
                long long l = lcms[sub];
                res = min(res, dfs(i - 1, j ^ sub) + (l - nums[i] % l) % l);
            }
            return res;
        };
        return dfs(n - 1, (1 << m) - 1);
    }
};
```

```go [sol-Go]
func minimumIncrements(nums []int, target []int) int {
	// 预处理 target 的所有子集的 LCM
	m := len(target)
	lcms := make([]int, 1<<m)
	lcms[0] = 1
	for i, t := range target {
		bit := 1 << i
		for mask, l := range lcms[:bit] {
			lcms[bit|mask] = lcm(t, l)
		}
	}

	n := len(nums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 1<<m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if j == 0 {
			return
		}
		if i < 0 { // 不能有剩余元素
			return math.MaxInt / 2 // 防止溢出
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		// 不修改 nums[i]
		res = dfs(i-1, j)
		// 枚举 j 的所有非空子集 sub，把 nums[i] 改成 lcms[sub] 的倍数
		for sub := j; sub > 0; sub = (sub - 1) & j {
			l := lcms[sub]
			res = min(res, dfs(i-1, j^sub)+(l-nums[i]%l)%l)
		}
		return
	}
	return dfs(n-1, 1<<m-1)
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
func lcm(a, b int) int { return a / gcd(a, b) * b }
```

#### 复杂度分析

同递推的复杂度分析。

## 四、1:1 翻译成递推

```py [sol-Python3]
class Solution:
    def minimumIncrements(self, nums: List[int], target: List[int]) -> int:
        # 预处理 target 的所有子集的 LCM
        m = len(target)
        lcms = [1] * (1 << m)
        for i, t in enumerate(target):
            bit = 1 << i
            for mask in range(bit):
                lcms[bit | mask] = lcm(t, lcms[mask])

        n = len(nums)
        f = [[0] * (1 << m) for _ in range(n + 1)]
        f[0][1:] = [inf] * ((1 << m) - 1)
        for i, x in enumerate(nums):
            for j in range(1, 1 << m):
                # 不修改 nums[i]
                f[i + 1][j] = f[i][j]
                # 枚举 j 的所有非空子集 sub，把 nums[i] 改成 lcms[sub] 的倍数
                sub = j
                while sub:
                    l = lcms[sub]
                    f[i + 1][j] = min(f[i + 1][j], f[i][j ^ sub] + (l - x % l) % l)
                    sub = (sub - 1) & j
        return f[n][-1]
```

```java [sol-Java]
class Solution {
    public int minimumIncrements(int[] nums, int[] target) {
        int n = nums.length;
        int m = target.length;

        // 预处理 target 的所有子集的 LCM
        long[] lcms = new long[1 << m];
        lcms[0] = 1;
        for (int i = 0; i < m; i++) {
            int bit = 1 << i;
            for (int mask = 0; mask < bit; mask++) {
                lcms[bit | mask] = lcm(target[i], lcms[mask]);
            }
        }

        long[][] f = new long[n + 1][1 << m];
        Arrays.fill(f[0], Long.MAX_VALUE / 2);
        f[0][0] = 0;
        for (int i = 0; i < n; i++) {
            for (int j = 1; j < (1 << m); j++) {
                // 不修改 nums[i]
                f[i + 1][j] = f[i][j];
                // 枚举 j 的所有非空子集 sub，把 nums[i] 改成 lcms[sub] 的倍数
                for (int sub = j; sub > 0; sub = (sub - 1) & j) {
                    long l = lcms[sub];
                    f[i + 1][j] = Math.min(f[i + 1][j], f[i][j ^ sub] + (l - nums[i] % l) % l);
                }
            }
        }
        return (int) f[n][(1 << m) - 1];
    }

    private long lcm(long a, long b) {
        return a / gcd(a, b) * b;
    }

    private long gcd(long a, long b) {
        return b == 0 ? a : gcd(b, a % b);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumIncrements(vector<int>& nums, vector<int>& target) {
        int n = nums.size();
        int m = target.size();

        // 预处理 target 的所有子集的 LCM
        vector<long long> lcms(1 << m);
        lcms[0] = 1;
        for (int i = 0; i < m; i++) {
            int bit = 1 << i;
            for (int mask = 0; mask < bit; mask++) {
                lcms[bit | mask] = lcm(target[i], lcms[mask]);
            }
        }

        vector f(n + 1, vector<long long>(1 << m));
        fill(f[0].begin() + 1, f[0].end(), LLONG_MAX / 2);
        for (int i = 0; i < n; i++) {
            for (int j = 1; j < (1 << m); j++) {
                // 不修改 nums[i]
                f[i + 1][j] = f[i][j];
                // 枚举 j 的所有非空子集 sub，把 nums[i] 改成 lcms[sub] 的倍数
                for (int sub = j; sub; sub = (sub - 1) & j) {
                    long long l = lcms[sub];
                    f[i + 1][j] = min(f[i + 1][j], f[i][j ^ sub] + (l - nums[i] % l) % l);
                }
            }
        }
        return f[n].back();
    }
};
```

```go [sol-Go]
func minimumIncrements(nums []int, target []int) int {
	// 预处理 target 的所有子集的 LCM
	m := len(target)
	lcms := make([]int, 1<<m)
	lcms[0] = 1
	for i, t := range target {
		bit := 1 << i
		for mask, l := range lcms[:bit] {
			lcms[bit|mask] = lcm(t, l)
		}
	}

	n := len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, 1<<m)
	}
	for j := 1; j < 1<<m; j++ {
		f[0][j] = math.MaxInt / 2
	}
	for i, x := range nums {
		for j := 1; j < 1<<m; j++ {
			// 不修改 nums[i]
			f[i+1][j] = f[i][j]
			// 枚举 j 的所有非空子集 sub，把 nums[i] 改成 lcms[sub] 的倍数
			for sub := j; sub > 0; sub = (sub - 1) & j {
				l := lcms[sub]
				f[i+1][j] = min(f[i+1][j], f[i][j^sub]+(l-x%l)%l)
			}
		}
	}
	return f[n][1<<m-1]
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
func lcm(a, b int) int { return a / gcd(a, b) * b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n3^m)$，其中 $n$ 是 $\textit{nums}$ 的长度，$m$ 是 $\textit{target}$ 的长度。我们枚举了 $j$ 的所有子集，由于元素个数为 $k$ 的集合有 $\binom m k$ 个，其子集有 $2^k$ 个，根据二项式定理，$\sum\limits_{k=0}^m \binom m k 2^k = (2+1)^m = 3^m$，所以内层循环的时间复杂度为 $\mathcal{O}(3^m)$。预处理 LCM 的时间复杂度为 $\mathcal{O}(2^m\log U)$（$U=\max(\textit{target})$），与 DP 相比可以忽略。其中 $2^m$ 是因为预处理的循环次数之和为 $2^0+2^1+\cdots+2^{m-1} = 2^m-1$。
- 空间复杂度：$\mathcal{O}(n2^m)$。

## 五、空间优化

观察上面的状态转移方程，在计算 $f[i+1]$ 时，只会用到 $f[i]$，不会用到比 $i$ 更早的状态。

因此可以像 [0-1 背包](https://www.bilibili.com/video/BV16Y411v7Y6/) 那样，去掉第一个维度，把 $f[i+1]$ 和 $f[i]$ 保存到**同一个数组**中。

```py [sol-Python3]
class Solution:
    def minimumIncrements(self, nums: List[int], target: List[int]) -> int:
        m = len(target)
        lcms = [1] * (1 << m)
        for i, t in enumerate(target):
            bit = 1 << i
            for mask in range(bit):
                lcms[bit | mask] = lcm(t, lcms[mask])

        n = len(nums)
        f = [inf] * (1 << m)
        f[0] = 0
        for x in nums:
            for j in range((1 << m) - 1, 0, -1):
                sub = j
                while sub:
                    l = lcms[sub]
                    f[j] = min(f[j], f[j ^ sub] + (l - x % l) % l)
                    sub = (sub - 1) & j
        return f[-1]
```

```java [sol-Java]
class Solution {
    public int minimumIncrements(int[] nums, int[] target) {
        int m = target.length;
        long[] lcms = new long[1 << m];
        lcms[0] = 1;
        for (int i = 0; i < m; i++) {
            int bit = 1 << i;
            for (int mask = 0; mask < bit; mask++) {
                lcms[bit | mask] = lcm(target[i], lcms[mask]);
            }
        }

        long[] f = new long[1 << m];
        Arrays.fill(f, Long.MAX_VALUE / 2);
        f[0] = 0;
        for (int x : nums) {
            for (int j = (1 << m) - 1; j > 0; j--) {
                for (int sub = j; sub > 0; sub = (sub - 1) & j) {
                    long l = lcms[sub];
                    f[j] = Math.min(f[j], f[j ^ sub] + (l - x % l) % l);
                }
            }
        }
        return (int) f[(1 << m) - 1];
    }

    private long lcm(long a, long b) {
        return a / gcd(a, b) * b;
    }

    private long gcd(long a, long b) {
        return b == 0 ? a : gcd(b, a % b);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumIncrements(vector<int>& nums, vector<int>& target) {
        int m = target.size();
        vector<long long> lcms(1 << m);
        lcms[0] = 1;
        for (int i = 0; i < m; i++) {
            int bit = 1 << i;
            for (int mask = 0; mask < bit; mask++) {
                lcms[bit | mask] = lcm(target[i], lcms[mask]);
            }
        }

        vector<long long> f(1 << m, LLONG_MAX / 2);
        f[0] = 0;
        for (int x : nums) {
            for (int j = (1 << m) - 1; j; j--) {
                for (int sub = j; sub; sub = (sub - 1) & j) {
                    long l = lcms[sub];
                    f[j] = min(f[j], f[j ^ sub] + (l - x % l) % l);
                }
            }
        }
        return f.back();
    }
};
```

```go [sol-Go]
func minimumIncrements(nums []int, target []int) int {
	m := len(target)
	lcms := make([]int, 1<<m)
	lcms[0] = 1
	for i, t := range target {
		bit := 1 << i
		for mask, l := range lcms[:bit] {
			lcms[bit|mask] = lcm(t, l)
		}
	}

	f := make([]int, 1<<m)
	for j := 1; j < 1<<m; j++ {
		f[j] = math.MaxInt / 2
	}
	for _, x := range nums {
		for j := 1<<m - 1; j > 0; j-- {
			for sub := j; sub > 0; sub = (sub - 1) & j {
				l := lcms[sub]
				f[j] = min(f[j], f[j^sub]+(l-x%l)%l)
			}
		}
	}
	return f[1<<m-1]
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
func lcm(a, b int) int { return a / gcd(a, b) * b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n3^m)$，其中 $n$ 是 $\textit{nums}$ 的长度，$m$ 是 $\textit{target}$ 的长度。
- 空间复杂度：$\mathcal{O}(2^m)$。

## 六、时间优化

对于每个非空子集的 $l=\text{LCM}(\textit{sub})$，计算前 $m$ 小的操作次数 $\text{op}(x,l)$（其中 $x$ 是 $\textit{nums}$ 中的数）。要修改的数一定在这 $m$ 个 $x$ 中。如果最小的被其他非空子集使用了，那么我们可以选次小的；如果次小也被其他非空子集使用了，那么选第三小的；如果第三小的也被其他非空子集使用了，那么选第四小的。由于 $\textit{target}$ 只有 $m\le 4$ 个数，所以只需要计算前 $m$ 小的操作次数。这可以用**最大堆**实现。

把这些 $x$ 在 $\textit{nums}$ 中的下标记录到一个哈希表 $\textit{candidateIndices}$ 中。注意这一共只有 $\mathcal{O}(m2^m)$ 个候选项，远远小于 $\textit{nums}$ 的长度。

在 $\textit{candidateIndices}$ 上跑状压 DP，只需要 $\mathcal{O}(m2^m\cdot 3^m)=\mathcal{O}(m6^m)$ 的时间。

**剪枝**：如果 $\text{LCM}(\textit{sub})$ 比 $m\cdot \max(\textit{nums})$ 还要大，并且也比 $\max(\textit{target})$ 大，那么这样的子集无需计算候选项，因为拆开算操作次数更优。

> 剪枝后，这个算法在随机数据下跑得更快（力扣出题喜欢出随机数据）。在随机数据下，有 $99\%$ 的概率，任选两个数的 LCM 符合上面的剪枝条件，所以期望只会遍历 $\mathcal{O}(m)$ 次 $\textit{nums}$。

```py [sol-Python3]
class Solution:
    def minimumIncrements(self, nums: List[int], target: List[int]) -> int:
        m = len(target)
        lcms = [1] * (1 << m)
        for i, t in enumerate(target):
            bit = 1 << i
            for mask in range(bit):
                lcms[bit | mask] = lcm(t, lcms[mask])

        max_lcm = max(max(nums) * m, max(target))
        candidate_indices = set()
        for l in set(lcms[1:]):
            if l > max_lcm:
                continue
            lst = nsmallest(m, (((l - x % l) % l, i) for i, x in enumerate(nums)))
            for _, i in lst:
                candidate_indices.add(i)

        f = [inf] * (1 << m)
        f[0] = 0
        for i in candidate_indices:
            x = nums[i]
            for j in range((1 << m) - 1, 0, -1):
                sub = j
                while sub:
                    l = lcms[sub]
                    f[j] = min(f[j], f[j ^ sub] + (l - x % l) % l)
                    sub = (sub - 1) & j
        return f[-1]
```

```java [sol-Java]
class Solution {
    public int minimumIncrements(int[] nums, int[] target) {
        int m = target.length;
        int u = 1 << m;
        long[] lcms = new long[u];
        lcms[0] = 1;
        for (int i = 0; i < m; i++) {
            int bit = 1 << i;
            for (int mask = 0; mask < bit; mask++) {
                lcms[bit | mask] = lcm(target[i], lcms[mask]);
            }
        }

        int maxLcm = 0;
        for (int x : nums) {
            maxLcm = Math.max(maxLcm, x);
        }
        maxLcm *= m;
        for (int x : target) {
            maxLcm = Math.max(maxLcm, x);
        }

        Set<Integer> candidateIndices = new HashSet<>();
        for (int i = 1; i < u; i++) {
            if (lcms[i] > maxLcm) {
                continue;
            }
            int l = (int) lcms[i];
            PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> b[0] - a[0]);
            for (int j = 0; j < nums.length; j++) {
                pq.offer(new int[]{(l - nums[j] % l) % l, j});
                if (pq.size() > m) {
                    pq.poll();
                }
            }
            while (!pq.isEmpty()) {
                candidateIndices.add(pq.poll()[1]);
            }
        }

        long[] f = new long[u];
        Arrays.fill(f, Long.MAX_VALUE / 2);
        f[0] = 0;
        for (int i : candidateIndices) {
            int x = nums[i];
            for (int j = u - 1; j > 0; j--) {
                for (int sub = j; sub > 0; sub = (sub - 1) & j) {
                    long l = lcms[sub];
                    f[j] = Math.min(f[j], f[j ^ sub] + (l - x % l) % l);
                }
            }
        }
        return (int) f[u - 1];
    }

    private long lcm(long a, long b) {
        return a / gcd(a, b) * b;
    }

    private long gcd(long a, long b) {
        return b == 0 ? a : gcd(b, a % b);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumIncrements(vector<int>& nums, vector<int>& target) {
        int m = target.size();
        int u = 1 << m;
        vector<long long> lcms(u);
        lcms[0] = 1;
        for (int i = 0; i < m; i++) {
            int bit = 1 << i;
            for (int mask = 0; mask < bit; mask++) {
                lcms[bit | mask] = lcm(target[i], lcms[mask]);
            }
        }

        int max_lcm = max(ranges::max(nums) * m, ranges::max(target));
        unordered_set<int> candidate_indices;
        for (int i = 1; i < u; i++) {
            if (lcms[i] > max_lcm) {
                continue;
            }
            int l = lcms[i];
            priority_queue<pair<int, int>> pq;
            for (int j = 0; j < nums.size(); j++) {
                pq.emplace((l - nums[j] % l) % l, j);
                if (pq.size() > m) {
                    pq.pop();
                }
            }
            while (!pq.empty()) {
                candidate_indices.insert(pq.top().second);
                pq.pop();
            }
        }

        vector<long long> f(u, LLONG_MAX / 2);
        f[0] = 0;
        for (int i : candidate_indices) {
            int x = nums[i];
            for (int j = u - 1; j; j--) {
                for (int sub = j; sub; sub = (sub - 1) & j) {
                    long l = lcms[sub];
                    f[j] = min(f[j], f[j ^ sub] + (l - x % l) % l);
                }
            }
        }
        return f.back();
    }
};
```

```go [sol-Go]
func minimumIncrements(nums []int, target []int) int {
	m := len(target)
	lcms := make([]int, 1<<m)
	lcms[0] = 1
	for i, t := range target {
		bit := 1 << i
		for mask, l := range lcms[:bit] {
			lcms[bit|mask] = lcm(t, l)
		}
	}

	maxLcm := max(slices.Max(nums)*m, slices.Max(target))
	candidateIndices := map[int]struct{}{}
	for _, l := range lcms[1:] {
		if l > maxLcm {
			continue
		}
		h := hp{}
		for i, x := range nums {
			p := pair{(l - x%l) % l, i}
			if len(h) < m {
				heap.Push(&h, p)
			} else {
				h.update(p)
			}
		}
		for _, p := range h {
			candidateIndices[p.i] = struct{}{}
		}
	}

	f := make([]int, 1<<m)
	for j := 1; j < 1<<m; j++ {
		f[j] = math.MaxInt / 2
	}
	for i := range candidateIndices {
		x := nums[i]
		for j := 1<<m - 1; j > 0; j-- {
			for sub := j; sub > 0; sub = (sub - 1) & j {
				l := lcms[sub]
				f[j] = min(f[j], f[j^sub]+(l-x%l)%l)
			}
		}
	}
	return f[1<<m-1]
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
func lcm(a, b int) int { return a / gcd(a, b) * b }

type pair struct{ op, i int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].op > h[j].op }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (hp) Pop() (_ any)         { return }
func (h *hp) update(p pair) {
	if p.op < (*h)[0].op {
		(*h)[0] = p
		heap.Fix(h, 0)
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n2^m\log m + m6^m)$，其中 $n$ 是 $\textit{nums}$ 的长度，$m$ 是 $\textit{target}$ 的长度。计算候选项 $\textit{candidateIndices}$ 的过程跑了 $\mathcal{O}(2^m)$ 次 $\mathcal{O}(n\log m)$ 的计算前 $m$ 小的算法，得到了 $\mathcal{O}(m2^m)$ 个候选项，所以计算 DP 的过程只需要 $\mathcal{O}(m2^m\cdot 3^m)=\mathcal{O}(m6^m)$ 的时间。注：如果使用快速选择算法代替最大堆，可以做到 $\mathcal{O}(n2^m + m6^m)$ 的时间复杂度。
- 空间复杂度：$\mathcal{O}(m2^m)$。

更多相似题目，见下面动态规划题单中的「§9.4 子集状压 DP」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
