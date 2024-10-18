请看 [本题视频讲解](https://www.bilibili.com/video/BV17w4m1e7Nw/) 第四题，欢迎点赞关注！

## 一、寻找子问题

看示例 1，我们需要构造长为 $n=3$ 且逆序对为 $2$ 的排列。

讨论最后一个数填什么：

- 填 $2$，那么左边不会有比 $2$ 大的数，这意味着剩下的 $n-1=2$ 个数的逆序对需要等于 $2-0=2$。（这是不存在的）
- 填 $1$，那么左边一定有 $1$ 个比 $1$ 大的数，这意味着剩下的 $n-1=2$ 个数的逆序对需要等于 $2-1=1$。（这只有 $1$ 种构造方式，即 $[2,0]$）
- 填 $0$，那么左边一定有 $2$ 个比 $0$ 大的数，这意味着剩下的 $n-1=2$ 个数的逆序对需要等于 $2-2=0$。（这只有 $1$ 种构造方式，即 $[1,2]$）

这些问题都是**和原问题相似的、规模更小的子问题**，可以用**递归**解决。

递归过程中，是否会有重复的子问题？在什么情况下，会出现重复的子问题？

比如 $n=4$（构造 $0,1,2,3$ 的排列），要求有 $4$ 个逆序对。考虑其中两种情况：

- 最右边的两个数是 $3,0$，它俩有 $1$ 个逆序对，再算上左边的 $1,2$ 和 $0$ 组成了 $2$ 个逆序对，一共组成了 $3$ 个逆序对。接下来要解决的子问题是：剩余 $2$ 个数，构造逆序对为 $4-3=1$ 的方案数。⚠**注意**：我们已经把左边尚未填入的 $1,2$ 和已填入的 $0$ 的逆序对去掉了（提前计算 $1,2$ 和 $0$ 的逆序对），所以**在解决子问题时，无需考虑当前填入的数字与右边已填入的数字的逆序对**。
- 最右边的两个数是 $2,1$，它俩有 $1$ 个逆序对，再算上左边的 $3$ 和 $2,1$ 组成了 $2$ 个逆序对，一共组成了 $3$ 个逆序对。接下来要解决的子问题是：剩余 $2$ 个数，构造逆序对为 $4-3=1$ 的方案数。⚠**注意**：我们已经把左边尚未填入的 $3$ 和已填入的 $2,1$ 的逆序对去掉了（提前计算 $3$ 和 $2,1$ 的逆序对），所以**在解决子问题时，无需考虑当前填入的数字与右边已填入的数字的逆序对**。

既然不考虑右边的数字，那么剩下的数字是 $1,2$ 还是 $0,3$ 其实是**等价**的！都等同于还剩下两个数，把小的数放左边，大的数放右边，就是 $0$ 个逆序对；把大的数放左边，小的数放右边，就是 $1$ 个逆序对。所以，在解决子问题时，无需知道具体还剩下哪些数，**只需要知道剩余数字的个数，以及剩余逆序对的个数**。用这两个个数去定义状态，就会产生重复的子问题，就可以用记忆化搜索解决。

> 注：动态规划有「选或不选」和「枚举选哪个」两种基本思考方式。在做题时，可根据题目要求，选择适合题目的一种来思考。本题用到的是「枚举选哪个」。

## 二、状态定义与状态转移方程

根据上面的讨论，定义 $\textit{dfs}(i,j)$ 表示 $\textit{perm}[0]$ 到 $\textit{perm}[i]$（剩余 $i+1$ 个数）逆序对为 $j$ 的排列个数。

设 $\textit{perm}[i]$ 和左边 $\textit{perm}[0]$ 到 $\textit{perm}[i-1]$ 组成了 $k$ 个逆序对：

- 枚举 $k=0,1,2,\cdots,\min(i,j)$。其中 $\min(i,j)$ 是因为左边只有 $i$ 个数，至多和 $\textit{perm}[i]$ 组成 $i$ 个逆序对。
- 组成了 $k$ 个逆序对（也就是左边有 $k$ 个数比 $\textit{perm}[i]$ 大），还剩下 $j-k$ 个逆序对，问题变成 $\textit{perm}[0]$ 到 $\textit{perm}[i-1]$ 的逆序对为 $j-k$ 的排列个数，即 $\textit{dfs}(i-1,j-k)$。

累加得

$$
\textit{dfs}(i,j) = \sum_{k=0}^{\min(i,j)}\textit{dfs}(i-1,j-k)
$$

⚠**注意**：我们**不需要知道每个位置具体填了什么数**。无论右边填了什么数，只要 $\textit{perm}[i]$ 填的是剩余元素的**最大值**，那么 $k$ 就是 $0$；只要 $\textit{perm}[i]$ 填的是剩余元素的**次大值**，那么 $k$ 就是 $1$；依此类推。

除此以外，设 $\textit{req}[i]$ 是 $\textit{perm}[0]$ 到 $\textit{perm}[i]$ 的逆序对个数（没有要求就是 $-1$），如果 $\textit{req}[i-1]\ge 0$，则无需枚举 $k$，分类讨论：

- 如果 $j<\textit{req}[i-1]$，由于 $j$ 只能变小不能变大，无法满足要求，所以 $\textit{dfs}(i,j) = 0$。
- 如果 $j-i>\textit{req}[i-1]$，即使当前填了最小的数，和左边 $i$ 个数组成了 $i$ 个逆序对，那么剩余的 $j-i$ 还是太大了，无法满足要求，所以 $\textit{dfs}(i,j) = 0$。
- 否则令上文中的 $k=j-\textit{req}[i-1]$，就可以把逆序对从 $j$ 减小到 $\textit{req}[i-1]$，从而满足要求。由于只有令 $k=j-\textit{req}[i-1]$ 一种方法，所以 $\textit{dfs}(i,j) = \textit{dfs}(i-1,j-k) = \textit{dfs}(i-1,\textit{req}[i-1])$。

**递归边界**：$\textit{dfs}(0,0)=1$，此时找到了一个符合要求的排列。

**递归入口**：$\textit{dfs}(n-1,\textit{req}[n-1])$，也就是答案。

根据题意，$\textit{req}[0]$ 一定为 $0$。代码实现时，可以在递归之前判断 $\textit{req}[0] > 0$ 的情况，如果成立则无解，直接返回 $0$。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,j)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

```py [sol-Python3]
class Solution:
    def numberOfPermutations(self, n: int, requirements: List[List[int]]) -> int:
        MOD = 1_000_000_007
        req = [-1] * n
        req[0] = 0
        for end, cnt in requirements:
            req[end] = cnt
        if req[0]:
            return 0

        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int) -> int:
            if i == 0:
                return 1
            r = req[i - 1]
            if r >= 0:
                return dfs(i - 1, r) if r <= j <= i + r else 0
            return sum(dfs(i - 1, j - k) for k in range(min(i, j) + 1)) % MOD
        return dfs(n - 1, req[-1])
```

```java [sol-Java]
class Solution {
    public int numberOfPermutations(int n, int[][] requirements) {
        int[] req = new int[n];
        Arrays.fill(req, -1);
        req[0] = 0;
        int m = 0;
        for (int[] p : requirements) {
            req[p[0]] = p[1];
            m = Math.max(m, p[1]);
        }
        if (req[0] > 0) {
            return 0;
        }

        int[][] memo = new int[n][m + 1];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return dfs(n - 1, req[n - 1], req, memo);
    }

    private int dfs(int i, int j, int[] req, int[][] memo) {
        if (i == 0) {
            return 1;
        }
        if (memo[i][j] != -1) { // 之前计算过
            return memo[i][j];
        }
        int res = 0;
        int r = req[i - 1];
        if (r >= 0) {
            if (j >= r && j - i <= r) {
                res = dfs(i - 1, r, req, memo);
            }
        } else {
            for (int k = 0; k <= Math.min(i, j); k++) {
                res = (res + dfs(i - 1, j - k, req, memo)) % 1_000_000_007;
            }
        }
        return memo[i][j] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1'000'000'007;
public:
    int numberOfPermutations(int n, vector<vector<int>>& requirements) {
        vector<int> req(n, -1);
        req[0] = 0;
        for (auto& p : requirements) {
            req[p[0]] = p[1];
        }
        if (req[0]) {
            return 0;
        }

        int m = ranges::max(req);
        vector<vector<int>> memo(n, vector<int>(m + 1, -1)); // -1 表示没有计算过
        auto dfs = [&](auto&& dfs, int i, int j) -> int {
            if (i == 0) {
                return 1;
            }
            int& res = memo[i][j]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            res = 0;
            if (int r = req[i - 1]; r >= 0) {
                if (j >= r && j - i <= r) {
                    res = dfs(dfs, i - 1, r);
                }
            } else {
                for (int k = 0; k <= min(i, j); k++) {
                    res = (res + dfs(dfs, i - 1, j - k)) % MOD;
                }
            }
            return res;
        };
        return dfs(dfs, n - 1, req[n - 1]);
    }
};
```

```go [sol-Go]
func numberOfPermutations(n int, requirements [][]int) int {
	const mod = 1_000_000_007
	req := make([]int, n)
	for i := 1; i < n; i++ {
		req[i] = -1
	}
	for _, p := range requirements {
		req[p[0]] = p[1]
	}
	if req[0] > 0 {
		return 0
	}

	m := slices.Max(req)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i == 0 {
			return 1
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if r := req[i-1]; r >= 0 {
			if j < r || j-i > r {
				return 0
			}
			return dfs(i-1, r)
		}
		for k := 0; k <= min(i, j); k++ {
			res += dfs(i-1, j-k)
		}
		return res % mod
	}
	return dfs(n-1, req[n-1])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm\min(n,m))$，其中 $m=\max(\textit{cnt}_i)$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(nm)$，单个状态的计算时间为 $\mathcal{O}(\min(n,m))$，所以动态规划的时间复杂度为 $\mathcal{O}(nm\min(n,m))$。
- 空间复杂度：$\mathcal{O}(nm)$。保存多少状态，就需要多少空间。

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i][j]$ 的定义和 $\textit{dfs}(i,j)$ 的定义是一样的，都表示 $\textit{perm}[0]$ 到 $\textit{perm}[i]$ 的逆序对为 $j$ 的排列个数。

如果 $\textit{req}[i-1]<0$，相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i][j] = \sum_{k=0}^{\min(i,j)}f[i-1][j-k]
$$

如果 $r=\textit{req}[i-1]\ge 0$，递推式为

$$
f[i][j] =
\begin{cases}
f[i-1][r],\ &r\le j \le i+r\\
0,\ &\text{otherwise}\\
\end{cases}
$$

初始值 $f[0][0]=1$，翻译自递归边界 $\textit{dfs}(0,0)=1$。

答案为 $f[n-1][\textit{req}[n-1]]$，翻译自递归入口 $\textit{dfs}(n-1,\textit{req}[n-1])$。

```py [sol-Python3]
class Solution:
    def numberOfPermutations(self, n: int, requirements: List[List[int]]) -> int:
        MOD = 1_000_000_007
        req = [-1] * n
        req[0] = 0
        for end, cnt in requirements:
            req[end] = cnt
        if req[0]:
            return 0

        m = max(req)
        f = [[0] * (m + 1) for _ in range(n)]
        f[0][0] = 1
        for i in range(1, n):
            mx = m if req[i] < 0 else req[i]
            r = req[i - 1]
            if r >= 0:
                for j in range(r, min(i + r, mx) + 1):
                    f[i][j] = f[i - 1][r]
            else:
                for j in range(mx + 1):
                    f[i][j] = sum(f[i - 1][j - k] for k in range(min(i, j) + 1)) % MOD
        return f[-1][req[-1]]
```

```java [sol-Java]
class Solution {
    public int numberOfPermutations(int n, int[][] requirements) {
        final int MOD = 1_000_000_007;
        int[] req = new int[n];
        Arrays.fill(req, -1);
        req[0] = 0;
        int m = 0;
        for (int[] p : requirements) {
            req[p[0]] = p[1];
            m = Math.max(m, p[1]);
        }
        if (req[0] > 0) {
            return 0;
        }

        int[][] f = new int[n][m + 1];
        f[0][0] = 1;
        for (int i = 1; i < n; i++) {
            int mx = req[i] < 0 ? m : req[i];
            int r = req[i - 1];
            if (r >= 0) {
                for (int j = r; j <= Math.min(i + r, mx); j++) {
                    f[i][j] = f[i - 1][r];
                }
            } else {
                for (int j = 0; j <= mx; j++) {
                    for (int k = 0; k <= Math.min(i, j); k++) {
                        f[i][j] = (f[i][j] + f[i - 1][j - k]) % MOD;
                    }
                }
            }
        }
        return f[n - 1][req[n - 1]];
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1'000'000'007;
public:
    int numberOfPermutations(int n, vector<vector<int>>& requirements) {
        vector<int> req(n, -1);
        req[0] = 0;
        for (auto& p : requirements) {
            req[p[0]] = p[1];
        }
        if (req[0]) {
            return 0;
        }

        int m = ranges::max(req);
        vector<vector<int>> f(n, vector<int>(m + 1));
        f[0][0] = 1;
        for (int i = 1; i < n; i++) {
            int mx = req[i] < 0 ? m : req[i];
            if (int r = req[i - 1]; r >= 0) {
                for (int j = r; j <= min(i + r, mx); j++) {
                    f[i][j] = f[i - 1][r];
                }
            } else {
                for (int j = 0; j <= mx; j++) {
                    for (int k = 0; k <= min(i, j); k++) {
                        f[i][j] = (f[i][j] + f[i - 1][j - k]) % MOD;
                    }
                }
            }
        }
        return f[n - 1][req[n - 1]];
    }
};
```

```go [sol-Go]
func numberOfPermutations(n int, requirements [][]int) int {
	const mod = 1_000_000_007
	req := make([]int, n)
	for i := 1; i < n; i++ {
		req[i] = -1
	}
	for _, p := range requirements {
		req[p[0]] = p[1]
	}
	if req[0] > 0 {
		return 0
	}

	m := slices.Max(req)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	f[0][0] = 1
	for i := 1; i < n; i++ {
		mx := m
		if req[i] >= 0 {
			mx = req[i]
		}
		if r := req[i-1]; r >= 0 {
			for j := r; j <= min(i+r, mx); j++ {
				f[i][j] = f[i-1][r]
			}
		} else {
			for j := 0; j <= mx; j++ {
				for k := 0; k <= min(i, j); k++ {
					f[i][j] = (f[i][j] + f[i-1][j-k]) % mod
				}
			}
		}
	}
	return f[n-1][req[n-1]]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm\min(n,m))$，其中 $m=\max(\textit{cnt}_i)$。
- 空间复杂度：$\mathcal{O}(nm)$。

## 五、终极优化

#### 1) 时间优化

和式

$$
\sum_{k=0}^{\min(i,j)}f[i-1][j-k]
$$

可以用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 优化。

#### 2) 空间优化

观察上面的状态转移方程，在计算 $f[i]$ 时，只会用到 $f[i-1]$，不会用到比 $i-1$ 更早的状态。

因此可以去掉第一个维度，反复利用同一个长为 $m+1$ 的一维数组。

代码实现时，前缀和可以直接保存在 $f$ 中。先计算前缀和，再利用前缀和计算和式（子数组和）。

关于取模的技巧，见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

```py [sol-Python3]
class Solution:
    def numberOfPermutations(self, n: int, requirements: List[List[int]]) -> int:
        MOD = 1_000_000_007
        req = [-1] * n
        req[0] = 0
        for end, cnt in requirements:
            req[end] = cnt
        if req[0]:
            return 0

        m = max(req)
        f = [0] * (m + 1)
        f[0] = 1
        for i in range(1, n):
            mx = m if req[i] < 0 else req[i]
            r = req[i - 1]
            if r >= 0:
                for j in range(m + 1):
                    f[j] = f[r] if r <= j <= min(i + r, mx) else 0
            else:
                for j in range(1, mx + 1):  # 计算前缀和
                    f[j] = (f[j] + f[j - 1]) % MOD
                for j in range(mx, i, -1):  # 计算子数组和
                    f[j] = (f[j] - f[j - i - 1]) % MOD
        return f[req[-1]]
```

```java [sol-Java]
class Solution {
    public int numberOfPermutations(int n, int[][] requirements) {
        final int MOD = 1_000_000_007;
        int[] req = new int[n];
        Arrays.fill(req, -1);
        req[0] = 0;
        int m = 0;
        for (int[] p : requirements) {
            req[p[0]] = p[1];
            m = Math.max(m, p[1]);
        }
        if (req[0] > 0) {
            return 0;
        }

        int[] f = new int[m + 1];
        f[0] = 1;
        for (int i = 1; i < n; i++) {
            int mx = req[i] < 0 ? m : req[i];
            int r = req[i - 1];
            if (r >= 0) {
                Arrays.fill(f, 0, r, 0);
                Arrays.fill(f, r + 1, Math.min(i + r, mx) + 1, f[r]);
                Arrays.fill(f, Math.min(i + r, mx) + 1, m + 1, 0);
            } else {
                for (int j = 1; j <= mx; j++) { // 计算前缀和
                    f[j] = (f[j] + f[j - 1]) % MOD;
                }
                for (int j = mx; j > i; j--) { // 计算子数组和
                    f[j] = (f[j] - f[j - i - 1] + MOD) % MOD;
                }
            }
        }
        return f[req[n - 1]];
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1'000'000'007;
public:
    int numberOfPermutations(int n, vector<vector<int>>& requirements) {
        vector<int> req(n, -1);
        req[0] = 0;
        for (auto& p : requirements) {
            req[p[0]] = p[1];
        }
        if (req[0]) {
            return 0;
        }

        int m = ranges::max(req);
        vector<int> f(m + 1);
        f[0] = 1;
        for (int i = 1; i < n; i++) {
            int mx = req[i] < 0 ? m : req[i];
            if (int r = req[i - 1]; r >= 0) {
                fill(f.begin(), f.begin() + r, 0);
                fill(f.begin() + r + 1, f.begin() + min(i + r, mx) + 1, f[r]);
                fill(f.begin() + min(i + r, mx) + 1, f.end(), 0);
            } else {
                for (int j = 1; j <= mx; j++) { // 计算前缀和
                    f[j] = (f[j] + f[j - 1]) % MOD;
                }
                for (int j = mx; j > i; j--) { // 计算子数组和
                    f[j] = (f[j] - f[j - i - 1] + MOD) % MOD;
                }
            }
        }
        return f[req[n - 1]];
    }
};
```

```go [sol-Go]
func numberOfPermutations(n int, requirements [][]int) int {
	const mod = 1_000_000_007
	req := make([]int, n)
	for i := 1; i < n; i++ {
		req[i] = -1
	}
	for _, p := range requirements {
		req[p[0]] = p[1]
	}
	if req[0] > 0 {
		return 0
	}

	m := slices.Max(req)
	f := make([]int, m+1)
	f[0] = 1
	for i := 1; i < n; i++ {
		mx := m
		if req[i] >= 0 {
			mx = req[i]
		}
		if r := req[i-1]; r >= 0 {
			clear(f[:r])
			for j := r + 1; j <= min(i+r, mx); j++ {
				f[j] = f[r]
			}
			clear(f[min(i+r, mx)+1:])
		} else {
			for j := 1; j <= mx; j++ { // 计算前缀和
				f[j] = (f[j] + f[j-1]) % mod
			}
			for j := mx; j > i; j-- { // 计算子数组和
				f[j] = (f[j] - f[j-i-1] + mod) % mod
			}
		}
	}
	return f[req[n-1]]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $m=\max(\textit{cnt}_i)$。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 相似题目

- [629. K 个逆序对数组](https://leetcode.cn/problems/k-inverse-pairs-array/)

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
