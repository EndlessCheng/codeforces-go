## 一、寻找子问题

规定空子序列的 GCD 等于 $0$。

考虑从右往左选数字，讨论 $\textit{nums}[n-1]$ **选或不选**：

- 不选，不把 $\textit{nums}[n-1]$ 加到任何子序列中，那么需要解决的问题为：在 $\textit{nums}[0]$ 到 $\textit{nums}[n-2]$ 中选数字，且已选的两个子序列的 GCD 分别为 $0,0$ 时的子序列对的个数。
- 选，把 $\textit{nums}[n-1]$ 加到第一个子序列中，那么需要解决的问题为：在 $\textit{nums}[0]$ 到 $\textit{nums}[n-2]$ 中选数字，且已选的两个子序列的 GCD 分别为 $\textit{nums}[n-1],0$ 时的子序列对的个数。
- 选，把 $\textit{nums}[n-1]$ 加到第二个子序列中，那么需要解决的问题为：在 $\textit{nums}[0]$ 到 $\textit{nums}[n-2]$ 中选数字，且已选的两个子序列的 GCD 分别为 $0,\textit{nums}[n-1]$ 时的子序列对的个数。

由于选或不选都会把原问题变成一个**和原问题相似的、规模更小的子问题**，所以可以用**递归**解决。

> 注 1：从右往左思考，主要是为了方便把递归翻译成递推。从左往右思考也是可以的。
>
> 注 2：动态规划有「**选或不选**」和「**枚举选哪个**」两种基本思考方式。子序列相邻无关一般是「选或不选」，子序列相邻相关（例如 LIS 问题）一般是「枚举选哪个」。本题用到的是「选或不选」。

## 二、状态定义与状态转移方程

根据上面的讨论，我们需要在递归过程中跟踪以下信息：

- $i$：当前考虑 $\textit{nums}[i]$ 选或不选。
- $j$：第一个子序列的 GCD 值。
- $k$：第二个子序列的 GCD 值。

因此，定义状态为 $\textit{dfs}(i,j,k)$，表示在 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 中选数字，且已选的两个子序列的 GCD 分别为 $j,k$ 时的子序列对的个数。

接下来，思考如何从一个状态转移到另一个状态。

讨论 $x=\textit{nums}[i]$ **选或不选**：

- 不选，不把 $x$ 加到任何子序列中，那么需要解决的问题为：在 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 中选数字，且已选的两个子序列的 GCD 分别为 $j,k$ 时的子序列对的个数，即 $\textit{dfs}(i-1,j,k)$。
- 选，把 $x$ 加到第一个子序列中，那么需要解决的问题为：在 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 中选数字，且已选的两个子序列的 GCD 分别为 $\text{GCD}(j,x),k$ 时的子序列对的个数，即 $\textit{dfs}(i-1,\text{GCD}(j,x),k)$。
- 选，把 $x$ 加到第二个子序列中，那么需要解决的问题为：在 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 中选数字，且已选的两个子序列的 GCD 分别为 $j,\text{GCD}(k,x)$ 时的子序列对的个数，即 $\textit{dfs}(i-1,j,\text{GCD}(k,x))$。

三种选法互斥，根据加法原理，三者相加得

$$
\textit{dfs}(i,j,k) = \textit{dfs}(i-1,j,k) + \textit{dfs}(i-1,\text{GCD}(j,x),k) + \textit{dfs}(i-1,j,\text{GCD}(k,x))
$$

**递归边界**：$\textit{dfs}(-1,j,j)=1,\ \textit{dfs}(-1,j,k)=0\ (j\ne k)$。

**递归入口**：$\textit{dfs}(n-1,0,0)-1$，也就是答案。减一是去掉两个子序列都为空的情况。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

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

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i+1][j][k]$ 的定义和 $\textit{dfs}(i,j,k)$ 的定义是一样的，都表示在 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 中选数字，且已选的两个子序列的 GCD 分别为 $j,k$ 时的子序列对的个数。这里 $+1$ 是为了把 $\textit{dfs}(-1,j,k)$ 这个状态也翻译过来，这样我们可以把 $f[0][j][k]$ 作为初始值。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i+1][j][k] = f[i][j][k] + f[i-1][\text{GCD}(j,x)][k] + f[i][j][\text{GCD}(k,x)]
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
