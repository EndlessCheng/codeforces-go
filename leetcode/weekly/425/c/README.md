[本题视频讲解](https://www.bilibili.com/video/BV1fFB4YGEZY/?t=14m36s)，欢迎点赞关注~

## 一、寻找子问题

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

## 二、状态定义与状态转移方程

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

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

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

## 四、1:1 翻译成递推

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

## 五、空间优化

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
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
