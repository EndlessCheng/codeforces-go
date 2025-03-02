## 一、寻找子问题

在示例 1 中，我们要解决的问题（原问题）是：

- 剩余元素下标为 $[0,n-1]$，移除所有元素的最小总成本。

分类讨论：

- 移除下标 $1,2$，需要解决的子问题为：剩余元素下标为 $[0]+[3,n-1]$，移除所有元素的最小总成本。
- 移除下标 $0,2$，需要解决的子问题为：剩余元素下标为 $[1]+[3,n-1]$，移除所有元素的最小总成本。
- 移除下标 $0,1$，需要解决的子问题为：剩余元素下标为 $[2]+[3,n-1]$，移除所有元素的最小总成本。

继续。假如现在剩余元素下标为 $[1]+[3,n-1]$，分类讨论：

- 移除下标 $3,4$，需要解决的子问题为：剩余元素下标为 $[1]+[5,n-1]$，移除所有元素的最小总成本。
- 移除下标 $1,4$，需要解决的子问题为：剩余元素下标为 $[3]+[5,n-1]$，移除所有元素的最小总成本。
- 移除下标 $1,3$，需要解决的子问题为：剩余元素下标为 $[4]+[5,n-1]$，移除所有元素的最小总成本。

这些问题都是**和原问题相似的、规模更小的子问题**，可以用**递归**解决。

可以发现，自始至终，至多有一个独立的下标，其他都是连续的后缀。为什么不能有两个独立的下标？如果要产生两个独立的下标，那么必须有一次操作（不是最后一次操作）只去掉了一个数，但这不符合题目要求。

## 二、状态定义与状态转移方程

根据上面的讨论，定义状态为 $\textit{dfs}(i,j)$，表示剩余元素下标为 $[j]+[i,n-1]$，移除所有元素的最小总成本。其中 $j<i$。

分类讨论：

- 移除下标 $i,i+1$，需要解决的子问题为：剩余元素下标为 $[j]+[i+2,n-1]$，移除所有元素的最小总成本，即 $\textit{dfs}(i+2,j)$。
- 移除下标 $j,i+1$，需要解决的子问题为：剩余元素下标为 $[i]+[i+2,n-1]$，移除所有元素的最小总成本，即 $\textit{dfs}(i+2,i)$。
- 移除下标 $j,i$，需要解决的子问题为：剩余元素下标为 $[i+1]+[i+2,n-1]$，移除所有元素的最小总成本，即 $\textit{dfs}(i+2,i+1)$。

这三种情况取最小值，就得到了 $\textit{dfs}(i,j)$，即

$$
\textit{dfs}(i,j) = \min \left\{
\begin{array}{l}
\textit{dfs}(i+2, j)+\max(b, c)   \\
\textit{dfs}(i+2, i)+\max(a, c)   \\
\textit{dfs}(i+2, i+1)+\max(a, b) \\
\end{array}
\right.
$$

其中 $a=\textit{nums}[j],\ b=\textit{nums}[i],\ c=\textit{nums}[i+1]$。

**递归边界**：

- $\textit{dfs}(n,j)=\textit{nums}[j]$。此时只剩下一个数。
- $\textit{dfs}(n-1,j)=\max(\textit{nums}[j],\textit{nums}[n-1])$。此时只剩下两个数。

**递归入口**：$\textit{dfs}(1,0)$。一开始，剩余元素下标可以视作 $[0] + [1,n-1]$。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,j)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。本题由于 $\textit{nums}[i]$ 都是正数，所以也可以初始化成 $0$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

[本题视频讲解](https://www.bilibili.com/video/BV1m39bYiEVV/?t=6m56s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minCost(self, nums: List[int]) -> int:
        n = len(nums)
        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(i: int, j: int) -> int:
            if i == n:
                return nums[j]
            if i == n - 1:
                return max(nums[j], nums[i])
            a, b, c = nums[j], nums[i], nums[i + 1]
            return min(dfs(i + 2, j) + max(b, c),
                       dfs(i + 2, i) + max(a, c),
                       dfs(i + 2, i + 1) + max(a, b))
        ans = dfs(1, 0)
        dfs.cache_clear()  # 避免超出内存限制
        return ans
```

```java [sol-Java]
class Solution {
    public int minCost(int[] nums) {
        int n = nums.length;
        int[][] memo = new int[n - 1][n - 1];
        return dfs(1, 0, nums, memo);
    }

    private int dfs(int i, int j, int[] nums, int[][] memo) {
        int n = nums.length;
        if (i == n) {
            return nums[j];
        }
        if (i == n - 1) {
            return Math.max(nums[j], nums[i]);
        }
        if (memo[i][j] > 0) { // 之前计算过
            return memo[i][j];
        }
        int a = nums[j], b = nums[i], c = nums[i + 1];
        int res1 = dfs(i + 2, j, nums, memo) + Math.max(b, c);
        int res2 = dfs(i + 2, i, nums, memo) + Math.max(a, c);
        int res3 = dfs(i + 2, i + 1, nums, memo) + Math.max(a, b);
        return memo[i][j] = Math.min(Math.min(res1, res2), res3); // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minCost(vector<int>& nums) {
        int n = nums.size();
        vector memo(n - 1, vector<int>(n - 1));
        auto dfs = [&](this auto&& dfs, int i, int j) -> int {
            if (i == n) {
                return nums[j];
            }
            if (i == n - 1) {
                return max(nums[j], nums[i]);
            }
            int& res = memo[i][j]; // 注意这里是引用
            if (res == 0) { // 没有计算过
                int a = nums[j], b = nums[i], c = nums[i + 1];
                res = min({dfs(i + 2, j) + max(b, c),
                           dfs(i + 2, i) + max(a, c),
                           dfs(i + 2, i + 1) + max(a, b)});
            }
            return res;
        };
        return dfs(1, 0);
    }
};
```

```go [sol-Go]
func minCost(nums []int) int {
	n := len(nums)
	memo := make([][]int, n-1)
	for i := range memo {
		memo[i] = make([]int, i)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == n {
			return nums[j]
		}
		if i == n-1 {
			return max(nums[j], nums[i])
		}
		p := &memo[i][j]
		a, b, c := nums[j], nums[i], nums[i+1]
		if *p == 0 { // 没有计算过
			*p = min(dfs(i+2, j)+max(b, c), dfs(i+2, i)+max(a, c), dfs(i+2, i+1)+max(a, b))
		}
		return *p
	}
	return dfs(1, 0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n^2)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n^2)$。保存多少状态，就需要多少空间。

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i][j]$ 的定义和 $\textit{dfs}(i,j)$ 的定义是完全一样的，都表示剩余元素下标为 $[j]+[i,n-1]$，移除所有元素的最小总成本。其中 $j<i$。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i][j] = \min \left\{
\begin{array}{l}
f[i+2][j]+\max(b, c)   \\
f[i+2][i]+\max(a, c)   \\
f[i+2][i+1]+\max(a, b) \\
\end{array}
\right.
$$

其中 $a=\textit{nums}[j],\ b=\textit{nums}[i],\ c=\textit{nums}[i+1]$。

**初始值**：

- $f[n][j]=\textit{nums}[j]$，翻译自递归边界 $\textit{dfs}(n,j)=\textit{nums}[j]$。
- $f[n-1][j]=\max(\textit{nums}[j],\textit{nums}[n-1])$，翻译自递归边界 $\textit{dfs}(n-1,j)=\max(\textit{nums}[j],\textit{nums}[n-1])$。

**答案**：$f[1][0]$，翻译自递归入口 $\textit{dfs}(1,0)$。

### 细节

- 如果 $n$ 是奇数，例如 $n=5$，那么在 $\textit{dfs}$ 中的递归顺序是 $i=1\to 3\to 5$，即最终递归到 $n$，最后一个非递归边界是 $i=n-2$。
- 如果 $n$ 是偶数，例如 $n=6$，那么在 $\textit{dfs}$ 中的递归顺序是 $i=1\to 3\to 5$，即最终递归到 $n-1$，最后一个非递归边界是 $i=n-3$。

所以在写递推的时候，可以从 $i=n-3+(n\bmod 2)$ 开始倒着递推，且每次把 $i$ 减少 $2$。

```py [sol-Python3]
class Solution:
    def minCost(self, nums: List[int]) -> int:
        n = len(nums)
        f = [[0] * i for i in range(n + 1)]
        f[n] = nums
        f[n - 1] = [max(x, nums[-1]) for x in nums]
        for i in range(n - 3 + n % 2, 0, -2):
            b, c = nums[i], nums[i + 1]
            for j in range(i):
                a = nums[j]
                f[i][j] = min(f[i + 2][j] + max(b, c),
                              f[i + 2][i] + max(a, c),
                              f[i + 2][i + 1] + max(a, b))
        return f[1][0]
```

```java [sol-Java]
class Solution {
    public int minCost(int[] nums) {
        int n = nums.length;
        int[][] f = new int[n + 1][n];
        f[n] = nums;
        for (int i = 0; i < n; i++) {
            f[n - 1][i] = Math.max(nums[i], nums[n - 1]);
        }
        for (int i = n - 3 + n % 2; i > 0; i -= 2) {
            int b = nums[i], c = nums[i + 1];
            for (int j = 0; j < i; j++) {
                int a = nums[j];
                int res1 = f[i + 2][j] + Math.max(b, c);
                int res2 = f[i + 2][i] + Math.max(a, c);
                int res3 = f[i + 2][i + 1] + Math.max(a, b);
                f[i][j] = Math.min(Math.min(res1, res2), res3);
            }
        }
        return f[1][0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minCost(vector<int>& nums) {
        int n = nums.size();
        vector f(n + 1, vector<int>(n));
        f[n] = nums;
        for (int i = 0; i < n; i++) {
            f[n - 1][i] = max(nums[i], nums[n - 1]);
        }
        for (int i = n - 3 + n % 2; i > 0; i -= 2) {
            int b = nums[i], c = nums[i + 1];
            for (int j = 0; j < i; j++) {
                int a = nums[j];
                f[i][j] = min({f[i + 2][j] + max(b, c),
                               f[i + 2][i] + max(a, c),
                               f[i + 2][i + 1] + max(a, b)});
            }
        }
        return f[1][0];
    }
};
```

```go [sol-Go]
func minCost(nums []int) int {
	n := len(nums)
	f := make([][]int, n+1)
	f[n] = nums
	f[n-1] = make([]int, n)
	for i, x := range nums {
		f[n-1][i] = max(x, nums[n-1])
	}
	for i := n - 3 + n%2; i > 0; i -= 2 {
		f[i] = make([]int, i)
		b, c := nums[i], nums[i+1]
		for j, a := range nums[:i] {
			f[i][j] = min(f[i+2][j]+max(b, c), f[i+2][i]+max(a, c), f[i+2][i+1]+max(a, b))
		}
	}
	return f[1][0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 五、空间优化

观察上面的状态转移方程，在计算 $f[i]$ 时，只会用到 $f[i+2]$，不会用到比 $i+2$ 更大的状态。

因此可以去掉第一个维度，把 $f[i+2]$ 和 $f[i]$ 保存到**同一个数组**中。

不了解这个技巧的同学可以看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

```py [sol-Python3]
class Solution:
    def minCost(self, nums: List[int]) -> int:
        n = len(nums)
        f = nums.copy() if n % 2 else [max(x, nums[-1]) for x in nums]
        for i in range(n - 3 + n % 2, 0, -2):
            b, c = nums[i], nums[i + 1]
            for j in range(i):
                a = nums[j]
                f[j] = min(f[j] + max(b, c), f[i] + max(a, c), f[i + 1] + max(a, b))
        return f[0]
```

```java [sol-Java]
class Solution {
    public int minCost(int[] nums) {
        int n = nums.length;
        int[] f;
        if (n % 2 > 0) {
            f = nums.clone();
        } else {
            f = new int[n];
            for (int i = 0; i < n; i++) {
                f[i] = Math.max(nums[i], nums[n - 1]);
            }
        }
        for (int i = n - 3 + n % 2; i > 0; i -= 2) {
            int b = nums[i], c = nums[i + 1];
            for (int j = 0; j < i; j++) {
                int a = nums[j];
                int res1 = f[j] + Math.max(b, c);
                int res2 = f[i] + Math.max(a, c);
                int res3 = f[i + 1] + Math.max(a, b);
                f[j] = Math.min(Math.min(res1, res2), res3);
            }
        }
        return f[0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minCost(vector<int>& nums) {
        int n = nums.size();
        vector<int> f;
        if (n % 2) {
            f = nums;
        } else {
            f.resize(n);
            for (int i = 0; i < n; i++) {
                f[i] = max(nums[i], nums[n - 1]);
            }
        }
        for (int i = n - 3 + n % 2; i > 0; i -= 2) {
            int b = nums[i], c = nums[i + 1];
            for (int j = 0; j < i; j++) {
                int a = nums[j];
                f[j] = min({f[j] + max(b, c), f[i] + max(a, c), f[i + 1] + max(a, b)});
            }
        }
        return f[0];
    }
};
```

```go [sol-Go]
func minCost(nums []int) int {
	n := len(nums)
	var f []int
	if n%2 == 1 {
		f = slices.Clone(nums)
	} else {
		f = make([]int, n)
		for i, x := range nums {
			f[i] = max(x, nums[n-1])
		}
	}
	for i := n - 3 + n%2; i > 0; i -= 2 {
		b, c := nums[i], nums[i+1]
		for j, a := range nums[:i] {
			f[j] = min(f[j]+max(b, c), f[i]+max(a, c), f[i+1]+max(a, b))
		}
	}
	return f[0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「**§7.5 多维 DP**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
