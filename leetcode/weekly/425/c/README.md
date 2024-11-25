## 方法一：动态规划

[视频讲解](https://www.bilibili.com/video/BV1fFB4YGEZY/?t=14m36s)，欢迎点赞关注~

### 一、寻找子问题

看示例 1，我们要解决的问题（原问题）是：

- 对于 $\textit{nums}[0]$ 到 $\textit{nums}[4]$，至多执行 $\textit{op}_1=1$ 次操作 1，$\textit{op}_2=1$ 次操作 2 后的所有元素的最小可能和。

从右往左考虑，分类讨论：

- 不操作，那么需要解决的问题为：对于 $\textit{nums}[0]$ 到 $\textit{nums}[3]$，至多执行 $1$ 次操作 1，$1$ 次操作 2 后的所有元素的最小可能和。
- 执行操作 1，那么需要解决的问题为：对于 $\textit{nums}[0]$ 到 $\textit{nums}[3]$，至多执行 $0$ 次操作 1，$1$ 次操作 2 后的所有元素的最小可能和。
- 执行操作 2，那么需要解决的问题为：对于 $\textit{nums}[0]$ 到 $\textit{nums}[3]$，至多执行 $1$ 次操作 1，$0$ 次操作 2 后的所有元素的最小可能和。
- 执行操作 1 和操作 2，那么需要解决的问题为：对于 $\textit{nums}[0]$ 到 $\textit{nums}[3]$，至多执行 $0$ 次操作 1，$0$ 次操作 2 后的所有元素的最小可能和。

这些问题都是**和原问题相似的、规模更小的子问题**，可以用**递归**解决。

> 注 1：从右往左思考，主要是为了方便把递归翻译成递推。从左往右思考也是可以的。
> 
> 注 2：动态规划有「**选或不选**」和「**枚举选哪个**」两种基本思考方式。子序列相邻无关一般是「选或不选」，子序列相邻相关（例如 LIS 问题）一般是「枚举选哪个」。本题用到的是「选或不选」。

### 二、状态定义与状态转移方程

根据上面的讨论，我们需要在递归过程中跟踪以下信息：

- $i$：当前需要操作的元素。
- $\textit{op}_1$：操作 1 的剩余次数。
- $\textit{op}_2$：操作 2 的剩余次数。

因此，定义状态为 $\textit{dfs}(i,\textit{op}_1,\textit{op}_2)$，表示对于 $\textit{nums}[0]$ 到 $\textit{nums}[i]$，至多执行 $\textit{op}_1$ 次操作 1，$\textit{op}_2$ 次操作 2 后的所有元素的最小可能和。

设 $x=\textit{nums}[i]$，分类讨论：

- 不操作，那么需要解决的问题为：对于 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$，至多执行 $\textit{op}_1$ 次操作 1，$\textit{op}_2$ 次操作 2 后的所有元素的最小可能和，即 $\textit{dfs}(i-1,\textit{op}_1,\textit{op}_2)$。由于 $x$ 没变，额外加上 $x$。
- 执行操作 1，那么需要解决的问题为：对于 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$，至多执行 $\textit{op}_1-1$ 次操作 1，$\textit{op}_2$ 次操作 2 后的所有元素的最小可能和，即 $\textit{dfs}(i-1,\textit{op}_1-1,\textit{op}_2)$。额外加上 $\left\lceil\dfrac{x}{2}\right\rceil = \left\lfloor\dfrac{x+1}{2}\right\rfloor$。
- 执行操作 2，那么需要解决的问题为：对于 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$，至多执行 $\textit{op}_1$ 次操作 1，$\textit{op}_2-1$ 次操作 2 后的所有元素的最小可能和，即 $\textit{dfs}(i-1,\textit{op}_1,\textit{op}_2-1)$。额外加上 $x-k$。
- 执行操作 1 和操作 2，那么需要解决的问题为：对于 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$，至多执行 $\textit{op}_1-1$ 次操作 1，$\textit{op}_2-1$ 次操作 2 后的所有元素的最小可能和，即 $\textit{dfs}(i-1,\textit{op}_1-1,\textit{op}_2-1)$。注意：如果能先除再减，那么先除再减更优，否则只能先减再除。

这四种情况取最小值，就得到了 $\textit{dfs}(i,\textit{op}_1,\textit{op}_2)$。

**递归边界**：$\textit{dfs}(-1,\textit{op}_1,\textit{op}_2)=0$。没有元素。

**递归入口**：$\textit{dfs}(n-1,\textit{op}_1,\textit{op}_2)$，这是原问题，也是答案。

### 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,\textit{op}_1,\textit{op}_2)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

```py [sol-Python3]
class Solution:
    def minArraySum(self, nums: List[int], k: int, op1: int, op2: int) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, op1: int, op2: int) -> int:
            if i < 0:
                return 0
            x = nums[i]
            res = dfs(i - 1, op1, op2) + x
            if op1:
                res = min(res, dfs(i - 1, op1 - 1, op2) + (x + 1) // 2)
            if op2 and x >= k:
                res = min(res, dfs(i - 1, op1, op2 - 1) + x - k)
                if op1:
                    y = (x + 1) // 2 - k if (x + 1) // 2 >= k else (x - k + 1) // 2
                    res = min(res, dfs(i - 1, op1 - 1, op2 - 1) + y)
            return res
        return dfs(len(nums) - 1, op1, op2)
```

```java [sol-Java]
class Solution {
    public int minArraySum(int[] nums, int k, int op1, int op2) {
        int n = nums.length;
        int[][][] memo = new int[n][op1 + 1][op2 + 1];
        for (int[][] mat : memo) {
            for (int[] row : mat) {
                Arrays.fill(row, -1); // -1 表示没有计算过
            }
        }
        return dfs(n - 1, op1, op2, k, nums, memo);
    }

    private int dfs(int i, int op1, int op2, int k, int[] nums, int[][][] memo) {
        if (i < 0) {
            return 0;
        }
        if (memo[i][op1][op2] != -1) { // 之前计算过
            return memo[i][op1][op2];
        }
        int x = nums[i];
        int res = dfs(i - 1, op1, op2, k, nums, memo) + x;
        if (op1 > 0) {
            res = Math.min(res, dfs(i - 1, op1 - 1, op2, k, nums, memo) + (x + 1) / 2);
        }
        if (op2 > 0 && x >= k) {
            res = Math.min(res, dfs(i - 1, op1, op2 - 1, k, nums, memo) + x - k);
            if (op1 > 0) {
                int y = (x + 1) / 2 >= k ? (x + 1) / 2 - k : (x - k + 1) / 2;
                res = Math.min(res, dfs(i - 1, op1 - 1, op2 - 1, k, nums, memo) + y);
            }
        }
        return memo[i][op1][op2] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minArraySum(vector<int>& nums, int k, int op1, int op2) {
        int n = nums.size();
        vector memo(n, vector(op1 + 1, vector<int>(op2 + 1, -1))); // -1 表示没有计算过
        auto dfs = [&](auto&& dfs, int i, int op1, int op2) -> int {
            if (i < 0) {
                return 0;
            }
            int& res = memo[i][op1][op2]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            int x = nums[i];
            res = dfs(dfs, i - 1, op1, op2) + x;
            if (op1) {
                res = min(res, dfs(dfs, i - 1, op1 - 1, op2) + (x + 1) / 2);
            }
            if (op2 && x >= k) {
                res = min(res, dfs(dfs, i - 1, op1, op2 - 1) + x - k);
                if (op1) {
                    int y = (x + 1) / 2 >= k ? (x + 1) / 2 - k : (x - k + 1) / 2;
                    res = min(res, dfs(dfs, i - 1, op1 - 1, op2 - 1) + y);
                }
            }
            return res;
        };
        return dfs(dfs, n - 1, op1, op2);
    }
};
```

```go [sol-Go]
func minArraySum(nums []int, k, op1, op2 int) int {
	n := len(nums)
	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, op1+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, op2+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1 // -1 表示没有计算过
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, op1, op2 int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i][op1][op2]
		if *p != -1 { // 之前计算过
			return *p
		}
		x := nums[i]
		res := dfs(i-1, op1, op2) + x
		if op1 > 0 {
			res = min(res, dfs(i-1, op1-1, op2)+(x+1)/2)
		}
		if op2 > 0 && x >= k {
			res = min(res, dfs(i-1, op1, op2-1)+x-k)
			if op1 > 0 {
                var y int
                if (x+1)/2 >= k {
                    y = (x+1)/2 - k // 先除再减更优
                } else {
                    y = (x - k + 1) / 2 // 只能先减再除
                }
				res = min(res, dfs(i-1, op1-1, op2-1)+y)
			}
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n-1, op1, op2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\cdot \textit{op}_1\cdot \textit{op}_2)$，其中 $n$ 为 $\textit{nums}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n\cdot \textit{op}_1\cdot \textit{op}_2)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(n\cdot \textit{op}_1\cdot \textit{op}_2)$。
- 空间复杂度：$\mathcal{O}(n\cdot \textit{op}_1\cdot \textit{op}_2)$。保存多少状态，就需要多少空间。

### 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

$f[i+1][\textit{op}_1][\textit{op}_2]$ 的定义和 $\textit{dfs}(i,\textit{op}_1,\textit{op}_2)$ 的定义是一样的，都表示对于 $\textit{nums}[0]$ 到 $\textit{nums}[i]$，至多执行 $\textit{op}_1$ 次操作 1，$\textit{op}_2$ 次操作 2 后的所有元素的最小可能和。

这里 $+1$ 是为了把避免 $i$ 变成负数。

初始值 $f[0][\textit{op}_1][\textit{op}_2]=0$，翻译自递归边界 $\textit{dfs}(-1,\textit{op}_1,\textit{op}_2)=0$。

答案为 $f[n][\textit{op}_1][\textit{op}_2]$，翻译自递归入口 $\textit{dfs}(n-1,\textit{op}_1,\textit{op}_2)$。

```py [sol-Python3]
class Solution:
    def minArraySum(self, nums: List[int], k: int, op1: int, op2: int) -> int:
        n = len(nums)
        f = [[[0] * (op2 + 1) for _ in range(op1 + 1)] for _ in range(n + 1)]
        for i, x in enumerate(nums):
            for p in range(op1 + 1):
                for q in range(op2 + 1):
                    res = f[i][p][q] + x
                    if p:
                        res = min(res, f[i][p - 1][q] + (x + 1) // 2)
                    if q and x >= k:
                        res = min(res, f[i][p][q - 1] + x - k)
                        if p:
                            y = (x + 1) // 2 - k if (x + 1) // 2 >= k else (x - k + 1) // 2
                            res = min(res, f[i][p - 1][q - 1] + y)
                    f[i + 1][p][q] = res
        return f[n][op1][op2]
```

```java [sol-Java]
class Solution {
    public int minArraySum(int[] nums, int k, int op1, int op2) {
        int n = nums.length;
        int[][][] f = new int[n + 1][op1 + 1][op2 + 1];
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            for (int p = 0; p <= op1; p++) {
                for (int q = 0; q <= op2; q++) {
                    int res = f[i][p][q] + x;
                    if (p > 0) {
                        res = Math.min(res, f[i][p - 1][q] + (x + 1) / 2);
                    }
                    if (q > 0 && x >= k) {
                        res = Math.min(res, f[i][p][q - 1] + x - k);
                        if (p > 0) {
                            int y = (x + 1) / 2 >= k ? (x + 1) / 2 - k : (x - k + 1) / 2;
                            res = Math.min(res, f[i][p - 1][q - 1] + y);
                        }
                    }
                    f[i + 1][p][q] = res;
                }
            }
        }
        return f[n][op1][op2];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minArraySum(vector<int>& nums, int k, int op1, int op2) {
        int n = nums.size();
        vector f(n + 1, vector(op1 + 1, vector<int>(op2 + 1)));
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            for (int p = 0; p <= op1; p++) {
                for (int q = 0; q <= op2; q++) {
                    int res = f[i][p][q] + x;
                    if (p) {
                        res = min(res, f[i][p - 1][q] + (x + 1) / 2);
                    }
                    if (q && x >= k) {
                        res = min(res, f[i][p][q - 1] + x - k);
                        if (p) {
                            int y = (x + 1) / 2 >= k ? (x + 1) / 2 - k : (x - k + 1) / 2;
                            res = min(res, f[i][p - 1][q - 1] + y);
                        }
                    }
                    f[i + 1][p][q] = res;
                }
            }
        }
        return f[n][op1][op2];
    }
};
```

```go [sol-Go]
func minArraySum(nums []int, k, op1, op2 int) int {
	n := len(nums)
	f := make([][][]int, n+1)
	for i := range f {
		f[i] = make([][]int, op1+1)
		for j := range f[i] {
			f[i][j] = make([]int, op2+1)
		}
	}
	for i, x := range nums {
		var y int
		if (x+1)/2 >= k {
			y = (x+1)/2 - k
		} else {
			y = (x - k + 1) / 2
		}
		for p := 0; p <= op1; p++ {
			for q := 0; q <= op2; q++ {
				res := f[i][p][q] + x
				if p > 0 {
					res = min(res, f[i][p-1][q]+(x+1)/2)
				}
				if q > 0 && x >= k {
					res = min(res, f[i][p][q-1]+x-k)
					if p > 0 {
						res = min(res, f[i][p-1][q-1]+y)
					}
				}
				f[i+1][p][q] = res
			}
		}
	}
	return f[n][op1][op2]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\cdot \textit{op}_1\cdot \textit{op}_2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n\cdot \textit{op}_1\cdot \textit{op}_2)$。

### 五、空间优化

观察上面的状态转移方程，在计算 $f[i+1]$ 时，只会用到 $f[i]$，不会用到比 $i$ 更早的状态。

因此可以像 [0-1 背包](https://www.bilibili.com/video/BV16Y411v7Y6/) 那样，去掉第一个维度，把 $f[i+1]$ 和 $f[i]$ 保存到**同一个数组**中。

```py [sol-Python3]
class Solution:
    def minArraySum(self, nums: List[int], k: int, op1: int, op2: int) -> int:
        f = [[0] * (op2 + 1) for _ in range(op1 + 1)]
        for x in nums:
            for p in range(op1, -1, -1):
                for q in range(op2, -1, -1):
                    res = f[p][q] + x
                    if p:
                        res = min(res, f[p - 1][q] + (x + 1) // 2)
                    if q and x >= k:
                        res = min(res, f[p][q - 1] + x - k)
                        if p:
                            y = (x + 1) // 2 - k if (x + 1) // 2 >= k else (x - k + 1) // 2
                            res = min(res, f[p - 1][q - 1] + y)
                    f[p][q] = res
        return f[op1][op2]
```

```java [sol-Java]
class Solution {
    public int minArraySum(int[] nums, int k, int op1, int op2) {
        int[][] f = new int[op1 + 1][op2 + 1];
        for (int x : nums) {
            for (int p = op1; p >= 0; p--) {
                for (int q = op2; q >= 0; q--) {
                    int res = f[p][q] + x;
                    if (p > 0) {
                        res = Math.min(res, f[p - 1][q] + (x + 1) / 2);
                    }
                    if (q > 0 && x >= k) {
                        res = Math.min(res, f[p][q - 1] + x - k);
                        if (p > 0) {
                            int y = (x + 1) / 2 >= k ? (x + 1) / 2 - k : (x - k + 1) / 2;
                            res = Math.min(res, f[p - 1][q - 1] + y);
                        }
                    }
                    f[p][q] = res;
                }
            }
        }
        return f[op1][op2];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minArraySum(vector<int>& nums, int k, int op1, int op2) {
        vector f(op1 + 1, vector<int>(op2 + 1));
        for (int x : nums) {
            for (int p = op1; p >= 0; p--) {
                for (int q = op2; q >= 0; q--) {
                    int res = f[p][q] + x;
                    if (p) {
                        res = min(res, f[p - 1][q] + (x + 1) / 2);
                    }
                    if (q && x >= k) {
                        res = min(res, f[p][q - 1] + x - k);
                        if (p) {
                            int y = (x + 1) / 2 >= k ? (x + 1) / 2 - k : (x - k + 1) / 2;
                            res = min(res, f[p - 1][q - 1] + y);
                        }
                    }
                    f[p][q] = res;
                }
            }
        }
        return f[op1][op2];
    }
};
```

```go [sol-Go]
func minArraySum(nums []int, k, op1, op2 int) int {
	f := make([][]int, op1+1)
	for i := range f {
		f[i] = make([]int, op2+1)
	}
	for _, x := range nums {
		var y int
		if (x+1)/2 >= k {
			y = (x+1)/2 - k
		} else {
			y = (x - k + 1) / 2
		}
		for p := op1; p >= 0; p-- {
			for q := op2; q >= 0; q-- {
				res := f[p][q] + x
				if p > 0 {
					res = min(res, f[p-1][q]+(x+1)/2)
				}
				if q > 0 && x >= k {
					res = min(res, f[p][q-1]+x-k)
					if p > 0 {
						res = min(res, f[p-1][q-1]+y)
					}
				}
				f[p][q] = res
			}
		}
	}
	return f[op1][op2]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\cdot \textit{op}_1\cdot \textit{op}_2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(\textit{op}_1\cdot \textit{op}_2)$。

更多相似题目，见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「**§7.5 多维 DP**」。

## 方法二：贪心

### 核心思路

1. 操作 1（除 $2$ 上取整）应用到大的数上，越大越好。
2. 操作 2（减 $k$）应用到 $\ge k$ 的数上。对于较小的数，需要细致地讨论。

### 分类讨论

- 在 $[2k-1,\infty)$ 中的数，同一个数可以执行两种操作各一次，其中先除再减比先减再除更优。注意 $2k-1$ 执行操作 1 后变成了 $k$，所以还可以再执行一次操作 2。
- 在 $[k,2k-2]$ 中的数，只能先减再除。
  - 对于两个数 $x$ 和 $y$，假设 $x\le y$。如果 $\textit{op}_1= \textit{op}_2=1$，应该怎么操作呢？
    - 操作 2 给 $x$，操作 1 给 $y$，得到 $(x-k) + \left\lceil\dfrac{y}{2}\right\rceil$。
    - 操作 2 给 $y$，操作 1 给 $x$，得到 $(y-k) + \left\lceil\dfrac{x}{2}\right\rceil \ge (x-k) + \left\lceil\dfrac{y}{2}\right\rceil$。
    - 操作 2 和 1 都给 $y$，得到 $x + \left\lceil\dfrac{y-k}{2}\right\rceil \ge (x-k) + \left\lceil\dfrac{y}{2}\right\rceil$。
    - 操作 2 和 1 都给 $x$，得到 $y + \left\lceil\dfrac{x-k}{2}\right\rceil \ge (x-k) + \left\lceil\dfrac{y}{2}\right\rceil$。
    - 所以操作 2 给 $x$，操作 1 给 $y$ 是最优的。
  - 一般地，对于 $[k,2k-2]$ 中的数，**把操作 2 应用到小的数上，操作 1 应用到大的数上**。
  - 如果这两类操作有交集呢？也就是同一个数执行两种操作各一次。
    - 如果 $k$ 是偶数，操作方法不变。
    - 如果 $k$ 是奇数，需要调整执行操作 2 的数，具体见下面的图解。
- 在 $[0,k-1]$ 中的数，只能执行操作 1。留到最后处理。

![lc3366-c.png](https://pic.leetcode.cn/1732495727-JVfVXC-lc3366-c.png)

### 调整执行操作 2 的数

如果 $k$ 是奇数：

- 用一个哈希表 $\textit{cnt}$ 记录 $[k,2k-2]$ 中执行了操作 2 后的数。这些数操作之前是偶数，操作之后是奇数。
- 用一个变量 $\textit{odd}$ 记录 $[k,2k-2]$ 中没有执行操作 2 的奇数的个数。

重新排序后，遍历要执行操作 1 的数 $x$，如果此时 $\textit{odd}>0$ 且 $x$ 在 $\textit{cnt}$ 中，那么我们可以调整（重新安排）一次操作 2，把答案减一。然后把 $x$ 从 $\textit{cnt}$ 中去掉，同时把 $\textit{odd}$ 减一。

最后，从大到小执行操作 1。

```py [sol-Python3]
class Solution:
    def minArraySum(self, nums: List[int], k: int, op1: int, op2: int) -> int:
        nums.sort()
        high = bisect_left(nums, k * 2 - 1)
        low = bisect_left(nums, k)

        # 在 [2k-1,∞) 中的数，直接先除再减（从大到小操作）
        for i in range(len(nums) - 1, high - 1, -1):
            if op1:
                nums[i] = (nums[i] + 1) // 2
                op1 -= 1
            if op2:
                nums[i] -= k
                op2 -= 1

        # 在 [k,2k-2] 中的数，先把小的数 -k
        cnt = defaultdict(int)
        odd = 0
        for i in range(low, high):
            if op2:
                nums[i] -= k
                if k % 2 and nums[i] % 2:
                    # nums[i] 原来是偶数，后面有机会把这次 -k 操作留给奇数，得到更小的答案
                    cnt[nums[i]] += 1
                op2 -= 1
            else:
                odd += nums[i] % 2  # 没有执行 -k 的奇数

        # 重新排序（注：这里可以改用合并两个有序数组的做法）
        nums[:high] = sorted(nums[:high])

        ans = 0
        if k % 2:
            # 调整，对于 [k,2k-2] 中 -k 后还要再 /2 的数，如果原来是偶数，改成给奇数 -k 再 /2，这样答案可以减一
            for i in range(high - op1, high):
                if odd == 0:
                    break
                x = nums[i]
                if x in cnt:
                    cnt[x] -= 1
                    if cnt[x] == 0:
                        del cnt[x]
                    odd -= 1
                    ans -= 1

        # 最后，从大到小执行操作 1
        for i in range(high - 1, -1, -1):
            if op1 == 0:
                break
            nums[i] = (nums[i] + 1) // 2
            op1 -= 1

        return ans + sum(nums)
```

```java [sol-Java]
class Solution {
    public int minArraySum(int[] nums, int k, int op1, int op2) {
        Arrays.sort(nums);
        int high = lowerBound(nums, k * 2 - 1);
        int low = lowerBound(nums, k);

        // 在 [2k-1,∞) 中的数，直接先除再减（从大到小操作）
        for (int i = nums.length - 1; i >= high; i--) {
            if (op1 > 0) {
                nums[i] = (nums[i] + 1) / 2;
                op1--;
            }
            if (op2 > 0) {
                nums[i] -= k;
                op2--;
            }
        }

        // 在 [k,2k-2] 中的数，先把小的数 -k
        Map<Integer, Integer> cnt = new HashMap<>();
        int odd = 0;
        for (int i = low; i < high; i++) {
            if (op2 > 0) {
                nums[i] -= k;
                if (k % 2 > 0 && nums[i] % 2 > 0) {
                    // nums[i] 原来是偶数，后面有机会把这次 -k 操作留给奇数，得到更小的答案
                    cnt.merge(nums[i], 1, Integer::sum); // cnt[nums[i]]++
                }
                op2--;
            } else {
                odd += nums[i] % 2; // 没有执行 -k 的奇数
            }
        }

        // 重新排序（注：这里可以改用合并两个有序数组的做法）
        Arrays.sort(nums, 0, high);

        int ans = 0;
        if (k % 2 > 0) {
            // 调整，对于 [k,2k-2] 中 -k 后还要再 /2 的数，如果原来是偶数，改成给奇数 -k 再 /2，这样答案可以减一
            for (int i = high - op1; i < high && odd > 0; i++) {
                int x = nums[i];
                if (cnt.containsKey(x)) {
                    if (cnt.merge(x, -1, Integer::sum) == 0) { // --cnt[x] == 0
                        cnt.remove(x);
                    }
                    odd--;
                    ans--;
                }
            }
        }

        // 最后，从大到小执行操作 1
        for (int i = high - 1; i >= 0 && op1 > 0; i--) {
            nums[i] = (nums[i] + 1) / 2;
            op1--;
        }

        for (int x : nums) {
            ans += x;
        }
        return ans;
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        int left = -1, right = nums.length;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (nums[mid] >= target) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minArraySum(vector<int>& nums, int k, int op1, int op2) {
        ranges::sort(nums);
        int high = ranges::lower_bound(nums, k * 2 - 1) - nums.begin();
        int low = ranges::lower_bound(nums, k) - nums.begin();

        // 在 [2k-1,∞) 中的数，直接先除再减（从大到小操作）
        for (int i = nums.size() - 1; i >= high; i--) {
            if (op1) {
                nums[i] = (nums[i] + 1) / 2;
                op1--;
            }
            if (op2) {
                nums[i] -= k;
                op2--;
            }
        }

        // 在 [k,2k-2] 中的数，先把小的数 -k
        unordered_multiset<int> st;
        int odd = 0;
        for (int i = low; i < high; i++) {
            if (op2) {
                nums[i] -= k;
                if (k % 2 && nums[i] % 2) {
                    // nums[i] 原来是偶数，后面有机会把这次 -k 操作留给奇数，得到更小的答案
                    st.insert(nums[i]);
                }
                op2--;
            } else {
                odd += nums[i] % 2; // 没有执行 -k 的奇数
            }
        }

        // 重新排序（注：这里可以改用合并两个有序数组的做法）
        sort(nums.begin(), nums.begin() + high);

        int ans = 0;
        if (k % 2) {
            // 调整，对于 [k,2k-2] 中 -k 后还要再 /2 的数，如果原来是偶数，改成给奇数 -k 再 /2，这样答案可以减一
            for (int i = high - op1; i < high && odd; i++) {
                int x = nums[i];
                auto it = st.find(x);
                if (it != st.end()) {
                    st.erase(it);
                    odd--;
                    ans--;
                }
            }
        }

        // 最后，从大到小执行操作 1
        for (int i = high - 1; i >= 0 && op1; i--) {
            nums[i] = (nums[i] + 1) / 2;
            op1--;
        }

        return ans + reduce(nums.begin(), nums.end(), 0);
    }
};
```

```go [sol-Go]
func minArraySum(nums []int, k, op1, op2 int) int {
	slices.Sort(nums)
	high := sort.SearchInts(nums, k*2-1)
	low := sort.SearchInts(nums, k)

	// 在 [2k-1,∞) 中的数，直接先除再减（从大到小操作）
	for i := len(nums) - 1; i >= high; i-- {
		if op1 > 0 {
			nums[i] = (nums[i] + 1) / 2
			op1--
		}
		if op2 > 0 {
			nums[i] -= k
			op2--
		}
	}

	// 在 [k,2k-2] 中的数，先把小的数 -k
	cnt := map[int]int{}
	odd := 0
	for i := low; i < high; i++ {
		if op2 > 0 {
			nums[i] -= k
			if k%2 > 0 && nums[i]%2 > 0 {
				// nums[i] 原来是偶数，后面有机会把这次 -k 操作留给奇数，得到更小的答案
				cnt[nums[i]]++
			}
			op2--
		} else {
			odd += nums[i] % 2 // 没有执行 -k 的奇数
		}
	}

	// 重新排序（注：这里可以改用合并两个有序数组的做法）
	slices.Sort(nums[:high])

	ans := 0
	if k%2 > 0 {
		// 调整，对于 [k,2k-2] 中 -k 后还要再 /2 的数，如果原来是偶数，改成给奇数 -k 再 /2，这样答案可以减一
		for i := high - op1; i < high && odd > 0; i++ {
			x := nums[i]
			if cnt[x] > 0 {
				cnt[x]--
				if cnt[x] == 0 {
					delete(cnt, x)
				}
				odd--
				ans--
			}
		}
	}

	// 最后，从大到小执行操作 1
	for i := high - 1; i >= 0 && op1 > 0; i-- {
		nums[i] = (nums[i] + 1) / 2
		op1--
	}

	for _, x := range nums {
		ans += x
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。如果细致地处理，结合快速选择算法可以做到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见 [贪心题单](https://leetcode.cn/circle/discuss/g6KTKL/) 中的「**§1.1 从最小/最大开始贪心**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
