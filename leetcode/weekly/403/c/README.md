## 前言

本题有两种定义状态的角度：

- 按照**划分型 DP** 定义。
- 按照**状态机 DP** 定义。

为方便描述，下文将 $\textit{nums}$ 简记为 $a$。

## 方法一：划分型 DP

### 一、寻找子问题

考虑每段子数组的符号：

- 子数组长度为 $1$：符号为 $+$。
- 子数组长度为 $2$：符号为 $+-$。
- 子数组长度为 $3$：符号为 $+-+$，这等价于两个更小的子数组，符号分别为 $+-$ 和 $+$。
- 子数组长度为 $4$：符号为 $+-+-$，这等价于两个更小的子数组，符号分别为 $+-$ 和 $+-$。
- 子数组长度为 $5$：符号为 $+-+-+$，这等价于三个更小的子数组，符号分别为 $+-$、$+-$ 和 $+$。
- 子数组长度为 $6$：符号为 $+-+-+-$，这等价于三个更小的子数组，符号分别为 $+-$、$+-$ 和 $+-$。
- ……

由上面的讨论可知，对于长度超过 $2$ 的子数组，可以继续分割为长度为 $2$ 和 $1$ 的子数组，而**不改变成本之和**。

所以我们只需要考虑把 $a$ 分成若干长为 $1$ 或者长为 $2$ 的子数组，符号分别为 $+$ 和 $+-$。

对于示例 1，从 $a[n-1]$ 开始思考。

分类讨论：

- 分成长为 $1$ 的子数组，即 $a[n-1]$ 单独作为一个长为 $1$ 的子数组，接下来需要解决的问题为：$a[0]$ 到 $a[n-2]$ 的最大成本和。
- 分成长为 $2$ 的子数组，即 $a[n-2]$ 和 $a[n-1]$ 作为一个长为 $2$ 的子数组，接下来需要解决的问题为：$a[0]$ 到 $a[n-3]$ 的最大成本和。

由于这两种情况都会把原问题变成一个**和原问题相似的、规模更小的子问题**，所以可以用**递归**解决。

> 注：从右往左思考，主要是为了方便把递归翻译成递推。从左往右思考也是可以的。

### 二、状态定义与状态转移方程

因为要解决的问题都形如「$a[0]$ 到 $a[i]$ 的最大成本和」，所以用它作为本题的状态定义 $\textit{dfs}(i)$。

分类讨论：

- 分成长为 $1$ 的子数组，即 $a[i]$ 单独作为一个长为 $1$ 的子数组，接下来需要解决的问题为：$a[0]$ 到 $a[i-1]$ 的最大成本和，即 $\textit{dfs}(i) = \textit{dfs}(i-1) + a[i]$。
- 分成长为 $2$ 的子数组，即 $a[i-1]$ 和 $a[i]$ 作为一个长为 $2$ 的子数组，接下来需要解决的问题为：$a[0]$ 到 $a[i-2]$ 的最大成本和，即 $\textit{dfs}(i) = \textit{dfs}(i-2) + a[i-1] - a[i]$。

这两种情况取最大值，就得到了 $\textit{dfs}(i)$，即

$$
\textit{dfs}(i) = \max(\textit{dfs}(i-1) + a[i], \textit{dfs}(i-2) + a[i-1] - a[i])
$$

递归边界：$\textit{dfs}(-1)=0,\ \textit{dfs}(0) = a[0]$。

递归入口：$\textit{dfs}(n-1)$，也就是答案。

### 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$ 或 $-\infty$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

```py [sol-Python3]
class Solution:
    def maximumTotalCost(self, a: List[int]) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int) -> int:
            if i < 0:
                return 0
            if i == 0:
                return a[0]
            return max(dfs(i - 1) + a[i], dfs(i - 2) + a[i - 1] - a[i])
        return dfs(len(a) - 1)
```

```java [sol-Java]
class Solution {
    public long maximumTotalCost(int[] a) {
        int n = a.length;
        long[] memo = new long[n];
        Arrays.fill(memo, Long.MIN_VALUE); // Long.MIN_VALUE 表示没有计算过
        return dfs(n - 1, a, memo);
    }

    private long dfs(int i, int[] a, long[] memo) {
        if (i < 0) {
            return 0;
        }
        if (i == 0) {
            return a[0];
        }
        if (memo[i] != Long.MIN_VALUE) { // 之前计算过
            return memo[i];
        }
        return memo[i] = Math.max(dfs(i - 1, a, memo) + a[i], dfs(i - 2, a, memo) + a[i - 1] - a[i]);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumTotalCost(vector<int>& a) {
        int n = a.size();
        vector<long long> memo(n, LLONG_MIN); // LLONG_MIN 表示没有计算过
        auto dfs = [&](auto&& dfs, int i) -> long long {
            if (i < 0) {
                return 0;
            }
            if (i == 0) {
                return a[0];
            }
            auto& res = memo[i]; // 注意这里是引用
            if (res != LLONG_MIN) { // 之前计算过
                return res;
            }
            return res = max(dfs(dfs, i - 1) + a[i], dfs(dfs, i - 2) + a[i - 1] - a[i]);
        };
        return dfs(dfs, n - 1);
    }
};
```

```go [sol-Go]
func maximumTotalCost(a []int) int64 {
	n := len(a)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = math.MinInt
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			return a[0]
		}
		p := &memo[i]
		if *p != math.MinInt {
			return *p
		}
		*p = max(dfs(i-1)+a[i], dfs(i-2)+a[i-1]-a[i])
		return *p
	}
	return int64(dfs(n - 1))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{a}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。保存多少状态，就需要多少空间。

### 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i]$ 的定义和 $\textit{dfs}(i)$ 的定义是一样的，都表示 $a[0]$ 到 $a[i]$ 的最大成本和。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i] = \max(f[i-1] + a[i], f[i-2] + a[i-1] - a[i])
$$

但是，这种定义方式没有状态能表示 $i=-1$ 的情况。

解决办法：在 $f$ 数组的最左边插入一个值为 $0$ 的状态，那么其余状态全部向右偏移一位，把 $f[i]$ 改为 $f[i+1]$，把 $f[i-1]$ 改为 $f[i]$，把 $f[i-2]$ 改为 $f[i-1]$。

修改后 $f[i+1]$ 表示 $a[0]$ 到 $a[i]$ 的最大成本和。

修改后的递推式为

$$
f[i+1] = \max(f[i] + a[i], f[i-1] + a[i-1] - a[i])
$$

> 问：为什么 $\textit{a}$ 的下标不用变？
>
> 答：既然是在 $f$ 的最左边插入一个状态，那么就只需要修改和 $f$ 有关的下标，其余任何逻辑都无需修改。或者说，如果把 $\textit{a}[i]$ 也改成 $a[i+1]$，那么当 $i=n-1$ 时 $a[i+1]=a[n]$ 会下标越界，这显然是错误的。

初始值 $f[0]=0,\ f[1]=a[0]$，翻译自递归边界 $\textit{dfs}(-1)=0,\ \textit{dfs}(0) = a[0]$。

答案为 $f[n]$，翻译自递归入口 $\textit{dfs}(n-1)$。

```py [sol-Python3]
class Solution:
    def maximumTotalCost(self, a: List[int]) -> int:
        n = len(a)
        f = [0] * (n + 1)
        f[1] = a[0]
        for i in range(1, n):
            f[i + 1] = max(f[i] + a[i], f[i - 1] + a[i - 1] - a[i])
        return f[n]
```

```java [sol-Java]
class Solution {
    public long maximumTotalCost(int[] a) {
        int n = a.length;
        long[] f = new long[n + 1];
        f[1] = a[0];
        for (int i = 1; i < n; i++) {
            f[i + 1] = Math.max(f[i] + a[i], f[i - 1] + a[i - 1] - a[i]);
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumTotalCost(vector<int>& a) {
        int n = a.size();
        vector<long long> f(n + 1);
        f[1] = a[0];
        for (int i = 1; i < n; i++) {
            f[i + 1] = max(f[i] + a[i], f[i - 1] + a[i - 1] - a[i]);
        }
        return f[n];
    }
};
```

```go [sol-Go]
func maximumTotalCost(a []int) int64 {
	n := len(a)
	f := make([]int, n+1)
	f[1] = a[0]
	for i := 1; i < n; i++ {
		f[i+1] = max(f[i]+a[i], f[i-1]+a[i-1]-a[i])
	}
	return int64(f[n])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{a}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

### 五、空间优化

观察上面的状态转移方程，在计算 $f[i+1]$ 时，只会用到 $f[i]$ 和 $f[i-1]$，不会用到比 $i-1$ 更早的状态。

因此可以像 [打家劫舍](https://www.bilibili.com/video/BV1Xj411K7oF/) 那样，反复利用两个变量。

状态转移方程改为

$$
\begin{align}
f_0 &= f_1\\
f_1 &= \max(f_1 + a[i], f_0 + a[i - 1] - a[i])\\
\end{align}
$$

注意这两个式子要**同时计算**。

初始值 $f_0=0,\ f_1=a[0]$。

答案为 $f_1$。

```py [sol-Python3]
class Solution:
    def maximumTotalCost(self, a: List[int]) -> int:
        f0, f1 = 0, a[0]
        for i in range(1, len(a)):
            f0, f1 = f1, max(f1 + a[i], f0 + a[i - 1] - a[i])
        return f1
```

```py [sol-Python3 pairwise]
class Solution:
    def maximumTotalCost(self, a: List[int]) -> int:
        f0, f1 = 0, a[0]
        for x, y in pairwise(a):
            f0, f1 = f1, max(f1 + y, f0 + x - y)
        return f1
```

```java [sol-Java]
class Solution {
    public long maximumTotalCost(int[] a) {
        long f0 = 0;
        long f1 = a[0];
        for (int i = 1; i < a.length; i++) {
            long tmp = f1;
            f1 = Math.max(f1 + a[i], f0 + a[i - 1] - a[i]);
            f0 = tmp;
        }
        return f1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumTotalCost(vector<int>& a) {
        long long f0 = 0, f1 = a[0];
        for (int i = 1; i < a.size(); i++) {
            long long tmp = f1;
            f1 = max(f1 + a[i], f0 + a[i - 1] - a[i]);
            f0 = tmp;
        }
        return f1;
    }
};
```

```go [sol-Go]
func maximumTotalCost(a []int) int64 {
	f0, f1 := 0, a[0]
	for i := 1; i < len(a); i++ {
		f0, f1 = f1, max(f1+a[i], f0+a[i-1]-a[i])
	}
	return int64(f1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{a}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：状态机 DP

### 一、寻找子问题

对于示例 1，$a[0]$ 一定不用变号（取相反数），我们从 $a[1]$ 开始思考。

分类讨论：

- 分割，把 $a[1]$ 作为子数组的第一个数，接下来需要解决的问题为：考虑 $a[2]$ 到 $a[n-1]$，且 $a[2]$ 如果不是子数组的第一个数，则需要变号，在这种情况下的最大成本和。
- 不分割，那么 $a[1]$ 需要变号，接下来需要解决的问题为：考虑 $a[2]$ 到 $a[n-1]$，且 $a[2]$ 如果不是子数组的第一个数，则无需变号，在这种情况下的最大成本和。

由于分割或不分割都会把原问题变成一个**和原问题相似的、规模更小的子问题**，所以可以用**递归**解决。

> 注：动态规划有「选或不选」和「枚举选哪个」两种基本思考方式。在做题时，可根据题目要求，选择适合题目的一种来思考。本题是「分割」这个操作「选或不选」。

### 二、状态定义与状态转移方程

因为要解决的问题都形如「考虑 $a[i]$ 到 $a[n-1]$，如果 $a[i]$ 不是子数组的第一个数，则不需要/需要变号，在这种情况下的最大成本和」，所以用它作为本题的状态定义 $\textit{dfs}(i,j)$。其中 $j=0$ 表示不变号，$j=1$ 表示变号。

定义

$$
x =
\begin{cases} 
a[i],\ &j=0\\
-a[i],\ &j=1
\end{cases}
$$

分类讨论：

- 分割，把 $a[i]$ 作为子数组的第一个数，接下来需要解决的问题为：考虑 $a[i+1]$ 到 $a[n-1]$，且 $a[i+1]$ 如果不是子数组的第一个数，则需要变号，在这种情况下的最大成本和，即 $\textit{dfs}(i,j) = \textit{dfs}(i+1, 1) + a[i]$。
- 不分割，接下来需要解决的问题为：考虑 $a[i+1]$ 到 $a[n-1]$，且 $a[i+1]$ 如果不是子数组的第一个数，则变号情况与 $a[i]$ 相反，在这种情况下的最大成本和，即 $\textit{dfs}(i,j) = \textit{dfs}(i+1, j\oplus 1) + x$。

这两种情况取最大值，就得到了 $\textit{dfs}(i,j)$，即

$$
\textit{dfs}(i,j) = \max(\textit{dfs}(i+1, 1) + a[i], \textit{dfs}(i+1, j\oplus 1) + x)
$$

递归边界：$\textit{dfs}(n,j)=0$。没有元素，成本和为 $0$。

递归入口：$\textit{dfs}(0,0)$，也就是答案。

### 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,j)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$ 或 $-\infty$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

```py [sol-Python3]
class Solution:
    def maximumTotalCost(self, a: List[int]) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int) -> int:
            if i == len(a):
                return 0
            return max(dfs(i + 1, 1) + a[i],
                       dfs(i + 1, j ^ 1) + (-a[i] if j else a[i]))
        return dfs(0, 0)
```

```java [sol-Java]
class Solution {
    public long maximumTotalCost(int[] a) {
        int n = a.length;
        long[][] memo = new long[n][2];
        for (long[] row : memo) {
            Arrays.fill(row, Long.MIN_VALUE);
        }
        return dfs(0, 0, a, memo);
    }

    private long dfs(int i, int j, int[] a, long[][] memo) {
        if (i == a.length) {
            return 0;
        }
        if (memo[i][j] != Long.MIN_VALUE) { // 之前计算过
            return memo[i][j];
        }
        return memo[i][j] = Math.max(dfs(i + 1, 1, a, memo) + a[i],
                dfs(i + 1, j ^ 1, a, memo) + (j == 0 ? a[i] : -a[i]));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumTotalCost(vector<int>& a) {
        int n = a.size();
        vector<array<long long, 2>> memo(n, {LLONG_MIN, LLONG_MIN});
        auto dfs = [&](auto&& dfs, int i, int j) -> long long {
            if (i == n) {
                return 0;
            }
            auto& res = memo[i][j]; // 注意这里是引用
            if (res != LLONG_MIN) { // 之前计算过
                return res;
            }
            return res = max(dfs(dfs, i + 1, 1) + a[i],
                             dfs(dfs, i + 1, j ^ 1) + (j == 0 ? a[i] : -a[i]));
        };
        return dfs(dfs, 0, 0);
    }
};
```

```go [sol-Go]
func maximumTotalCost(a []int) int64 {
	n := len(a)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{math.MinInt, math.MinInt}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == n {
			return 0
		}
		p := &memo[i][j]
		if *p != math.MinInt { // 之前计算过
			return *p
		}
		res := dfs(i+1, 1) + a[i] // 分割
		r := dfs(i+1, j^1) // 不分割
		if j == 0 {
			r += a[i]
		} else {
			r -= a[i]
		}
		res = max(res, r)
		*p = res // 记忆化
		return res
	}
	return int64(dfs(0, 0))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $a$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。保存多少状态，就需要多少空间。

### 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i][j]$ 的定义和 $\textit{dfs}(i,j)$ 的定义是一样的，都表示考虑 $a[i]$ 到 $a[n-1]$，如果 $a[i]$ 不是子数组的第一个数，则不需要/需要变号，在这种情况下的最大成本和。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i][j] = \max(f[i+1][1] + a[i], f[i+1][j\oplus 1] + x)
$$

即

$$
\begin{align}
&f[i][0] = f[i+1][1] + a[i]\\
&f[i][1] = \max(f[i+1][1]+a[i], f[i+1][0]-a[i])\\
\end{align}
$$

初始值 $f[n][j]=0$，翻译自递归边界 $\textit{dfs}(n,j)=0$。

答案为 $f[0][0]$，翻译自递归入口 $\textit{dfs}(0,0)$。

```py [sol-Python3]
class Solution:
    def maximumTotalCost(self, a: List[int]) -> int:
        n = len(a)
        f = [[0, 0] for _ in range(n + 1)]
        for i in range(n - 1, -1, -1):
            f[i][0] = f[i + 1][1] + a[i]
            f[i][1] = max(f[i + 1][1] + a[i], f[i + 1][0] - a[i])
        return f[0][0]
```

```java [sol-Java]
class Solution {
    public long maximumTotalCost(int[] a) {
        int n = a.length;
        long[][] f = new long[n + 1][2];
        for (int i = n - 1; i >= 0; i--) {
            f[i][0] = f[i + 1][1] + a[i];
            f[i][1] = Math.max(f[i + 1][1] + a[i], f[i + 1][0] - a[i]);
        }
        return f[0][0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumTotalCost(vector<int>& a) {
        int n = a.size();
        vector<array<long long, 2>> f(n + 1);
        for (int i = n - 1; i >= 0; i--) {
            f[i][0] = f[i + 1][1] + a[i];
            f[i][1] = max(f[i + 1][1] + a[i], f[i + 1][0] - a[i]);
        }
        return f[0][0];
    }
};
```

```go [sol-Go]
func maximumTotalCost(a []int) int64 {
	n := len(a)
	f := make([][2]int, n+1)
	for i := n - 1; i >= 0; i-- {
		f[i][0] = f[i+1][1] + a[i]
		f[i][1] = max(f[i+1][1]+a[i], f[i+1][0]-a[i])
	}
	return int64(f[0][0])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $a$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

### 五、空间优化

观察上面的状态转移方程，在计算 $f[i]$ 时，只会用到 $f[i+1]$，不会用到比 $i+1$ 更大的状态。

因此可以像 [买卖股票的最佳时机](https://www.bilibili.com/video/BV1ho4y1W7QK/) 那样，反复利用两个变量。

状态转移方程改为

$$
\begin{align}
&f_0 = f_1 + a[i]\\
&f_1 = \max(f_1+a[i], f_0-a[i])\\
\end{align}
$$

注意这两个式子要**同时计算**。

初始值 $f_0=f_1=0$。

答案为 $f_0$。

```py [sol-Python3]
class Solution:
    def maximumTotalCost(self, a: List[int]) -> int:
        f0 = f1 = 0
        for x in reversed(a):
            f0, f1 = f1 + x, max(f1 + x, f0 - x)
        return f0
```

```java [sol-Java]
class Solution {
    public long maximumTotalCost(int[] a) {
        long f0 = 0;
        long f1 = 0;
        for (int i = a.length - 1; i >= 0; i--) {
            long newF0 = f1 + a[i];
            f1 = Math.max(f1 + a[i], f0 - a[i]);
            f0 = newF0;
        }
        return f0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumTotalCost(vector<int>& a) {
        long long f0 = 0, f1 = 0;
        for (int i = a.size() - 1; i >= 0; i--) {
            long long new_f0 = f1 + a[i];
            f1 = max(f1 + a[i], f0 - a[i]);
            f0 = new_f0;
        }
        return f0;
    }
};
```

```go [sol-Go]
func maximumTotalCost(a []int) int64 {
	f0, f1 := 0, 0
	for i := len(a) - 1; i >= 0; i-- {
		f0, f1 = f1+a[i], max(f1+a[i], f0-a[i])
	}
	return int64(f0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $a$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
