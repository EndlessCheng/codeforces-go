[视频讲解](https://www.bilibili.com/video/BV1tw411q7VZ/) 第三题。

## 前置知识

请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://b23.tv/72onpYq)

## 写法一：记忆化搜索

把大于 $k$ 的元素视作 $k$。

由于大于 $3$ 的子数组必然包含等于 $3$ 的子数组，问题转换成：

- 每个长为 $3$ 的子数组，都需要包含至少一个 $k$。

考虑最后一个元素「选或不选」，即是否增大：

- 增大到 $k$：那么对于左边那个数来说，它右边就有一个 $k$ 了。
- 不增大：那么对于左边那个数来说，它右边有一个没有增大的数。

进一步地，如果倒数第二个数也不增大，那么对于倒数第三个数，它右边就有两个没有增大的数，那么它一定要增大（不用管右边那两个数是否为 $k$，因为下面的 $\textit{dfs}$ 会考虑到所有的情况，取最小值）。

因此，用 $i$ 表示当前位置，$j=0,1,2$ 表示下标 $i$ 的右边有（连续）零个/一个/两个没有增大的数。定义 $\textit{dfs}(i,j)$ 表示在该状态下对于前 $i$ 个数的最小递增运算数。

考虑 $\textit{nums}[i]$ 是否增大到 $k$：

- 增大，即 $\textit{dfs}(i-1,0) + \max(k-\textit{nums}[i], 0)$。
- 如果 $j<2$，则可以不增大，即 $\textit{dfs}(i-1,j+1)$。

这两种情况取最小值，就得到了 $\textit{dfs}(i,j)$。

递归边界：当 $i<0$ 时返回 $0$。

递归入口：$\textit{dfs}(n-1,0)$，即答案。

```py [sol-Python3]
class Solution:
    def minIncrementOperations(self, nums: List[int], k: int) -> int:
        @cache
        def dfs(i: int, j: int) -> int:
            if i < 0:
                return 0
            res = dfs(i - 1, 0) + max(k - nums[i], 0)  # nums[i] 增大
            if j < 2:
                res = min(res, dfs(i - 1, j + 1))  # nums[i] 不增大
            return res
        return dfs(len(nums) - 1, 0)
```

```java [sol-Java]
class Solution {
    public long minIncrementOperations(int[] nums, int k) {
        int n = nums.length;
        long[][] memo = new long[n][3];
        for (long[] m : memo) {
            Arrays.fill(m, -1); // -1 表示没有计算过
        }
        return dfs(n - 1, 0, memo, nums, k);
    }
     
    private long dfs(int i, int j, long[][] memo, int[] nums, int k) {
        if (i < 0) { // 递归边界
            return 0;
        }
        if (memo[i][j] != -1) { // 之前计算过
            return memo[i][j];
        }
        long res = dfs(i - 1, 0, memo, nums, k) + Math.max(k - nums[i], 0); // nums[i] 增大
        if (j < 2) res = Math.min(res, dfs(i - 1, j + 1, memo, nums, k)); // nums[i] 不增大
        return memo[i][j] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<array<long long, 3>> memo;
    
    long long dfs(vector<int> &nums, int i, int j, int k) {
        if (i < 0) {
            return 0;
        }
        auto &res = memo[i][j]; // 注意这里是引用
        if (res != -1) { // 之前计算过
            return res;
        }
        res = dfs(nums, i - 1, 0, k) + max(k - nums[i], 0); // nums[i] 增大
        if (j < 2) res = min(res, dfs(nums, i - 1, j + 1, k)); // nums[i] 不增大
        return res;
    }

public:
    long long minIncrementOperations(vector<int> &nums, int k) {
        int n = nums.size();
        memo = vector<array<long long, 3>>(n, {-1, -1, -1}); // -1 表示没有计算过
        return dfs(nums, n - 1, 0, k);
    }
};
```

```go [sol-Go]
func minIncrementOperations(nums []int, k int) int64 {
	n := len(nums)
	memo := make([][3]int, n)
	for i := range memo {
		memo[i] = [3]int{-1, -1, -1}
	}
	var dfs func(int, int) int
	dfs = func(i, j int)  int {
		if i < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		res := dfs(i-1, 0) + max(k-nums[i], 0) // nums[i] 增大
		if j < 2 {
			res = min(res, dfs(i-1, j+1)) // nums[i] 不增大
		}
		*p = res
		return res
	}
	return int64(dfs(n-1, 0))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二：递推

按照视频中讲的方法，1:1 翻译成递推。

由于 $f[i]$ 只会用到 $f[i-1]$，所以可以像 0-1 背包那样压缩空间，具体请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)

```py [sol-Python3]
class Solution:
    def minIncrementOperations(self, nums: List[int], k: int) -> int:
        f0 = f1 = f2 = 0
        for x in nums:
            inc = f0 + max(k - x, 0)
            f0 = min(inc, f1)
            f1 = min(inc, f2)
            f2 = inc
        return f0
```

```cpp [sol-C++]
class Solution {
public:
    long long minIncrementOperations(vector<int> &nums, int k) {
        long long f0 = 0, f1 = 0, f2 = 0;
        for (int x : nums) {
            long long inc = f0 + max(k - x, 0);
            f0 = min(inc, f1);
            f1 = min(inc, f2);
            f2 = inc;
        }
        return f0;
    }
};
```

```java [sol-Java]
class Solution {
    public long minIncrementOperations(int[] nums, int k) {
        long f0 = 0, f1 = 0, f2 = 0;
        for (int x : nums) {
            long inc = f0 + Math.max(k - x, 0);
            f0 = Math.min(inc, f1);
            f1 = Math.min(inc, f2);
            f2 = inc;
        }
        return f0;
    }
}
```

```go [sol-Go]
func minIncrementOperations(nums []int, k int) int64 {
	var f0, f1, f2 int
	for _, x := range nums {
		inc := f0 + max(k-x, 0)
		f0 = min(inc, f1)
		f1 = min(inc, f2)
		f2 = inc
	}
	return int64(f0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
